package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Have a channel where every time you work on a new project, area, or technology
//or develop an interest in learning about it you send a message in this channel.
//It's public, so others can see what you're working on or interested in
//but also searchable via a slackbot, so other people can find experience people.

//In Slack user types:
//tech: scala, circleci, newthing2.0
//learn: devops, cooking

//Then to query user can use slackbot:
//  /collab tech: go
// to find users who have previously said they're using Go.

//Doesn't require a weekly standup or check in, but allows to be quickly updated
//whenever a new thing crops up without maintaining a separate system than Slack
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
