# FinancialObservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObservationId** | **string** | Stable identity of the source, observation date, instrument and quote dimensions. | 
**SourceId** | **string** |  | 
**Date** | **string** | Source observation date; not an inferred publication date. | 
**Val** | **float64** | Official numerical value in unit, retaining source precision. | 
**Currency** | **string** |  | 
**Issuer** | **string** | Issuer, official benchmark administrator, or explicitly identified market/curve basket. | 
**InstrumentId** | **string** |  | 
**InstrumentName** | Pointer to **NullableString** |  | [optional] 
**SourceSeriesId** | Pointer to **NullableString** |  | [optional] 
**Measure** | **string** |  | 
**Unit** | **string** |  | 
**SourceUrl** | **string** |  | 
**Basis** | Pointer to **NullableString** |  | [optional] 
**CurveType** | Pointer to **NullableString** |  | [optional] 
**Tenor** | Pointer to **NullableString** |  | [optional] 
**MaturityYears** | Pointer to **NullableFloat64** |  | [optional] 
**MaturityDate** | Pointer to **NullableString** |  | [optional] 
**PriceSide** | Pointer to **NullableString** |  | [optional] 
**PriceBasis** | Pointer to **NullableString** |  | [optional] 
**PriceDenominator** | Pointer to **NullableFloat64** |  | [optional] 
**CouponRate** | Pointer to **NullableFloat64** |  | [optional] 
**QuoteType** | Pointer to **NullableString** |  | [optional] 
**YieldConvention** | Pointer to **NullableString** |  | [optional] 
**ReferenceRate** | Pointer to **NullableString** |  | [optional] 
**InflationIndex** | Pointer to **NullableString** |  | [optional] 
**Compounding** | Pointer to **NullableString** |  | [optional] 
**AveragingPeriodMonths** | Pointer to **NullableFloat64** |  | [optional] 
**Percentile** | Pointer to **NullableFloat64** |  | [optional] 
**PublicationDate** | Pointer to **NullableString** |  | [optional] 
**SourcePublicationDate** | Pointer to **NullableString** |  | [optional] 
**SourceValueDate** | Pointer to **NullableString** |  | [optional] 
**RatePeriodStart** | Pointer to **NullableString** |  | [optional] 
**RatePeriodEnd** | Pointer to **NullableString** |  | [optional] 
**ObservationPeriodStart** | Pointer to **NullableString** |  | [optional] 
**ObservationPeriodEnd** | Pointer to **NullableString** |  | [optional] 
**ObservationStatus** | Pointer to **NullableString** |  | [optional] 
**AnnouncementDatetime** | Pointer to **NullableInt64** | Publication epoch seconds when available. A scheduled hour is explicitly assumed. | [optional] 
**PublicationTimeConfirmed** | **bool** |  | 
**ReleaseTimeAssumed** | **bool** |  | 
**PointInTimeSafe** | **bool** | True only when the source supplied exact publication-time evidence. | 

## Methods

### NewFinancialObservation

`func NewFinancialObservation(observationId string, sourceId string, date string, val float64, currency string, issuer string, instrumentId string, measure string, unit string, sourceUrl string, publicationTimeConfirmed bool, releaseTimeAssumed bool, pointInTimeSafe bool, ) *FinancialObservation`

NewFinancialObservation instantiates a new FinancialObservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFinancialObservationWithDefaults

`func NewFinancialObservationWithDefaults() *FinancialObservation`

NewFinancialObservationWithDefaults instantiates a new FinancialObservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservationId

`func (o *FinancialObservation) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *FinancialObservation) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *FinancialObservation) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.


### GetSourceId

`func (o *FinancialObservation) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *FinancialObservation) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *FinancialObservation) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetDate

`func (o *FinancialObservation) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FinancialObservation) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FinancialObservation) SetDate(v string)`

SetDate sets Date field to given value.


### GetVal

`func (o *FinancialObservation) GetVal() float64`

GetVal returns the Val field if non-nil, zero value otherwise.

### GetValOk

`func (o *FinancialObservation) GetValOk() (*float64, bool)`

GetValOk returns a tuple with the Val field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVal

`func (o *FinancialObservation) SetVal(v float64)`

