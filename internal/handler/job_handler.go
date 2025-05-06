package handler

import (
	"go-web-scraper/internal/model"
)

type JobParser interface {
	Parse(html string) []model.Job
}

type JobSink interface {
	Write(jobs []model.Job)
}

type JobProvider interface {
	Get(path string) (string, error)
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

func (h *JobHandler) Handle() error {
	html, err := h.Provider.Get("../node-web-fetcher/rendered.html")
	if err != nil {
		return err
	}

	jobs := h.Parser.Parse(html)
	h.Sink.Write(jobs)

	return nil
}
