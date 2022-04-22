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

type PolicySubject struct {
	Filter string `json:"filter,omitempty"`
	Format []string `json:"format,omitempty"`
	MatchAttribute string `json:"matchAttribute,omitempty"`
	MatchType *PolicySubjectMatchType `json:"matchType,omitempty"`
	UserNameTemplate *PolicyUserNameTemplate `json:"userNameTemplate,omitempty"`
}
