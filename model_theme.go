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

type Theme struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	BackgroundImage string `json:"backgroundImage,omitempty"`
	EmailTemplateTouchPointVariant *EmailTemplateTouchPointVariant `json:"emailTemplateTouchPointVariant,omitempty"`
	EndUserDashboardTouchPointVariant *EndUserDashboardTouchPointVariant `json:"endUserDashboardTouchPointVariant,omitempty"`
	ErrorPageTouchPointVariant *ErrorPageTouchPointVariant `json:"errorPageTouchPointVariant,omitempty"`
	PrimaryColorContrastHex string `json:"primaryColorContrastHex,omitempty"`
	PrimaryColorHex string `json:"primaryColorHex,omitempty"`
	SecondaryColorContrastHex string `json:"secondaryColorContrastHex,omitempty"`
	SecondaryColorHex string `json:"secondaryColorHex,omitempty"`
	SignInPageTouchPointVariant *SignInPageTouchPointVariant `json:"signInPageTouchPointVariant,omitempty"`
}
