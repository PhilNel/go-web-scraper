package sink

import (
	"context"
	"go-web-scraper/internal/logging"
	"go-web-scraper/internal/model"
)

type JobStore interface {
	Store(ctx context.Context, job model.Job) error
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

func (s *StoreSink) Write(ctx context.Context, jobs []model.Job) error {
	for _, job := range jobs {
		err := s.store.Store(ctx, job)
		if err != nil {
			s.log.WithError(err).Error("Error writing job to store")
			return err
		}
	}
	return nil
}
