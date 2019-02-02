package empleoPonferradaOrg

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"jobdigger/digger"
	"jobdigger/offer"
	"jobdigger/rss"
	"net/http"
)

type Digger struct {
	digger.Digger

	TargetUri string

	errors []string
}

func New(targetUri string) *Digger {
	return &Digger{
		TargetUri: targetUri,
		errors:    []string{},
	}
}

func (w *Digger) GetErrors() []string {
	return w.errors
}

func (w *Digger) addError(message string) {
	w.errors = append(w.errors, message)
}

func (w *Digger) Parse(in *bytes.Reader) []offer.Offer {
	result := rss.New()

	content, err := ioutil.ReadAll(in)

	if err != nil {
		w.addError(err.Error())
		return nil
	}

	parserErr := xml.Unmarshal(content, &result)

	if parserErr != nil {
		w.addError(parserErr.Error())
		return nil
	}

	var offers []offer.Offer

	for _, item := range result.Channel.Items {
		offers = append(offers, offer.Offer{
			Title: item.Title,
		})
	}

	return offers
}

func (w *Digger) fetch() *bytes.Reader {
	resp, err := http.Get(w.TargetUri)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("error on GET %s", w.TargetUri)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err == nil {
		fmt.Println(err)
	}

	reader := bytes.NewReader(body)
	return reader
}

func (w *Digger) FetchAll() []offer.Offer {
	content := w.fetch()

	if content == nil {
		fmt.Println("no content")
	}

	offers := w.Parse(content)

	return offers
}

func (w *Digger) FetchNew() []*offer.Offer {
	var result []*offer.Offer

	return result
}
