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

	apiRouter.HandleFunc("/status", api.GetAlerts).Methods("GET")
	apiRouter.HandleFunc("/status/{id}", api.GetAlertById).Methods("GET")

	apiRouter.HandleFunc("/check/{id}", api.GetAlertStatusResume).Methods("GET")
	apiRouter.HandleFunc("/check_details/{id}", api.GetAlertStatusDetails).Methods("GET")

	apiRouter.HandleFunc("/alerts/{id}", api.UpdateAlert).Methods("PUT")

	/* 	apiRouter.HandleFunc("/alerts/{id}", api.DeleteAlert).Methods("DELETE")
	   	apiRouter.HandleFunc("/alerts", api.CreateAlert).Methods("POST")
	*/
	fmt.Printf("Server running at port %s", port)
	http.ListenAndServe(":"+port, router)
}
