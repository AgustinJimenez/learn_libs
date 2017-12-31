package main

import (
	"fmt"

	xorm "./xorm"
	"./xorm/entities"
)

func main() {
	DB := xorm.GetInstance()
	_ = DB.Sync2(new(Entities.User))
	new_user := Entities.User{}
	fmt.Println("\n\n\n123 = ", new_user.ValidateHash("123", new_user.GetHash("123")))
	fmt.Println("123 = ", new_user.ValidateHash("123", new_user.GetHash("123")))
	fmt.Println("123 = ", new_user.ValidateHash("123", new_user.GetHash("1234")))
	fmt.Println("123 = ", new_user.ValidateHash("123", new_user.GetHash("123")))
	fmt.Println("123 = ", new_user.ValidateHash("123", new_user.GetHash("123")))
}