SetVal sets Val field to given value.


### GetCurrency

`func (o *FinancialObservation) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FinancialObservation) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FinancialObservation) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetIssuer

`func (o *FinancialObservation) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *FinancialObservation) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *FinancialObservation) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetInstrumentId

`func (o *FinancialObservation) GetInstrumentId() string`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *FinancialObservation) GetInstrumentIdOk() (*string, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *FinancialObservation) SetInstrumentId(v string)`

SetInstrumentId sets InstrumentId field to given value.


### GetInstrumentName

`func (o *FinancialObservation) GetInstrumentName() string`

GetInstrumentName returns the InstrumentName field if non-nil, zero value otherwise.

### GetInstrumentNameOk

`func (o *FinancialObservation) GetInstrumentNameOk() (*string, bool)`

GetInstrumentNameOk returns a tuple with the InstrumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentName

`func (o *FinancialObservation) SetInstrumentName(v string)`

SetInstrumentName sets InstrumentName field to given value.

### HasInstrumentName

`func (o *FinancialObservation) HasInstrumentName() bool`

HasInstrumentName returns a boolean if a field has been set.

### SetInstrumentNameNil

`func (o *FinancialObservation) SetInstrumentNameNil(b bool)`

 SetInstrumentNameNil sets the value for InstrumentName to be an explicit nil

### UnsetInstrumentName
`func (o *FinancialObservation) UnsetInstrumentName()`

UnsetInstrumentName ensures that no value is present for InstrumentName, not even an explicit nil
### GetSourceSeriesId

`func (o *FinancialObservation) GetSourceSeriesId() string`

GetSourceSeriesId returns the SourceSeriesId field if non-nil, zero value otherwise.

### GetSourceSeriesIdOk

`func (o *FinancialObservation) GetSourceSeriesIdOk() (*string, bool)`

GetSourceSeriesIdOk returns a tuple with the SourceSeriesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSeriesId

`func (o *FinancialObservation) SetSourceSeriesId(v string)`

SetSourceSeriesId sets SourceSeriesId field to given value.

### HasSourceSeriesId

`func (o *FinancialObservation) HasSourceSeriesId() bool`

HasSourceSeriesId returns a boolean if a field has been set.

### SetSourceSeriesIdNil

`func (o *FinancialObservation) SetSourceSeriesIdNil(b bool)`

 SetSourceSeriesIdNil sets the value for SourceSeriesId to be an explicit nil

### UnsetSourceSeriesId
`func (o *FinancialObservation) UnsetSourceSeriesId()`

UnsetSourceSeriesId ensures that no value is present for SourceSeriesId, not even an explicit nil
### GetMeasure

`func (o *FinancialObservation) GetMeasure() string`

GetMeasure returns the Measure field if non-nil, zero value otherwise.

### GetMeasureOk

`func (o *FinancialObservation) GetMeasureOk() (*string, bool)`

GetMeasureOk returns a tuple with the Measure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasure

`func (o *FinancialObservation) SetMeasure(v string)`

SetMeasure sets Measure field to given value.


### GetUnit

`func (o *FinancialObservation) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *FinancialObservation) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *FinancialObservation) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetSourceUrl

