package handler

import (
	"context"
	"go-web-scraper/internal/model"
)

type JobParser interface {
	Parse(html string) []model.Job
}

type JobSink interface {
	Write(jobs []model.Job)
}

type JobProvider interface {
	Get(ctx context.Context) (string, error)
}

type JobHandler struct {
	Provider JobProvider
	Parser   JobParser
	Sink     JobSink
}

func NewJobHandler(p JobProvider, pr JobParser, s JobSink) *JobHandler {
	return &JobHandler{
		Provider: p,
		Parser:   pr,
		Sink:     s,
	}
}

func (h *JobHandler) Handle(ctx context.Context) error {
	html, err := h.Provider.Get(ctx)
	if err != nil {
		return err
	}

	jobs := h.Parser.Parse(html)
	h.Sink.Write(jobs)

	return nil
}
