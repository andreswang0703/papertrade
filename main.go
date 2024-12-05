package main

import (
	"fmt"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"marketdata/internal"
	"time"
)

func main() {

	request := marketdata.GetCryptoBarsRequest{
		TimeFrame: marketdata.OneDay,
		Start:     time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
		End:       time.Date(2024, 12, 5, 0, 0, 0, 0, time.UTC),
	}

	bars, err := internal.MarketDataClient().GetCryptoBars("BTC/USD", request)
	if err != nil {
		panic(err)
	}
	for _, bar := range bars {
		fmt.Printf("%+v\n", bar)
	}

	stockRequest := marketdata.GetBarsRequest{
		TimeFrame: marketdata.OneDay,
		Start:     time.Date(2024, 12, 1, 0, 0, 0, 0, time.UTC),
		End:       time.Date(2024, 12, 5, 0, 0, 0, 0, time.UTC),
	}
	stockBars, err := internal.MarketDataClient().GetBars("AAPL", stockRequest)
	if err != nil {
		panic(err)
	}
	for _, bar := range stockBars {
		fmt.Printf("%+v\n", bar)
	}
}