`func (o *FinancialObservation) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *FinancialObservation) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *FinancialObservation) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.


### GetBasis

`func (o *FinancialObservation) GetBasis() string`

GetBasis returns the Basis field if non-nil, zero value otherwise.

### GetBasisOk

`func (o *FinancialObservation) GetBasisOk() (*string, bool)`

GetBasisOk returns a tuple with the Basis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasis

`func (o *FinancialObservation) SetBasis(v string)`

SetBasis sets Basis field to given value.

### HasBasis

`func (o *FinancialObservation) HasBasis() bool`

HasBasis returns a boolean if a field has been set.

### SetBasisNil

`func (o *FinancialObservation) SetBasisNil(b bool)`

 SetBasisNil sets the value for Basis to be an explicit nil

### UnsetBasis
`func (o *FinancialObservation) UnsetBasis()`

UnsetBasis ensures that no value is present for Basis, not even an explicit nil
### GetCurveType

`func (o *FinancialObservation) GetCurveType() string`

GetCurveType returns the CurveType field if non-nil, zero value otherwise.

### GetCurveTypeOk

`func (o *FinancialObservation) GetCurveTypeOk() (*string, bool)`

GetCurveTypeOk returns a tuple with the CurveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurveType

`func (o *FinancialObservation) SetCurveType(v string)`

SetCurveType sets CurveType field to given value.

### HasCurveType

`func (o *FinancialObservation) HasCurveType() bool`

HasCurveType returns a boolean if a field has been set.

### SetCurveTypeNil

`func (o *FinancialObservation) SetCurveTypeNil(b bool)`

 SetCurveTypeNil sets the value for CurveType to be an explicit nil

### UnsetCurveType
`func (o *FinancialObservation) UnsetCurveType()`

UnsetCurveType ensures that no value is present for CurveType, not even an explicit nil
### GetTenor

`func (o *FinancialObservation) GetTenor() string`

GetTenor returns the Tenor field if non-nil, zero value otherwise.

### GetTenorOk

`func (o *FinancialObservation) GetTenorOk() (*string, bool)`

GetTenorOk returns a tuple with the Tenor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenor

`func (o *FinancialObservation) SetTenor(v string)`

SetTenor sets Tenor field to given value.

### HasTenor

`func (o *FinancialObservation) HasTenor() bool`

HasTenor returns a boolean if a field has been set.

### SetTenorNil

`func (o *FinancialObservation) SetTenorNil(b bool)`

 SetTenorNil sets the value for Tenor to be an explicit nil

### UnsetTenor
`func (o *FinancialObservation) UnsetTenor()`

UnsetTenor ensures that no value is present for Tenor, not even an explicit nil
### GetMaturityYears

`func (o *FinancialObservation) GetMaturityYears() float64`

GetMaturityYears returns the MaturityYears field if non-nil, zero value otherwise.

### GetMaturityYearsOk

`func (o *FinancialObservation) GetMaturityYearsOk() (*float64, bool)`

GetMaturityYearsOk returns a tuple with the MaturityYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityYears

`func (o *FinancialObservation) SetMaturityYears(v float64)`

SetMaturityYears sets MaturityYears field to given value.

### HasMaturityYears

`func (o *FinancialObservation) HasMaturityYears() bool`

HasMaturityYears returns a boolean if a field has been set.

### SetMaturityYearsNil

`func (o *FinancialObservation) SetMaturityYearsNil(b bool)`

 SetMaturityYearsNil sets the value for MaturityYears to be an explicit nil

### UnsetMaturityYears
`func (o *FinancialObservation) UnsetMaturityYears()`

UnsetMaturityYears ensures that no value is present for MaturityYears, not even an explicit nil
### GetMaturityDate

`func (o *FinancialObservation) GetMaturityDate() string`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *FinancialObservation) GetMaturityDateOk() (*string, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *FinancialObservation) SetMaturityDate(v string)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *FinancialObservation) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### SetMaturityDateNil

`func (o *FinancialObservation) SetMaturityDateNil(b bool)`

 SetMaturityDateNil sets the value for MaturityDate to be an explicit nil

### UnsetMaturityDate
`func (o *FinancialObservation) UnsetMaturityDate()`

UnsetMaturityDate ensures that no value is present for MaturityDate, not even an explicit nil
### GetPriceSide

`func (o *FinancialObservation) GetPriceSide() string`

GetPriceSide returns the PriceSide field if non-nil, zero value otherwise.

### GetPriceSideOk

`func (o *FinancialObservation) GetPriceSideOk() (*string, bool)`

GetPriceSideOk returns a tuple with the PriceSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSide

`func (o *FinancialObservation) SetPriceSide(v string)`

SetPriceSide sets PriceSide field to given value.

### HasPriceSide

`func (o *FinancialObservation) HasPriceSide() bool`

HasPriceSide returns a boolean if a field has been set.

### SetPriceSideNil

`func (o *FinancialObservation) SetPriceSideNil(b bool)`

 SetPriceSideNil sets the value for PriceSide to be an explicit nil

### UnsetPriceSide
`func (o *FinancialObservation) UnsetPriceSide()`

UnsetPriceSide ensures that no value is present for PriceSide, not even an explicit nil
### GetPriceBasis

`func (o *FinancialObservation) GetPriceBasis() string`

GetPriceBasis returns the PriceBasis field if non-nil, zero value otherwise.

### GetPriceBasisOk

`func (o *FinancialObservation) GetPriceBasisOk() (*string, bool)`

GetPriceBasisOk returns a tuple with the PriceBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBasis

`func (o *FinancialObservation) SetPriceBasis(v string)`

SetPriceBasis sets PriceBasis field to given value.

### HasPriceBasis

`func (o *FinancialObservation) HasPriceBasis() bool`

HasPriceBasis returns a boolean if a field has been set.

### SetPriceBasisNil

`func (o *FinancialObservation) SetPriceBasisNil(b bool)`

 SetPriceBasisNil sets the value for PriceBasis to be an explicit nil

### UnsetPriceBasis
`func (o *FinancialObservation) UnsetPriceBasis()`

UnsetPriceBasis ensures that no value is present for PriceBasis, not even an explicit nil
### GetPriceDenominator

`func (o *FinancialObservation) GetPriceDenominator() float64`

GetPriceDenominator returns the PriceDenominator field if non-nil, zero value otherwise.

### GetPriceDenominatorOk

`func (o *FinancialObservation) GetPriceDenominatorOk() (*float64, bool)`

GetPriceDenominatorOk returns a tuple with the PriceDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDenominator

`func (o *FinancialObservation) SetPriceDenominator(v float64)`

SetPriceDenominator sets PriceDenominator field to given value.

### HasPriceDenominator

`func (o *FinancialObservation) HasPriceDenominator() bool`

HasPriceDenominator returns a boolean if a field has been set.

### SetPriceDenominatorNil

`func (o *FinancialObservation) SetPriceDenominatorNil(b bool)`

 SetPriceDenominatorNil sets the value for PriceDenominator to be an explicit nil

### UnsetPriceDenominator
`func (o *FinancialObservation) UnsetPriceDenominator()`

UnsetPriceDenominator ensures that no value is present for PriceDenominator, not even an explicit nil
### GetCouponRate

`func (o *FinancialObservation) GetCouponRate() float64`

GetCouponRate returns the CouponRate field if non-nil, zero value otherwise.

### GetCouponRateOk

`func (o *FinancialObservation) GetCouponRateOk() (*float64, bool)`

GetCouponRateOk returns a tuple with the CouponRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponRate

`func (o *FinancialObservation) SetCouponRate(v float64)`

SetCouponRate sets CouponRate field to given value.

### HasCouponRate

`func (o *FinancialObservation) HasCouponRate() bool`

HasCouponRate returns a boolean if a field has been set.

### SetCouponRateNil

`func (o *FinancialObservation) SetCouponRateNil(b bool)`

 SetCouponRateNil sets the value for CouponRate to be an explicit nil

### UnsetCouponRate
`func (o *FinancialObservation) UnsetCouponRate()`

UnsetCouponRate ensures that no value is present for CouponRate, not even an explicit nil
### GetQuoteType

`func (o *FinancialObservation) GetQuoteType() string`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *FinancialObservation) GetQuoteTypeOk() (*string, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *FinancialObservation) SetQuoteType(v string)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *FinancialObservation) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.

### SetQuoteTypeNil

`func (o *FinancialObservation) SetQuoteTypeNil(b bool)`

 SetQuoteTypeNil sets the value for QuoteType to be an explicit nil

### UnsetQuoteType
`func (o *FinancialObservation) UnsetQuoteType()`

UnsetQuoteType ensures that no value is present for QuoteType, not even an explicit nil
### GetYieldConvention

`func (o *FinancialObservation) GetYieldConvention() string`

GetYieldConvention returns the YieldConvention field if non-nil, zero value otherwise.

### GetYieldConventionOk

`func (o *FinancialObservation) GetYieldConventionOk() (*string, bool)`

GetYieldConventionOk returns a tuple with the YieldConvention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYieldConvention

`func (o *FinancialObservation) SetYieldConvention(v string)`

SetYieldConvention sets YieldConvention field to given value.

### HasYieldConvention

`func (o *FinancialObservation) HasYieldConvention() bool`

HasYieldConvention returns a boolean if a field has been set.

### SetYieldConventionNil

`func (o *FinancialObservation) SetYieldConventionNil(b bool)`

 SetYieldConventionNil sets the value for YieldConvention to be an explicit nil

### UnsetYieldConvention
`func (o *FinancialObservation) UnsetYieldConvention()`

UnsetYieldConvention ensures that no value is present for YieldConvention, not even an explicit nil
### GetReferenceRate

`func (o *FinancialObservation) GetReferenceRate() string`

GetReferenceRate returns the ReferenceRate field if non-nil, zero value otherwise.

### GetReferenceRateOk

`func (o *FinancialObservation) GetReferenceRateOk() (*string, bool)`

GetReferenceRateOk returns a tuple with the ReferenceRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceRate

`func (o *FinancialObservation) SetReferenceRate(v string)`

SetReferenceRate sets ReferenceRate field to given value.

### HasReferenceRate

`func (o *FinancialObservation) HasReferenceRate() bool`

HasReferenceRate returns a boolean if a field has been set.

### SetReferenceRateNil

`func (o *FinancialObservation) SetReferenceRateNil(b bool)`

 SetReferenceRateNil sets the value for ReferenceRate to be an explicit nil

### UnsetReferenceRate
`func (o *FinancialObservation) UnsetReferenceRate()`

UnsetReferenceRate ensures that no value is present for ReferenceRate, not even an explicit nil
### GetInflationIndex

`func (o *FinancialObservation) GetInflationIndex() string`

GetInflationIndex returns the InflationIndex field if non-nil, zero value otherwise.

### GetInflationIndexOk

`func (o *FinancialObservation) GetInflationIndexOk() (*string, bool)`

GetInflationIndexOk returns a tuple with the InflationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInflationIndex

`func (o *FinancialObservation) SetInflationIndex(v string)`

SetInflationIndex sets InflationIndex field to given value.

### HasInflationIndex

`func (o *FinancialObservation) HasInflationIndex() bool`

HasInflationIndex returns a boolean if a field has been set.

### SetInflationIndexNil

`func (o *FinancialObservation) SetInflationIndexNil(b bool)`

 SetInflationIndexNil sets the value for InflationIndex to be an explicit nil

### UnsetInflationIndex
`func (o *FinancialObservation) UnsetInflationIndex()`

UnsetInflationIndex ensures that no value is present for InflationIndex, not even an explicit nil
### GetCompounding

`func (o *FinancialObservation) GetCompounding() string`

GetCompounding returns the Compounding field if non-nil, zero value otherwise.

### GetCompoundingOk

`func (o *FinancialObservation) GetCompoundingOk() (*string, bool)`

GetCompoundingOk returns a tuple with the Compounding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompounding

`func (o *FinancialObservation) SetCompounding(v string)`

SetCompounding sets Compounding field to given value.

### HasCompounding

`func (o *FinancialObservation) HasCompounding() bool`

HasCompounding returns a boolean if a field has been set.

### SetCompoundingNil

`func (o *FinancialObservation) SetCompoundingNil(b bool)`

 SetCompoundingNil sets the value for Compounding to be an explicit nil

### UnsetCompounding
`func (o *FinancialObservation) UnsetCompounding()`

UnsetCompounding ensures that no value is present for Compounding, not even an explicit nil
### GetAveragingPeriodMonths

`func (o *FinancialObservation) GetAveragingPeriodMonths() float64`

GetAveragingPeriodMonths returns the AveragingPeriodMonths field if non-nil, zero value otherwise.

### GetAveragingPeriodMonthsOk

`func (o *FinancialObservation) GetAveragingPeriodMonthsOk() (*float64, bool)`

GetAveragingPeriodMonthsOk returns a tuple with the AveragingPeriodMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragingPeriodMonths

`func (o *FinancialObservation) SetAveragingPeriodMonths(v float64)`

SetAveragingPeriodMonths sets AveragingPeriodMonths field to given value.

### HasAveragingPeriodMonths

`func (o *FinancialObservation) HasAveragingPeriodMonths() bool`

HasAveragingPeriodMonths returns a boolean if a field has been set.

### SetAveragingPeriodMonthsNil

`func (o *FinancialObservation) SetAveragingPeriodMonthsNil(b bool)`

 SetAveragingPeriodMonthsNil sets the value for AveragingPeriodMonths to be an explicit nil

### UnsetAveragingPeriodMonths
`func (o *FinancialObservation) UnsetAveragingPeriodMonths()`

UnsetAveragingPeriodMonths ensures that no value is present for AveragingPeriodMonths, not even an explicit nil
### GetPercentile

`func (o *FinancialObservation) GetPercentile() float64`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *FinancialObservation) GetPercentileOk() (*float64, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *FinancialObservation) SetPercentile(v float64)`

