package parser

import (
	"strings"

	"go-web-scraper/internal/model"

	"github.com/PuerkitoBio/goquery"
)

type PostHogParser struct{}

func NewPostHogParser() *PostHogParser {
	return &PostHogParser{}
}

func (p *PostHogParser) Parse(html string) ([]model.Job, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var jobs []model.Job

	doc.Find("button.flex.flex-col.text-left.px-2.py-1").Each(func(_ int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("span").Eq(0).Text())
		dept := strings.TrimSpace(s.Find("span").Eq(1).Text())

		if title != "" && dept != "" {
			jobs = append(jobs, model.Job{
				Title:      title,
				Department: dept,
				Company:    "PostHog",
			})
		}
	})

	return jobs, nil
}
