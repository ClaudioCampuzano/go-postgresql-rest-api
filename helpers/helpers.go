package helpers

import (
	"encoding/json"
	"go-postgresql-rest-api/models"
	"net/http"
	"strings"
	"time"
)

func DecodeBody(req *http.Request) (models.Status, bool) {
	var status models.Status
	err := json.NewDecoder(req.Body).Decode(&status)
	if err != nil {
		return models.Status{}, false
	}

	return status, true
}

func IsValidEntry(entry string) bool {
	desc := strings.TrimSpace(entry)
	if len(desc) == 0 {
		return false
	}
	return true
}

func IsValidEntryTimestamp(entry string) bool {
	desc := strings.TrimSpace(entry)
	_, err := time.Parse("2006-01-02 15:04:05.999", desc)
	if err != nil {
		return false
	}

	return true
}
