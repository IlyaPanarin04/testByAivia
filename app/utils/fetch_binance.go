package utils

import (
	"context"
	"fmt"
	"github.com/aiviaio/go-binance/v2"
	"sync"
)

func GetSymbols(ctx context.Context, client *binance.Client) ([]string, error) {
	exchangeInfo, err := client.NewExchangeInfoService().Do(ctx)
	if err != nil {
		return nil, err
	}

	var symbols []string
	for i, symbol := range exchangeInfo.Symbols {
		symbols = append(symbols, symbol.Symbol)
		if i == 4 {
			break
		}
	}

	return symbols, nil
}

func GetPrice(ctx context.Context, client *binance.Client, symbol string, ch chan map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()

	ticker, err := client.NewListPricesService().Symbol(symbol).Do(ctx)
	if err != nil {
		fmt.Printf(`Error fetching price for %v: %v`, symbol, err)
		return
	}

	ch <- map[string]string{symbol: ticker[0].Price}
}
