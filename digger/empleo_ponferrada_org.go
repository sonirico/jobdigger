package digger

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"jobdigger/offer"
	"net/http"
)

type ItemNode struct {
	Title string `xml:"title"`
}

type ChannelNode struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Items       []ItemNode `xml:"item"`
}

type RSSNode struct {
	Channel ChannelNode `xml:"channel"`
}

type EmpleoPonferradaOrgDigger struct {
	Digger

	TargetUri string

	errors []string
}

func New(targetUri string) *EmpleoPonferradaOrgDigger {
	return &EmpleoPonferradaOrgDigger{
		TargetUri: targetUri,
		errors:    []string{},
	}
}

func (w *EmpleoPonferradaOrgDigger) GetErrors() []string {
	return w.errors
}

func (w *EmpleoPonferradaOrgDigger) addError(message string) {
	w.errors = append(w.errors, message)
}

func (w *EmpleoPonferradaOrgDigger) Parse(in *bytes.Reader) []offer.Offer {
	result := RSSNode{}

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

func (w *EmpleoPonferradaOrgDigger) fetch() *bytes.Reader {
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

func (w *EmpleoPonferradaOrgDigger) FetchAll() []offer.Offer {
	content := w.fetch()

	if content == nil {
		fmt.Println("no content")
	}

	offers := w.Parse(content)

	return offers
}

func (w *EmpleoPonferradaOrgDigger) FetchNew() []*offer.Offer {
	var result []*offer.Offer

	return result
}
