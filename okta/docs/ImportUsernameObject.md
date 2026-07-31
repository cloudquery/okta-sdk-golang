# ImportUsernameObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserNameExpression** | Pointer to **string** | For &#x60;userNameFormat&#x3D;CUSTOM&#x60;, specifies the Okta Expression Language statement for a username format that imported users use to sign in to Okta | [optional] 
**UserNameFormat** | Pointer to **string** | Determines the username format when users sign in to Okta | [optional] [default to "EMAIL"]

## Methods

### NewImportUsernameObject

`func NewImportUsernameObject() *ImportUsernameObject`

NewImportUsernameObject instantiates a new ImportUsernameObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportUsernameObjectWithDefaults

`func NewImportUsernameObjectWithDefaults() *ImportUsernameObject`

NewImportUsernameObjectWithDefaults instantiates a new ImportUsernameObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserNameExpression

`func (o *ImportUsernameObject) GetUserNameExpression() string`

GetUserNameExpression returns the UserNameExpression field if non-nil, zero value otherwise.

### GetUserNameExpressionOk

`func (o *ImportUsernameObject) GetUserNameExpressionOk() (*string, bool)`

GetUserNameExpressionOk returns a tuple with the UserNameExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameExpression

`func (o *ImportUsernameObject) SetUserNameExpression(v string)`

SetUserNameExpression sets UserNameExpression field to given value.

### HasUserNameExpression

`func (o *ImportUsernameObject) HasUserNameExpression() bool`

HasUserNameExpression returns a boolean if a field has been set.

### GetUserNameFormat

`func (o *ImportUsernameObject) GetUserNameFormat() string`

GetUserNameFormat returns the UserNameFormat field if non-nil, zero value otherwise.

### GetUserNameFormatOk

`func (o *ImportUsernameObject) GetUserNameFormatOk() (*string, bool)`

GetUserNameFormatOk returns a tuple with the UserNameFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameFormat

`func (o *ImportUsernameObject) SetUserNameFormat(v string)`

SetUserNameFormat sets UserNameFormat field to given value.

### HasUserNameFormat

`func (o *ImportUsernameObject) HasUserNameFormat() bool`

HasUserNameFormat returns a boolean if a field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


