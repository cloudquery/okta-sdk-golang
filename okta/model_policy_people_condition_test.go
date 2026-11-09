package okta

import (
	"encoding/json"
	"testing"
)

const accessPolicyRulesWithPartialPeople = `[{"id":"rul1","name":"Groups only","type":"ACCESS_POLICY","status":"ACTIVE","conditions":{"people":{"groups":{"include":["00g1"]}}}},{"id":"rul2","name":"Users only","type":"ACCESS_POLICY","status":"ACTIVE","conditions":{"people":{"users":{"exclude":["00u1"]}}}}]`

func TestListPolicyRulesAccessPolicyWithPartialPeopleCondition(t *testing.T) {
	var items []ListPolicyRules200ResponseInner
	if err := json.Unmarshal([]byte(accessPolicyRulesWithPartialPeople), &items); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("got %d items, want 2", len(items))
	}

	groupsOnly := items[0].AccessPolicyRule
	if groupsOnly == nil {
		t.Fatal("first AccessPolicyRule is nil")
	}
	groupsPeople := groupsOnly.Conditions.GetPeople()
	if groupsPeople.HasUsers() {
		t.Fatal("users condition is set, want unset")
	}
	groups := groupsPeople.GetGroups()
	if got := groups.GetInclude(); len(got) != 1 || got[0] != "00g1" {
		t.Fatalf("groups.include = %v, want [00g1]", got)
	}
	if groups.HasExclude() {
		t.Fatal("groups.exclude is set, want unset")
	}

	usersOnly := items[1].AccessPolicyRule
	if usersOnly == nil {
		t.Fatal("second AccessPolicyRule is nil")
	}
	usersPeople := usersOnly.Conditions.GetPeople()
	if usersPeople.HasGroups() {
		t.Fatal("groups condition is set, want unset")
	}
	users := usersPeople.GetUsers()
	if got := users.GetExclude(); len(got) != 1 || got[0] != "00u1" {
		t.Fatalf("users.exclude = %v, want [00u1]", got)
	}
	if users.HasInclude() {
		t.Fatal("users.include is set, want unset")
	}
}
