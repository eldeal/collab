package data

import (
	"log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//type List []string

//Mongo holds the details of an established session with MongoDB
type Mongo struct {
	Session  *mgo.Session
	Database string
}

var DB *Mongo

//StartSession initializes a connection to MongoDB for use throughout the application.
func StartSession() {
	mongoURI := os.Getenv("MONGODB_URI")
	db := os.Getenv("DB_NAME")
	s, err := mgo.Dial(mongoURI)
	if err != nil {
		panic(err)
	}
	DB = &Mongo{Session: s, Database: db}
}

//FindTechnology uses the name of the technology to return the JSON document
//stored for it.
func (m *Mongo) FindTechnology(tech string) *Technology {
	s := m.Session.Clone()
	defer s.Close()
	c := s.DB(m.Database).C("technology")

	result := &Technology{}

	if err := c.Find(bson.M{"name": tech}).One(&result); err != nil {
		log.Fatal(err)
	}
	return result
}

//UpdateTechnology takes a Technology and upserts the relevant JSON docuemnt
func (m *Mongo) UpdateTechnology(tech *Technology) *mgo.ChangeInfo {
	s := m.Session.Clone()
	defer s.Close()
	c := s.DB(m.Database).C("technology")

	info, err := c.UpsertId(tech.ID, tech)
	if err != nil {
		log.Fatal(err)
	}
	return info
}

//NewTechnology adds a new JSON document for the technology name, adding it's
//first user or learner determined by the triggering slack command.
func (m *Mongo) NewTechnology(tech string, user string, trigger string) {
	new := &Technology{Name: tech}
	if trigger == "tech:" {
		new.Users = append(new.Users, user)
	} else if trigger == "learn:" {
		new.Learners = append(new.Learners, user)
	}

	s := m.Session.Clone()
	defer s.Close()
	c := s.DB(m.Database).C("technology")
	if err := c.Insert(new); err != nil {
		log.Fatal(err)
	}

}

/*
func AllRecords() {
	s := getSession()
	defer s.Close()
	c := s.DB("collab").C("technology")

	results := &[]Technology{}

	err := c.Find(nil).Sort("-timestamp").All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
} */
