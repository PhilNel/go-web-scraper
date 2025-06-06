package parser

import (
	"strings"

	"go-web-scraper/internal/model"

	"github.com/PuerkitoBio/goquery"
)

type DuckDuckGoParser struct{}

func NewDuckDuckGoParser() *DuckDuckGoParser {
	return &DuckDuckGoParser{}
}

func (p *DuckDuckGoParser) Parse(html string) ([]model.Job, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var jobs []model.Job
	var currentDept string

	doc.Find("h2, article").Each(func(_ int, s *goquery.Selection) {
		if isDepartmentHeader(s) {
			currentDept = strings.TrimSpace(s.Text())
			return
		}

		if isJobPosting(s) {
			title := strings.TrimSpace(s.Find("h3").Text())
			jobs = append(jobs, model.Job{
				Title:      title,
				Department: currentDept,
				Company:    "DuckDuckGo",
			})
		}
	})

	return jobs, nil
}

func isDepartmentHeader(s *goquery.Selection) bool {
	return s.Is("h2.openPositions_department__WDYK7")
}

func isJobPosting(s *goquery.Selection) bool {
	return s.Is("article[typeof='JobPosting']")
}
