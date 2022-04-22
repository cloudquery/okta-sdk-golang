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

type GroupRulePeopleCondition struct {
	Groups *GroupRuleGroupCondition `json:"groups,omitempty"`
	Users *GroupRuleUserCondition `json:"users,omitempty"`
}
