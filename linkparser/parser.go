package linkparser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link will represent an anchor html element
type Link struct {
	Href string
	Text string
}

// Parse will take a golang reader and will return a list with the detected anchor element
func Parse(reader io.Reader) ([]Link, error) {
	root, error := html.Parse(reader)
	if error != nil {
		return nil, error
	}

	links := getLinksFromSubtree(root, []Link{})

	return links, nil
}

// This can be refactored making both steps (getting link nodes, and getting Link structure in separated functions)
func getLinksFromSubtree(element *html.Node, links []Link) []Link {
	if element.Data == "a" {
		var linkInfo Link
		for _, attr := range element.Attr {
			if attr.Key == "href" {
				linkInfo.Href = attr.Val
				break
			}
		}

		linkInfo.Text = strings.Join(strings.Fields(buildLinkText(element.FirstChild)), " ")
		links = append(links, linkInfo)
	}

	if element.FirstChild != nil {
		links = getLinksFromSubtree(element.FirstChild, links)
	}

	if element.NextSibling != nil {
		links = getLinksFromSubtree(element.NextSibling, links)
	}

	return links
}

func buildLinkText(element *html.Node) string {
	if element.Data == "a" || element == nil {
		return ""
	}

	if element.Type == html.TextNode {
		return element.Data
	}

	var linkText string = ""
	for node := element.FirstChild; node != nil; node = element.NextSibling {
		linkText += buildLinkText(node)
	}

	return linkText
}
