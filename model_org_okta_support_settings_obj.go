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
import (
	"time"
)

type OrgOktaSupportSettingsObj struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	Expiration time.Time `json:"expiration,omitempty"`
	Support *OrgOktaSupportSetting `json:"support,omitempty"`
}
