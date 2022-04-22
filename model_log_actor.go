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

type LogActor struct {
	AlternateId string `json:"alternateId,omitempty"`
	Detail map[string]interface{} `json:"detail,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Id string `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
}
