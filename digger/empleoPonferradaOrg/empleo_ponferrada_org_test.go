package empleoPonferradaOrg

import (
	"jobdigger/offer"
	"testing"
)

type TestOffer struct {
	Title string
	Link string
	Description string
	PublishedAt string
}

func checkParserErrors(t *testing.T, w *Digger) {
	if len(w.GetErrors()) > 0 {
		t.Errorf("parser has errors: %d", len(w.GetErrors()))
		for _, errorMessage := range w.GetErrors() {
			t.Errorf(errorMessage)
		}
		t.FailNow()
	}
}

func testOffer(t *testing.T, o offer.Offer, e TestOffer) bool {
	if o.Title != e.Title {
		t.Fatalf("expected offer title to be '%s'. Got '%s'",
			e.Title, o.Title)
		return false
	}
	if o.Description != e.Description {
		t.Fatalf("expected offer description to be '%s'. Got '%s'",
			e.Description, o.Description)
		return false
	}
	if o.Link != e.Link {
		t.Fatalf("expected offer link to be '%s'. Got '%s'",
			e.Link, o.Link)
		return false
	}
	if o.PublishedAt() != e.PublishedAt {
		t.Fatalf("expected offer date to be '%s'. Got '%s'",
			e.PublishedAt, o.PublishedAt())
		return false
	}
	return true
}

func TestDiggerSeveralResults_Parse(t *testing.T) {
	payload := []byte(`
		<?xml version="1.0" encoding="utf-8"?>
		<rss version="2.0">
			<channel>
				<title>INSTITUTO MUNICIPAL PARA LA FORMACIÓN Y EL EMPLEO RSS</title>
				<link>http://empleo.ponferrada.org</link>
				<description>Ofertas, novedades y artículos de INSTITUTO MUNICIPAL PARA LA FORMACIÓN Y EL EMPLEO</description>
				<ttl>60</ttl>
				<item>
					<title>Fisioterapeuta u osteópata.</title>
					<link>http://empleo.ponferrada.org/ofertas/ver/d3ce2632</link>
					<description>Fisioterapeuta u osteópata.</description>
					<pubDate>Wed, 30 Jan 2019 08:11:23 GMT</pubDate>
				</item>
				<item>
					<title>Encofrador (duplicada)</title>
					<link>http://empleo.ponferrada.org/ofertas/ver/f66d3a27</link>
					<description>Encofrador oficial de 1ª</description>
					<pubDate>Wed, 30 Jan 2019 08:11:31 GMT</pubDate>
				</item>
			</channel>
		</rss>
	`)
	digger := New("https://empleo.ponferrada.org/rss")
	offers := digger.Parse(payload)
	checkParserErrors(t, digger)

	if len(offers) != 2 {
		t.Fatalf("expected %d offers. got %d", 2, len(offers))
	}

	expected := []TestOffer{
		{
			Title: "Fisioterapeuta u osteópata.",
			Link: "http://empleo.ponferrada.org/ofertas/ver/d3ce2632",
			Description: "Fisioterapeuta u osteópata.",
			PublishedAt: "2019-01-30 08:11:23 GMT",
		},
		{
			Title: "Encofrador (duplicada)",
			Link: "http://empleo.ponferrada.org/ofertas/ver/f66d3a27",
			Description: "Encofrador oficial de 1ª",
			PublishedAt: "2019-01-30 08:11:31 GMT",
		},
	}

	for index, expectedOffer := range expected {
		testOffer(t, offers[index], expectedOffer)
	}
}

func TestDiggerNoneResults_Parse(t *testing.T) {
	payload := []byte(`
		<?xml version="1.0" encoding="utf-8"?>
		<rss version="2.0">
			<channel>
				<title>INSTITUTO MUNICIPAL PARA LA FORMACIÓN Y EL EMPLEO RSS</title>
				<link>http://empleo.ponferrada.org</link>
				<description>Ofertas, novedades y artículos de INSTITUTO MUNICIPAL PARA LA FORMACIÓN Y EL EMPLEO</description>
				<ttl>60</ttl>
			</channel>
		</rss>
	`)
	digger := New("https://empleo.ponferrada.org/rss")
	offers := digger.Parse(payload)
	checkParserErrors(t, digger)

	if len(offers) != 0 {
		t.Fatalf("expected %d offers. got %d", 0, len(offers))
	}
}


func TestDiggerEmptyPayload_Parse(t *testing.T) {
	payload := []byte(``)
	digger := New("https://empleo.ponferrada.org/rss")
	digger.Parse(payload)

	if len(digger.GetErrors()) < 1 {
		t.Fatalf("expected digger to have errors.")
	}
}
