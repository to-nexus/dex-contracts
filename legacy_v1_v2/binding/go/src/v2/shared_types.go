// Code generated - DO NOT EDIT.
// This file contains shared type definitions.

package binding

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IPairOrder struct {
	Side   uint8
	Owner  common.Address
	FeeBps uint32
	Price  *big.Int
	Amount *big.Int
}


type IPairConfig struct {
	QUOTE       common.Address
	BASE        common.Address
	DENOMINATOR *big.Int
}


type PendingDefaultAdminDelayOutput struct {
	NewDelay *big.Int
	Schedule *big.Int
}


type AccountReservesOutput struct {
	Base  *big.Int
	Quote *big.Int
}


type PendingDefaultAdminOutput struct {
	NewAdmin common.Address
	Schedule *big.Int
}


type AllPairsOutput struct {
	Bases []common.Address
	Pairs []common.Address
}


type TicksOutput struct {
	SellPrices []*big.Int
	BuyPrices  []*big.Int
}


type AllMarketsOutput struct {
	Quotes  []common.Address
	Markets []common.Address
}


type TickSizesOutput struct {
	Tick *big.Int
	Lot  *big.Int
}


