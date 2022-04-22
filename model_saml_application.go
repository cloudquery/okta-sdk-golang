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

type SamlApplication struct {
	Embedded map[string]interface{} `json:"_embedded,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	Accessibility *ApplicationAccessibility `json:"accessibility,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Credentials *ApplicationCredentials `json:"credentials,omitempty"`
	Features []string `json:"features,omitempty"`
	Id string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Licensing *ApplicationLicensing `json:"licensing,omitempty"`
	Name string `json:"name,omitempty"`
	Profile map[string]interface{} `json:"profile,omitempty"`
	Settings *SamlApplicationSettings `json:"settings,omitempty"`
	SignOnMode *ApplicationSignOnMode `json:"signOnMode,omitempty"`
	Status *ApplicationLifecycleStatus `json:"status,omitempty"`
	Visibility *ApplicationVisibility `json:"visibility,omitempty"`
}
