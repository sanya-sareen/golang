package main

import ("context"
		"fmt"
)

type PriceFetcher interface{
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string)(float64, error){
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 2000.00,
	"ETH": 200.9,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error){
	price, ok := priceMocks[ticker]

	if !ok{
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker )
	}

	return price, nil
}