package spaciv

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	apiKeyTokenPath      = "/v1/auth/api-key-token"
	apiKeyTokenOperation = "AuthenticationAPIService.CreateAuthApiKeyToken"
	tokenExpirySkew      = 30 * time.Second
)

func NewAPIKeyClient(apiKey string) *APIClient {
	return NewAPIKeyClientWithConfig(apiKey, nil)
}

func NewAPIKeyClientWithConfig(apiKey string, cfg *Configuration) *APIClient {
	if cfg == nil {
		cfg = NewConfiguration()
	}
	copied := *cfg
	cfg = &copied
	base := cfg.HTTPClient
	if base == nil {
		base = http.DefaultClient
	}
	authenticated := *base
	authenticated.Transport = newAPIKeyTransport(apiKey, apiKeyTokenURL(cfg), configuredOrigins(cfg), base.Transport)
	cfg.HTTPClient = &authenticated
	return NewAPIClient(cfg)
}

var contextAuthenticatedOperation = contextKey("authenticatedOperation")

type operationOrigins struct {
	fallback    string
	byOperation map[string]string
}

func (o operationOrigins) of(operation string) string {
	if origin, ok := o.byOperation[operation]; ok {
		return origin
	}
	return o.fallback
}

func newAPIKeyTransport(apiKey, authURL string, origins operationOrigins, base http.RoundTripper) *apiKeyTransport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &apiKeyTransport{apiKey: apiKey, authURL: authURL, origins: origins, base: base, now: time.Now}
}

func apiKeyTokenURL(cfg *Configuration) string {
	serverURL, err := cfg.ServerURLWithContext(context.Background(), apiKeyTokenOperation)
	if err != nil {
		return ""
	}
	u, err := url.Parse(serverURL + apiKeyTokenPath)
	if err != nil {
		return ""
	}
	applyOverrides(cfg, u)
	return u.String()
}

func configuredOrigins(cfg *Configuration) operationOrigins {
	origins := operationOrigins{fallback: defaultOrigin(cfg, cfg.Servers), byOperation: map[string]string{}}
	for operation, servers := range cfg.OperationServers {
		origins.byOperation[operation] = defaultOrigin(cfg, servers)
	}
	return origins
}

func defaultOrigin(cfg *Configuration, servers ServerConfigurations) string {
	if len(servers) == 0 {
		return ""
	}
	u, err := url.Parse(servers[0].URL)
	if err != nil || u.Host == "" {
		return ""
	}
	applyOverrides(cfg, u)
	return originOf(u)
}

func applyOverrides(cfg *Configuration, u *url.URL) {
	if cfg.Host != "" {
		u.Host = cfg.Host
	}
	if cfg.Scheme != "" {
		u.Scheme = cfg.Scheme
	}
}

func originOf(u *url.URL) string {
	return strings.ToLower(u.Scheme + "://" + u.Host)
}

type apiKeyTransport struct {
	apiKey  string
	authURL string
	origins operationOrigins
	base    http.RoundTripper
	now     func() time.Time

	mu        sync.Mutex
	token     string
	expiresAt time.Time
}

func (t *apiKeyTransport) authenticates(req *http.Request) bool {
	operation, ok := req.Context().Value(contextAuthenticatedOperation).(string)
	return ok && req.Response == nil && originOf(req.URL) == t.origins.of(operation)
}

func (t *apiKeyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if !t.authenticates(req) {
		return t.base.RoundTrip(req)
	}
	token, err := t.currentToken(req, "")
	if err != nil {
		return nil, err
	}
	res, err := t.send(req, token)
	if err != nil || res.StatusCode != http.StatusUnauthorized {
		return res, err
	}
	if req.Body != nil && req.GetBody == nil {
		return res, nil
	}

	res.Body.Close()
	token, err = t.currentToken(req, token)
	if err != nil {
		return nil, err
	}
	retry := req.Clone(req.Context())
	if req.GetBody != nil {
		if retry.Body, err = req.GetBody(); err != nil {
			return nil, err
		}
	}
	return t.send(retry, token)
}

func (t *apiKeyTransport) send(req *http.Request, token string) (*http.Response, error) {
	authorised := req.Clone(req.Context())
	authorised.Header.Set("Authorization", "Bearer "+token)
	return t.base.RoundTrip(authorised)
}

func (t *apiKeyTransport) currentToken(req *http.Request, rejected string) (string, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.token != "" && t.token != rejected && t.now().Add(tokenExpirySkew).Before(t.expiresAt) {
		return t.token, nil
	}

	token, err := t.exchange(req)
	if err != nil {
		return "", err
	}
	t.token = token
	t.expiresAt = tokenExpiry(token, t.now())
	return token, nil
}

func (t *apiKeyTransport) exchange(req *http.Request) (string, error) {
	body, err := json.Marshal(map[string]string{"key": t.apiKey})
	if err != nil {
		return "", err
	}
	exchangeReq, err := http.NewRequestWithContext(req.Context(), http.MethodPost, t.authURL, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	exchangeReq.Header.Set("Content-Type", "application/json")
	exchangeReq.Header.Set("Accept", "application/json")

	res, err := t.base.RoundTrip(exchangeReq)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("spaciv: exchanging API key for an access token: %s: %s", res.Status, strings.TrimSpace(string(resBody)))
	}

	var parsed AuthAPIKeyTokenResBody
	if err := json.Unmarshal(resBody, &parsed); err != nil {
		return "", fmt.Errorf("spaciv: decoding access token response: %w", err)
	}
	return parsed.AccessToken, nil
}

func tokenExpiry(token string, now time.Time) time.Time {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return now
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return now
	}
	var claims struct {
		Exp int64 `json:"exp"`
	}
	if err := json.Unmarshal(payload, &claims); err != nil || claims.Exp == 0 {
		return now
	}
	return time.Unix(claims.Exp, 0)
}
