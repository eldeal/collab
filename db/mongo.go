package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//var session *mgo.Session
var coll *mgo.Collection

func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func Find() {

}

func FindLearners() {

}

type Technology struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Users    []string `json:"users"`
	Learners []string `json:"learners"`
}

func FindTechnology(tech string) *Technology {
	s := getSession()
	defer s.Close()
	c := s.DB("collab").C("technology")

	result := &Technology{}

	if err := c.Find(bson.M{"name": tech}).One(&result); err != nil {
		log.Fatal(err)
	}
	return result
}

func NewTechnology(tech string, user string, trigger string) {
	new := &Technology{Name: tech}
	if trigger == "tech:" {
		new.Users = append(new.Users, user)
	} else if trigger == "learn:" {
		new.Learners = append(new.Learners, user)
	}

	s := getSession()
	defer s.Close()
	c := s.DB("collab").C("technology")
	if err := c.Insert(new); err != nil {
		panic(err)
	}

}

func NewLearning() {

}

func AllRecords() {
	err := c.Find(nil).Sort("-timestamp").All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
}
