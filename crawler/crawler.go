package crawler

import (
	"dictio-scrapper/model"
	"dictio-scrapper/parser"
	"dictio-scrapper/persistence"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type HttpGetter interface {
	Get(url string) (string, error)
}

type HttpGetterImpl struct {
	HttpGetter
}

type Crawler interface {
	Process(url string) error
}

type CrawlerImpl struct {
	Crawler
	getter     HttpGetter
	listParser parser.ListParser
	wordParser parser.DefinitionParser
	db         persistence.DB
	alphabet   []string
}

func New(getter HttpGetter, listParser parser.ListParser, wordParser parser.DefinitionParser, db persistence.DB, alphabet []string) Crawler {
	return CrawlerImpl{
		getter:     getter,
		listParser: listParser,
		wordParser: wordParser,
		db:         db,
		alphabet:   alphabet,
	}
}

func (c CrawlerImpl) Process(url string) error {
	logrus.Info("==================================================")
	logrus.Info("= Starting crawler process                       =")
	logrus.Info("==================================================")

	for _, letter := range c.alphabet {
		logrus.Infof("Processing letter [%s]", letter)
		uri := strings.Replace(url, "${LETTER}", letter, -1)

		logrus.Infof("Request data from %s", uri)
		listContent, err := c.getter.Get(uri)

		if err != nil {
			return fmt.Errorf("Error getting data from %s: %s", uri, err)
		}

		list := c.listParser.Parse(listContent)
		logrus.Infof("Parsed %d words with letter %s", len(list), letter)

		for _, word := range list {
			logrus.Infof("Parsing definition for word [%s]", word.Name)

			definitionContent, err := c.getter.Get(word.URL)

			if err != nil {
				return fmt.Errorf("Error getting data from %s: %s", word.URL, err)
			}

			definition := c.wordParser.Parse(definitionContent)

			if len(definition) > 0 {
				logrus.Infof("Processed definition of %s: %s", word.Name, definition)

				data := model.Entry{
					Word:       word.Name,
					Definition: definition,
					Letters:    []string{strings.ToLower(word.Name)[0:1]},
				}
				err = c.db.Save(data)

				if err != nil {
					return err
				}
			} else {
				logrus.Infof("Discarding word [%s] with empty definition", word)
			}
		}
	}

	return nil
}

func (h HttpGetterImpl) Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	buffer, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", nil
	}

	body := string(buffer)

	if res.StatusCode != 200 {
		return "", fmt.Errorf("Error %d from HTTP response: %s", res.StatusCode, body)
	}

	return body, nil
}
