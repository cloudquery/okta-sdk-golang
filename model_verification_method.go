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

type VerificationMethod struct {
	Constraints []AccessPolicyConstraints `json:"constraints,omitempty"`
	FactorMode string `json:"factorMode,omitempty"`
	ReauthenticateIn string `json:"reauthenticateIn,omitempty"`
	Type_ string `json:"type,omitempty"`
}
