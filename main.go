package main

import (
	"log"
	"net/http"

	"Assignment2/helper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/orders", helper.GetOrders).Methods("GET")
	r.HandleFunc("/orders/create", helper.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/update", helper.UpdateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", helper.DeleteOrder).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
