package main

import (

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinh98/convictional/engineering-interview/handlers"
)

func main() {
	_, err := handlers.GetProducts()
	if err != nil {
		log.Println(err)
	}

	listenAddr := ":443"
	rtr := mux.NewRouter()
	rtr.HandleFunc("/product/{id}", handlers.FindProductById).Methods(http.MethodGet)

	http.Handle("/", rtr)
	http.ListenAndServe(listenAddr, nil)
}
