package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/team").Subrouter()
	s.Methods("POST").HandleFunc("/technology", record.NewTechnology)
	s.Methods("POST").HandleFunc("/learning", record.NewLearning)
	s.Methods("GET").HandleFunc("/user/{name}/", search.ByUser)
	s.Methods("GET").HandleFunc("/technology/{name}", search.ByTechnology)
	s.Methods("GET").HandleFunc("/learning/{name}", search.ByLearning)
	http.Handle("/", r)
}
