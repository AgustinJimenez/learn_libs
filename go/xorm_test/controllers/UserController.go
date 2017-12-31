package controllers

import (
	"fmt"
	/*
		"../xorm"
		"../entities"
		"../repositories"
	*/
	"../entities"
)

type UserController struct {
	//Service services.UserService
	User entities.User
}

func (this *UserController) Create() {
	fmt.Printf("\n==============================\n")

	//this.User.Firstname = "here"
	//fmt.Println(this.User)

	var tmp string

	fmt.Printf("\nInsert user first name:\n")
	fmt.Scanln(&tmp)
	this.User.Firstname = tmp

	fmt.Printf("\nInsert user last name:\n")
	fmt.Scanln(&tmp)
	this.User.Lastname = tmp

	fmt.Printf("\nInsert user email:\n")
	fmt.Scanln(&tmp)
	this.User.Email = tmp

	fmt.Printf("\nInsert user username:\n")
	fmt.Scanln(&tmp)
	this.User.Username = tmp

	fmt.Printf("\nInsert user password:\n")
	fmt.Scanln(&tmp)
	this.User.Password = tmp
	fmt.Printf("\n==============================\n")
	this.Store()
}

func (this *UserController) Store() {
	fmt.Println(this.User)
}
