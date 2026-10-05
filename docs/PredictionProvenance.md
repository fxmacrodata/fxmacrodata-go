# PredictionProvenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValueOrigin** | **string** | Whether the value is a publisher historical vintage or an FXMacroData model output. | 
**CaptureMode** | **string** | Whether FXMacroData recorded the value when it was published or later imported it from the publisher&#39;s official archive. | 
**ArchivedInRealTimeByFxmacrodata** | **bool** | True only when FXMacroData captured the prediction at the time it was generated or published. | 
**GeneratedAtBasis** | **string** | The source fact used to construct &#x60;generated_at&#x60;. | 
**GeneratedAtPrecision** | **string** | Precision of the underlying source timestamp. &#x60;date&#x60; means an official source date was normalized to a documented source-local time. | 
**ReconstructedWithLaterData** | **bool** | Whether the prediction value was recomputed using later data or a newer model version. | 
**VintageSelection** | Pointer to **NullableString** | Which publisher or model vintages are retained in this feed. | [optional] 

## Methods

### NewPredictionProvenance

`func NewPredictionProvenance(valueOrigin string, captureMode string, archivedInRealTimeByFxmacrodata bool, generatedAtBasis string, generatedAtPrecision string, reconstructedWithLaterData bool, ) *PredictionProvenance`

NewPredictionProvenance instantiates a new PredictionProvenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionProvenanceWithDefaults

`func NewPredictionProvenanceWithDefaults() *PredictionProvenance`

NewPredictionProvenanceWithDefaults instantiates a new PredictionProvenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValueOrigin

`func (o *PredictionProvenance) GetValueOrigin() string`

GetValueOrigin returns the ValueOrigin field if non-nil, zero value otherwise.

### GetValueOriginOk

`func (o *PredictionProvenance) GetValueOriginOk() (*string, bool)`

GetValueOriginOk returns a tuple with the ValueOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOrigin

`func (o *PredictionProvenance) SetValueOrigin(v string)`

SetValueOrigin sets ValueOrigin field to given value.


### GetCaptureMode

`func (o *PredictionProvenance) GetCaptureMode() string`

GetCaptureMode returns the CaptureMode field if non-nil, zero value otherwise.

### GetCaptureModeOk

`func (o *PredictionProvenance) GetCaptureModeOk() (*string, bool)`

GetCaptureModeOk returns a tuple with the CaptureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMode

`func (o *PredictionProvenance) SetCaptureMode(v string)`

SetCaptureMode sets CaptureMode field to given value.


### GetArchivedInRealTimeByFxmacrodata

`func (o *PredictionProvenance) GetArchivedInRealTimeByFxmacrodata() bool`

GetArchivedInRealTimeByFxmacrodata returns the ArchivedInRealTimeByFxmacrodata field if non-nil, zero value otherwise.

### GetArchivedInRealTimeByFxmacrodataOk

`func (o *PredictionProvenance) GetArchivedInRealTimeByFxmacrodataOk() (*bool, bool)`

GetArchivedInRealTimeByFxmacrodataOk returns a tuple with the ArchivedInRealTimeByFxmacrodata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedInRealTimeByFxmacrodata

`func (o *PredictionProvenance) SetArchivedInRealTimeByFxmacrodata(v bool)`

SetArchivedInRealTimeByFxmacrodata sets ArchivedInRealTimeByFxmacrodata field to given value.


### GetGeneratedAtBasis

`func (o *PredictionProvenance) GetGeneratedAtBasis() string`

GetGeneratedAtBasis returns the GeneratedAtBasis field if non-nil, zero value otherwise.

### GetGeneratedAtBasisOk

`func (o *PredictionProvenance) GetGeneratedAtBasisOk() (*string, bool)`

GetGeneratedAtBasisOk returns a tuple with the GeneratedAtBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAtBasis

`func (o *PredictionProvenance) SetGeneratedAtBasis(v string)`

SetGeneratedAtBasis sets GeneratedAtBasis field to given value.


### GetGeneratedAtPrecision

`func (o *PredictionProvenance) GetGeneratedAtPrecision() string`

GetGeneratedAtPrecision returns the GeneratedAtPrecision field if non-nil, zero value otherwise.

### GetGeneratedAtPrecisionOk

`func (o *PredictionProvenance) GetGeneratedAtPrecisionOk() (*string, bool)`

GetGeneratedAtPrecisionOk returns a tuple with the GeneratedAtPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAtPrecision

`func (o *PredictionProvenance) SetGeneratedAtPrecision(v string)`

SetGeneratedAtPrecision sets GeneratedAtPrecision field to given value.


### GetReconstructedWithLaterData

`func (o *PredictionProvenance) GetReconstructedWithLaterData() bool`

GetReconstructedWithLaterData returns the ReconstructedWithLaterData field if non-nil, zero value otherwise.

### GetReconstructedWithLaterDataOk

`func (o *PredictionProvenance) GetReconstructedWithLaterDataOk() (*bool, bool)`

GetReconstructedWithLaterDataOk returns a tuple with the ReconstructedWithLaterData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconstructedWithLaterData

`func (o *PredictionProvenance) SetReconstructedWithLaterData(v bool)`

SetReconstructedWithLaterData sets ReconstructedWithLaterData field to given value.


### GetVintageSelection

`func (o *PredictionProvenance) GetVintageSelection() string`

GetVintageSelection returns the VintageSelection field if non-nil, zero value otherwise.

### GetVintageSelectionOk

`func (o *PredictionProvenance) GetVintageSelectionOk() (*string, bool)`

GetVintageSelectionOk returns a tuple with the VintageSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVintageSelection

`func (o *PredictionProvenance) SetVintageSelection(v string)`

SetVintageSelection sets VintageSelection field to given value.

### HasVintageSelection

`func (o *PredictionProvenance) HasVintageSelection() bool`

HasVintageSelection returns a boolean if a field has been set.

### SetVintageSelectionNil

`func (o *PredictionProvenance) SetVintageSelectionNil(b bool)`

 SetVintageSelectionNil sets the value for VintageSelection to be an explicit nil

### UnsetVintageSelection
`func (o *PredictionProvenance) UnsetVintageSelection()`

UnsetVintageSelection ensures that no value is present for VintageSelection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


