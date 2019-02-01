package digger

import (
	"bytes"
	"jobdigger/offer"
)

type Digger interface {
	GetErrors() []string

	Parse(in *bytes.Reader) []offer.Offer

	FetchNew() []*offer.Offer

	FetchAll() []*offer.Offer
}
