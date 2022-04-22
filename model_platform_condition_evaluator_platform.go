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

type PlatformConditionEvaluatorPlatform struct {
	Os *PlatformConditionEvaluatorPlatformOperatingSystem `json:"os,omitempty"`
	Type_ *PolicyPlatformType `json:"type,omitempty"`
}
