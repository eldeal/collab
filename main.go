package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eldeal/collab/db"
	"github.com/gorilla/mux"
)

type SlackMessage struct {
	ID      string `json:"user_id"`
	Name    string `json:"user_name"`
	Text    string `json:"text"`         //googlebot: What is the air-speed velocity of an unladen swallow?
	Trigger string `json:"trigger_word"` //googlebot:
}

type User struct {
	id    string
	tech  []Thing
	learn []Thing
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/collab").Subrouter()
	s.Methods("POST").HandleFunc("/add", add)
	//	s.Methods("POST").HandleFunc("/learning", learn)
	s.Methods("GET").HandleFunc("/user/{name}/", findUser)
	s.Methods("GET").HandleFunc("/technology/{name}", findTechnology)
	s.Methods("GET").HandleFunc("/learning/{name}", findLearners)
	http.Handle("/", r)
}

func add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	var msg SlackMessage
	if err := decoder.Decode(&msg); err != nil {
		panic()
	}

	// break msg.text into list of individual technologies
	techs := msg.split()

List:
	for _, tech := range techs {
		doc := db.FindTechnology(tech)
		if doc == nil {
			db.NewTechnology(tech, msg.Name, msg.Trigger)
			continue
		}

		if msg.Trigger == "tech:" {
			if contains(doc.Users, msg.Name) {
				fmt.Println(msg.Name + " is already using " + tech)
			} else {
				doc.Users = append(doc.Users, msg.Name)
				db.UpdateTechnology(doc)
			}
		} else if msg.Trigger == "learn:" {
			if contains(doc.Learners, msg.Name) {
				fmt.Println(msg.Name + " is already learning " + tech)
			} else {
				doc.Learners = append(doc.Learners, msg.Name)
				db.UpdateTechnology(doc)
			}
		}

	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (*SlackMessage) split() []string {
	s := string.TrimPrefix(msg.Text, msg.Trigger)
	s = string.TrimSpace(s)
	return string.Split(s, ",")
}

func findUser(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	name := args["name"]
}
