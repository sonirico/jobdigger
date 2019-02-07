package digger

import (
	"jobdigger/offer"
)

type Digger interface {
	GetErrors() []string

	GetOffers() (*Result, error)

	Parse(in []byte) []offer.Offer
}

type Result struct {
	Ok     int
	Failed int
	Offers []offer.Offer
}

func NewResult(offers []offer.Offer, ok int, failed int) *Result {
	return &Result{Offers: offers, Ok: ok, Failed: failed}
}

func (dr *Result) Total() int {
	return dr.Ok + dr.Failed
}
