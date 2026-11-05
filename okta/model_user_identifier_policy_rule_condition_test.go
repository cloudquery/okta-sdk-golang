package okta

import (
	"encoding/json"
	"testing"
)

const idpDiscoveryRuleWithoutUserIdentifierPatterns = `[{"id":"rul3","name":"IdP routing","type":"IDP_DISCOVERY","status":"ACTIVE","conditions":{"userIdentifier":{"attribute":"login"}}}]`

func TestListPolicyRulesIdpDiscoveryWithoutUserIdentifierPatterns(t *testing.T) {
	var items []ListPolicyRules200ResponseInner
	if err := json.Unmarshal([]byte(idpDiscoveryRuleWithoutUserIdentifierPatterns), &items); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("got %d items, want 1", len(items))
	}
	rule := items[0].IdpDiscoveryPolicyRule
	if rule == nil {
		t.Fatal("IdpDiscoveryPolicyRule is nil")
	}
	userIdentifier := rule.Conditions.GetUserIdentifier()
	if got := userIdentifier.GetAttribute(); got != "login" {
		t.Fatalf("attribute = %q, want login", got)
	}
	if userIdentifier.HasPatterns() {
		t.Fatal("patterns is set, want unset")
	}
}
