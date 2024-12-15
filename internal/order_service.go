package internal

import (
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"github.com/shopspring/decimal"
	"log"
)

func PlaceMarketOrder(client *alpaca.Client, symbol string, side alpaca.Side, quantity float64, intent alpaca.PositionIntent) {
	qty := decimal.NewFromFloat(quantity)
	if _, err := client.PlaceOrder(alpaca.PlaceOrderRequest{
		Symbol:         symbol,
		Qty:            &qty,
		Side:           side,
		Type:           alpaca.Market,
		TimeInForce:    alpaca.Day,
		PositionIntent: intent,
	}); err != nil {
		log.Fatalf("failed to place an order: %v", err)
	}
	log.Printf("%s order sent for %s with quantity of %f\n", side, symbol, quantity)
}
