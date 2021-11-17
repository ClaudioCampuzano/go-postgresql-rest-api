package main

import (
	"fmt"
	"go-postgresql-rest-api/api"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var port string = "2711"

	router := mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/omia/").Subrouter()
	apiRouter.HandleFunc("/alerts", api.GetAlerts).Methods("GET")
	apiRouter.HandleFunc("/alerts/{id}", api.GetAlert).Methods("GET")
	apiRouter.HandleFunc("/status/{id}", api.GetAlertStatus).Methods("GET")
	apiRouter.HandleFunc("/alerts/{id}", api.UpdateAlert).Methods("PUT")

	/* 	apiRouter.HandleFunc("/alerts/{id}", api.DeleteAlert).Methods("DELETE")
	   	apiRouter.HandleFunc("/alerts", api.CreateAlert).Methods("POST")
	*/
	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)
}
