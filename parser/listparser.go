package parser

import (
	"dictio-scrapper/model"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	WORDS_SECTION = "#columna_resultados_generales"
	WORD_BLOCK    = "#palabra_resultado"
)

// ListParser Data type for words list file parsing
type ListParser interface {
	Parse(string) []model.Word
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
func (p ListParserImpl) Parse(content string) []model.Word {
	result := []model.Word{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return result
	}

	words := doc.Find(WORDS_SECTION).First()

	words.Find(WORD_BLOCK).Each(func(_ int, node *goquery.Selection) {
		title, _ := node.Attr("title")
		wordsList := strings.Split(title, " ")
		word := wordsList[len(wordsList)-1]
		word = strings.Trim(word, "\n")
		if len(word) > 3 {
			url, _ := node.Attr("href")
			url = strings.Trim(url, "\n")
			wordObject := model.Word{
				Name: word,
				URL:  url,
			}
			result = append(result, wordObject)
		}
	})

	return result
}
