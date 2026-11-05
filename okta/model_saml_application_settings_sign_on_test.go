package okta

import (
	"encoding/json"
	"testing"
)

const samlApplicationWithOverridesOnly = `[{"id":"0oa3","name":"example_saml","label":"SAML app","status":"ACTIVE","signOnMode":"SAML_2_0","settings":{"signOn":{"defaultRelayState":null,"ssoAcsUrlOverride":null,"audienceOverride":null,"recipientOverride":null,"destinationOverride":null,"attributeStatements":[]}}}]`

func TestListApplicationsSamlWithoutSignOnRequiredFields(t *testing.T) {
	var items []ListApplications200ResponseInner
	if err := json.Unmarshal([]byte(samlApplicationWithOverridesOnly), &items); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	app := items[0].SamlApplication
	if app == nil {
		t.Fatal("SamlApplication is nil")
	}
	signOn := app.Settings.GetSignOn()
	if signOn.HasSsoAcsUrl() {
		t.Fatalf("ssoAcsUrl = %q, want unset", signOn.GetSsoAcsUrl())
	}
}
