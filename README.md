# FXMacroData SDK for Go

Client for the [FXMacroData API](https://fxmacrodata.com/documentation/reference?utm_source=github&utm_medium=referral&utm_campaign=fxmacrodata-go&utm_content=readme): macroeconomic announcements with release timestamps, release calendars, FX reference rates, rate differentials, yield curves, COT positioning and commodities, sourced from central banks and statistics offices across 22 currencies.

USD endpoints work without an API key. Other currencies need a [subscription](https://fxmacrodata.com/subscribe?utm_source=github&utm_medium=referral&utm_campaign=fxmacrodata-go&utm_content=subscribe).

## Install

```bash
go get github.com/fxmacrodata/fxmacrodata-go
```

## Quickstart

```go
// Quickstart: latest US policy rate, release calendar and catalogue.
// USD works without a key; set FXMACRODATA_API_KEY for other currencies.
package main

import (
	"context"
	"fmt"
	"os"

	fxmacrodata "github.com/fxmacrodata/fxmacrodata-go"
)

func main() {
	client := fxmacrodata.NewAPIClient(fxmacrodata.NewConfiguration())
	ctx := context.Background()
	if key := os.Getenv("FXMACRODATA_API_KEY"); key != "" {
		ctx = context.WithValue(ctx, fxmacrodata.ContextAPIKeys, map[string]fxmacrodata.APIKey{
			"APIKeyHeader": {Key: key},
		})
	}

	rates, _, err := client.FXMacroDataAPI.GetAnnouncements(ctx, "USD", "policy_rate").Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "announcements:", err)
		os.Exit(1)
	}
	if len(rates.Data) == 0 {
		fmt.Fprintln(os.Stderr, "announcements: no observations returned")
		os.Exit(1)
	}
	latest := rates.Data[0]
	fmt.Printf("USD policy rate: %.2f on %s\n", latest.GetVal(), latest.Date)

	if _, _, err := client.FXMacroDataAPI.GetCalendar(ctx, "USD").Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "calendar:", err)
		os.Exit(1)
	}

	catalogue, _, err := client.FXMacroDataAPI.GetDataCatalogue(ctx, "USD").Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "catalogue:", err)
		os.Exit(1)
	}
	fmt.Printf("USD catalogue: %d indicators\n", len(*catalogue))
}
```

The same file is in this repository at `examples/quickstart/main.go`.

## Authentication

The key is sent in the `X-API-Key` header and never in the URL. The quickstart reads it from the `FXMACRODATA_API_KEY` environment variable.

## API reference

All URIs are relative to *https://api.fxmacrodata.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FXMacroDataAPI* | [**GetAnnouncementChanges**](docs/FXMacroDataAPI.md#getannouncementchanges) | **Get** /v1/announcements/changes | Poll recent live announcement release changes
*FXMacroDataAPI* | [**GetAnnouncements**](docs/FXMacroDataAPI.md#getannouncements) | **Get** /v1/announcements/{currency}/{indicator} | Get macroeconomic indicator data
*FXMacroDataAPI* | [**GetCalendar**](docs/FXMacroDataAPI.md#getcalendar) | **Get** /v1/calendar/{currency} | Get upcoming economic data release dates
*FXMacroDataAPI* | [**GetCommodities**](docs/FXMacroDataAPI.md#getcommodities) | **Get** /v1/commodities/{indicator} | Get commodity indicator data
*FXMacroDataAPI* | [**GetCot**](docs/FXMacroDataAPI.md#getcot) | **Get** /v1/cot/{currency} | Get CFTC Commitment of Traders (COT) data
*FXMacroDataAPI* | [**GetCurves**](docs/FXMacroDataAPI.md#getcurves) | **Get** /v1/curves/{currency} | Get curve analytics by currency
*FXMacroDataAPI* | [**GetDataCatalogue**](docs/FXMacroDataAPI.md#getdatacatalogue) | **Get** /v1/data_catalogue/{currency} | Get Currently Available Macroeconomic Indicators
*FXMacroDataAPI* | [**GetFactor**](docs/FXMacroDataAPI.md#getfactor) | **Get** /v1/factors/{currency}/{factor} | Get a precomputed FX factor for one currency
*FXMacroDataAPI* | [**GetFinancialPrices**](docs/FXMacroDataAPI.md#getfinancialprices) | **Get** /v1/financial_prices/{currency} | Official financial prices and yield curves
*FXMacroDataAPI* | [**GetForex**](docs/FXMacroDataAPI.md#getforex) | **Get** /v1/forex/{base}/{quote} | Get FX spot rates
*FXMacroDataAPI* | [**GetFxIntradayReferenceRates**](docs/FXMacroDataAPI.md#getfxintradayreferencerates) | **Get** /v1/fx/intraday-reference-rates/{base}/{quote} | Get subscriber intraday FX reference rates
*FXMacroDataAPI* | [**GetFxReferenceRates**](docs/FXMacroDataAPI.md#getfxreferencerates) | **Get** /v1/fx/reference-rates/{base}/{quote} | Get timestamped official FX reference rates
*FXMacroDataAPI* | [**GetFxSessionRates**](docs/FXMacroDataAPI.md#getfxsessionrates) | **Get** /v1/fx/session-rates/{base}/{quote} | Get official FX session-close reference rates
*FXMacroDataAPI* | [**GetFxSourceUniverse**](docs/FXMacroDataAPI.md#getfxsourceuniverse) | **Get** /v1/fx/source-universe | List public FX reference-rate pair universe
*FXMacroDataAPI* | [**GetFxSources**](docs/FXMacroDataAPI.md#getfxsources) | **Get** /v1/fx/sources | List public FX reference-rate sources
*FXMacroDataAPI* | [**GetHealth**](docs/FXMacroDataAPI.md#gethealth) | **Get** /v1/health | Ping the API
*FXMacroDataAPI* | [**GetLatestAnnouncements**](docs/FXMacroDataAPI.md#getlatestannouncements) | **Get** /v1/announcements/{currency}/latest | Get latest values for all indicators in a currency
*FXMacroDataAPI* | [**GetLatestCommodities**](docs/FXMacroDataAPI.md#getlatestcommodities) | **Get** /v1/commodities/latest | Get latest values for all commodity indicators
*FXMacroDataAPI* | [**GetMarketSessions**](docs/FXMacroDataAPI.md#getmarketsessions) | **Get** /v1/market_sessions | FX market session timetable
*FXMacroDataAPI* | [**GetPredictionCoverage**](docs/FXMacroDataAPI.md#getpredictioncoverage) | **Get** /v1/predictions/coverage | Which forecast publishers cover which currency (no API key)
*FXMacroDataAPI* | [**GetPredictionCoverageForCurrency**](docs/FXMacroDataAPI.md#getpredictioncoverageforcurrency) | **Get** /v1/predictions/coverage/{currency} | Per-indicator forecast coverage for one currency (no API key)
*FXMacroDataAPI* | [**GetPredictions**](docs/FXMacroDataAPI.md#getpredictions) | **Get** /v1/predictions/{currency}/{indicator} | List predictions for one indicator of a currency
*FXMacroDataAPI* | [**GetPressReleases**](docs/FXMacroDataAPI.md#getpressreleases) | **Get** /v1/press-releases/{currency} | Recent central bank press releases
*FXMacroDataAPI* | [**GetRateDifferentials**](docs/FXMacroDataAPI.md#getratedifferentials) | **Get** /v1/rate_differentials/{base}/{quote} | Get pair rate differentials
*FXMacroDataAPI* | [**GetRiskSentiment**](docs/FXMacroDataAPI.md#getrisksentiment) | **Get** /v1/risk_sentiment | Global risk-on / risk-off indicator
*FXMacroDataAPI* | [**Ping**](docs/FXMacroDataAPI.md#ping) | **Get** /v1/ping | Ping the API
*FXMacroDataAPI* | [**ResearchPanel**](docs/FXMacroDataAPI.md#researchpanel) | **Post** /v1/research/panel | Build a panel of values known at each decision time


## Documentation For Models

 - [AnnouncementChangeEvent](docs/AnnouncementChangeEvent.md)
 - [AnnouncementChangesResponse](docs/AnnouncementChangesResponse.md)
 - [AnnouncementChangesScope](docs/AnnouncementChangesScope.md)
 - [AnnouncementDataPoint](docs/AnnouncementDataPoint.md)
 - [AnnouncementPredictions](docs/AnnouncementPredictions.md)
 - [AnnouncementProvenance](docs/AnnouncementProvenance.md)
 - [AnnouncementRemap](docs/AnnouncementRemap.md)
 - [AnnouncementResponse](docs/AnnouncementResponse.md)
 - [CBTargetEntry](docs/CBTargetEntry.md)
 - [CBTargetInfo](docs/CBTargetInfo.md)
 - [COTDataPoint](docs/COTDataPoint.md)
 - [COTFXOverlay](docs/COTFXOverlay.md)
 - [COTPaginationInfo](docs/COTPaginationInfo.md)
 - [COTProvenance](docs/COTProvenance.md)
 - [COTResponse](docs/COTResponse.md)
 - [CalendarReleaseRow](docs/CalendarReleaseRow.md)
 - [CalendarResponse](docs/CalendarResponse.md)
 - [CatalogueCoverage](docs/CatalogueCoverage.md)
 - [CatalogueSeriesVariant](docs/CatalogueSeriesVariant.md)
 - [CatalogueSupportedOptions](docs/CatalogueSupportedOptions.md)
 - [CommoditiesLatestResponse](docs/CommoditiesLatestResponse.md)
 - [CommodityCoverage](docs/CommodityCoverage.md)
 - [CommodityCoverageSegment](docs/CommodityCoverageSegment.md)
 - [CommodityDataPoint](docs/CommodityDataPoint.md)
 - [CommodityLatestItem](docs/CommodityLatestItem.md)
 - [CommodityMissingWeekdayRange](docs/CommodityMissingWeekdayRange.md)
 - [CommodityObservation](docs/CommodityObservation.md)
 - [CommodityResponse](docs/CommodityResponse.md)
 - [CurveAnalyticsResponse](docs/CurveAnalyticsResponse.md)
 - [DataCatalogueItem](docs/DataCatalogueItem.md)
 - [DataQuality](docs/DataQuality.md)
 - [FactorDataPoint](docs/FactorDataPoint.md)
 - [FactorResponse](docs/FactorResponse.md)
 - [FinancialCoverage](docs/FinancialCoverage.md)
 - [FinancialObservation](docs/FinancialObservation.md)
 - [FinancialPagination](docs/FinancialPagination.md)
 - [FinancialPriceResponse](docs/FinancialPriceResponse.md)
 - [FinancialSource](docs/FinancialSource.md)
 - [ForexDataPoint](docs/ForexDataPoint.md)
 - [ForexPairMetadata](docs/ForexPairMetadata.md)
 - [ForexPointSource](docs/ForexPointSource.md)
 - [ForexProvenance](docs/ForexProvenance.md)
 - [ForexResponse](docs/ForexResponse.md)
 - [FreemiumDelayNotice](docs/FreemiumDelayNotice.md)
 - [FreemiumWindow](docs/FreemiumWindow.md)
 - [FxIntradayReferenceRatePoint](docs/FxIntradayReferenceRatePoint.md)
 - [FxIntradayReferenceRatesResponse](docs/FxIntradayReferenceRatesResponse.md)
 - [FxReferencePoint](docs/FxReferencePoint.md)
 - [FxReferenceResponse](docs/FxReferenceResponse.md)
 - [FxSessionRatePoint](docs/FxSessionRatePoint.md)
 - [FxSessionRateResponse](docs/FxSessionRateResponse.md)
 - [FxSource](docs/FxSource.md)
 - [FxSourceLeg](docs/FxSourceLeg.md)
 - [FxSourcePolicy](docs/FxSourcePolicy.md)
 - [FxSourceUniverseEntry](docs/FxSourceUniverseEntry.md)
 - [FxSourceUniversePair](docs/FxSourceUniversePair.md)
 - [FxSourceUniverseResponse](docs/FxSourceUniverseResponse.md)
 - [FxSourcesResponse](docs/FxSourcesResponse.md)
 - [HTTPValidationError](docs/HTTPValidationError.md)
 - [IndicatorPaginationInfo](docs/IndicatorPaginationInfo.md)
 - [LatestAnnouncementItem](docs/LatestAnnouncementItem.md)
 - [LatestAnnouncementProvenance](docs/LatestAnnouncementProvenance.md)
 - [LatestAnnouncementValue](docs/LatestAnnouncementValue.md)
 - [LatestAnnouncementsProvenance](docs/LatestAnnouncementsProvenance.md)
 - [LatestAnnouncementsResponse](docs/LatestAnnouncementsResponse.md)
 - [MarketSession](docs/MarketSession.md)
 - [MarketSessionOverlap](docs/MarketSessionOverlap.md)
 - [MarketSessionsResponse](docs/MarketSessionsResponse.md)
 - [OfficialForwardSourceSupport](docs/OfficialForwardSourceSupport.md)
 - [Pagination](docs/Pagination.md)
 - [PointInTimeCompleteness](docs/PointInTimeCompleteness.md)
 - [PolicyFamilyEntry](docs/PolicyFamilyEntry.md)
 - [PredictionItem](docs/PredictionItem.md)
 - [PredictionProvenance](docs/PredictionProvenance.md)
 - [PredictionsResponse](docs/PredictionsResponse.md)
 - [PressRelease](docs/PressRelease.md)
 - [PressReleaseRatePath](docs/PressReleaseRatePath.md)
 - [PressReleasesResponse](docs/PressReleasesResponse.md)
 - [RateDifferentialPoint](docs/RateDifferentialPoint.md)
 - [RateDifferentialResponse](docs/RateDifferentialResponse.md)
 - [RateDifferentialSources](docs/RateDifferentialSources.md)
 - [RatePathMatch](docs/RatePathMatch.md)
 - [ReleaseDeliveryAnnouncement](docs/ReleaseDeliveryAnnouncement.md)
 - [ReleaseTiming](docs/ReleaseTiming.md)
 - [ResearchPanelRequest](docs/ResearchPanelRequest.md)
 - [ResearchPanelResponse](docs/ResearchPanelResponse.md)
 - [ResearchSeries](docs/ResearchSeries.md)
 - [RevisionEntry](docs/RevisionEntry.md)
 - [RiskSentimentAliases](docs/RiskSentimentAliases.md)
 - [RiskSentimentComponentCoverage](docs/RiskSentimentComponentCoverage.md)
 - [RiskSentimentComponentMetadata](docs/RiskSentimentComponentMetadata.md)
 - [RiskSentimentPoint](docs/RiskSentimentPoint.md)
 - [RiskSentimentResponse](docs/RiskSentimentResponse.md)
 - [SelectedSeries](docs/SelectedSeries.md)
 - [SeriesFilters](docs/SeriesFilters.md)
 - [ServiceStatusResponse](docs/ServiceStatusResponse.md)
 - [SourceProvenance](docs/SourceProvenance.md)
 - [ValidationError](docs/ValidationError.md)

## Versioning

This client is generated from the FXMacroData OpenAPI specification, so it tracks the API one to one. Report problems at https://github.com/fxmacrodata/fxmacrodata-go/issues.

## Licence

MIT
