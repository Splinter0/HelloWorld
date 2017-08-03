package main

import "encoding/xml"

type Query struct {
	XMLName xml.Name `xml:"urlset"`
	Pars    []string `xml:"url>p"`
	Chapts  []string `xml:"url>h1"`
}

type Line struct {
	Text string
	Id   string
}

func conLine(t string, id string) *Line {
	l := &Line{
		Text: t,
		Id:   id,
	}
	return l
}

func main() {

}
