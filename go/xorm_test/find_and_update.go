package main

import (
	"fmt"

	xorm "./xorm"
	"./xorm/entities"
)

func main() {
	DB := xorm.GetInstance()
	_ = DB.Sync2(new(Entities.User))

	user := &Entities.User{}

	has, err := DB.Id(3).Get(user)

	if err != nil {
		fmt.Println("DATABASE ERROR!!!")
	} else if has {
		fmt.Println("USER FOUNDED!!!", user)
	} else {
		fmt.Println("USER NOT FOUNDED!!!")
	}

	user.Firstname = "Agustin"
	user.Lastname = "Jimenez"
	affected, err := DB.Update(user)

	if err != nil {
		fmt.Println("DATABASE ERROR!!!")
	} else if affected > 0 {
		fmt.Println("USER UPDATED!!!", affected)
	} else {
		fmt.Println("USER NOT UPDATED!!!")
	}

}
