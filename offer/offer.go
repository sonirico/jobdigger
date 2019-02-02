package offer

import "time"

type Offer struct {
	Title       string
	link        string
	description string
	publishedAt time.Time
	updatedAt   *time.Time
}

func New() *Offer {
	return &Offer{}
}

func (o *Offer) GetTitle() string {
	return o.Title
}

func (o *Offer) Equals(other *Offer) bool {
	if o.Title != other.GetTitle() {
		return false
	}
	return true
}
