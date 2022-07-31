package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/towfeeq/bookstore-with-db/pkg/route"
)

func main() {
	r := mux.NewRouter()

	route.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	r.HandleFunc("/test", healthcheck).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App is up and runninng")
}
