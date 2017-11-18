package main

import (
	"fmt"
)

//interfaces
type User interface {
	Permissions() int
	Name() string
}

func auth(this User) string {

	if this.Permissions() > 5 {
		return this.Name() + " is admin"
	} else if this.Permissions() <= 6 {
		return this.Name() + " is editor"
	} else {
		return "invalid op"
	}

}

type Admin struct {
	name string
}

func (this Admin) Permissions() int {
	return 10
}
func (this Admin) Name() string {
	return this.name
}

type Editor struct {
	Admin
}

func (this Editor) Permissions() int {
	return 5
}

func main() {

	users := []User{Admin{"John"}, Editor{Admin{"Johnny"}}}

	for index, user := range users {
		fmt.Println(auth(user))
	}

}
