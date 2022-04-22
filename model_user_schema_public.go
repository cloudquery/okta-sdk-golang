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

type UserSchemaPublic struct {
	Id string `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	Properties map[string]UserSchemaAttribute `json:"properties,omitempty"`
	Required []string `json:"required,omitempty"`
}
