package data

import "fmt"

//Technology collects the slack user ID of people who use or want to learn a
//given technology, identified by Name
type Technology struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Users    []string `json:"users"`
	Learners []string `json:"learners"`
}

//AddUser ...
func (d *Technology) AddUser(user string, tech string) {
	if listContains(d.Users, user) {
		fmt.Println(user + " is already using " + tech)
		return
	}

	d.Users = append(d.Users, user)
	DB.UpdateTechnology(d)

}

//AddLearner ...
func (d *Technology) AddLearner(user string, tech string) {
	if listContains(d.Learners, user) {
		fmt.Println(user + " is already learning " + tech)
		return
	}

	d.Learners = append(d.Learners, user)
	DB.UpdateTechnology(d)

}

func listContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
