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
