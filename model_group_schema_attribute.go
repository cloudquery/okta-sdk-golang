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

type GroupSchemaAttribute struct {
	Description string `json:"description,omitempty"`
	Enum []string `json:"enum,omitempty"`
	ExternalName string `json:"externalName,omitempty"`
	ExternalNamespace string `json:"externalNamespace,omitempty"`
	Items *UserSchemaAttributeItems `json:"items,omitempty"`
	Master *UserSchemaAttributeMaster `json:"master,omitempty"`
	MaxLength int32 `json:"maxLength,omitempty"`
	MinLength int32 `json:"minLength,omitempty"`
	Mutability string `json:"mutability,omitempty"`
	OneOf []UserSchemaAttributeEnum `json:"oneOf,omitempty"`
	Permissions []UserSchemaAttributePermission `json:"permissions,omitempty"`
	Required bool `json:"required,omitempty"`
	Scope *UserSchemaAttributeScope `json:"scope,omitempty"`
	Title string `json:"title,omitempty"`
	Type_ *UserSchemaAttributeType `json:"type,omitempty"`
	Union *UserSchemaAttributeUnion `json:"union,omitempty"`
	Unique string `json:"unique,omitempty"`
}
