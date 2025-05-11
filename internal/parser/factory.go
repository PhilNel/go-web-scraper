package parser

import (
	"fmt"
	"go-web-scraper/internal/model"
)

type JobParser interface {
	Parse(html string) []model.Job
}

func GetParserForCompany(company string) (JobParser, error) {
	switch company {
	case "duckduckgo":
		return NewDuckDuckGoParser(), nil
	default:
		return nil, fmt.Errorf("no parser available for company: %s", company)
	}
}
