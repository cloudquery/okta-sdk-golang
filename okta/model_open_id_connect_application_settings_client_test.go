package okta

import (
	"encoding/json"
	"testing"
)

const oidcApplicationWithoutGrantTypes = `[{"id":"0oa1","name":"oidc_client","label":"OIDC app","status":"ACTIVE","signOnMode":"OPENID_CONNECT","credentials":{"oauthClient":{"client_id":"0oa1","autoKeyRotation":true,"token_endpoint_auth_method":"client_secret_basic"}},"settings":{"oauthClient":{"application_type":"web","redirect_uris":["https://example.com/callback"],"response_types":["code"]}}}]`

func TestListApplicationsOpenIdConnectWithoutGrantTypes(t *testing.T) {
	var items []ListApplications200ResponseInner
	if err := json.Unmarshal([]byte(oidcApplicationWithoutGrantTypes), &items); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	app := items[0].OpenIdConnectApplication
	if app == nil {
		t.Fatal("OpenIdConnectApplication is nil")
	}
	if app.GetId() != "0oa1" {
		t.Fatalf("id = %q, want 0oa1", app.GetId())
	}
	if grantTypes := app.Settings.OauthClient.GetGrantTypes(); len(grantTypes) != 0 {
		t.Fatalf("grant_types = %v, want empty", grantTypes)
	}
}
