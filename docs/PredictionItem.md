# PredictionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PredictedValue** | **float64** |  | 
**EventCompatible** | Pointer to **bool** |  | [optional] [default to false]
**EventCompatibilityReason** | Pointer to **NullableString** |  | [optional] 
**SourceUrl** | Pointer to **NullableString** |  | [optional] 
**TargetPeriod** | Pointer to **NullableString** |  | [optional] 
**PeriodBasis** | Pointer to **NullableString** |  | [optional] 
**Unit** | Pointer to **NullableString** |  | [optional] 
**Statistic** | Pointer to **NullableString** |  | [optional] 
**PublicationDate** | Pointer to **NullableString** |  | [optional] 
**PublicationDatetime** | Pointer to **NullableString** |  | [optional] 
**PublicationPrecision** | Pointer to **NullableString** |  | [optional] 
**ForecastObservationId** | Pointer to **NullableString** |  | [optional] 
**SeasonalAdjustment** | Pointer to **NullableString** |  | [optional] 
**Geography** | Pointer to **NullableString** |  | [optional] 
**AnnouncementMappingVerified** | Pointer to **bool** |  | [optional] [default to false]
**PredictionClass** | Pointer to **NullableString** | How the forecast was produced, and the field to filter on. One of &#x60;&#x60;compiled_consensus&#x60;&#x60; (a publisher&#39;s compilation of independent forecasters), &#x60;&#x60;forecaster_survey&#x60;&#x60; (an official survey of forecasters or market participants), &#x60;&#x60;market_implied&#x60;&#x60; (derived from traded prices), &#x60;&#x60;model_nowcast&#x60;&#x60; (a published statistical model), &#x60;&#x60;central_bank_projection&#x60;&#x60; (the central bank&#39;s own projection), &#x60;&#x60;institutional_projection&#x60;&#x60; (IMF, OECD), or &#x60;&#x60;fxmacrodata&#x60;&#x60; (an FXMacroData estimate, never a market consensus). | [optional] 
**PredictionClassLabel** | Pointer to **NullableString** | Human-readable name for &#x60;&#x60;prediction_class&#x60;&#x60;, for display. Filter on &#x60;&#x60;prediction_class&#x60;&#x60;, not on this. | [optional] 
**PredictionType** | Pointer to **NullableString** | Original per-source category, retained unchanged for backward compatibility: &#x60;&#x60;market_consensus&#x60;&#x60;, &#x60;&#x60;market_prediction&#x60;&#x60;, &#x60;&#x60;model_nowcast&#x60;&#x60;, &#x60;&#x60;survey&#x60;&#x60;, &#x60;&#x60;central_bank_forecast&#x60;&#x60;, &#x60;&#x60;central_bank_projection&#x60;&#x60;, &#x60;&#x60;imf_weo&#x60;&#x60;, &#x60;&#x60;oecd_eo&#x60;&#x60; or &#x60;&#x60;fxmacrodata&#x60;&#x60;. These slugs split identical instruments (the ECB and Philadelphia Fed Surveys of Professional Forecasters landed in different values) and two of them name a publisher rather than a category. Prefer &#x60;&#x60;prediction_class&#x60;&#x60;. | [optional] 
**PredictionSource** | Pointer to **NullableString** | Stable source identifier slug (e.g. &#x60;&#x60;philly_fed_spf&#x60;&#x60;, &#x60;&#x60;cleveland_fed_nowcast&#x60;&#x60;, &#x60;&#x60;fomc_sep&#x60;&#x60;, &#x60;&#x60;atlanta_fed_gdpnow&#x60;&#x60;, &#x60;&#x60;ny_fed_sme&#x60;&#x60;, &#x60;&#x60;ecb_spf&#x60;&#x60;, &#x60;&#x60;imf_weo&#x60;&#x60;, &#x60;&#x60;rbnz_mps&#x60;&#x60;, &#x60;&#x60;fxmacrodata_blended&#x60;&#x60;). Use this value when filtering via the &#x60;&#x60;prediction_source&#x60;&#x60; query parameter. | [optional] 
**PredictionSourceLabel** | Pointer to **NullableString** | Human-readable name for &#x60;&#x60;prediction_source&#x60;&#x60; (e.g. \&quot;RBA Statement on Monetary Policy (Table J1)\&quot;, \&quot;Philadelphia Fed Survey of Professional Forecasters\&quot;, \&quot;IMF World Economic Outlook\&quot;). Suitable for display in UI surfaces; use &#x60;&#x60;prediction_source&#x60;&#x60; for filtering and joins. | [optional] 
**GeneratedAt** | Pointer to **NullableInt64** | Unix epoch seconds (UTC) for the prediction&#39;s known-at marker. Inspect &#x60;provenance.generated_at_basis&#x60; and &#x60;provenance.generated_at_precision&#x60;: publisher archives may provide a date that FXMacroData normalizes to a documented source-local publication time rather than an original intraday timestamp. | [optional] 
**Provenance** | Pointer to [**NullablePredictionProvenance**](PredictionProvenance.md) | Documented origin, capture, timestamp, reconstruction, and vintage semantics for this source. Omitted for sources whose provenance is not yet documented. | [optional] 
**IsPreRelease** | Pointer to **NullableBool** | Whether &#x60;generated_at&#x60; is strictly earlier than the linked &#x60;announcement_datetime&#x60;. Null when the source did not provide a generation timestamp. | [optional] 
**Confidence** | Pointer to **NullableFloat64** | Confidence score in &#x60;&#x60;[0.0, 1.0]&#x60;&#x60; indicating how strong this prediction is. &#x60;&#x60;1.0&#x60;&#x60; is maximally confident, &#x60;&#x60;0.0&#x60;&#x60; is no confidence. When the source did not record an explicit value, a per-source default is returned (central-bank publications and official surveys score higher than the FXMacroData blended forecast). Use this to weight or filter predictions in client applications. | [optional] 
**PredictionReason** | Pointer to **NullableString** | Short human-readable explanation of how this prediction was derived. Currently populated for FXMacroData-generated predictions (&#x60;&#x60;prediction_source&#x3D;fxmacrodata_blended&#x60;&#x60;); it summarises the persistence anchor, momentum/trend extrapolation, cross-indicator pressure, and any external forecast that was blended in. Other prediction sources may omit this field. | [optional] 

## Methods

### NewPredictionItem

`func NewPredictionItem(predictedValue float64, ) *PredictionItem`

NewPredictionItem instantiates a new PredictionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPredictionItemWithDefaults

`func NewPredictionItemWithDefaults() *PredictionItem`

NewPredictionItemWithDefaults instantiates a new PredictionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredictedValue

`func (o *PredictionItem) GetPredictedValue() float64`

GetPredictedValue returns the PredictedValue field if non-nil, zero value otherwise.

### GetPredictedValueOk

`func (o *PredictionItem) GetPredictedValueOk() (*float64, bool)`

GetPredictedValueOk returns a tuple with the PredictedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedValue

`func (o *PredictionItem) SetPredictedValue(v float64)`

SetPredictedValue sets PredictedValue field to given value.


### GetEventCompatible

`func (o *PredictionItem) GetEventCompatible() bool`

GetEventCompatible returns the EventCompatible field if non-nil, zero value otherwise.

### GetEventCompatibleOk

`func (o *PredictionItem) GetEventCompatibleOk() (*bool, bool)`

GetEventCompatibleOk returns a tuple with the EventCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCompatible

`func (o *PredictionItem) SetEventCompatible(v bool)`

SetEventCompatible sets EventCompatible field to given value.

### HasEventCompatible

`func (o *PredictionItem) HasEventCompatible() bool`

HasEventCompatible returns a boolean if a field has been set.

### GetEventCompatibilityReason

`func (o *PredictionItem) GetEventCompatibilityReason() string`

GetEventCompatibilityReason returns the EventCompatibilityReason field if non-nil, zero value otherwise.

### GetEventCompatibilityReasonOk

`func (o *PredictionItem) GetEventCompatibilityReasonOk() (*string, bool)`

GetEventCompatibilityReasonOk returns a tuple with the EventCompatibilityReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCompatibilityReason

`func (o *PredictionItem) SetEventCompatibilityReason(v string)`

SetEventCompatibilityReason sets EventCompatibilityReason field to given value.

### HasEventCompatibilityReason

`func (o *PredictionItem) HasEventCompatibilityReason() bool`

HasEventCompatibilityReason returns a boolean if a field has been set.

### SetEventCompatibilityReasonNil

`func (o *PredictionItem) SetEventCompatibilityReasonNil(b bool)`

 SetEventCompatibilityReasonNil sets the value for EventCompatibilityReason to be an explicit nil

### UnsetEventCompatibilityReason
`func (o *PredictionItem) UnsetEventCompatibilityReason()`

UnsetEventCompatibilityReason ensures that no value is present for EventCompatibilityReason, not even an explicit nil
### GetSourceUrl

`func (o *PredictionItem) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *PredictionItem) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *PredictionItem) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *PredictionItem) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *PredictionItem) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *PredictionItem) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetTargetPeriod

`func (o *PredictionItem) GetTargetPeriod() string`

GetTargetPeriod returns the TargetPeriod field if non-nil, zero value otherwise.

### GetTargetPeriodOk

`func (o *PredictionItem) GetTargetPeriodOk() (*string, bool)`

GetTargetPeriodOk returns a tuple with the TargetPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPeriod

`func (o *PredictionItem) SetTargetPeriod(v string)`

SetTargetPeriod sets TargetPeriod field to given value.

### HasTargetPeriod

`func (o *PredictionItem) HasTargetPeriod() bool`

HasTargetPeriod returns a boolean if a field has been set.

### SetTargetPeriodNil

`func (o *PredictionItem) SetTargetPeriodNil(b bool)`

 SetTargetPeriodNil sets the value for TargetPeriod to be an explicit nil

### UnsetTargetPeriod
`func (o *PredictionItem) UnsetTargetPeriod()`

UnsetTargetPeriod ensures that no value is present for TargetPeriod, not even an explicit nil
### GetPeriodBasis

`func (o *PredictionItem) GetPeriodBasis() string`

GetPeriodBasis returns the PeriodBasis field if non-nil, zero value otherwise.

### GetPeriodBasisOk

`func (o *PredictionItem) GetPeriodBasisOk() (*string, bool)`

GetPeriodBasisOk returns a tuple with the PeriodBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodBasis

`func (o *PredictionItem) SetPeriodBasis(v string)`

SetPeriodBasis sets PeriodBasis field to given value.

### HasPeriodBasis

`func (o *PredictionItem) HasPeriodBasis() bool`

HasPeriodBasis returns a boolean if a field has been set.

### SetPeriodBasisNil

`func (o *PredictionItem) SetPeriodBasisNil(b bool)`

 SetPeriodBasisNil sets the value for PeriodBasis to be an explicit nil

### UnsetPeriodBasis
`func (o *PredictionItem) UnsetPeriodBasis()`

UnsetPeriodBasis ensures that no value is present for PeriodBasis, not even an explicit nil
### GetUnit

`func (o *PredictionItem) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PredictionItem) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PredictionItem) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PredictionItem) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *PredictionItem) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *PredictionItem) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetStatistic

`func (o *PredictionItem) GetStatistic() string`

GetStatistic returns the Statistic field if non-nil, zero value otherwise.

### GetStatisticOk

`func (o *PredictionItem) GetStatisticOk() (*string, bool)`

GetStatisticOk returns a tuple with the Statistic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistic

`func (o *PredictionItem) SetStatistic(v string)`

SetStatistic sets Statistic field to given value.

### HasStatistic

`func (o *PredictionItem) HasStatistic() bool`

HasStatistic returns a boolean if a field has been set.

### SetStatisticNil

`func (o *PredictionItem) SetStatisticNil(b bool)`

 SetStatisticNil sets the value for Statistic to be an explicit nil

### UnsetStatistic
`func (o *PredictionItem) UnsetStatistic()`

UnsetStatistic ensures that no value is present for Statistic, not even an explicit nil
### GetPublicationDate

`func (o *PredictionItem) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *PredictionItem) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *PredictionItem) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *PredictionItem) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### SetPublicationDateNil

`func (o *PredictionItem) SetPublicationDateNil(b bool)`

 SetPublicationDateNil sets the value for PublicationDate to be an explicit nil

### UnsetPublicationDate
`func (o *PredictionItem) UnsetPublicationDate()`

UnsetPublicationDate ensures that no value is present for PublicationDate, not even an explicit nil
### GetPublicationDatetime

`func (o *PredictionItem) GetPublicationDatetime() string`

GetPublicationDatetime returns the PublicationDatetime field if non-nil, zero value otherwise.

### GetPublicationDatetimeOk

`func (o *PredictionItem) GetPublicationDatetimeOk() (*string, bool)`

GetPublicationDatetimeOk returns a tuple with the PublicationDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDatetime

`func (o *PredictionItem) SetPublicationDatetime(v string)`

SetPublicationDatetime sets PublicationDatetime field to given value.

### HasPublicationDatetime

`func (o *PredictionItem) HasPublicationDatetime() bool`

HasPublicationDatetime returns a boolean if a field has been set.

### SetPublicationDatetimeNil

`func (o *PredictionItem) SetPublicationDatetimeNil(b bool)`

 SetPublicationDatetimeNil sets the value for PublicationDatetime to be an explicit nil

### UnsetPublicationDatetime
`func (o *PredictionItem) UnsetPublicationDatetime()`

UnsetPublicationDatetime ensures that no value is present for PublicationDatetime, not even an explicit nil
### GetPublicationPrecision

`func (o *PredictionItem) GetPublicationPrecision() string`

GetPublicationPrecision returns the PublicationPrecision field if non-nil, zero value otherwise.

### GetPublicationPrecisionOk

`func (o *PredictionItem) GetPublicationPrecisionOk() (*string, bool)`

GetPublicationPrecisionOk returns a tuple with the PublicationPrecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationPrecision

`func (o *PredictionItem) SetPublicationPrecision(v string)`

SetPublicationPrecision sets PublicationPrecision field to given value.

### HasPublicationPrecision

`func (o *PredictionItem) HasPublicationPrecision() bool`

HasPublicationPrecision returns a boolean if a field has been set.

### SetPublicationPrecisionNil

`func (o *PredictionItem) SetPublicationPrecisionNil(b bool)`

 SetPublicationPrecisionNil sets the value for PublicationPrecision to be an explicit nil

### UnsetPublicationPrecision
`func (o *PredictionItem) UnsetPublicationPrecision()`

UnsetPublicationPrecision ensures that no value is present for PublicationPrecision, not even an explicit nil
### GetForecastObservationId

`func (o *PredictionItem) GetForecastObservationId() string`

GetForecastObservationId returns the ForecastObservationId field if non-nil, zero value otherwise.

### GetForecastObservationIdOk

`func (o *PredictionItem) GetForecastObservationIdOk() (*string, bool)`

GetForecastObservationIdOk returns a tuple with the ForecastObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastObservationId

`func (o *PredictionItem) SetForecastObservationId(v string)`

SetForecastObservationId sets ForecastObservationId field to given value.

### HasForecastObservationId

`func (o *PredictionItem) HasForecastObservationId() bool`

HasForecastObservationId returns a boolean if a field has been set.

### SetForecastObservationIdNil

`func (o *PredictionItem) SetForecastObservationIdNil(b bool)`

 SetForecastObservationIdNil sets the value for ForecastObservationId to be an explicit nil

### UnsetForecastObservationId
`func (o *PredictionItem) UnsetForecastObservationId()`

UnsetForecastObservationId ensures that no value is present for ForecastObservationId, not even an explicit nil
### GetSeasonalAdjustment

`func (o *PredictionItem) GetSeasonalAdjustment() string`

GetSeasonalAdjustment returns the SeasonalAdjustment field if non-nil, zero value otherwise.

### GetSeasonalAdjustmentOk

`func (o *PredictionItem) GetSeasonalAdjustmentOk() (*string, bool)`

GetSeasonalAdjustmentOk returns a tuple with the SeasonalAdjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonalAdjustment

`func (o *PredictionItem) SetSeasonalAdjustment(v string)`

SetSeasonalAdjustment sets SeasonalAdjustment field to given value.

### HasSeasonalAdjustment

`func (o *PredictionItem) HasSeasonalAdjustment() bool`

HasSeasonalAdjustment returns a boolean if a field has been set.

### SetSeasonalAdjustmentNil

`func (o *PredictionItem) SetSeasonalAdjustmentNil(b bool)`

 SetSeasonalAdjustmentNil sets the value for SeasonalAdjustment to be an explicit nil

### UnsetSeasonalAdjustment
`func (o *PredictionItem) UnsetSeasonalAdjustment()`

UnsetSeasonalAdjustment ensures that no value is present for SeasonalAdjustment, not even an explicit nil
### GetGeography

`func (o *PredictionItem) GetGeography() string`

GetGeography returns the Geography field if non-nil, zero value otherwise.

### GetGeographyOk

`func (o *PredictionItem) GetGeographyOk() (*string, bool)`

GetGeographyOk returns a tuple with the Geography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeography

`func (o *PredictionItem) SetGeography(v string)`

SetGeography sets Geography field to given value.

### HasGeography

`func (o *PredictionItem) HasGeography() bool`

HasGeography returns a boolean if a field has been set.

### SetGeographyNil

`func (o *PredictionItem) SetGeographyNil(b bool)`

 SetGeographyNil sets the value for Geography to be an explicit nil

### UnsetGeography
`func (o *PredictionItem) UnsetGeography()`

UnsetGeography ensures that no value is present for Geography, not even an explicit nil
### GetAnnouncementMappingVerified

`func (o *PredictionItem) GetAnnouncementMappingVerified() bool`

GetAnnouncementMappingVerified returns the AnnouncementMappingVerified field if non-nil, zero value otherwise.

### GetAnnouncementMappingVerifiedOk

`func (o *PredictionItem) GetAnnouncementMappingVerifiedOk() (*bool, bool)`

GetAnnouncementMappingVerifiedOk returns a tuple with the AnnouncementMappingVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementMappingVerified

`func (o *PredictionItem) SetAnnouncementMappingVerified(v bool)`

SetAnnouncementMappingVerified sets AnnouncementMappingVerified field to given value.

### HasAnnouncementMappingVerified

`func (o *PredictionItem) HasAnnouncementMappingVerified() bool`

HasAnnouncementMappingVerified returns a boolean if a field has been set.

### GetPredictionClass

`func (o *PredictionItem) GetPredictionClass() string`

GetPredictionClass returns the PredictionClass field if non-nil, zero value otherwise.

### GetPredictionClassOk

`func (o *PredictionItem) GetPredictionClassOk() (*string, bool)`

GetPredictionClassOk returns a tuple with the PredictionClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionClass

`func (o *PredictionItem) SetPredictionClass(v string)`

SetPredictionClass sets PredictionClass field to given value.

### HasPredictionClass

`func (o *PredictionItem) HasPredictionClass() bool`

HasPredictionClass returns a boolean if a field has been set.

### SetPredictionClassNil

`func (o *PredictionItem) SetPredictionClassNil(b bool)`

 SetPredictionClassNil sets the value for PredictionClass to be an explicit nil

### UnsetPredictionClass
`func (o *PredictionItem) UnsetPredictionClass()`

UnsetPredictionClass ensures that no value is present for PredictionClass, not even an explicit nil
### GetPredictionClassLabel

`func (o *PredictionItem) GetPredictionClassLabel() string`

GetPredictionClassLabel returns the PredictionClassLabel field if non-nil, zero value otherwise.

### GetPredictionClassLabelOk

`func (o *PredictionItem) GetPredictionClassLabelOk() (*string, bool)`

GetPredictionClassLabelOk returns a tuple with the PredictionClassLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionClassLabel

`func (o *PredictionItem) SetPredictionClassLabel(v string)`

SetPredictionClassLabel sets PredictionClassLabel field to given value.

### HasPredictionClassLabel

`func (o *PredictionItem) HasPredictionClassLabel() bool`

HasPredictionClassLabel returns a boolean if a field has been set.

### SetPredictionClassLabelNil

`func (o *PredictionItem) SetPredictionClassLabelNil(b bool)`

 SetPredictionClassLabelNil sets the value for PredictionClassLabel to be an explicit nil

### UnsetPredictionClassLabel
`func (o *PredictionItem) UnsetPredictionClassLabel()`

UnsetPredictionClassLabel ensures that no value is present for PredictionClassLabel, not even an explicit nil
### GetPredictionType

`func (o *PredictionItem) GetPredictionType() string`

GetPredictionType returns the PredictionType field if non-nil, zero value otherwise.

### GetPredictionTypeOk

`func (o *PredictionItem) GetPredictionTypeOk() (*string, bool)`

GetPredictionTypeOk returns a tuple with the PredictionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionType

`func (o *PredictionItem) SetPredictionType(v string)`

SetPredictionType sets PredictionType field to given value.

### HasPredictionType

`func (o *PredictionItem) HasPredictionType() bool`

HasPredictionType returns a boolean if a field has been set.

### SetPredictionTypeNil

`func (o *PredictionItem) SetPredictionTypeNil(b bool)`

 SetPredictionTypeNil sets the value for PredictionType to be an explicit nil

### UnsetPredictionType
`func (o *PredictionItem) UnsetPredictionType()`

UnsetPredictionType ensures that no value is present for PredictionType, not even an explicit nil
### GetPredictionSource

`func (o *PredictionItem) GetPredictionSource() string`

GetPredictionSource returns the PredictionSource field if non-nil, zero value otherwise.

### GetPredictionSourceOk

`func (o *PredictionItem) GetPredictionSourceOk() (*string, bool)`

GetPredictionSourceOk returns a tuple with the PredictionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionSource

`func (o *PredictionItem) SetPredictionSource(v string)`

SetPredictionSource sets PredictionSource field to given value.

### HasPredictionSource

`func (o *PredictionItem) HasPredictionSource() bool`

HasPredictionSource returns a boolean if a field has been set.

### SetPredictionSourceNil

`func (o *PredictionItem) SetPredictionSourceNil(b bool)`

 SetPredictionSourceNil sets the value for PredictionSource to be an explicit nil

### UnsetPredictionSource
`func (o *PredictionItem) UnsetPredictionSource()`

UnsetPredictionSource ensures that no value is present for PredictionSource, not even an explicit nil
### GetPredictionSourceLabel

`func (o *PredictionItem) GetPredictionSourceLabel() string`

GetPredictionSourceLabel returns the PredictionSourceLabel field if non-nil, zero value otherwise.

### GetPredictionSourceLabelOk

`func (o *PredictionItem) GetPredictionSourceLabelOk() (*string, bool)`

GetPredictionSourceLabelOk returns a tuple with the PredictionSourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionSourceLabel

`func (o *PredictionItem) SetPredictionSourceLabel(v string)`

SetPredictionSourceLabel sets PredictionSourceLabel field to given value.

### HasPredictionSourceLabel

`func (o *PredictionItem) HasPredictionSourceLabel() bool`

HasPredictionSourceLabel returns a boolean if a field has been set.

### SetPredictionSourceLabelNil

`func (o *PredictionItem) SetPredictionSourceLabelNil(b bool)`

 SetPredictionSourceLabelNil sets the value for PredictionSourceLabel to be an explicit nil

### UnsetPredictionSourceLabel
`func (o *PredictionItem) UnsetPredictionSourceLabel()`

UnsetPredictionSourceLabel ensures that no value is present for PredictionSourceLabel, not even an explicit nil
### GetGeneratedAt

`func (o *PredictionItem) GetGeneratedAt() int64`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *PredictionItem) GetGeneratedAtOk() (*int64, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *PredictionItem) SetGeneratedAt(v int64)`

