package persistence

import "dictio-scrapper/model"

type DB interface {
	Connect()
	Save(entry model.Entry) error
}
