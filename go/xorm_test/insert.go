package main

import (
	"fmt"
	"time"

	xorm "./xorm"
	"./xorm/entities"
)

func main() {
	DB := xorm.GetInstance()
	_ = DB.Sync2(new(Entities.User))
	new_user := Entities.User{}

	new_user.Firstname = "agustin"
	new_user.Lastname = "jimenez"
	new_user.Email = "agus.jimenez.caba@mail.com"
	new_user.Username = "agusjim"
	new_user.Password = "123"
	new_user.GeneratePassword()
	new_user.CreatedAt = time.Now()

	affected, err := DB.Insert(&new_user)
	if err != nil {
		fmt.Println("ERROR ON INSERT", err)
	}

	if affected > 0 {
		fmt.Println("affected")
		results, _ := DB.Query("select * from users")
		fmt.Println(results, affected)
	} else {
		fmt.Printf("not affected " /*+err.Error()*/)
	}
}
