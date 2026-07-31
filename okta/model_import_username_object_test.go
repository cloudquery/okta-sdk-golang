package okta

import (
	"encoding/json"
	"testing"
)

const inboundProvisioningFeature = `[{"name":"INBOUND_PROVISIONING","status":"ENABLED","description":"In-bound provisioning settings for provisioning users from an application to Okta","capabilities":{"importSettings":{"username":{"userNameFormat":"EMAIL"},"schedule":{"status":"DISABLED","fullImport":null,"incrementalImport":null}},"importRules":{"userCreateAndMatch":{"exactMatchCriteria":"EMAIL","allowPartialMatch":true,"autoConfirmPartialMatch":false,"autoConfirmExactMatch":false,"autoConfirmNewUsers":false,"autoActivateNewUsers":false}}}}]`

func TestListFeaturesForApplicationInboundProvisioning(t *testing.T) {
	var items []ListFeaturesForApplication200ResponseInner
	if err := json.Unmarshal([]byte(inboundProvisioningFeature), &items); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	feature := items[0].InboundProvisioningApplicationFeature
	if feature == nil {
		t.Fatal("InboundProvisioningApplicationFeature is nil")
	}
	got := feature.Capabilities.ImportSettings.Username.GetUserNameFormat()
	if got != "EMAIL" {
		t.Fatalf("userNameFormat = %q, want EMAIL", got)
	}
}
