package main

import (
	"dictio-scrapper/config"
	"dictio-scrapper/crawler"
	"dictio-scrapper/parser"
	"strings"
)

func main() {
	config.LoadConfig()

	getter := crawler.HttpGetterImpl{}
	listParser := parser.NewListParser()
	definitionParser := parser.NewDefinitionParser()

	alphabet := strings.Split(config.GlobalConfig.Alphabet, ",")

	c := crawler.New(getter, listParser, definitionParser, nil, alphabet)
	c.Process(config.GlobalConfig.URL)
}
