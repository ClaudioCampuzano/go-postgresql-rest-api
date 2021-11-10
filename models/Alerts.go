package models

import (
	"database/sql"
	"go-postgresql-rest-api/database"
	"log"
)

type Status struct {
	Cc_name      string `json:"cc_name"`
	Ds_status    string `json:"ds_status"`
	Faust_status string `json:"faust_status"`
	Timestamp    string `json:"timestamp"`
}

func Get(id string) (Status, bool) {
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM server_status WHERE cc_name = $1", id)

	var cc_name sql.NullString
	var ds_status sql.NullString
	var faust_status sql.NullString
	var timestamp sql.NullString

	err := row.Scan(&cc_name, &ds_status, &faust_status, &timestamp)
	if err != nil {
		return Status{}, false
	}

	return Status{cc_name.String, ds_status.String, faust_status.String, timestamp.String}, true
}

func GetAll() []Status {
	db := database.GetConnection()
	rows, err := db.Query("SELECT * FROM server_status")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var status []Status
	for rows.Next() {
		t := Status{}

		var cc_name sql.NullString
		var ds_status sql.NullString
		var faust_status sql.NullString
		var timestamp sql.NullString

		err := rows.Scan(&cc_name, &ds_status, &faust_status, &timestamp)
		if err != nil {
			log.Fatal(err)
		}

		t.Cc_name = cc_name.String
		t.Ds_status = ds_status.String
		t.Faust_status = faust_status.String
		t.Timestamp = timestamp.String

		status = append(status, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return status
}

func Update(id string, ds_status string, faust_status string, timestamp string) (Status, bool) {
	db := database.GetConnection()

	var cc_id sql.NullString
	db.QueryRow("UPDATE server_status SET ds_status = $1, faust_status = $2, timestamp = $3 WHERE cc_name = $4 RETURNING cc_name", ds_status, faust_status, timestamp, id).Scan(&cc_id)
	if cc_id.String == "" {
		return Status{}, false
	}

	return Status{cc_id.String, ds_status, faust_status, timestamp}, true
}

/* func Delete(id string) (Todo, bool) {
	db := database.GetConnection()

	var todo_id int
	db.QueryRow("DELETE FROM todos WHERE id = $1 RETURNING id", id).Scan(&todo_id)
	if todo_id == 0 {
		return Todo{}, false
	}

	return Todo{todo_id, ""}, true
}

func Insert(description string) (Todo, bool) {
	db := database.GetConnection()

	var todo_id int
	db.QueryRow("INSERT INTO todos(description) VALUES($1) RETURNING id", description).Scan(&todo_id)

	if todo_id == 0 {
		return Todo{}, false
	}

	return Todo{todo_id, ""}, true
}
*/
