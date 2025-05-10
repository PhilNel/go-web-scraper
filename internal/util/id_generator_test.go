package util

import (
	"testing"

	"go-web-scraper/internal/model"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateJobID_OutputsExpected(t *testing.T) {
	job := model.Job{
		Title:      "Senior Backend Engineer",
		Department: "Engineering - Backend Engineering",
		Company:    "DuckDuckGo",
	}
	expected := "f44efb505a4969fb0dcc565f95fbfd05ad822127c56ca3a63a31b441eb2c376a"

	actual := GenerateJobID(job)

	assert.Equal(t, expected, actual)
}
