package api

import (
	"encoding/json"
	"go-postgresql-rest-api/models"
	"net/http"
	"strconv"

	"go-postgresql-rest-api/helpers"
	"strings"

	"github.com/gorilla/mux"
)

type Data struct {
	Success bool            `json:"success"`
	Data    []models.Status `json:"data"`
	Errors  []string        `json:"errors"`
}

type StatusResume struct {
	Status   string `json:"status"`
	Elements string `json:"elements"`
	Success  bool   `json:"success"`
}

func GetAlerts(w http.ResponseWriter, req *http.Request) {
	var Status []models.Status = models.GetAll()
	var data = Data{true, Status, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetAlertById(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var status models.Status
	var success bool
	status, success = models.Get(id)
	if !success {
		data.Success = false
		data.Errors = append(data.Errors, "not found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true
	data.Data = append(data.Data, status)
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetAlertStatusDetails(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var status models.Status
	var success bool
	status, success = models.Get(id)
	if !success {
		data.Success = false
		data.Errors = append(data.Errors, "not found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	errorStatus := strconv.Itoa(1)
	var dataStatus StatusResume
	dataStatus.Success = true
	dataStatus.Elements = ""

	if status.Ds_status_flujo == errorStatus || status.Ds_status_aforo == errorStatus || status.Faust_status_flujo == errorStatus || status.Faust_status_aforo == errorStatus || helpers.IsTimeInInterval(status.Timestamp, 15) {
		dataStatus.Status = "Error"

		if status.Ds_status_flujo == errorStatus {
			dataStatus.Elements += "Ds_FLUJO "
		}

		if status.Ds_status_aforo == errorStatus {
			dataStatus.Elements += "Ds_AFORO "
		}

		if status.Faust_status_flujo == errorStatus {
			dataStatus.Elements += "faust_FLUJO "
		}

		if status.Faust_status_aforo == errorStatus {
			dataStatus.Elements += "faust_AFORO "
		}
		if helpers.IsTimeInInterval(status.Timestamp, 15) {
			dataStatus.Elements += "appliances without connection (diff > 15min)"
		}

	} else {
		dataStatus.Status = "Ok"
	}

	json, _ := json.Marshal(dataStatus)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetAlertStatusResume(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var status models.Status
	var success bool
	status, success = models.Get(id)
	if !success {
		data.Success = false
		data.Errors = append(data.Errors, "not found")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	errorStatus := strconv.Itoa(1)
	var dataStatus StatusResume
	dataStatus.Success = true
	if status.Ds_status_flujo == errorStatus || status.Ds_status_aforo == errorStatus || status.Faust_status_flujo == errorStatus || status.Faust_status_aforo == errorStatus || helpers.IsTimeInInterval(status.Timestamp, 15) {
		dataStatus.Status = "Error"

		if status.Ds_status_flujo == errorStatus || status.Ds_status_aforo == errorStatus {
			dataStatus.Elements = "Ds"
		}

		if status.Faust_status_flujo == errorStatus || status.Faust_status_aforo == errorStatus {
			if dataStatus.Elements != "" {
				dataStatus.Elements += ", faust"
			} else {
				dataStatus.Elements = "faust"
			}
		}
		if helpers.IsTimeInInterval(status.Timestamp, 15) {
			if dataStatus.Elements != "" {
				dataStatus.Elements += ", appliances without connection (diff > 15min)"
			} else {
				dataStatus.Elements = "appliances without connection (diff > 15min)"
			}
		}

	} else {
		dataStatus.Status = "Ok"
	}

	json, _ := json.Marshal(dataStatus)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func UpdateAlert(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	cc_id := vars["id"]

	bodyStatus, success := helpers.DecodeBody(req)
	if !success {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}
	var data Data = Data{Errors: make([]string, 0)}
	bodyStatus.Ds_status_flujo = strings.TrimSpace(bodyStatus.Ds_status_flujo)
	bodyStatus.Ds_status_aforo = strings.TrimSpace(bodyStatus.Ds_status_aforo)
	bodyStatus.Faust_status_flujo = strings.TrimSpace(bodyStatus.Faust_status_flujo)
	bodyStatus.Faust_status_aforo = strings.TrimSpace(bodyStatus.Faust_status_aforo)
	bodyStatus.Timestamp = strings.TrimSpace(bodyStatus.Timestamp)

	if !(helpers.IsValidEntry(bodyStatus.Ds_status_flujo) && helpers.IsValidEntry(bodyStatus.Ds_status_aforo) && helpers.IsValidEntry(bodyStatus.Faust_status_flujo) && helpers.IsValidEntry(bodyStatus.Faust_status_aforo) && helpers.IsValidEntryTimestamp(bodyStatus.Timestamp)) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid payload")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}

	status, success := models.Update(cc_id, bodyStatus.Ds_status_flujo, bodyStatus.Ds_status_aforo, bodyStatus.Faust_status_flujo, bodyStatus.Faust_status_aforo, bodyStatus.Timestamp)
	if !success {
		data.Errors = append(data.Errors, "could not update alert")
	}

	data.Success = success
	data.Data = append(data.Data, status)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
