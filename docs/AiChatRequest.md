# AiChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**Messages** | [**[]AiChatRequestMessagesInner**](AiChatRequestMessagesInner.md) |  | 

## Methods

### NewAiChatRequest

`func NewAiChatRequest(messages []AiChatRequestMessagesInner, ) *AiChatRequest`

NewAiChatRequest instantiates a new AiChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiChatRequestWithDefaults

`func NewAiChatRequestWithDefaults() *AiChatRequest`

NewAiChatRequestWithDefaults instantiates a new AiChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiChatRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiChatRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiChatRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiChatRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMessages

`func (o *AiChatRequest) GetMessages() []AiChatRequestMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AiChatRequest) GetMessagesOk() (*[]AiChatRequestMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AiChatRequest) SetMessages(v []AiChatRequestMessagesInner)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


