package okta

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDoWithRetriesKeepsRequestQueryOnReauthorization(t *testing.T) {
	const query = "limit=200&filter=status+eq+%22ACTIVE%22"
	var gotQueries []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/oauth2/v1/token" {
			w.Header().Set("Content-Type", "application/json")
			require.NoError(t, json.NewEncoder(w).Encode(&RequestAccessToken{
				AccessToken: "test-access-token",
				ExpiresIn:   3600,
				TokenType:   "Bearer",
			}))
			return
		}

		gotQueries = append(gotQueries, r.URL.RawQuery)
		if len(gotQueries) == 1 {
			w.Header().Set("Date", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
			w.Header().Set("X-Rate-Limit-Reset", "1")
			w.Header().Set("X-Okta-Request-Id", "test-request-id")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	cfg := &Configuration{
		HTTPClient: &http.Client{Timeout: time.Minute},
		Host:       strings.TrimPrefix(server.URL, "http://"),
		Scheme:     "http",
	}
	cfg.Okta.Client.RateLimit.MaxRetries = 3
	cfg.Okta.Client.RateLimit.MaxBackoff = 1
	cfg.Okta.Client.AuthorizationMode = "PrivateKey"
	cfg.Okta.Client.OrgUrl = server.URL
	cfg.Okta.Client.ClientId = "test-client-id"
	privateKeyPem, err := generatePrivateKeyPem()
	require.NoError(t, err)
	cfg.Okta.Client.PrivateKey = string(privateKeyPem)

	client := NewAPIClient(cfg)

	req, err := http.NewRequest(http.MethodGet, server.URL+"/api/v1/apps?"+query, nil)
	require.NoError(t, err)

	resp, err := client.doWithRetries(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, []string{query, query}, gotQueries)
	require.Equal(t, query, req.URL.RawQuery)
}
