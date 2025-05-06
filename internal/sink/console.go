package sink

import (
	"fmt"
	"go-web-scraper/internal/model"
)

type ConsoleSink struct{}

func NewConsoleSink() *ConsoleSink {
	return &ConsoleSink{}
}

func (s *ConsoleSink) Write(jobs []model.Job) {
	for _, job := range jobs {
		fmt.Printf("📌 %s %s\n", job.Title, job.Department)
	}
}
