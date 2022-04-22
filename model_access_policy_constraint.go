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

type AccessPolicyConstraint struct {
	Methods []string `json:"methods,omitempty"`
	ReauthenticateIn string `json:"reauthenticateIn,omitempty"`
	Types []string `json:"types,omitempty"`
}
