package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var members = []Member{Member{"user1", "user1@example.com"}}

func main() {
	http.HandleFunc("/members", membersHandler)
	http.ListenAndServe(":8080", nil)
}

type Member struct {
	Login string
	Email string
}

func membersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		j, _ := json.Marshal(members)

		w.Write(j)
	}

	if r.Method == "POST" {
		var m Member
		b, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(b, &m)

		members = append(members, m)

		j, _ := json.Marshal(m)
		w.Write(j)
	}
}

/*
GET or POST url:
http://localhost:8080/members

JSON structure:
{"Login": "user1","Email": "user1@example.com"}
*/
