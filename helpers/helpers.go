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
	return len(desc) != 0
}

func IsValidEntryTimestamp(entry string) bool {
	desc := strings.TrimSpace(entry)
	_, err := time.Parse("2006-01-02 15:04:05.999", desc)
	return err == nil
}

func IsTimeInInterval(entry string, interval int) bool {
	loc, _ := time.LoadLocation("America/Santiago")
	desc := strings.TrimSpace(entry)
	tMsg, err := time.ParseInLocation("2006-01-02 15:04:05.999", desc, loc)
	tNow := time.Now()
	if err != nil {
		return false
	}
	diff := tNow.Sub(tMsg)

	return diff > time.Minute*time.Duration(interval)
}
