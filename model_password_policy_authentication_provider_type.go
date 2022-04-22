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

type PasswordPolicyAuthenticationProviderType string

// List of PasswordPolicyAuthenticationProviderType
const (
	ACTIVE_DIRECTORY_PasswordPolicyAuthenticationProviderType PasswordPolicyAuthenticationProviderType = "ACTIVE_DIRECTORY"
	ANY_PasswordPolicyAuthenticationProviderType PasswordPolicyAuthenticationProviderType = "ANY"
	LDAP_PasswordPolicyAuthenticationProviderType PasswordPolicyAuthenticationProviderType = "LDAP"
	OKTA_PasswordPolicyAuthenticationProviderType PasswordPolicyAuthenticationProviderType = "OKTA"
)
