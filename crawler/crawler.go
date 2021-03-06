package crawler

import (
	"dictio-scrapper/model"
	"dictio-scrapper/parser"
	"dictio-scrapper/persistence"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"

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

	status, err := c.db.LoadStatus()
	if err != nil {
		return err
	}

	logrus.Infof("Last saved status: %v", status)
	letterFound := false

	for _, letter := range c.alphabet {
		if !letterFound && status.Letter != "" && status.Letter != letter {
			continue
		}
		letterFound = true
		status.Letter = letter

		logrus.Infof("Processing letter [%s]", letter)
		uri := strings.Replace(url, "${LETTER}", letter, -1)

		logrus.Infof("Request data from %s", uri)
		listContent, err := c.getter.Get(uri)

		if err != nil {
			return fmt.Errorf("Error getting data from %s: %s", uri, err)
		}

		list := c.listParser.Parse(listContent)
		logrus.Infof("Parsed %d words with letter %s", len(list), letter)

		wordFound := false

		for _, word := range list {
			if !wordFound {
				if status.Word == "" || status.Word == word.Name {
					wordFound = true
					if status.Word == word.Name {
						continue
					}
				} else {
					continue
				}
			}

			logrus.Infof("Parsing definition for word [%s] from %s", word.Name, word.URL)

			urlParts := strings.Split(word.URL, " ")

			if len(urlParts) > 1 {
				logrus.Infof("Discarding invalid URL for %s: %s", word.Name, word.URL)
				continue
			}

			definitionContent, err := c.getter.Get(word.URL)

			if err != nil {
				return fmt.Errorf("Error getting data from %s: %s", word.URL, err)
			}

			definition := c.wordParser.Parse(definitionContent)

			if len(definition) > 0 {
				logrus.Infof("Processed definition of %s: %s", word.Name, definition)

				letters := []string{getLetter(strings.ToLower(word.Name))}

				logrus.Info(letters)

				data := model.Entry{
					Word:       word.Name,
					Definition: definition,
					Letters:    letters,
				}
				err = c.db.Save(data)

				if err != nil {
					return err
				}
			} else {
				logrus.Infof("Discarding word [%s] with empty definition", word)
			}

			status.Word = word.Name
			status.Timestamp = time.Now().UTC().Format(time.RFC3339)
			err = c.db.SaveStatus(status)
			if err != nil {
				return err
			}
		}

		status.Word = ""
	}

	logrus.Info("==================================================")
	logrus.Info("= Finished crawler process                       =")
	logrus.Info("==================================================")

	return nil
}

func getLetter(s string) string {
	letter := s[0:1]
	if utf8.ValidString(letter) {
		return string(letter[0])
	} else {
		tildeMap := map[byte]string{
			"??"[1]: "a",
			"??"[1]: "e",
			"??"[1]: "i",
			"??"[1]: "o",
			"??"[1]: "u",
		}
		return tildeMap[s[1]]
	}
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
