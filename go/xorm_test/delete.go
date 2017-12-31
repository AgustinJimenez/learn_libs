package main

import (
	"fmt"

	xorm "./xorm"
	"./xorm/entities"
)

func main() {
	DB := xorm.GetInstance()
	_ = DB.Sync2(new(Entities.User))

	affected, err := DB.Delete(&Entities.User{Id: 2})
	if err != nil {
		fmt.Println("DATABASE ERROR!!!")
	} else if affected > 0 {
		fmt.Println("USER DELETED!!!----->", affected)
	} else {
		fmt.Println("USER NOT FOUNDED!!!")
	}

}
