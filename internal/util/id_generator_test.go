package util

import (
	"testing"

	"go-web-scraper/internal/model"

	"github.com/stretchr/testify/assert"
)

func Test_GenerateJobID_OutputsExpected(t *testing.T) {
	job := model.Job{
		Title:      "Software Engineer",
		Department: "Engineering",
		Company:    "DuckDuckGo",
	}
	expected := "8a118a0579d79aafd2e59d1928adb5abeb55a8af899cb9aebdd8e3eecc569b4e"

	actual := GenerateJobID(job)

	assert.Equal(t, expected, actual)
}
