package handler

import (
	"context"
	"fmt"
	"go-web-scraper/internal/logging"
	"go-web-scraper/internal/model"
	"go-web-scraper/internal/parser"
	"go-web-scraper/internal/util"
)

type JobSink interface {
	Write(ctx context.Context, jobs []model.Job) error
}

type JobProvider interface {
	Get(ctx context.Context, path string) (string, error)
}

type JobHandler struct {
	Provider JobProvider
	Sink     JobSink
	log      logging.Logger
}

func NewJobHandler(p JobProvider, s JobSink) *JobHandler {
	return &JobHandler{
		Provider: p,
		Sink:     s,
		log:      logging.GetLogger("JobHandler"),
	}
}

func (h *JobHandler) Handle(ctx context.Context, event model.EventBridgeEvent) error {
	eventKey := event.Detail.Object.Key
	h.log.WithFields(map[string]interface{}{
		"key": eventKey,
	}).Info("Obtained event from event bridge")

	html, err := h.Provider.Get(ctx, eventKey)
	if err != nil {
		return err
	}

	company, err := util.ParseCompanyNameFromKey(eventKey)
	if err != nil {
		return fmt.Errorf("failed to parse company from key: %w", err)
	}

	parser, err := parser.GetParserForCompany(company)
	if err != nil {
		return fmt.Errorf("parser lookup failed: %w", err)
	}
	jobs := parser.Parse(html)

	err = h.Sink.Write(ctx, jobs)
	if err != nil {
		return err
	}

	return nil
}
