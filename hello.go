package hello

import (
	"rsc.io/quote/v3"
)

// Hello returns a hello world message
func Hello() string {
	return quote.HelloV3()
}

// Proverb returns a Go proverb
func Proverb() string {
	return quote.Concurrency()
}
