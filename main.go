package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/eldeal/collab/db"
	"github.com/gorilla/mux"
)

//SlackMessage is a subset of the information sent in the message body of a
//request from Slack
type SlackMessage struct {
	ID      string `json:"user_id"`
	Name    string `json:"user_name"`
	Text    string `json:"text"`         //googlebot: What is the air-speed velocity of an unladen swallow?
	Trigger string `json:"trigger_word"` //googlebot:
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/collab").Subrouter()
	s.HandleFunc("/add", add).Methods("POST")
	//s.Methods("POST").HandleFunc("/user/{name}/", findUser)
	//s.Methods("POST").HandleFunc("/technology/{name}", findTechnology)
	//s.Methods("POST").HandleFunc("/learning/{name}", findLearners)
	http.Handle("/", r)
}

func add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var msg SlackMessage
	if err := decoder.Decode(&msg); err != nil {
		panic(err)
	}

	// break msg.text into list of individual technologies
	techs := msg.split()

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

func (msg *SlackMessage) split() []string {
	s := strings.TrimPrefix(msg.Text, msg.Trigger)
	s = strings.TrimSpace(s)
	return strings.Split(s, ",")
}