SetPercentile sets Percentile field to given value.

### HasPercentile

`func (o *FinancialObservation) HasPercentile() bool`

HasPercentile returns a boolean if a field has been set.

### SetPercentileNil

`func (o *FinancialObservation) SetPercentileNil(b bool)`

 SetPercentileNil sets the value for Percentile to be an explicit nil

### UnsetPercentile
`func (o *FinancialObservation) UnsetPercentile()`

UnsetPercentile ensures that no value is present for Percentile, not even an explicit nil
### GetPublicationDate

`func (o *FinancialObservation) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *FinancialObservation) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *FinancialObservation) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *FinancialObservation) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### SetPublicationDateNil

`func (o *FinancialObservation) SetPublicationDateNil(b bool)`

 SetPublicationDateNil sets the value for PublicationDate to be an explicit nil

### UnsetPublicationDate
`func (o *FinancialObservation) UnsetPublicationDate()`

UnsetPublicationDate ensures that no value is present for PublicationDate, not even an explicit nil
### GetSourcePublicationDate

`func (o *FinancialObservation) GetSourcePublicationDate() string`

GetSourcePublicationDate returns the SourcePublicationDate field if non-nil, zero value otherwise.

### GetSourcePublicationDateOk

`func (o *FinancialObservation) GetSourcePublicationDateOk() (*string, bool)`

GetSourcePublicationDateOk returns a tuple with the SourcePublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePublicationDate

`func (o *FinancialObservation) SetSourcePublicationDate(v string)`

SetSourcePublicationDate sets SourcePublicationDate field to given value.

### HasSourcePublicationDate

`func (o *FinancialObservation) HasSourcePublicationDate() bool`

HasSourcePublicationDate returns a boolean if a field has been set.

### SetSourcePublicationDateNil

`func (o *FinancialObservation) SetSourcePublicationDateNil(b bool)`

 SetSourcePublicationDateNil sets the value for SourcePublicationDate to be an explicit nil

### UnsetSourcePublicationDate
`func (o *FinancialObservation) UnsetSourcePublicationDate()`

UnsetSourcePublicationDate ensures that no value is present for SourcePublicationDate, not even an explicit nil
### GetSourceValueDate

`func (o *FinancialObservation) GetSourceValueDate() string`

GetSourceValueDate returns the SourceValueDate field if non-nil, zero value otherwise.

### GetSourceValueDateOk

`func (o *FinancialObservation) GetSourceValueDateOk() (*string, bool)`

GetSourceValueDateOk returns a tuple with the SourceValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceValueDate

`func (o *FinancialObservation) SetSourceValueDate(v string)`

SetSourceValueDate sets SourceValueDate field to given value.

### HasSourceValueDate

`func (o *FinancialObservation) HasSourceValueDate() bool`

HasSourceValueDate returns a boolean if a field has been set.

### SetSourceValueDateNil

`func (o *FinancialObservation) SetSourceValueDateNil(b bool)`

 SetSourceValueDateNil sets the value for SourceValueDate to be an explicit nil

### UnsetSourceValueDate
`func (o *FinancialObservation) UnsetSourceValueDate()`

UnsetSourceValueDate ensures that no value is present for SourceValueDate, not even an explicit nil
### GetRatePeriodStart

`func (o *FinancialObservation) GetRatePeriodStart() string`

GetRatePeriodStart returns the RatePeriodStart field if non-nil, zero value otherwise.

### GetRatePeriodStartOk

`func (o *FinancialObservation) GetRatePeriodStartOk() (*string, bool)`

GetRatePeriodStartOk returns a tuple with the RatePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePeriodStart

`func (o *FinancialObservation) SetRatePeriodStart(v string)`

SetRatePeriodStart sets RatePeriodStart field to given value.

### HasRatePeriodStart

`func (o *FinancialObservation) HasRatePeriodStart() bool`

HasRatePeriodStart returns a boolean if a field has been set.

### SetRatePeriodStartNil

`func (o *FinancialObservation) SetRatePeriodStartNil(b bool)`

 SetRatePeriodStartNil sets the value for RatePeriodStart to be an explicit nil

### UnsetRatePeriodStart
`func (o *FinancialObservation) UnsetRatePeriodStart()`

UnsetRatePeriodStart ensures that no value is present for RatePeriodStart, not even an explicit nil
### GetRatePeriodEnd

`func (o *FinancialObservation) GetRatePeriodEnd() string`

GetRatePeriodEnd returns the RatePeriodEnd field if non-nil, zero value otherwise.

### GetRatePeriodEndOk

`func (o *FinancialObservation) GetRatePeriodEndOk() (*string, bool)`

GetRatePeriodEndOk returns a tuple with the RatePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePeriodEnd

`func (o *FinancialObservation) SetRatePeriodEnd(v string)`

SetRatePeriodEnd sets RatePeriodEnd field to given value.

### HasRatePeriodEnd

`func (o *FinancialObservation) HasRatePeriodEnd() bool`

HasRatePeriodEnd returns a boolean if a field has been set.

### SetRatePeriodEndNil

`func (o *FinancialObservation) SetRatePeriodEndNil(b bool)`

 SetRatePeriodEndNil sets the value for RatePeriodEnd to be an explicit nil

### UnsetRatePeriodEnd
`func (o *FinancialObservation) UnsetRatePeriodEnd()`

UnsetRatePeriodEnd ensures that no value is present for RatePeriodEnd, not even an explicit nil
### GetObservationPeriodStart

`func (o *FinancialObservation) GetObservationPeriodStart() string`

GetObservationPeriodStart returns the ObservationPeriodStart field if non-nil, zero value otherwise.

### GetObservationPeriodStartOk

`func (o *FinancialObservation) GetObservationPeriodStartOk() (*string, bool)`

GetObservationPeriodStartOk returns a tuple with the ObservationPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationPeriodStart

`func (o *FinancialObservation) SetObservationPeriodStart(v string)`

SetObservationPeriodStart sets ObservationPeriodStart field to given value.

### HasObservationPeriodStart

`func (o *FinancialObservation) HasObservationPeriodStart() bool`

HasObservationPeriodStart returns a boolean if a field has been set.

### SetObservationPeriodStartNil

`func (o *FinancialObservation) SetObservationPeriodStartNil(b bool)`

 SetObservationPeriodStartNil sets the value for ObservationPeriodStart to be an explicit nil

### UnsetObservationPeriodStart
`func (o *FinancialObservation) UnsetObservationPeriodStart()`

UnsetObservationPeriodStart ensures that no value is present for ObservationPeriodStart, not even an explicit nil
### GetObservationPeriodEnd

`func (o *FinancialObservation) GetObservationPeriodEnd() string`

GetObservationPeriodEnd returns the ObservationPeriodEnd field if non-nil, zero value otherwise.

### GetObservationPeriodEndOk

`func (o *FinancialObservation) GetObservationPeriodEndOk() (*string, bool)`

GetObservationPeriodEndOk returns a tuple with the ObservationPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationPeriodEnd

`func (o *FinancialObservation) SetObservationPeriodEnd(v string)`

SetObservationPeriodEnd sets ObservationPeriodEnd field to given value.

### HasObservationPeriodEnd

`func (o *FinancialObservation) HasObservationPeriodEnd() bool`

HasObservationPeriodEnd returns a boolean if a field has been set.

### SetObservationPeriodEndNil

`func (o *FinancialObservation) SetObservationPeriodEndNil(b bool)`

 SetObservationPeriodEndNil sets the value for ObservationPeriodEnd to be an explicit nil

### UnsetObservationPeriodEnd
`func (o *FinancialObservation) UnsetObservationPeriodEnd()`

UnsetObservationPeriodEnd ensures that no value is present for ObservationPeriodEnd, not even an explicit nil
### GetObservationStatus

`func (o *FinancialObservation) GetObservationStatus() string`

GetObservationStatus returns the ObservationStatus field if non-nil, zero value otherwise.

### GetObservationStatusOk

`func (o *FinancialObservation) GetObservationStatusOk() (*string, bool)`

GetObservationStatusOk returns a tuple with the ObservationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationStatus

`func (o *FinancialObservation) SetObservationStatus(v string)`

SetObservationStatus sets ObservationStatus field to given value.

### HasObservationStatus

`func (o *FinancialObservation) HasObservationStatus() bool`

HasObservationStatus returns a boolean if a field has been set.

### SetObservationStatusNil

`func (o *FinancialObservation) SetObservationStatusNil(b bool)`

 SetObservationStatusNil sets the value for ObservationStatus to be an explicit nil

### UnsetObservationStatus
`func (o *FinancialObservation) UnsetObservationStatus()`

UnsetObservationStatus ensures that no value is present for ObservationStatus, not even an explicit nil
### GetAnnouncementDatetime

`func (o *FinancialObservation) GetAnnouncementDatetime() int64`

GetAnnouncementDatetime returns the AnnouncementDatetime field if non-nil, zero value otherwise.

### GetAnnouncementDatetimeOk

`func (o *FinancialObservation) GetAnnouncementDatetimeOk() (*int64, bool)`

GetAnnouncementDatetimeOk returns a tuple with the AnnouncementDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncementDatetime

`func (o *FinancialObservation) SetAnnouncementDatetime(v int64)`

SetAnnouncementDatetime sets AnnouncementDatetime field to given value.

### HasAnnouncementDatetime

`func (o *FinancialObservation) HasAnnouncementDatetime() bool`

HasAnnouncementDatetime returns a boolean if a field has been set.

### SetAnnouncementDatetimeNil

`func (o *FinancialObservation) SetAnnouncementDatetimeNil(b bool)`

 SetAnnouncementDatetimeNil sets the value for AnnouncementDatetime to be an explicit nil

### UnsetAnnouncementDatetime
`func (o *FinancialObservation) UnsetAnnouncementDatetime()`

UnsetAnnouncementDatetime ensures that no value is present for AnnouncementDatetime, not even an explicit nil
### GetPublicationTimeConfirmed

`func (o *FinancialObservation) GetPublicationTimeConfirmed() bool`

GetPublicationTimeConfirmed returns the PublicationTimeConfirmed field if non-nil, zero value otherwise.

### GetPublicationTimeConfirmedOk

`func (o *FinancialObservation) GetPublicationTimeConfirmedOk() (*bool, bool)`

GetPublicationTimeConfirmedOk returns a tuple with the PublicationTimeConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationTimeConfirmed

`func (o *FinancialObservation) SetPublicationTimeConfirmed(v bool)`

SetPublicationTimeConfirmed sets PublicationTimeConfirmed field to given value.


### GetReleaseTimeAssumed

`func (o *FinancialObservation) GetReleaseTimeAssumed() bool`

GetReleaseTimeAssumed returns the ReleaseTimeAssumed field if non-nil, zero value otherwise.

### GetReleaseTimeAssumedOk

`func (o *FinancialObservation) GetReleaseTimeAssumedOk() (*bool, bool)`

GetReleaseTimeAssumedOk returns a tuple with the ReleaseTimeAssumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTimeAssumed

`func (o *FinancialObservation) SetReleaseTimeAssumed(v bool)`

SetReleaseTimeAssumed sets ReleaseTimeAssumed field to given value.


### GetPointInTimeSafe

`func (o *FinancialObservation) GetPointInTimeSafe() bool`

GetPointInTimeSafe returns the PointInTimeSafe field if non-nil, zero value otherwise.

### GetPointInTimeSafeOk

`func (o *FinancialObservation) GetPointInTimeSafeOk() (*bool, bool)`

GetPointInTimeSafeOk returns a tuple with the PointInTimeSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointInTimeSafe

`func (o *FinancialObservation) SetPointInTimeSafe(v bool)`

SetPointInTimeSafe sets PointInTimeSafe field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


