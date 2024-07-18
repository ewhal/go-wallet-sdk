package evmos

import (
	"github.com/ewhal/go-wallet-sdk/coins/cosmos"
)

const (
	HRP = "evmos"
)

// NewAddress method of eth is used
func NewAddress(privateKey string) (string, error) {
	return cosmos.NewAddress(privateKey, HRP, true)
}

func ValidateAddress(address string) bool {
	return cosmos.ValidateAddress(address, HRP)
}
