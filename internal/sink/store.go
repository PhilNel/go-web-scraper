package sink

import (
	"go-web-scraper/internal/logging"
	"go-web-scraper/internal/model"
)

type JobStore interface {
	Store(job model.Job) error
}

type StoreSink struct {
	log   logging.Logger
	store JobStore
}

func NewStoreSink(store JobStore) *StoreSink {
	return &StoreSink{
		log:   logging.GetLogger("StoreSink"),
		store: store,
	}
}

func (s *StoreSink) Write(jobs []model.Job) {
	for _, job := range jobs {
		err := s.store.Store(job)
		if err != nil {
			s.log.WithError(err).Error("Error writing job to store")
		}
	}
}
