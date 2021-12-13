package mapper

import (
	"math/big"
	"strconv"

	"github.com/coinbase/rosetta-sdk-go/types"
)

func Amount(value *big.Int, currency *types.Currency) *types.Amount {
	if value == nil {
		return nil
	}
	return &types.Amount{
		Value:    value.String(),
		Currency: currency,
	}
}

func FeeAmount(value int64) *types.Amount {
	return &types.Amount{
		Value:    strconv.FormatInt(value, 10), //nolint:gomnd
		Currency: ActiveCurrency,
	}
}

func ActiveAmount(value *big.Int) *types.Amount {
	return Amount(value, ActiveCurrency)
}
