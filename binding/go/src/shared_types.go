// Code generated - DO NOT EDIT.
// This file contains shared type definitions.

package binding

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type IPairConfig struct {
	QUOTE       common.Address
	BASE        common.Address
	DENOMINATOR *big.Int
}


type IPairOrder struct {
	Side   uint8
	Owner  common.Address
	FeeBps uint32
	Price  *big.Int
	Amount *big.Int
}


