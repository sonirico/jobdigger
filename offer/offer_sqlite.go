package offer

import (
	"bytes"
	"text/template"
)

const (
	DB_TABLE = "offer"
)

type SQliteOffer struct {
	Table string
	Offer *Offer
}

func NewSQliteOffer (offer *Offer) *SQliteOffer {
	return &SQliteOffer{Offer: offer, Table: DB_TABLE}
}

func (so *SQliteOffer) Insert() string {
	tplStr := `
		INSERT INTO {{.Table}}
		(title, permalink, description, published_at)
		VALUES ("{{.Offer.Title}}", "{{.Offer.Link}}", "{{.Offer.Description}}", {{.Offer.PublishedAtUnix}})
	`
	tpl, err := template.New("test").Parse(tplStr)
	if err != nil {
		panic(err)
	}
	var buffer bytes.Buffer
	err = tpl.Execute(&buffer, so)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}