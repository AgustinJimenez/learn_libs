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
		return "is admin"
	} else if this.Permissions() <= 6 {
		return "is editor"
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

	new_admin := Admin{"John"}
	new_editor := Editor{Admin{"Johnny"}}

	fmt.Println(new_admin.Name(), auth(new_admin), new_editor.Name(), auth(new_editor))

}
