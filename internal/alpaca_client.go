package internal

import (
	"github.com/alpacahq/alpaca-trade-api-go/v3/marketdata"
	"os"
	"sync"
)

var (
	marketDataClient *marketdata.Client
	once             sync.Once
)

type APIKeySecret struct {
	APIKey    string
	APISecret string
}

func MarketDataClient() *marketdata.Client {
	once.Do(
		func() {
			marketDataClient = marketdata.NewClient(marketdata.ClientOpts{
				APIKey:    os.Getenv("ALPACA_API_KEY"),
				APISecret: os.Getenv("ALPACA_SECRET_KEY"),
			})
		})
	return marketDataClient
}
