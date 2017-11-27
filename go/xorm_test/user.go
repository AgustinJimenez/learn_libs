package main

import (
	"fmt"
	"time"

	xorm "./xorm"
)

type Users struct {
	Id       int64
	Name     string
	LastName string
	Age      int
	Passwd   string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
}

func init() {

}

func (this Users) CreatedFormated() string {
	return this.Created.Format("2006-01-02 15:04:05")
}
func (this Users) UpdatedFormated() string {
	return this.Updated.Format("2006-01-02 15:04:05")
}

func main() {

	DB := xorm.GetInstance()
	_ = DB.Sync2(new(Users))
	new_user := Users{}
	new_user.Name = "agustin"
	new_user.LastName = "jimenez"
	new_user.Age = 21
	new_user.Passwd = "helloworld"
	new_user.Created = time.Now()

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

	//fmt.Println(new_user.CreatedFormated())
}
