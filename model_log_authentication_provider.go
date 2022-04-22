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

type LogAuthenticationProvider string

// List of LogAuthenticationProvider
const (
	OKTA_AUTHENTICATION_PROVIDER_LogAuthenticationProvider LogAuthenticationProvider = "OKTA_AUTHENTICATION_PROVIDER"
	ACTIVE_DIRECTORY_LogAuthenticationProvider LogAuthenticationProvider = "ACTIVE_DIRECTORY"
	LDAP_LogAuthenticationProvider LogAuthenticationProvider = "LDAP"
	FEDERATION_LogAuthenticationProvider LogAuthenticationProvider = "FEDERATION"
	SOCIAL_LogAuthenticationProvider LogAuthenticationProvider = "SOCIAL"
	FACTOR_PROVIDER_LogAuthenticationProvider LogAuthenticationProvider = "FACTOR_PROVIDER"
)
