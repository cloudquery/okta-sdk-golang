package okta

import (
	"encoding/json"
	"testing"
)

const swaThreeFieldApplication = `[{"id":"0oa2","name":"template_swa3field","label":"SWA 3 field app","status":"ACTIVE","signOnMode":"BROWSER_PLUGIN","settings":{"app":{"targetURL":"https://example.com/login","buttonSelector":"#login","passwordSelector":"#password","userNameSelector":"#username","extraFieldSelector":"#org","extraFieldValue":"acme","loginUrlRegex":null}}}]`

func TestListApplicationsBrowserPluginWithoutSwaRequiredFields(t *testing.T) {
	var items []ListApplications200ResponseInner
	if err := json.Unmarshal([]byte(swaThreeFieldApplication), &items); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	app := items[0].BrowserPluginApplication
	if app == nil {
		t.Fatal("BrowserPluginApplication is nil")
	}
	appSettings := app.Settings.GetApp()
	if got := appSettings.GetTargetURL(); got != "https://example.com/login" {
		t.Fatalf("targetURL = %q, want https://example.com/login", got)
	}
	if appSettings.HasUrl() {
		t.Fatalf("url = %q, want unset", appSettings.GetUrl())
	}
}
