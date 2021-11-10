package api

import (
	"encoding/json"
	"go-postgresql-rest-api/models"
	"net/http"

	"go-postgresql-rest-api/helpers"
	"strings"

	"github.com/gorilla/mux"
)

type Data struct {
	Success bool            `json:"success"`
	Data    []models.Status `json:"data"`
	Errors  []string        `json:"errors"`
}

func GetAlerts(w http.ResponseWriter, req *http.Request) {
	var Status []models.Status = models.GetAll()

	var data = Data{true, Status, make([]string, 0)}
	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetAlert(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data

	var status models.Status
	var success bool
	status, success = models.Get(id)
	if success != true {
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

func UpdateAlert(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	cc_id := vars["id"]

	bodyStatus, success := helpers.DecodeBody(req)
	if success != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyStatus.Ds_status = strings.TrimSpace(bodyStatus.Ds_status)
	bodyStatus.Faust_status = strings.TrimSpace(bodyStatus.Faust_status)
	bodyStatus.Timestamp = strings.TrimSpace(bodyStatus.Timestamp)

	if !(helpers.IsValidEntry(bodyStatus.Ds_status) && helpers.IsValidEntry(bodyStatus.Faust_status) && helpers.IsValidEntryTimestamp(bodyStatus.Timestamp)) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid payload")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	status, success := models.Update(cc_id, bodyStatus.Ds_status, bodyStatus.Faust_status, bodyStatus.Timestamp)
	if success != true {
		data.Errors = append(data.Errors, "could not update alert")
	}

	data.Success = success
	data.Data = append(data.Data, status)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
}

/*

func DeleteTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	var data Data = Data{Errors: make([]string, 0)}

	todo, success := models.Delete(id)
	if success != true {
		data.Errors = append(data.Errors, "could not delete todo")
	}

	data.Success = success
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	bodyTodo, success := helpers.DecodeBody(req)
	if success != true {
		http.Error(w, "could not decode body", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "invalid description")

		json, _ := json.Marshal(data)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	todo, success := models.Insert(bodyTodo.Description)
	if success != true {
		data.Errors = append(data.Errors, "could not create todo")
	}

	data.Success = success
	data.Data = append(data.Data, todo)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
	return
} */
