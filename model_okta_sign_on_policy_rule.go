/*
 * Okta API
 *
 * Allows customers to easily access the Okta API
 *
 * API version: 2.10.0
 * Contact: devex-public@okta.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type OktaSignOnPolicyRule struct {
	Created time.Time `json:"created,omitempty"`
	Id string `json:"id,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Priority int32 `json:"priority,omitempty"`
	Status *LifecycleStatus `json:"status,omitempty"`
	System bool `json:"system,omitempty"`
	Type_ *PolicyRuleType `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Conditions *OktaSignOnPolicyRuleConditions `json:"conditions,omitempty"`
	Actions *OktaSignOnPolicyRuleActions `json:"actions,omitempty"`
}
