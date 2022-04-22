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

type OAuth2Scope struct {
	Consent *OAuth2ScopeConsentType `json:"consent,omitempty"`
	Default_ bool `json:"default,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Id string `json:"id,omitempty"`
	MetadataPublish *OAuth2ScopeMetadataPublish `json:"metadataPublish,omitempty"`
	Name string `json:"name,omitempty"`
	System bool `json:"system,omitempty"`
}
