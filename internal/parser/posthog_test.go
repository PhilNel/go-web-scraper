package parser

import (
	"testing"

	"go-web-scraper/internal/model"

	"github.com/stretchr/testify/assert"
)

const testPostHogHTML = `
<button class="w-full flex flex-col text-left px-2 py-1 rounded border border-b-3 hover:bg-light/50 hover:dark:bg-dark/50 border-transparent md:hover:border-light dark:md:hover:border-dark hover:translate-y-[-1px] active:translate-y-[1px] active:transition-all">
	<span class="font-semibold text-[15px] ">
		Product Engineer
	</span>
	<span class="text-[13px] text-black/50 dark:text-white/50">
		Multiple teams
	</span>
</button>
`

func TestPostHogParser_ParseJobs_ValidJobs_ParsesCorrectly(t *testing.T) {
	expected := []model.Job{
		{
			Title:      "Product Engineer",
			Department: "Multiple teams",
		},
	}
	parser := NewPostHogParser()

	jobs, err := parser.Parse(testPostHogHTML)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(jobs), "should parse exactly 1 jobs")
	for i, job := range jobs {
		assert.Equal(t, expected[i].Title, job.Title, "title mismatch at index %d", i)
		assert.Equal(t, expected[i].Department, job.Department, "department mismatch at index %d", i)
	}
}
