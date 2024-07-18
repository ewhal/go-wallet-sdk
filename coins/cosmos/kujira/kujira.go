package kujira

import "github.com/ewhal/go-wallet-sdk/coins/cosmos"

const (
	HRP = "kujira"
)

func NewAddress(privateKeyHex string) (string, error) {
	return cosmos.NewAddress(privateKeyHex, HRP, false)
}

func ValidateAddress(address string) bool {
	return cosmos.ValidateAddress(address, HRP)
}
