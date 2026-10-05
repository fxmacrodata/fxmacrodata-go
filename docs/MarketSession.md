# MarketSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Currencies** | Pointer to **[]string** |  | [optional] 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**OpenUtc** | Pointer to **NullableString** |  | [optional] 
**CloseUtc** | Pointer to **NullableString** |  | [optional] 
**OpenUnix** | Pointer to **NullableInt64** |  | [optional] 
**CloseUnix** | Pointer to **NullableInt64** |  | [optional] 
**IsOpen** | Pointer to **NullableBool** |  | [optional] 
**SecondsToOpen** | Pointer to **NullableInt64** |  | [optional] 
**SecondsToClose** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewMarketSession

`func NewMarketSession() *MarketSession`

NewMarketSession instantiates a new MarketSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketSessionWithDefaults

`func NewMarketSessionWithDefaults() *MarketSession`

NewMarketSessionWithDefaults instantiates a new MarketSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MarketSession) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketSession) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketSession) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketSession) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MarketSession) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MarketSession) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDisplayName

`func (o *MarketSession) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MarketSession) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MarketSession) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MarketSession) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MarketSession) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MarketSession) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *MarketSession) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MarketSession) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MarketSession) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MarketSession) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MarketSession) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MarketSession) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCurrencies

`func (o *MarketSession) GetCurrencies() []string`

GetCurrencies returns the Currencies field if non-nil, zero value otherwise.

### GetCurrenciesOk

`func (o *MarketSession) GetCurrenciesOk() (*[]string, bool)`

GetCurrenciesOk returns a tuple with the Currencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencies

`func (o *MarketSession) SetCurrencies(v []string)`

SetCurrencies sets Currencies field to given value.

### HasCurrencies

`func (o *MarketSession) HasCurrencies() bool`

HasCurrencies returns a boolean if a field has been set.

### SetCurrenciesNil

`func (o *MarketSession) SetCurrenciesNil(b bool)`

 SetCurrenciesNil sets the value for Currencies to be an explicit nil

### UnsetCurrencies
`func (o *MarketSession) UnsetCurrencies()`

UnsetCurrencies ensures that no value is present for Currencies, not even an explicit nil
### GetTimezone

`func (o *MarketSession) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MarketSession) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MarketSession) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *MarketSession) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *MarketSession) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *MarketSession) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetOpenUtc

`func (o *MarketSession) GetOpenUtc() string`

GetOpenUtc returns the OpenUtc field if non-nil, zero value otherwise.

### GetOpenUtcOk

`func (o *MarketSession) GetOpenUtcOk() (*string, bool)`

GetOpenUtcOk returns a tuple with the OpenUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenUtc

`func (o *MarketSession) SetOpenUtc(v string)`

SetOpenUtc sets OpenUtc field to given value.

### HasOpenUtc

`func (o *MarketSession) HasOpenUtc() bool`

HasOpenUtc returns a boolean if a field has been set.

### SetOpenUtcNil

`func (o *MarketSession) SetOpenUtcNil(b bool)`

 SetOpenUtcNil sets the value for OpenUtc to be an explicit nil

### UnsetOpenUtc
`func (o *MarketSession) UnsetOpenUtc()`

UnsetOpenUtc ensures that no value is present for OpenUtc, not even an explicit nil
### GetCloseUtc

`func (o *MarketSession) GetCloseUtc() string`

GetCloseUtc returns the CloseUtc field if non-nil, zero value otherwise.

### GetCloseUtcOk

`func (o *MarketSession) GetCloseUtcOk() (*string, bool)`

GetCloseUtcOk returns a tuple with the CloseUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseUtc

`func (o *MarketSession) SetCloseUtc(v string)`

SetCloseUtc sets CloseUtc field to given value.

### HasCloseUtc

`func (o *MarketSession) HasCloseUtc() bool`

HasCloseUtc returns a boolean if a field has been set.

### SetCloseUtcNil

`func (o *MarketSession) SetCloseUtcNil(b bool)`

 SetCloseUtcNil sets the value for CloseUtc to be an explicit nil

### UnsetCloseUtc
`func (o *MarketSession) UnsetCloseUtc()`

UnsetCloseUtc ensures that no value is present for CloseUtc, not even an explicit nil
### GetOpenUnix

`func (o *MarketSession) GetOpenUnix() int64`

GetOpenUnix returns the OpenUnix field if non-nil, zero value otherwise.

### GetOpenUnixOk

`func (o *MarketSession) GetOpenUnixOk() (*int64, bool)`

GetOpenUnixOk returns a tuple with the OpenUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenUnix

