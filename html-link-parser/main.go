package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	fName := flag.String("file", "ex1.html", "HTML file with links to be parsed")
	flag.Parse()
	file, err := os.Open(*fName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	links, err := parseHtml(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(links)
}

func parseHtml(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	var f func(n *html.Node)
	var links []Link
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			links = append(links, Link{
				Href: getHref(n),
				Text: getLinkText(n.FirstChild)})
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}

func getLinkText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getLinkText(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func getHref(n *html.Node) string {
	for _, a := range n.Attr {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}
