package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/eldeal/collab/data"
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
	data.StartSession()
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

	for _, tech := range msg.split() {
		doc := data.DB.FindTechnology(tech)
		if doc == nil {
			data.DB.NewTechnology(tech, msg.Name, msg.Trigger)
			continue
		}

		switch msg.Trigger {
		case "tech:":
			doc.AddUser(msg.Name, tech)
		case "learn:":
			doc.AddLearner(msg.Name, tech)
		default:
			fmt.Println("How very invalid of you. Try 'tech:' or 'learn:', I don't understand anything else... yet!")
		}

	}
}

func (msg *SlackMessage) split() []string {
	s := strings.TrimPrefix(msg.Text, msg.Trigger)
	s = strings.TrimSpace(s)
	return strings.Split(s, ",")
}