SetGeneratedAt sets GeneratedAt field to given value.

### HasGeneratedAt

`func (o *PredictionItem) HasGeneratedAt() bool`

HasGeneratedAt returns a boolean if a field has been set.

### SetGeneratedAtNil

`func (o *PredictionItem) SetGeneratedAtNil(b bool)`

 SetGeneratedAtNil sets the value for GeneratedAt to be an explicit nil

### UnsetGeneratedAt
`func (o *PredictionItem) UnsetGeneratedAt()`

UnsetGeneratedAt ensures that no value is present for GeneratedAt, not even an explicit nil
### GetProvenance

`func (o *PredictionItem) GetProvenance() PredictionProvenance`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *PredictionItem) GetProvenanceOk() (*PredictionProvenance, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *PredictionItem) SetProvenance(v PredictionProvenance)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *PredictionItem) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### SetProvenanceNil

`func (o *PredictionItem) SetProvenanceNil(b bool)`

 SetProvenanceNil sets the value for Provenance to be an explicit nil

### UnsetProvenance
`func (o *PredictionItem) UnsetProvenance()`

UnsetProvenance ensures that no value is present for Provenance, not even an explicit nil
### GetIsPreRelease

`func (o *PredictionItem) GetIsPreRelease() bool`

GetIsPreRelease returns the IsPreRelease field if non-nil, zero value otherwise.

