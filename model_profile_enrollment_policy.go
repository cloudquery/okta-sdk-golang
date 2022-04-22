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

type ProfileEnrollmentPolicy struct {
	Embedded map[string]interface{} `json:"_embedded,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Name string `json:"name,omitempty"`
	Priority int32 `json:"priority,omitempty"`
	Status *LifecycleStatus `json:"status,omitempty"`
	System bool `json:"system,omitempty"`
	Type_ *PolicyType `json:"type,omitempty"`
}
