package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ListParser Data type for words list file parsing
type ListParser interface {
	Parse(string) []string
}

// ListParserImpl Implementation of words list parser type
type ListParserImpl struct {
	ListParser
}

// NewListParser Constructor of words list parser type
func NewListParser() ListParser {
	return ListParserImpl{}
}

// Parse function to parse list of words from html string
func (p ListParserImpl) Parse(content string) []string {
	result := []string{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return result
	}

	doc.Find("#palabra_resultado").Each(func(_ int, node *goquery.Selection) {
		title, _ := node.Attr("title")
		wordsList := strings.Split(title, " ")
		word := wordsList[len(wordsList)-1]
		word = strings.Trim(word, "\n")
		if len(word) > 3 {
			result = append(result, word)
		}
	})

	return result
}
