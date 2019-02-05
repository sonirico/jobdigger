package empleoPonferradaOrg

import (
	"encoding/xml"
	"io/ioutil"
	"jobdigger/digger"
	"jobdigger/fetcher"
	"jobdigger/offer"
	"jobdigger/rss"
	"time"
)

type Digger struct {
	digger.Digger

	Fetcher *fetcher.BaseFetcher

	errors []string
}

func New(targetUri string) *Digger {
	return &Digger{
		Fetcher: fetcher.New(targetUri),
		errors:  []string{},
	}
}

func (d *Digger) parsePubDate(rawDate rss.PubDate) *time.Time {
	parsed, err := time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", string(rawDate))

	if err != nil {
		d.addError(err.Error())
		return nil
	}
	return &parsed
}

func (d *Digger) addError(message string) {
	d.errors = append(d.errors, message)
}

func (d *Digger) GetErrors() []string {
	return d.errors
}

func (d *Digger) Parse(payload []byte) []offer.Offer {
	result := rss.New()
	parserErr := xml.Unmarshal(payload, &result)

	if parserErr != nil {
		d.addError(parserErr.Error())
		return nil
	}

	var offers []offer.Offer

	for _, item := range result.Channel.Items {
		offers = append(offers, offer.Offer{
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			PubDate:     d.parsePubDate(item.PubDate),
		})
	}

	return offers
}

func (d *Digger) Get() []offer.Offer {
	reader := d.Fetcher.Fetch()
	content, err := ioutil.ReadAll(reader)

	if err != nil {
		d.addError(err.Error())
		return nil
	}

	offers := d.Parse(content)

	if len(d.GetErrors()) > 0 {
		return nil
	}

	return offers
}
