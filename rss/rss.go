package rss

type PubDate string

type ItemNode struct {
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	Description string  `xml:"description"`
	PubDate     PubDate `xml:"pubDate"`
}

type ChannelNode struct {
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Items       []ItemNode `xml:"item"`
}

type RootNode struct {
	Channel ChannelNode `xml:"channel"`
}

func New() *RootNode {
	return &RootNode{}
}
