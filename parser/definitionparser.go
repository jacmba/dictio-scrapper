package parser

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

const definitionPrefix = "Versión electrónica 23.4 del «Diccionario de la lengua española», obra lexicográfica académica por excelencia."

// DefinitionParser type that parses word definitions
type DefinitionParser interface {
	Parse(text string) string
}

// DefinitionParserImpl Implementation of definition parser type
type DefinitionParserImpl struct {
	DefinitionParser
}

// NewDefinitionParser Definition parser constructor
func NewDefinitionParser() DefinitionParser {
	return DefinitionParserImpl{}
}

// Parse method that parses html and extracts word definition
func (p DefinitionParserImpl) Parse(text string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(text))
	if err != nil {
		logrus.Errorf("Error parsing html: %v", err)
		return ""
	}

	var result string

	doc.Find("meta").Each(func(_ int, meta *goquery.Selection) {
		name, _ := meta.Attr("name")

		if name == "description" {
			content, _ := meta.Attr("content")

			content = strings.TrimPrefix(content, definitionPrefix)
			contentParts := strings.Split(content, ".")

			if len(content) < 2 {
				result = ""
				return
			}

			result = contentParts[len(contentParts)-2]
			result = strings.Trim(result, " ")
		}
	})
	return result
}
