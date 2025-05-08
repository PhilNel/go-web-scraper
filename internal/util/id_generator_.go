package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-web-scraper/internal/model"
)

func GenerateJobID(job model.Job) string {
	hash := sha256.New()
	data := fmt.Sprintf("%s|%s|%s", job.Title, job.Department, job.Company)
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
