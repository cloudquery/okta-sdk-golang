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

type ApplicationSignOnMode string

// List of ApplicationSignOnMode
const (
	BOOKMARK_ApplicationSignOnMode ApplicationSignOnMode = "BOOKMARK"
	BASIC_AUTH_ApplicationSignOnMode ApplicationSignOnMode = "BASIC_AUTH"
	BROWSER_PLUGIN_ApplicationSignOnMode ApplicationSignOnMode = "BROWSER_PLUGIN"
	SECURE_PASSWORD_STORE_ApplicationSignOnMode ApplicationSignOnMode = "SECURE_PASSWORD_STORE"
	AUTO_LOGIN_ApplicationSignOnMode ApplicationSignOnMode = "AUTO_LOGIN"
	WS_FEDERATION_ApplicationSignOnMode ApplicationSignOnMode = "WS_FEDERATION"
	SAML_2_0_ApplicationSignOnMode ApplicationSignOnMode = "SAML_2_0"
	OPENID_CONNECT_ApplicationSignOnMode ApplicationSignOnMode = "OPENID_CONNECT"
	SAML_1_1_ApplicationSignOnMode ApplicationSignOnMode = "SAML_1_1"
)
