package models

import (
	"database/sql"
	"go-postgresql-rest-api/database"
	"log"
)

type Status struct {
	Cc_name            string `json:"cc_name"`
	Ds_status_flujo    string `json:"ds_status_flujo"`
	Ds_status_aforo    string `json:"ds_status_aforo"`
	Faust_status_flujo string `json:"faust_status_flujo"`
	Faust_status_aforo string `json:"faust_status_aforo"`
	Timestamp          string `json:"timestamp"`
}

func Get(id string) (Status, bool) {
	db := database.GetConnection()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM server_statusV2 WHERE cc_name = $1", id)

	var cc_name sql.NullString
	var ds_status_flujo sql.NullString
	var ds_status_aforo sql.NullString
	var faust_status_flujo sql.NullString
	var faust_status_aforo sql.NullString
	var timestamp sql.NullString

	err := row.Scan(&cc_name, &ds_status_flujo, &ds_status_aforo, &faust_status_flujo, &faust_status_aforo, &timestamp)
	if err != nil {
		return Status{}, false
	}

	return Status{cc_name.String, ds_status_flujo.String, ds_status_aforo.String, faust_status_flujo.String, faust_status_aforo.String, timestamp.String}, true
}

func GetAll() []Status {
	db := database.GetConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM server_statusV2")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var status []Status
	for rows.Next() {
		t := Status{}

		var cc_name sql.NullString
		var ds_status_flujo sql.NullString
		var ds_status_aforo sql.NullString
		var faust_status_flujo sql.NullString
		var faust_status_aforo sql.NullString
		var timestamp sql.NullString

		err := rows.Scan(&cc_name, &ds_status_flujo, &ds_status_aforo, &faust_status_flujo, &faust_status_aforo, &timestamp)
		if err != nil {
			log.Fatal(err)
		}

		t.Cc_name = cc_name.String
		t.Ds_status_flujo = ds_status_flujo.String
		t.Ds_status_aforo = ds_status_aforo.String
		t.Faust_status_flujo = faust_status_flujo.String
		t.Faust_status_aforo = faust_status_aforo.String
		t.Timestamp = timestamp.String

		status = append(status, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return status
}

func Update(id string, ds_status_flujo string, ds_status_aforo string, faust_status_flujo string, faust_status_aforo string, timestamp string) (Status, bool) {
	db := database.GetConnection()
	defer db.Close()

	var cc_id sql.NullString
	db.QueryRow("UPDATE server_statusV2 SET ds_status_flujo = $1, ds_status_aforo = $2, faust_status_flujo = $3, faust_status_aforo = $4, timestamp = $5 WHERE cc_name = $6 RETURNING cc_name",
		ds_status_flujo, ds_status_aforo, faust_status_flujo, faust_status_aforo, timestamp, id).Scan(&cc_id)
	if cc_id.String == "" {
		return Status{}, false
	}

	return Status{cc_id.String, ds_status_flujo, ds_status_aforo, faust_status_flujo, faust_status_aforo, timestamp}, true
}