`func (o *MarketSession) SetOpenUnix(v int64)`

SetOpenUnix sets OpenUnix field to given value.

### HasOpenUnix

`func (o *MarketSession) HasOpenUnix() bool`

HasOpenUnix returns a boolean if a field has been set.

### SetOpenUnixNil

`func (o *MarketSession) SetOpenUnixNil(b bool)`

 SetOpenUnixNil sets the value for OpenUnix to be an explicit nil

### UnsetOpenUnix
`func (o *MarketSession) UnsetOpenUnix()`

UnsetOpenUnix ensures that no value is present for OpenUnix, not even an explicit nil
### GetCloseUnix

`func (o *MarketSession) GetCloseUnix() int64`

GetCloseUnix returns the CloseUnix field if non-nil, zero value otherwise.

### GetCloseUnixOk

`func (o *MarketSession) GetCloseUnixOk() (*int64, bool)`

GetCloseUnixOk returns a tuple with the CloseUnix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseUnix

`func (o *MarketSession) SetCloseUnix(v int64)`

SetCloseUnix sets CloseUnix field to given value.

### HasCloseUnix

`func (o *MarketSession) HasCloseUnix() bool`

HasCloseUnix returns a boolean if a field has been set.

### SetCloseUnixNil

`func (o *MarketSession) SetCloseUnixNil(b bool)`

 SetCloseUnixNil sets the value for CloseUnix to be an explicit nil

### UnsetCloseUnix
`func (o *MarketSession) UnsetCloseUnix()`

UnsetCloseUnix ensures that no value is present for CloseUnix, not even an explicit nil
### GetIsOpen

`func (o *MarketSession) GetIsOpen() bool`

GetIsOpen returns the IsOpen field if non-nil, zero value otherwise.

### GetIsOpenOk

`func (o *MarketSession) GetIsOpenOk() (*bool, bool)`

GetIsOpenOk returns a tuple with the IsOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOpen

`func (o *MarketSession) SetIsOpen(v bool)`

SetIsOpen sets IsOpen field to given value.

### HasIsOpen

`func (o *MarketSession) HasIsOpen() bool`

HasIsOpen returns a boolean if a field has been set.

### SetIsOpenNil

`func (o *MarketSession) SetIsOpenNil(b bool)`

 SetIsOpenNil sets the value for IsOpen to be an explicit nil

### UnsetIsOpen
`func (o *MarketSession) UnsetIsOpen()`

UnsetIsOpen ensures that no value is present for IsOpen, not even an explicit nil
### GetSecondsToOpen

`func (o *MarketSession) GetSecondsToOpen() int64`

GetSecondsToOpen returns the SecondsToOpen field if non-nil, zero value otherwise.

### GetSecondsToOpenOk

`func (o *MarketSession) GetSecondsToOpenOk() (*int64, bool)`

GetSecondsToOpenOk returns a tuple with the SecondsToOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToOpen

`func (o *MarketSession) SetSecondsToOpen(v int64)`

SetSecondsToOpen sets SecondsToOpen field to given value.

### HasSecondsToOpen

`func (o *MarketSession) HasSecondsToOpen() bool`

HasSecondsToOpen returns a boolean if a field has been set.

### SetSecondsToOpenNil

`func (o *MarketSession) SetSecondsToOpenNil(b bool)`

 SetSecondsToOpenNil sets the value for SecondsToOpen to be an explicit nil

### UnsetSecondsToOpen
`func (o *MarketSession) UnsetSecondsToOpen()`

UnsetSecondsToOpen ensures that no value is present for SecondsToOpen, not even an explicit nil
### GetSecondsToClose

`func (o *MarketSession) GetSecondsToClose() int64`

GetSecondsToClose returns the SecondsToClose field if non-nil, zero value otherwise.

### GetSecondsToCloseOk

`func (o *MarketSession) GetSecondsToCloseOk() (*int64, bool)`

GetSecondsToCloseOk returns a tuple with the SecondsToClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToClose

`func (o *MarketSession) SetSecondsToClose(v int64)`

SetSecondsToClose sets SecondsToClose field to given value.

### HasSecondsToClose

`func (o *MarketSession) HasSecondsToClose() bool`

HasSecondsToClose returns a boolean if a field has been set.

### SetSecondsToCloseNil

`func (o *MarketSession) SetSecondsToCloseNil(b bool)`

 SetSecondsToCloseNil sets the value for SecondsToClose to be an explicit nil

### UnsetSecondsToClose
`func (o *MarketSession) UnsetSecondsToClose()`

UnsetSecondsToClose ensures that no value is present for SecondsToClose, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


