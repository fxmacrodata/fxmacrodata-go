# \FXMacroDataAPI

All URIs are relative to *https://api.fxmacrodata.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnnouncementChanges**](FXMacroDataAPI.md#GetAnnouncementChanges) | **Get** /v1/announcements/changes | Poll recent live announcement release changes
[**GetAnnouncements**](FXMacroDataAPI.md#GetAnnouncements) | **Get** /v1/announcements/{currency}/{indicator} | Get macroeconomic indicator data
[**GetCalendar**](FXMacroDataAPI.md#GetCalendar) | **Get** /v1/calendar/{currency} | Get upcoming economic data release dates
[**GetCommodities**](FXMacroDataAPI.md#GetCommodities) | **Get** /v1/commodities/{indicator} | Get commodity indicator data
[**GetCot**](FXMacroDataAPI.md#GetCot) | **Get** /v1/cot/{currency} | Get CFTC Commitment of Traders (COT) data
[**GetCurves**](FXMacroDataAPI.md#GetCurves) | **Get** /v1/curves/{currency} | Get curve analytics by currency
[**GetDataCatalogue**](FXMacroDataAPI.md#GetDataCatalogue) | **Get** /v1/data_catalogue/{currency} | Get Currently Available Macroeconomic Indicators
[**GetFactor**](FXMacroDataAPI.md#GetFactor) | **Get** /v1/factors/{currency}/{factor} | Get a precomputed FX factor for one currency
[**GetFinancialPrices**](FXMacroDataAPI.md#GetFinancialPrices) | **Get** /v1/financial_prices/{currency} | Official financial prices and yield curves
[**GetForex**](FXMacroDataAPI.md#GetForex) | **Get** /v1/forex/{base}/{quote} | Get FX spot rates
[**GetFxIntradayReferenceRates**](FXMacroDataAPI.md#GetFxIntradayReferenceRates) | **Get** /v1/fx/intraday-reference-rates/{base}/{quote} | Get subscriber intraday FX reference rates
[**GetFxReferenceRates**](FXMacroDataAPI.md#GetFxReferenceRates) | **Get** /v1/fx/reference-rates/{base}/{quote} | Get timestamped official FX reference rates
[**GetFxSessionRates**](FXMacroDataAPI.md#GetFxSessionRates) | **Get** /v1/fx/session-rates/{base}/{quote} | Get official FX session-close reference rates
[**GetFxSourceUniverse**](FXMacroDataAPI.md#GetFxSourceUniverse) | **Get** /v1/fx/source-universe | List public FX reference-rate pair universe
[**GetFxSources**](FXMacroDataAPI.md#GetFxSources) | **Get** /v1/fx/sources | List public FX reference-rate sources
[**GetHealth**](FXMacroDataAPI.md#GetHealth) | **Get** /v1/health | Ping the API
[**GetLatestAnnouncements**](FXMacroDataAPI.md#GetLatestAnnouncements) | **Get** /v1/announcements/{currency}/latest | Get latest values for all indicators in a currency
[**GetLatestCommodities**](FXMacroDataAPI.md#GetLatestCommodities) | **Get** /v1/commodities/latest | Get latest values for all commodity indicators
[**GetMarketSessions**](FXMacroDataAPI.md#GetMarketSessions) | **Get** /v1/market_sessions | FX market session timetable
[**GetPredictionCoverage**](FXMacroDataAPI.md#GetPredictionCoverage) | **Get** /v1/predictions/coverage | Which forecast publishers cover which currency (no API key)
[**GetPredictionCoverageForCurrency**](FXMacroDataAPI.md#GetPredictionCoverageForCurrency) | **Get** /v1/predictions/coverage/{currency} | Per-indicator forecast coverage for one currency (no API key)
[**GetPredictions**](FXMacroDataAPI.md#GetPredictions) | **Get** /v1/predictions/{currency}/{indicator} | List predictions for one indicator of a currency
[**GetPressReleases**](FXMacroDataAPI.md#GetPressReleases) | **Get** /v1/press-releases/{currency} | Recent central bank press releases
[**GetRateDifferentials**](FXMacroDataAPI.md#GetRateDifferentials) | **Get** /v1/rate_differentials/{base}/{quote} | Get pair rate differentials
[**GetRiskSentiment**](FXMacroDataAPI.md#GetRiskSentiment) | **Get** /v1/risk_sentiment | Global risk-on / risk-off indicator
[**Ping**](FXMacroDataAPI.md#Ping) | **Get** /v1/ping | Ping the API
[**ResearchPanel**](FXMacroDataAPI.md#ResearchPanel) | **Post** /v1/research/panel | Build a panel of values known at each decision time



## GetAnnouncementChanges

> AnnouncementChangesResponse GetAnnouncementChanges(ctx).Currencies(currencies).Indicators(indicators).Since(since).Limit(limit).Payload(payload).Execute()

Poll recent live announcement release changes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currencies := "currencies_example" // string | Optional query parameter. Comma-separated currency codes, `g10`, or `all`. Without a paid API key the result is restricted to USD. (optional)
	indicators := "indicators_example" // string | Optional query parameter. Optional comma-separated announcement indicator slugs. Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages. (optional)
	since := "since_example" // string | Optional query parameter. Cursor returned as `next_cursor` by the previous response. ISO 8601 timestamps and Unix epoch seconds are also accepted. (optional)
	limit := int64(789) // int64 | Optional query parameter. Maximum number of matching change events to return. (optional) (default to 100)
	payload := "payload_example" // string | Optional query parameter. Use `compact` for trigger payloads or `full` for row fields. Allowed values: compact, full. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetAnnouncementChanges(context.Background()).Currencies(currencies).Indicators(indicators).Since(since).Limit(limit).Payload(payload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetAnnouncementChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnnouncementChanges`: AnnouncementChangesResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetAnnouncementChanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnouncementChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currencies** | **string** | Optional query parameter. Comma-separated currency codes, &#x60;g10&#x60;, or &#x60;all&#x60;. Without a paid API key the result is restricted to USD. | 
 **indicators** | **string** | Optional query parameter. Optional comma-separated announcement indicator slugs. Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages. | 
 **since** | **string** | Optional query parameter. Cursor returned as &#x60;next_cursor&#x60; by the previous response. ISO 8601 timestamps and Unix epoch seconds are also accepted. | 
 **limit** | **int64** | Optional query parameter. Maximum number of matching change events to return. | [default to 100]
 **payload** | **string** | Optional query parameter. Use &#x60;compact&#x60; for trigger payloads or &#x60;full&#x60; for row fields. Allowed values: compact, full. | 

### Return type

[**AnnouncementChangesResponse**](AnnouncementChangesResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnnouncements

> AnnouncementResponse GetAnnouncements(ctx, currency, indicator).StartDate(startDate).EndDate(endDate).SeriesMode(seriesMode).Limit(limit).Offset(offset).Page(page).Seasonality(seasonality).Frequency(frequency).Annualization(annualization).PeriodAggregation(periodAggregation).Revisions(revisions).Basis(basis).OfficialOnly(officialOnly).AsOf(asOf).Availability(availability).DatasetVersion(datasetVersion).ValueMode(valueMode).Execute()

Get macroeconomic indicator data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "AUD" // string | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD.
	indicator := "afe_dollar_index" // string | Required path parameter. Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages.
	startDate := "startDate_example" // string | Optional query parameter. Start date YYYY-MM-DD (optional)
	endDate := "endDate_example" // string | Optional query parameter. End date YYYY-MM-DD (optional)
	seriesMode := "seriesMode_example" // string | Optional query parameter. Optional series rendering mode. `canonical` applies configured remap rules (for continuity across methodology transitions). `raw` returns the direct indicator series without remapping. (optional) (default to "raw")
	limit := int64(789) // int64 | Optional query parameter. Maximum number of rows to return in `data[]`. Rows are ordered most-recent-first, so `limit=10` returns the 10 most recent observations within the requested window. Use together with `offset` to page through long histories. Defaults to 20. Maximum 100. (optional) (default to 20)
	offset := int64(789) // int64 | Optional query parameter. Number of rows to skip from the start of `data[]` (after most-recent-first ordering). Use with `limit` for pagination. Defaults to 0. (optional) (default to 0)
	page := int64(789) // int64 | Optional query parameter. One-based page number for callers that prefer page-based pagination. When supplied, `offset` is derived as `(page - 1) * limit`. (optional)
	seasonality := "seasonality_example" // string | Optional query parameter. Optional seasonal adjustment selector. Supported public values are `sa` and `nsa`; check `supported_options.seasonality` in the response/catalogue for values available on this endpoint. Allowed values: sa, nsa. (optional)
	frequency := "frequency_example" // string | Optional query parameter. Optional returned change basis. Supported public values are `level`, `yoy`, `qoq`, and `mom`; check `supported_options.frequency` for values available on this endpoint. Allowed values: level, yoy, qoq, mom. (optional)
	annualization := "annualization_example" // string | Optional query parameter. Optional annualization selector. Supported public values are `not_annualized` and `saar`; check `supported_options.annualization` for values available on this endpoint. Allowed values: not_annualized, saar. (optional)
	periodAggregation := "periodAggregation_example" // string | Optional query parameter. Optional period-aggregation selector. Supported public values are `monthly` and `quarterly_average`; check `supported_options.period_aggregation` for values available on this endpoint. Allowed values: monthly, quarterly_average. (optional)
	revisions := "revisions_example" // string | Optional query parameter. Optional revision view. Supported public values are `latest`, `first`, `final`, and `all`; unsupported views return `unsupported_series_option` with supported values. Allowed values: latest, first, final, all. (optional)
	basis := "basis_example" // string | Optional query parameter. Optional real/nominal selector. Supported public values are `real` and `nominal`; only endpoints with known basis metadata accept it. Allowed values: real, nominal. (optional)
	officialOnly := true // bool | Optional query parameter. When true, return only official-source rows and remove explicit fallback observations. Prohibited private aggregator rows are always removed regardless of this setting. (optional) (default to true)
	asOf := "asOf_example" // string | Optional query parameter. (optional)
	availability := "availability_example" // string | Optional query parameter. (optional) (default to "public")
	datasetVersion := "datasetVersion_example" // string | Optional query parameter. (optional)
	valueMode := "valueMode_example" // string | Optional query parameter. (optional) (default to "normalized")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetAnnouncements(context.Background(), currency, indicator).StartDate(startDate).EndDate(endDate).SeriesMode(seriesMode).Limit(limit).Offset(offset).Page(page).Seasonality(seasonality).Frequency(frequency).Annualization(annualization).PeriodAggregation(periodAggregation).Revisions(revisions).Basis(basis).OfficialOnly(officialOnly).AsOf(asOf).Availability(availability).DatasetVersion(datasetVersion).ValueMode(valueMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetAnnouncements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAnnouncements`: AnnouncementResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetAnnouncements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD. | 
**indicator** | **string** | Required path parameter. Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnnouncementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **string** | Optional query parameter. Start date YYYY-MM-DD | 
 **endDate** | **string** | Optional query parameter. End date YYYY-MM-DD | 
 **seriesMode** | **string** | Optional query parameter. Optional series rendering mode. &#x60;canonical&#x60; applies configured remap rules (for continuity across methodology transitions). &#x60;raw&#x60; returns the direct indicator series without remapping. | [default to &quot;raw&quot;]
 **limit** | **int64** | Optional query parameter. Maximum number of rows to return in &#x60;data[]&#x60;. Rows are ordered most-recent-first, so &#x60;limit&#x3D;10&#x60; returns the 10 most recent observations within the requested window. Use together with &#x60;offset&#x60; to page through long histories. Defaults to 20. Maximum 100. | [default to 20]
 **offset** | **int64** | Optional query parameter. Number of rows to skip from the start of &#x60;data[]&#x60; (after most-recent-first ordering). Use with &#x60;limit&#x60; for pagination. Defaults to 0. | [default to 0]
 **page** | **int64** | Optional query parameter. One-based page number for callers that prefer page-based pagination. When supplied, &#x60;offset&#x60; is derived as &#x60;(page - 1) * limit&#x60;. | 
 **seasonality** | **string** | Optional query parameter. Optional seasonal adjustment selector. Supported public values are &#x60;sa&#x60; and &#x60;nsa&#x60;; check &#x60;supported_options.seasonality&#x60; in the response/catalogue for values available on this endpoint. Allowed values: sa, nsa. | 
 **frequency** | **string** | Optional query parameter. Optional returned change basis. Supported public values are &#x60;level&#x60;, &#x60;yoy&#x60;, &#x60;qoq&#x60;, and &#x60;mom&#x60;; check &#x60;supported_options.frequency&#x60; for values available on this endpoint. Allowed values: level, yoy, qoq, mom. | 
 **annualization** | **string** | Optional query parameter. Optional annualization selector. Supported public values are &#x60;not_annualized&#x60; and &#x60;saar&#x60;; check &#x60;supported_options.annualization&#x60; for values available on this endpoint. Allowed values: not_annualized, saar. | 
 **periodAggregation** | **string** | Optional query parameter. Optional period-aggregation selector. Supported public values are &#x60;monthly&#x60; and &#x60;quarterly_average&#x60;; check &#x60;supported_options.period_aggregation&#x60; for values available on this endpoint. Allowed values: monthly, quarterly_average. | 
 **revisions** | **string** | Optional query parameter. Optional revision view. Supported public values are &#x60;latest&#x60;, &#x60;first&#x60;, &#x60;final&#x60;, and &#x60;all&#x60;; unsupported views return &#x60;unsupported_series_option&#x60; with supported values. Allowed values: latest, first, final, all. | 
 **basis** | **string** | Optional query parameter. Optional real/nominal selector. Supported public values are &#x60;real&#x60; and &#x60;nominal&#x60;; only endpoints with known basis metadata accept it. Allowed values: real, nominal. | 
 **officialOnly** | **bool** | Optional query parameter. When true, return only official-source rows and remove explicit fallback observations. Prohibited private aggregator rows are always removed regardless of this setting. | [default to true]
 **asOf** | **string** | Optional query parameter. | 
 **availability** | **string** | Optional query parameter. | [default to &quot;public&quot;]
 **datasetVersion** | **string** | Optional query parameter. | 
 **valueMode** | **string** | Optional query parameter. | [default to &quot;normalized&quot;]

### Return type

[**AnnouncementResponse**](AnnouncementResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCalendar

> CalendarResponse GetCalendar(ctx, currency).Indicator(indicator).StartDate(startDate).EndDate(endDate).Timezone(timezone).Execute()

Get upcoming economic data release dates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "AUD" // string | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, COMM, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD.
	indicator := "afe_dollar_index" // string | Optional query parameter. Filter by indicator name (e.g., gdp, inflation, unemployment) Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages. (optional)
	startDate := "startDate_example" // string | Optional query parameter. Optional inclusive lower bound in YYYY-MM-DD format. (optional)
	endDate := "endDate_example" // string | Optional query parameter. Optional inclusive upper bound in YYYY-MM-DD format. (optional)
	timezone := "timezone_example" // string | Optional query parameter. Optional IANA timezone for an additional converted timestamp field, for example America/Sao_Paulo. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetCalendar(context.Background(), currency).Indicator(indicator).StartDate(startDate).EndDate(endDate).Timezone(timezone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetCalendar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCalendar`: CalendarResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetCalendar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, COMM, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCalendarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **indicator** | **string** | Optional query parameter. Filter by indicator name (e.g., gdp, inflation, unemployment) Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages. | 
 **startDate** | **string** | Optional query parameter. Optional inclusive lower bound in YYYY-MM-DD format. | 
 **endDate** | **string** | Optional query parameter. Optional inclusive upper bound in YYYY-MM-DD format. | 
 **timezone** | **string** | Optional query parameter. Optional IANA timezone for an additional converted timestamp field, for example America/Sao_Paulo. | 

### Return type

[**CalendarResponse**](CalendarResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommodities

> CommodityResponse GetCommodities(ctx, indicator).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Page(page).Execute()

Get commodity indicator data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	indicator := "crude_oil_inventories" // string | Required path parameter. Commodity indicator slug Allowed values: crude_oil_inventories, gold, natural_gas, natural_gas_storage, oil_brent, oil_wti, platinum, silver.
	startDate := "startDate_example" // string | Optional query parameter. Start date YYYY-MM-DD (optional)
	endDate := "endDate_example" // string | Optional query parameter. End date YYYY-MM-DD (optional)
	limit := int64(789) // int64 | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. (optional) (default to 20)
	offset := int64(789) // int64 | Optional query parameter. Number of rows to skip after most-recent-first ordering. (optional) (default to 0)
	page := int64(789) // int64 | Optional query parameter. One-based page number. When supplied, it takes precedence over offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetCommodities(context.Background(), indicator).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetCommodities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommodities`: CommodityResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetCommodities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**indicator** | **string** | Required path parameter. Commodity indicator slug Allowed values: crude_oil_inventories, gold, natural_gas, natural_gas_storage, oil_brent, oil_wti, platinum, silver. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommoditiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Optional query parameter. Start date YYYY-MM-DD | 
 **endDate** | **string** | Optional query parameter. End date YYYY-MM-DD | 
 **limit** | **int64** | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. | [default to 20]
 **offset** | **int64** | Optional query parameter. Number of rows to skip after most-recent-first ordering. | [default to 0]
 **page** | **int64** | Optional query parameter. One-based page number. When supplied, it takes precedence over offset. | 

### Return type

[**CommodityResponse**](CommodityResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCot

> COTResponse GetCot(ctx, currency).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Page(page).Execute()

Get CFTC Commitment of Traders (COT) data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "AUD" // string | Required path parameter. Allowed values: AUD, CAD, CHF, EUR, GBP, JPY, MXN, NZD, USD, XAU.
	startDate := "startDate_example" // string | Optional query parameter. Start date YYYY-MM-DD (optional)
	endDate := "endDate_example" // string | Optional query parameter. End date YYYY-MM-DD (optional)
	limit := int64(789) // int64 | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. (optional) (default to 20)
	offset := int64(789) // int64 | Optional query parameter. Number of rows to skip after most-recent-first ordering. (optional) (default to 0)
	page := int64(789) // int64 | Optional query parameter. One-based page number. When supplied, it takes precedence over offset. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetCot(context.Background(), currency).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetCot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCot`: COTResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetCot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. Allowed values: AUD, CAD, CHF, EUR, GBP, JPY, MXN, NZD, USD, XAU. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | Optional query parameter. Start date YYYY-MM-DD | 
 **endDate** | **string** | Optional query parameter. End date YYYY-MM-DD | 
 **limit** | **int64** | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. | [default to 20]
 **offset** | **int64** | Optional query parameter. Number of rows to skip after most-recent-first ordering. | [default to 0]
 **page** | **int64** | Optional query parameter. One-based page number. When supplied, it takes precedence over offset. | 

### Return type

[**COTResponse**](COTResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurves

> CurveAnalyticsResponse GetCurves(ctx, currency).CurveFamily(curveFamily).Metric(metric).View(view).Method(method).Date(date).AsOf(asOf).Availability(availability).MaxStalenessDays(maxStalenessDays).RequireCommonDate(requireCommonDate).Execute()

Get curve analytics by currency



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "AUD" // string | Required path parameter. Currency code Allowed values: AUD, CAD, CHF, CNH, CNY, DKK, EUR, GBP, ILS, JPY, KRW, MYR, NOK, NZD, PEN, SEK, THB, USD.
	curveFamily := "government_nominal" // string | Optional query parameter. Official curve family to return. Allowed values: government_nominal. (optional) (default to "government_nominal")
	metric := "spot" // string | Optional query parameter. Curve metric to return. Allowed values: spot. (optional) (default to "spot")
	view := "nodes" // string | Optional query parameter. Curve view to return: nodes, slopes, or forwards. Allowed values: nodes, slopes, forwards. (optional) (default to "nodes")
	method := "derived_from_spot_nodes" // string | Optional query parameter. Forward derivation method used when view=forwards. Allowed values: derived_from_spot_nodes. (optional) (default to "derived_from_spot_nodes")
	date := "date_example" // string | Optional query parameter. Optional target date YYYY-MM-DD. The endpoint returns the latest available node per maturity up to this date, using a recent lookback window. (optional)
	asOf := "asOf_example" // string | Optional query parameter. Explicit UTC information cutoff for verified vintage replay; date means start of day. (optional)
	availability := "availability_example" // string | Optional query parameter. Source-public or captured-value availability clock. (optional) (default to "public")
	maxStalenessDays := int64(789) // int64 | Optional query parameter. Maximum reference-date age; defaults to 7 with as_of, otherwise 60. (optional)
	requireCommonDate := true // bool | Optional query parameter. Require all populated curve nodes to share a reference date; defaults true with as_of. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetCurves(context.Background(), currency).CurveFamily(curveFamily).Metric(metric).View(view).Method(method).Date(date).AsOf(asOf).Availability(availability).MaxStalenessDays(maxStalenessDays).RequireCommonDate(requireCommonDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetCurves``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurves`: CurveAnalyticsResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetCurves`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. Currency code Allowed values: AUD, CAD, CHF, CNH, CNY, DKK, EUR, GBP, ILS, JPY, KRW, MYR, NOK, NZD, PEN, SEK, THB, USD. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurvesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **curveFamily** | **string** | Optional query parameter. Official curve family to return. Allowed values: government_nominal. | [default to &quot;government_nominal&quot;]
 **metric** | **string** | Optional query parameter. Curve metric to return. Allowed values: spot. | [default to &quot;spot&quot;]
 **view** | **string** | Optional query parameter. Curve view to return: nodes, slopes, or forwards. Allowed values: nodes, slopes, forwards. | [default to &quot;nodes&quot;]
 **method** | **string** | Optional query parameter. Forward derivation method used when view&#x3D;forwards. Allowed values: derived_from_spot_nodes. | [default to &quot;derived_from_spot_nodes&quot;]
 **date** | **string** | Optional query parameter. Optional target date YYYY-MM-DD. The endpoint returns the latest available node per maturity up to this date, using a recent lookback window. | 
 **asOf** | **string** | Optional query parameter. Explicit UTC information cutoff for verified vintage replay; date means start of day. | 
 **availability** | **string** | Optional query parameter. Source-public or captured-value availability clock. | [default to &quot;public&quot;]
 **maxStalenessDays** | **int64** | Optional query parameter. Maximum reference-date age; defaults to 7 with as_of, otherwise 60. | 
 **requireCommonDate** | **bool** | Optional query parameter. Require all populated curve nodes to share a reference date; defaults true with as_of. | 

### Return type

[**CurveAnalyticsResponse**](CurveAnalyticsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataCatalogue

> map[string]DataCatalogueItem GetDataCatalogue(ctx, currency).IncludeCapabilities(includeCapabilities).IncludeCoverage(includeCoverage).Indicator(indicator).Execute()

Get Currently Available Macroeconomic Indicators



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "AUD" // string | Required path parameter. The 3-letter currency code (e.g., USD, AUD, EUR, GBP, CHF) Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD.
	includeCapabilities := true // bool | Optional query parameter. Include currency-scoped endpoint/auth metadata that replaces the deprecated /v1/capabilities response. (optional) (default to false)
	includeCoverage := true // bool | Optional query parameter. Also include the legacy currency-scoped coverage list that replaces the deprecated /v1/coverage_matrix response. Per-indicator coverage is included in catalogue entries by default. (optional) (default to false)
	indicator := "indicator_example" // string | Optional query parameter. Optional indicator filter when include_coverage=true. Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetDataCatalogue(context.Background(), currency).IncludeCapabilities(includeCapabilities).IncludeCoverage(includeCoverage).Indicator(indicator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetDataCatalogue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataCatalogue`: map[string]DataCatalogueItem
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetDataCatalogue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. The 3-letter currency code (e.g., USD, AUD, EUR, GBP, CHF) Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataCatalogueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCapabilities** | **bool** | Optional query parameter. Include currency-scoped endpoint/auth metadata that replaces the deprecated /v1/capabilities response. | [default to false]
 **includeCoverage** | **bool** | Optional query parameter. Also include the legacy currency-scoped coverage list that replaces the deprecated /v1/coverage_matrix response. Per-indicator coverage is included in catalogue entries by default. | [default to false]
 **indicator** | **string** | Optional query parameter. Optional indicator filter when include_coverage&#x3D;true. Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages. | 

### Return type

[**map[string]DataCatalogueItem**](DataCatalogueItem.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFactor

> FactorResponse GetFactor(ctx, currency, factor).StartDate(startDate).EndDate(endDate).IncludeComponents(includeComponents).IncludeSources(includeSources).Limit(limit).Offset(offset).Execute()

Get a precomputed FX factor for one currency



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "AUD" // string | Required path parameter. Allowed values: AUD, CAD, CHF, EUR, GBP, JPY, NZD, USD.
	factor := "monetary_stance" // string | Required path parameter. Factor slug, for example monetary_stance. Allowed values: monetary_stance.
	startDate := "startDate_example" // string | Optional query parameter. Inclusive start date YYYY-MM-DD. With no dates, returns latest only. (optional)
	endDate := "endDate_example" // string | Optional query parameter. Inclusive end date YYYY-MM-DD. Defaults to today when start_date is supplied. (optional)
	includeComponents := true // bool | Optional query parameter. Include component scores and source observations. (optional) (default to false)
	includeSources := true // bool | Optional query parameter. Include FXMacroData source endpoint URLs. (optional) (default to false)
	limit := int64(789) // int64 | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. (optional) (default to 20)
	offset := int64(789) // int64 | Optional query parameter. Number of rows to skip after most-recent-first ordering. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetFactor(context.Background(), currency, factor).StartDate(startDate).EndDate(endDate).IncludeComponents(includeComponents).IncludeSources(includeSources).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetFactor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFactor`: FactorResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetFactor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. Allowed values: AUD, CAD, CHF, EUR, GBP, JPY, NZD, USD. | 
**factor** | **string** | Required path parameter. Factor slug, for example monetary_stance. Allowed values: monetary_stance. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFactorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **string** | Optional query parameter. Inclusive start date YYYY-MM-DD. With no dates, returns latest only. | 
 **endDate** | **string** | Optional query parameter. Inclusive end date YYYY-MM-DD. Defaults to today when start_date is supplied. | 
 **includeComponents** | **bool** | Optional query parameter. Include component scores and source observations. | [default to false]
 **includeSources** | **bool** | Optional query parameter. Include FXMacroData source endpoint URLs. | [default to false]
 **limit** | **int64** | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. | [default to 20]
 **offset** | **int64** | Optional query parameter. Number of rows to skip after most-recent-first ordering. | [default to 0]

### Return type

[**FactorResponse**](FactorResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFinancialPrices

> FinancialPriceResponse GetFinancialPrices(ctx, currency).Source(source).StartDate(startDate).EndDate(endDate).Measure(measure).InstrumentId(instrumentId).Issuer(issuer).Limit(limit).Offset(offset).Execute()

Official financial prices and yield curves



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "currency_example" // string | Required path parameter.
	source := "source_example" // string | Optional query parameter. (optional)
	startDate := time.Now() // string | Optional query parameter. (optional)
	endDate := time.Now() // string | Optional query parameter. (optional)
	measure := "measure_example" // string | Optional query parameter. (optional)
	instrumentId := "instrumentId_example" // string | Optional query parameter. (optional)
	issuer := "issuer_example" // string | Optional query parameter. (optional)
	limit := int64(789) // int64 | Optional query parameter. (optional) (default to 100)
	offset := int64(789) // int64 | Optional query parameter. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetFinancialPrices(context.Background(), currency).Source(source).StartDate(startDate).EndDate(endDate).Measure(measure).InstrumentId(instrumentId).Issuer(issuer).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetFinancialPrices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFinancialPrices`: FinancialPriceResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetFinancialPrices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFinancialPricesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **string** | Optional query parameter. | 
 **startDate** | **string** | Optional query parameter. | 
 **endDate** | **string** | Optional query parameter. | 
 **measure** | **string** | Optional query parameter. | 
 **instrumentId** | **string** | Optional query parameter. | 
 **issuer** | **string** | Optional query parameter. | 
 **limit** | **int64** | Optional query parameter. | [default to 100]
 **offset** | **int64** | Optional query parameter. | [default to 0]

### Return type

[**FinancialPriceResponse**](FinancialPriceResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForex

> ForexResponse GetForex(ctx, base, quote).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Page(page).Indicators(indicators).DatasetVersion(datasetVersion).Execute()

Get FX spot rates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	base := "AUD" // string | Required path parameter. Base currency code Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD.
	quote := "AUD" // string | Required path parameter. Quote currency code Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD.
	startDate := "startDate_example" // string | Optional query parameter. Start date YYYY-MM-DD (optional)
	endDate := "endDate_example" // string | Optional query parameter. End date YYYY-MM-DD (optional)
	limit := int64(789) // int64 | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. (optional) (default to 20)
	offset := int64(789) // int64 | Optional query parameter. Number of rows to skip after most-recent-first ordering. (optional) (default to 0)
	page := int64(789) // int64 | Optional query parameter. One-based page number. When supplied, it takes precedence over offset. (optional)
	indicators := "indicators_example" // string | Optional query parameter. Comma-separated technical indicators to compute (e.g. sma_20,rsi_14,macd,bollinger_bands). Use 'all' to request every supported indicator. (optional)
	datasetVersion := "datasetVersion_example" // string | Optional query parameter. Pin the content version returned by the first page; changed history returns 409. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetForex(context.Background(), base, quote).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Page(page).Indicators(indicators).DatasetVersion(datasetVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetForex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForex`: ForexResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetForex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**base** | **string** | Required path parameter. Base currency code Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD. | 
**quote** | **string** | Required path parameter. Quote currency code Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startDate** | **string** | Optional query parameter. Start date YYYY-MM-DD | 
 **endDate** | **string** | Optional query parameter. End date YYYY-MM-DD | 
 **limit** | **int64** | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. | [default to 20]
 **offset** | **int64** | Optional query parameter. Number of rows to skip after most-recent-first ordering. | [default to 0]
 **page** | **int64** | Optional query parameter. One-based page number. When supplied, it takes precedence over offset. | 
 **indicators** | **string** | Optional query parameter. Comma-separated technical indicators to compute (e.g. sma_20,rsi_14,macd,bollinger_bands). Use &#39;all&#39; to request every supported indicator. | 
 **datasetVersion** | **string** | Optional query parameter. Pin the content version returned by the first page; changed history returns 409. | 

### Return type

[**ForexResponse**](ForexResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFxIntradayReferenceRates

> FxIntradayReferenceRatesResponse GetFxIntradayReferenceRates(ctx, base, quote).StartTime(startTime).EndTime(endTime).Execute()

Get subscriber intraday FX reference rates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	base := "base_example" // string | Required path parameter. Base currency code
	quote := "quote_example" // string | Required path parameter. Quote currency code
	startTime := time.Now() // time.Time | Optional query parameter. Inclusive RFC 3339 timestamp. Defaults to 24 hours before end_time. (optional)
	endTime := time.Now() // time.Time | Optional query parameter. Inclusive RFC 3339 source timestamp. Defaults to the current time. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetFxIntradayReferenceRates(context.Background(), base, quote).StartTime(startTime).EndTime(endTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetFxIntradayReferenceRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFxIntradayReferenceRates`: FxIntradayReferenceRatesResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetFxIntradayReferenceRates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**base** | **string** | Required path parameter. Base currency code | 
**quote** | **string** | Required path parameter. Quote currency code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFxIntradayReferenceRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **time.Time** | Optional query parameter. Inclusive RFC 3339 timestamp. Defaults to 24 hours before end_time. | 
 **endTime** | **time.Time** | Optional query parameter. Inclusive RFC 3339 source timestamp. Defaults to the current time. | 

### Return type

[**FxIntradayReferenceRatesResponse**](FxIntradayReferenceRatesResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFxReferenceRates

> FxReferenceResponse GetFxReferenceRates(ctx, base, quote).Date(date).StartDate(startDate).EndDate(endDate).Source(source).TimestampType(timestampType).NativeOnly(nativeOnly).IncludeSourceLegs(includeSourceLegs).Order(order).Limit(limit).Offset(offset).Page(page).DatasetVersion(datasetVersion).Execute()

Get timestamped official FX reference rates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	base := "base_example" // string | Required path parameter. Base currency code
	quote := "quote_example" // string | Required path parameter. Quote currency code
	date := "date_example" // string | Optional query parameter. Single reference date YYYY-MM-DD. Overrides start_date/end_date. (optional)
	startDate := "startDate_example" // string | Optional query parameter. Start reference date YYYY-MM-DD. (optional)
	endDate := "endDate_example" // string | Optional query parameter. End reference date YYYY-MM-DD. (optional)
	source := "source_example" // string | Optional query parameter. Optional source id filter, for example ecb_eurofxref_daily. (optional)
	timestampType := "timestampType_example" // string | Optional query parameter. Optional timestamp type filter, for example fixing or published. (optional)
	nativeOnly := true // bool | Optional query parameter. Return only rows published directly for the requested pair, excluding inverse and cross-derived rows. (optional) (default to false)
	includeSourceLegs := true // bool | Optional query parameter. Include derivation source legs when a rate is inverted or crossed. (optional) (default to true)
	order := "order_example" // string | Optional query parameter. Result ordering. (optional) (default to "timestamp_asc")
	limit := int64(789) // int64 | Optional query parameter. Maximum rows to return. Maximum 100. (optional) (default to 100)
	offset := int64(789) // int64 | Optional query parameter. Number of rows to skip. (optional) (default to 0)
	page := int64(789) // int64 | Optional query parameter. One-based page number. (optional)
	datasetVersion := "datasetVersion_example" // string | Optional query parameter. Pin first-page content version; changed history returns 409. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetFxReferenceRates(context.Background(), base, quote).Date(date).StartDate(startDate).EndDate(endDate).Source(source).TimestampType(timestampType).NativeOnly(nativeOnly).IncludeSourceLegs(includeSourceLegs).Order(order).Limit(limit).Offset(offset).Page(page).DatasetVersion(datasetVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetFxReferenceRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFxReferenceRates`: FxReferenceResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetFxReferenceRates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**base** | **string** | Required path parameter. Base currency code | 
**quote** | **string** | Required path parameter. Quote currency code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFxReferenceRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **date** | **string** | Optional query parameter. Single reference date YYYY-MM-DD. Overrides start_date/end_date. | 
 **startDate** | **string** | Optional query parameter. Start reference date YYYY-MM-DD. | 
 **endDate** | **string** | Optional query parameter. End reference date YYYY-MM-DD. | 
 **source** | **string** | Optional query parameter. Optional source id filter, for example ecb_eurofxref_daily. | 
 **timestampType** | **string** | Optional query parameter. Optional timestamp type filter, for example fixing or published. | 
 **nativeOnly** | **bool** | Optional query parameter. Return only rows published directly for the requested pair, excluding inverse and cross-derived rows. | [default to false]
 **includeSourceLegs** | **bool** | Optional query parameter. Include derivation source legs when a rate is inverted or crossed. | [default to true]
 **order** | **string** | Optional query parameter. Result ordering. | [default to &quot;timestamp_asc&quot;]
 **limit** | **int64** | Optional query parameter. Maximum rows to return. Maximum 100. | [default to 100]
 **offset** | **int64** | Optional query parameter. Number of rows to skip. | [default to 0]
 **page** | **int64** | Optional query parameter. One-based page number. | 
 **datasetVersion** | **string** | Optional query parameter. Pin first-page content version; changed history returns 409. | 

### Return type

[**FxReferenceResponse**](FxReferenceResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFxSessionRates

> FxSessionRateResponse GetFxSessionRates(ctx, base, quote).Session(session).Date(date).StartDate(startDate).EndDate(endDate).Order(order).Limit(limit).Offset(offset).Page(page).DatasetVersion(datasetVersion).Execute()

Get official FX session-close reference rates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	base := "base_example" // string | Required path parameter. Base currency code
	quote := "quote_example" // string | Required path parameter. Quote currency code
	session := "session_example" // string | Required query parameter. Session name: sydney, tokyo, london, or new_york.
	date := "date_example" // string | Optional query parameter. Single reference date YYYY-MM-DD. Overrides start_date/end_date. (optional)
	startDate := "startDate_example" // string | Optional query parameter. Start reference date YYYY-MM-DD. (optional)
	endDate := "endDate_example" // string | Optional query parameter. End reference date YYYY-MM-DD. (optional)
	order := "order_example" // string | Optional query parameter. Result ordering. (optional) (default to "date_asc")
	limit := int64(789) // int64 | Optional query parameter. Maximum rows to return. Maximum 100. (optional) (default to 100)
	offset := int64(789) // int64 | Optional query parameter. Number of rows to skip. (optional) (default to 0)
	page := int64(789) // int64 | Optional query parameter. One-based page number. (optional)
	datasetVersion := "datasetVersion_example" // string | Optional query parameter. Pin first-page content version; changed history returns 409. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetFxSessionRates(context.Background(), base, quote).Session(session).Date(date).StartDate(startDate).EndDate(endDate).Order(order).Limit(limit).Offset(offset).Page(page).DatasetVersion(datasetVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetFxSessionRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFxSessionRates`: FxSessionRateResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetFxSessionRates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**base** | **string** | Required path parameter. Base currency code | 
**quote** | **string** | Required path parameter. Quote currency code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFxSessionRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **session** | **string** | Required query parameter. Session name: sydney, tokyo, london, or new_york. | 
 **date** | **string** | Optional query parameter. Single reference date YYYY-MM-DD. Overrides start_date/end_date. | 
 **startDate** | **string** | Optional query parameter. Start reference date YYYY-MM-DD. | 
 **endDate** | **string** | Optional query parameter. End reference date YYYY-MM-DD. | 
 **order** | **string** | Optional query parameter. Result ordering. | [default to &quot;date_asc&quot;]
 **limit** | **int64** | Optional query parameter. Maximum rows to return. Maximum 100. | [default to 100]
 **offset** | **int64** | Optional query parameter. Number of rows to skip. | [default to 0]
 **page** | **int64** | Optional query parameter. One-based page number. | 
 **datasetVersion** | **string** | Optional query parameter. Pin first-page content version; changed history returns 409. | 

### Return type

[**FxSessionRateResponse**](FxSessionRateResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFxSourceUniverse

> FxSourceUniverseResponse GetFxSourceUniverse(ctx).Currency(currency).Source(source).Execute()

List public FX reference-rate pair universe

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "currency_example" // string | Optional query parameter. Optional currency filter, for example USD. (optional)
	source := "source_example" // string | Optional query parameter. Optional source id filter, for example ecb_eurofxref_daily. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetFxSourceUniverse(context.Background()).Currency(currency).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetFxSourceUniverse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFxSourceUniverse`: FxSourceUniverseResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetFxSourceUniverse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFxSourceUniverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** | Optional query parameter. Optional currency filter, for example USD. | 
 **source** | **string** | Optional query parameter. Optional source id filter, for example ecb_eurofxref_daily. | 

### Return type

[**FxSourceUniverseResponse**](FxSourceUniverseResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFxSources

> FxSourcesResponse GetFxSources(ctx).Execute()

List public FX reference-rate sources



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetFxSources(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetFxSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFxSources`: FxSourcesResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetFxSources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFxSourcesRequest struct via the builder pattern


### Return type

[**FxSourcesResponse**](FxSourcesResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHealth

> ServiceStatusResponse GetHealth(ctx).Execute()

Ping the API



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHealth`: ServiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetHealth`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHealthRequest struct via the builder pattern


### Return type

[**ServiceStatusResponse**](ServiceStatusResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestAnnouncements

> LatestAnnouncementsResponse GetLatestAnnouncements(ctx, currency).Execute()

Get latest values for all indicators in a currency

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "AUD" // string | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetLatestAnnouncements(context.Background(), currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetLatestAnnouncements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestAnnouncements`: LatestAnnouncementsResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetLatestAnnouncements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestAnnouncementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LatestAnnouncementsResponse**](LatestAnnouncementsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestCommodities

> CommoditiesLatestResponse GetLatestCommodities(ctx).Execute()

Get latest values for all commodity indicators

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetLatestCommodities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetLatestCommodities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestCommodities`: CommoditiesLatestResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetLatestCommodities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestCommoditiesRequest struct via the builder pattern


### Return type

[**CommoditiesLatestResponse**](CommoditiesLatestResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketSessions

> MarketSessionsResponse GetMarketSessions(ctx).At(at).Execute()

FX market session timetable



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	at := "at_example" // string | Optional query parameter. Optional ISO-8601 UTC timestamp (e.g. 2026-03-05T12:00:00Z) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetMarketSessions(context.Background()).At(at).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetMarketSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketSessions`: MarketSessionsResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetMarketSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **at** | **string** | Optional query parameter. Optional ISO-8601 UTC timestamp (e.g. 2026-03-05T12:00:00Z) | 

### Return type

[**MarketSessionsResponse**](MarketSessionsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPredictionCoverage

> interface{} GetPredictionCoverage(ctx).Execute()

Which forecast publishers cover which currency (no API key)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetPredictionCoverage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetPredictionCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPredictionCoverage`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetPredictionCoverage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPredictionCoverageRequest struct via the builder pattern


### Return type

**interface{}**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPredictionCoverageForCurrency

> interface{} GetPredictionCoverageForCurrency(ctx, currency).Execute()

Per-indicator forecast coverage for one currency (no API key)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "currency_example" // string | Required path parameter.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetPredictionCoverageForCurrency(context.Background(), currency).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetPredictionCoverageForCurrency``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPredictionCoverageForCurrency`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetPredictionCoverageForCurrency`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPredictionCoverageForCurrencyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPredictions

> PredictionsResponse GetPredictions(ctx, currency, indicator).PredictionClass(predictionClass).PredictionType(predictionType).PredictionSource(predictionSource).OfficialLimit(officialLimit).PreReleaseOnly(preReleaseOnly).Seasonality(seasonality).Frequency(frequency).Annualization(annualization).PeriodAggregation(periodAggregation).Basis(basis).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Page(page).BeforeDate(beforeDate).AsOf(asOf).EventCompatibleOnly(eventCompatibleOnly).Execute()

List predictions for one indicator of a currency



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "AUD" // string | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD.
	indicator := "afe_dollar_index" // string | Required path parameter. Indicator slug Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages.
	predictionClass := "predictionClass_example" // string | Optional query parameter. Filter by how the forecast was produced. Prefer this over `prediction_type`: one class can span several legacy types, which is how `forecaster_survey` returns both the ECB and Philadelphia Fed Surveys of Professional Forecasters. (optional)
	predictionType := "predictionType_example" // string | Optional query parameter. Optional legacy prediction-type filter, retained unchanged for existing clients. Superseded by `prediction_class`. (optional)
	predictionSource := "predictionSource_example" // string | Optional query parameter. Optional source identifier filter. (optional)
	officialLimit := int64(789) // int64 | Optional query parameter. Maximum rows in ``official_forecasts`` (native-period official forecasts and surveys). 0 omits them. (optional) (default to 500)
	preReleaseOnly := true // bool | Optional query parameter. Return only forecasts with `generated_at` strictly before the linked `announcement_datetime` (default: true). FXMacroData only publishes forecasts generated before the announcement, so this flag is informational: setting it to false does not return additional rows. (optional) (default to true)
	seasonality := "seasonality_example" // string | Optional query parameter. Optional seasonal adjustment selector, matching /v1/announcements. Allowed values: sa, nsa. (optional)
	frequency := "frequency_example" // string | Optional query parameter. Optional returned change basis selector, matching /v1/announcements. Allowed values: level, yoy, qoq, mom. (optional)
	annualization := "annualization_example" // string | Optional query parameter. Optional annualization selector, matching /v1/announcements. Allowed values: not_annualized, saar. (optional)
	periodAggregation := "periodAggregation_example" // string | Optional query parameter. Optional period aggregation selector, matching /v1/announcements. Allowed values: monthly, quarterly_average. (optional)
	basis := "basis_example" // string | Optional query parameter. Optional real/nominal selector, matching /v1/announcements. Allowed values: real, nominal. (optional)
	startDate := "startDate_example" // string | Optional query parameter. Start date YYYY-MM-DD (period date). Defaults to 180 days before today when omitted. (optional)
	endDate := "endDate_example" // string | Optional query parameter. End date YYYY-MM-DD (period date). Defaults to 90 days ahead when omitted. (optional)
	limit := int64(789) // int64 | Optional query parameter. Maximum number of prediction groups to return in `data[]`. Use with `offset` to page through long histories. Defaults to 20. Maximum 100. (optional) (default to 20)
	offset := int64(789) // int64 | Optional query parameter. Number of prediction groups to skip from the start of `data[]`. Use with `limit` for pagination. Defaults to 0. (optional) (default to 0)
	page := int64(789) // int64 | Optional query parameter. One-based page number. When supplied, it takes precedence over offset. (optional)
	beforeDate := "beforeDate_example" // string | Optional query parameter. Return groups strictly before this YYYY-MM-DD period date. Use the previous response's `next_cursor` for bounded pagination. (optional)
	asOf := "asOf_example" // string | Optional query parameter. (optional)
	eventCompatibleOnly := true // bool | Optional query parameter. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetPredictions(context.Background(), currency, indicator).PredictionClass(predictionClass).PredictionType(predictionType).PredictionSource(predictionSource).OfficialLimit(officialLimit).PreReleaseOnly(preReleaseOnly).Seasonality(seasonality).Frequency(frequency).Annualization(annualization).PeriodAggregation(periodAggregation).Basis(basis).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Page(page).BeforeDate(beforeDate).AsOf(asOf).EventCompatibleOnly(eventCompatibleOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetPredictions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPredictions`: PredictionsResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetPredictions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD. | 
**indicator** | **string** | Required path parameter. Indicator slug Allowed values: afe_dollar_index, average_hourly_earnings, average_hourly_earnings_mom, balance_on_goods, balance_on_services, bgcr, breakeven_inflation_rate, broad_dollar_index, broad_money, building_approvals, building_permits, business_confidence, capital_account_balance, cb_assets, commodity_price_energy, commodity_price_ex_energy, commodity_price_index, commodity_prices, consumer_confidence, core_inflation, core_inflation_median, core_inflation_mom, core_inflation_trim, core_pce, core_pce_mom, credit_growth, crude_oil_inventories, current_account_balance, dairy_exports, deposit_rates, durable_goods_orders, effr, eme_dollar_index, employment, exports, financial_account_balance, foreign_reserves, full_time_employment, fx_reserves, gdp, gdp_growth_q4_yoy, gold_reserves, gov_bond_10y, gov_bond_15y, gov_bond_1_5m, gov_bond_1m, gov_bond_1y, gov_bond_20y, gov_bond_25y, gov_bond_2m, gov_bond_2y, gov_bond_30y, gov_bond_3m, gov_bond_3y, gov_bond_40y, gov_bond_4m, gov_bond_4y, gov_bond_5y, gov_bond_6m, gov_bond_6y, gov_bond_7y, gov_bond_8y, gov_bond_9y, government_debt, house_price_index, house_prices, household_credit, household_spending, housing_starts, imports, inflation, inflation_linked_bond, inflation_mom, initial_jobless_claims, international_assets, international_liabilities, job_openings, m1, m2, m3, nairu, natural_gas_storage, net_foreign_asset_position, non_farm_payrolls, non_farm_payrolls_change, obfr, part_time_employment, participation_rate, pce, pce_mom, policy_rate, policy_rate_midpoint, policy_rate_mlf, policy_rate_mro, policy_rate_target_lower, ppi, primary_income_balance, real_yield_20y, real_yield_30y, real_yield_5y, real_yield_7y, retail_sales, retail_sales_control_group, retail_sales_ex_autos, retail_sales_ex_autos_and_gas, risk_free_rate, secondary_income_balance, sight_deposits, snb_balance_sheet, sofr_average_180d, sofr_average_30d, sofr_average_90d, sofr_index, terms_of_trade, tgcr, trade_balance, trade_weighted_index, trimmed_mean_inflation, unemployment, wage_price_index, wages. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPredictionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **predictionClass** | **string** | Optional query parameter. Filter by how the forecast was produced. Prefer this over &#x60;prediction_type&#x60;: one class can span several legacy types, which is how &#x60;forecaster_survey&#x60; returns both the ECB and Philadelphia Fed Surveys of Professional Forecasters. | 
 **predictionType** | **string** | Optional query parameter. Optional legacy prediction-type filter, retained unchanged for existing clients. Superseded by &#x60;prediction_class&#x60;. | 
 **predictionSource** | **string** | Optional query parameter. Optional source identifier filter. | 
 **officialLimit** | **int64** | Optional query parameter. Maximum rows in &#x60;&#x60;official_forecasts&#x60;&#x60; (native-period official forecasts and surveys). 0 omits them. | [default to 500]
 **preReleaseOnly** | **bool** | Optional query parameter. Return only forecasts with &#x60;generated_at&#x60; strictly before the linked &#x60;announcement_datetime&#x60; (default: true). FXMacroData only publishes forecasts generated before the announcement, so this flag is informational: setting it to false does not return additional rows. | [default to true]
 **seasonality** | **string** | Optional query parameter. Optional seasonal adjustment selector, matching /v1/announcements. Allowed values: sa, nsa. | 
 **frequency** | **string** | Optional query parameter. Optional returned change basis selector, matching /v1/announcements. Allowed values: level, yoy, qoq, mom. | 
 **annualization** | **string** | Optional query parameter. Optional annualization selector, matching /v1/announcements. Allowed values: not_annualized, saar. | 
 **periodAggregation** | **string** | Optional query parameter. Optional period aggregation selector, matching /v1/announcements. Allowed values: monthly, quarterly_average. | 
 **basis** | **string** | Optional query parameter. Optional real/nominal selector, matching /v1/announcements. Allowed values: real, nominal. | 
 **startDate** | **string** | Optional query parameter. Start date YYYY-MM-DD (period date). Defaults to 180 days before today when omitted. | 
 **endDate** | **string** | Optional query parameter. End date YYYY-MM-DD (period date). Defaults to 90 days ahead when omitted. | 
 **limit** | **int64** | Optional query parameter. Maximum number of prediction groups to return in &#x60;data[]&#x60;. Use with &#x60;offset&#x60; to page through long histories. Defaults to 20. Maximum 100. | [default to 20]
 **offset** | **int64** | Optional query parameter. Number of prediction groups to skip from the start of &#x60;data[]&#x60;. Use with &#x60;limit&#x60; for pagination. Defaults to 0. | [default to 0]
 **page** | **int64** | Optional query parameter. One-based page number. When supplied, it takes precedence over offset. | 
 **beforeDate** | **string** | Optional query parameter. Return groups strictly before this YYYY-MM-DD period date. Use the previous response&#39;s &#x60;next_cursor&#x60; for bounded pagination. | 
 **asOf** | **string** | Optional query parameter. | 
 **eventCompatibleOnly** | **bool** | Optional query parameter. | [default to false]

### Return type

[**PredictionsResponse**](PredictionsResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPressReleases

> PressReleasesResponse GetPressReleases(ctx, currency).Limit(limit).Offset(offset).Execute()

Recent central bank press releases



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	currency := "AUD" // string | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CZK, DKK, EUR, GBP, HKD, ILS, INR, JPY, NGN, NOK, NZD, PEN, PLN, SEK, USD, ZAR.
	limit := int64(789) // int64 | Optional query parameter. Maximum number of press releases to return (optional) (default to 20)
	offset := int64(789) // int64 | Optional query parameter. Zero-based offset into the full press-release feed before pagination. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetPressReleases(context.Background(), currency).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetPressReleases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPressReleases`: PressReleasesResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetPressReleases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**currency** | **string** | Required path parameter. Allowed values: AUD, BRL, CAD, CHF, CZK, DKK, EUR, GBP, HKD, ILS, INR, JPY, NGN, NOK, NZD, PEN, PLN, SEK, USD, ZAR. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPressReleasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Optional query parameter. Maximum number of press releases to return | [default to 20]
 **offset** | **int64** | Optional query parameter. Zero-based offset into the full press-release feed before pagination. | [default to 0]

### Return type

[**PressReleasesResponse**](PressReleasesResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRateDifferentials

> RateDifferentialResponse GetRateDifferentials(ctx, base, quote).Measure(measure).RateType(rateType).CurveFamily(curveFamily).StartTenorYears(startTenorYears).EndTenorYears(endTenorYears).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).AsOf(asOf).Availability(availability).Execute()

Get pair rate differentials



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	base := "AUD" // string | Required path parameter. Base currency code Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD.
	quote := "AUD" // string | Required path parameter. Quote currency code Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD.
	measure := "auto" // string | Optional query parameter. Shared rate measure to compare across the pair. Allowed values: auto, risk_free_rate, gov_bond_2y, gov_bond_10y, policy_rate. (optional) (default to "auto")
	rateType := "spot" // string | Optional query parameter. Differential type: spot for shared rate measures, forward for derived curve segments. Allowed values: spot, forward. (optional) (default to "spot")
	curveFamily := "government_nominal" // string | Optional query parameter. Curve family used when rate_type=forward. Allowed values: government_nominal. (optional) (default to "government_nominal")
	startTenorYears := float64(1.2) // float64 | Optional query parameter. Start tenor in years when rate_type=forward. (optional) (default to 2.0)
	endTenorYears := float64(1.2) // float64 | Optional query parameter. End tenor in years when rate_type=forward. (optional) (default to 5.0)
	startDate := "startDate_example" // string | Optional query parameter. Inclusive start date YYYY-MM-DD. Defaults to 365 days ago. (optional)
	endDate := "endDate_example" // string | Optional query parameter. Inclusive end date YYYY-MM-DD. Defaults to today. (optional)
	limit := int64(789) // int64 | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. (optional) (default to 20)
	offset := int64(789) // int64 | Optional query parameter. Number of rows to skip after most-recent-first ordering. (optional) (default to 0)
	asOf := "asOf_example" // string | Optional query parameter. Explicit UTC information cutoff for verified vintage replay; date means start of day. (optional)
	availability := "availability_example" // string | Optional query parameter. Source-public or captured-value availability clock. (optional) (default to "public")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetRateDifferentials(context.Background(), base, quote).Measure(measure).RateType(rateType).CurveFamily(curveFamily).StartTenorYears(startTenorYears).EndTenorYears(endTenorYears).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).AsOf(asOf).Availability(availability).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetRateDifferentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRateDifferentials`: RateDifferentialResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetRateDifferentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**base** | **string** | Required path parameter. Base currency code Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD. | 
**quote** | **string** | Required path parameter. Quote currency code Allowed values: AUD, BRL, CAD, CHF, CNH, CNY, DKK, EUR, GBP, HUF, ILS, JPY, KRW, MYR, NGN, NOK, NZD, PEN, SEK, THB, TWD, USD. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRateDifferentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **measure** | **string** | Optional query parameter. Shared rate measure to compare across the pair. Allowed values: auto, risk_free_rate, gov_bond_2y, gov_bond_10y, policy_rate. | [default to &quot;auto&quot;]
 **rateType** | **string** | Optional query parameter. Differential type: spot for shared rate measures, forward for derived curve segments. Allowed values: spot, forward. | [default to &quot;spot&quot;]
 **curveFamily** | **string** | Optional query parameter. Curve family used when rate_type&#x3D;forward. Allowed values: government_nominal. | [default to &quot;government_nominal&quot;]
 **startTenorYears** | **float64** | Optional query parameter. Start tenor in years when rate_type&#x3D;forward. | [default to 2.0]
 **endTenorYears** | **float64** | Optional query parameter. End tenor in years when rate_type&#x3D;forward. | [default to 5.0]
 **startDate** | **string** | Optional query parameter. Inclusive start date YYYY-MM-DD. Defaults to 365 days ago. | 
 **endDate** | **string** | Optional query parameter. Inclusive end date YYYY-MM-DD. Defaults to today. | 
 **limit** | **int64** | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. | [default to 20]
 **offset** | **int64** | Optional query parameter. Number of rows to skip after most-recent-first ordering. | [default to 0]
 **asOf** | **string** | Optional query parameter. Explicit UTC information cutoff for verified vintage replay; date means start of day. | 
 **availability** | **string** | Optional query parameter. Source-public or captured-value availability clock. | [default to &quot;public&quot;]

### Return type

[**RateDifferentialResponse**](RateDifferentialResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskSentiment

> RiskSentimentResponse GetRiskSentiment(ctx).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()

Global risk-on / risk-off indicator



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	startDate := "startDate_example" // string | Optional query parameter. Start date YYYY-MM-DD (optional)
	endDate := "endDate_example" // string | Optional query parameter. End date YYYY-MM-DD (optional)
	limit := int64(789) // int64 | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. (optional) (default to 20)
	offset := int64(789) // int64 | Optional query parameter. Number of rows to skip after most-recent-first ordering. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.GetRiskSentiment(context.Background()).StartDate(startDate).EndDate(endDate).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.GetRiskSentiment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRiskSentiment`: RiskSentimentResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.GetRiskSentiment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskSentimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** | Optional query parameter. Start date YYYY-MM-DD | 
 **endDate** | **string** | Optional query parameter. End date YYYY-MM-DD | 
 **limit** | **int64** | Optional query parameter. Maximum rows to return. Defaults to 20. Maximum 100. | [default to 20]
 **offset** | **int64** | Optional query parameter. Number of rows to skip after most-recent-first ordering. | [default to 0]

### Return type

[**RiskSentimentResponse**](RiskSentimentResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> ServiceStatusResponse Ping(ctx).Execute()

Ping the API



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.Ping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Ping`: ServiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

[**ServiceStatusResponse**](ServiceStatusResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResearchPanel

> ResearchPanelResponse ResearchPanel(ctx).ResearchPanelRequest(researchPanelRequest).Execute()

Build a panel of values known at each decision time

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	researchPanelRequest := *openapiclient.NewResearchPanelRequest([]openapiclient.ResearchSeries{*openapiclient.NewResearchSeries("Currency_example", "Indicator_example")}, []string{"DecisionTimes_example"}, time.Now(), time.Now()) // ResearchPanelRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FXMacroDataAPI.ResearchPanel(context.Background()).ResearchPanelRequest(researchPanelRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FXMacroDataAPI.ResearchPanel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResearchPanel`: ResearchPanelResponse
	fmt.Fprintf(os.Stdout, "Response from `FXMacroDataAPI.ResearchPanel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResearchPanelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **researchPanelRequest** | [**ResearchPanelRequest**](ResearchPanelRequest.md) |  | 

### Return type

[**ResearchPanelResponse**](ResearchPanelResponse.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

