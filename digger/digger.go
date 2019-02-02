package digger

import (
	"bytes"
	"jobdigger/fetcher"
	"jobdigger/offer"
)

type Digger interface {
	fetcher.Fetcher

	GetErrors() []string

	Parse(in *bytes.Reader) []offer.Offer
}
