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

type AuthorizationServer struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	Audiences []string `json:"audiences,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Credentials *AuthorizationServerCredentials `json:"credentials,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id,omitempty"`
	Issuer string `json:"issuer,omitempty"`
	IssuerMode *IssuerMode `json:"issuerMode,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Name string `json:"name,omitempty"`
	Status *LifecycleStatus `json:"status,omitempty"`
}
