package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var members = []Member{Member{"user1", "user1@example.com"}}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/members", getMembersHandler).Methods("GET")
	r.HandleFunc("/members", postMembersHandler).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

type Member struct {
	Login string
	Email string
}

func getMembersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	j, _ := json.Marshal(members)
	w.Write(j)
}

func postMembersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m Member
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &m)

	members = append(members, m)

	j, _ := json.Marshal(m)
	w.Write(j)
}

/*
GET or POST url:
http://localhost:8080/members

JSON structure:
{"Login": "user1","Email": "user1@example.com"}
*/
