package offer

import "time"

type Offer struct {
	Title       string
	Link        string
	Description string
	PubDate     *time.Time
}

func New() *Offer {
	return &Offer{}
}

func (o *Offer) Equals(other *Offer) bool {
	if o.Title != other.Title {
		return false
	}
	if o.Description != other.Description {
		return false
	}
	if o.Link != other.Link {
		return false
	}
	return true
}

func (o *Offer) PublishedAt() string {
	if o.PubDate == nil {
		return ""
	}
	return o.PubDate.Format("2006-01-02 15:04:05 GMT")
}
