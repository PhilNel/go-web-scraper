package parser

import (
	"testing"

	"go-web-scraper/internal/model"

	"github.com/stretchr/testify/assert"
)

const testHTML = `
<h2 class="openPositions_department__WDYK7">Engineering - Android</h2>
<article typeof="JobPosting">
  <header>
    <h3 class="openPositions_title__KiGPr">Senior Android Engineer</h3>
  </header>
</article>
<h2 class="openPositions_department__WDYK7">Engineering - Backend Engineering</h2>
<article typeof="JobPosting">
  <header>
    <h3>Senior Backend Engineer</h3>
  </header>
</article>
`

func TestDuckDuckGoParser_ParseJobs(t *testing.T) {
	expected := []model.Job{
		{
			Title:      "Senior Android Engineer",
			Department: "Engineering - Android",
		},
		{
			Title:      "Senior Backend Engineer",
			Department: "Engineering - Backend Engineering",
		},
	}
	parser := NewDuckDuckGoParser()

	jobs := parser.Parse(testHTML)

	assert.Equal(t, 2, len(jobs), "should parse exactly two jobs")
	for i, job := range jobs {
		assert.Equal(t, expected[i].Title, job.Title, "title mismatch at index %d", i)
		assert.Equal(t, expected[i].Department, job.Department, "department mismatch at index %d", i)
	}
}
