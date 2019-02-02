package empleoPonferradaOrg

import (
	"bytes"
	"jobdigger/offer"
	"testing"
)

func checkParserErrors(t *testing.T, w *Digger) {
	if len(w.GetErrors()) > 0 {
		t.Errorf("parser has errors: %d", len(w.GetErrors()))
		for _, errorMessage := range w.GetErrors() {
			t.Errorf(errorMessage)
		}
		t.FailNow()
	}
}
func testOffer(t *testing.T, o offer.Offer, e offer.Offer) bool {
	if o.GetTitle() != e.GetTitle() {
		t.Fatalf("expected offer title to be '%s'. Got '%s'",
			e.GetTitle(), o.GetTitle())
		return false
	}
	return true
}

func TestDigger_Parse(t *testing.T) {
	bytesFeed := []byte(`
		<?xml version="1.0" encoding="utf-8"?>
		<rss version="2.0">
			<channel>
				<title>INSTITUTO MUNICIPAL PARA LA FORMACIÓN Y EL EMPLEO RSS</title>
				<link>http://empleo.ponferrada.org</link>
				<description>Ofertas, novedades y artículos de INSTITUTO MUNICIPAL PARA LA FORMACIÓN Y EL EMPLEO</description>
				<ttl>60</ttl>
				<item>
					<title>Fisioterapeuta u osteópata.</title>
					<link>http://empleo.ponferrada.org/ofertas/ver/d3ce2632-00fe-4f1e-b932-19b90740c369</link>
					<description>Fisioterapeuta u osteópata.</description>
					<pubDate>Wed, 30 Jan 2019 08:11:23 GMT</pubDate>
				</item>
				<item>
					<title>Encofrador (duplicada)</title>
					<link>http://empleo.ponferrada.org/ofertas/ver/f66d3a27-10b2-468b-a70d-afa591501270</link>
					<description>Encofrador oficial de 1ª</description>
					<pubDate>Wed, 30 Jan 2019 08:11:31 GMT</pubDate>
				</item>
			</channel>
		</rss>
	`)
	digger := New("https://empleo.ponferrada.org/rss")
	reader := bytes.NewReader(bytesFeed)
	offers := digger.Parse(reader)
	checkParserErrors(t, digger)

	if len(offers) != 2 {
		t.Fatalf("expected %d offers. got %d", 2, len(offers))
	}

	expected := []offer.Offer{
		{Title: "Fisioterapeuta u osteópata."},
		{Title: "Encofrador (duplicada)"},
	}

	for index, expectedOffer := range expected {
		testOffer(t, offers[index], expectedOffer)
	}
}
