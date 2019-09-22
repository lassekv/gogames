package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func docParser(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		l := Link{
			Href: extractHref(n.Attr),
			Text: extractText(n),
		}
		*links = append(*links, l)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		docParser(c, links)
	}

}

func extractText(n *html.Node) string {
	var text []string
	extractTextComponents(n, &text)
	return strings.TrimSpace(strings.Join(text, ""))
}

func extractTextComponents(n *html.Node, ss *[]string) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			*ss = append(*ss, c.Data)
		}
		if c.Type == html.ElementNode {
			extractTextComponents(c, ss)
		}
	}
}

func extractHref(attributes []html.Attribute) string {
	for _, a := range attributes {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func ParseHTML(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return make([]Link, 0), err
	}
	var links []Link
	docParser(doc, &links)
	return links, nil
}