### GetIsPreReleaseOk

`func (o *PredictionItem) GetIsPreReleaseOk() (*bool, bool)`

GetIsPreReleaseOk returns a tuple with the IsPreRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreRelease

`func (o *PredictionItem) SetIsPreRelease(v bool)`

SetIsPreRelease sets IsPreRelease field to given value.

### HasIsPreRelease

`func (o *PredictionItem) HasIsPreRelease() bool`

HasIsPreRelease returns a boolean if a field has been set.

### SetIsPreReleaseNil

`func (o *PredictionItem) SetIsPreReleaseNil(b bool)`

 SetIsPreReleaseNil sets the value for IsPreRelease to be an explicit nil

### UnsetIsPreRelease
`func (o *PredictionItem) UnsetIsPreRelease()`

UnsetIsPreRelease ensures that no value is present for IsPreRelease, not even an explicit nil
### GetConfidence

`func (o *PredictionItem) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *PredictionItem) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *PredictionItem) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.

### HasConfidence

`func (o *PredictionItem) HasConfidence() bool`

HasConfidence returns a boolean if a field has been set.

### SetConfidenceNil

`func (o *PredictionItem) SetConfidenceNil(b bool)`

 SetConfidenceNil sets the value for Confidence to be an explicit nil

### UnsetConfidence
`func (o *PredictionItem) UnsetConfidence()`

UnsetConfidence ensures that no value is present for Confidence, not even an explicit nil
### GetPredictionReason

`func (o *PredictionItem) GetPredictionReason() string`

GetPredictionReason returns the PredictionReason field if non-nil, zero value otherwise.

### GetPredictionReasonOk

`func (o *PredictionItem) GetPredictionReasonOk() (*string, bool)`

GetPredictionReasonOk returns a tuple with the PredictionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictionReason

`func (o *PredictionItem) SetPredictionReason(v string)`

SetPredictionReason sets PredictionReason field to given value.

### HasPredictionReason

`func (o *PredictionItem) HasPredictionReason() bool`

HasPredictionReason returns a boolean if a field has been set.

### SetPredictionReasonNil

`func (o *PredictionItem) SetPredictionReasonNil(b bool)`

 SetPredictionReasonNil sets the value for PredictionReason to be an explicit nil

### UnsetPredictionReason
`func (o *PredictionItem) UnsetPredictionReason()`

UnsetPredictionReason ensures that no value is present for PredictionReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


