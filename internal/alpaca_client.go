package internal

import (
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"log"
	"os"
	"sync"
)

var (
	marketDataClient     *marketdata.Client
	onceMarketDataClient sync.Once

	alpacaClient     *alpaca.Client
	onceAlpacaClient sync.Once
)

type APIKeySecret struct {
	APIKey    string
	APISecret string
}

func init() {
	if os.Getenv("ALPACA_API_KEY") == "" {
		log.Fatalln("please provide ALPACA_API_KEY as environment variable")
	}
	if os.Getenv("ALPACA_SECRET_KEY") == "" {
		log.Fatalln("please provide ALPACA_SECRET_KEY as environment variable")
	}
}

func MarketDataClient() *marketdata.Client {
	onceMarketDataClient.Do(
		func() {
			marketDataClient = marketdata.NewClient(marketdata.ClientOpts{
				APIKey:    os.Getenv("ALPACA_API_KEY"),
				APISecret: os.Getenv("ALPACA_SECRET_KEY"),
			})
		})
	return marketDataClient
}

func AlpacaClient() *alpaca.Client {
	onceAlpacaClient.Do(
		func() {
			alpacaClient = alpaca.NewClient(alpaca.ClientOpts{
				APIKey:    os.Getenv("ALPACA_API_KEY"),
				APISecret: os.Getenv("ALPACA_SECRET_KEY"),
				BaseURL:   "https://paper-api.alpaca.markets",
			})
		},
	)
	return alpacaClient
}
