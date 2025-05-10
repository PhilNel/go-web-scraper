package sink

import (
	"context"
	"fmt"
	"go-web-scraper/internal/logging"
	"go-web-scraper/internal/model"
)

type ConsoleSink struct {
	log logging.Logger
}

func NewConsoleSink() *ConsoleSink {
	return &ConsoleSink{
		log: logging.GetLogger("ConsoleSink"),
	}
}

func (s *ConsoleSink) Write(_ context.Context, jobs []model.Job) error {
	for _, job := range jobs {
		s.log.Info(fmt.Sprintf("📌 %s — %s", job.Title, job.Department))
	}
	return nil
}
