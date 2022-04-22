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

type FactorType string

// List of FactorType
const (
	CALL_FactorType FactorType = "call"
	EMAIL_FactorType FactorType = "email"
	HOTP_FactorType FactorType = "hotp"
	PUSH_FactorType FactorType = "push"
	QUESTION_FactorType FactorType = "question"
	SMS_FactorType FactorType = "sms"
	TOKENHARDWARE_FactorType FactorType = "token:hardware"
	TOKENHOTP_FactorType FactorType = "token:hotp"
	TOKENSOFTWARETOTP_FactorType FactorType = "token:software:totp"
	TOKEN_FactorType FactorType = "token"
	U2F_FactorType FactorType = "u2f"
	WEB_FactorType FactorType = "web"
	WEBAUTHN_FactorType FactorType = "webauthn"
)
