package repositories

import (
	"fmt"

	"../entities"
	"../xorm"
)

type UserRepository interface {
	Store(new_user entities.User) (entities.User, error)
	GetByUsernameAndPassword(username string, password string) (entities.User, bool)
	GetById(id int64) (entities.User, bool)
}
type userRepository struct {
	User entities.User
	xorm xorm.xorm_mysql.Engine
}

func (this *userRepository) Store(new_user entities.User) (entities.User, error) {
	DB := this.xorm.New()
	err := DB.Sync2(new(entities.User))

	/*affected*/
	_, err2 := DB.Insert(&this.User)

	if err != nil {
		return this.User, err
	} else {
		return this.User, err2
	}
}
func (this *userRepository) GetByUsernameAndPassword(username string, password string) (entities.User, bool) {
	valid_password := false

	DB := xorm_mysql.GetInstance()
	_ = DB.Sync2(new(entities.User))

	user := &entities.User{}
	has /*err*/, _ := DB.Where("username = ?", username).Get(user)

	if has {
		valid_password = user.ValidatePassword(password)
	}

	return *user, (has && valid_password)
}

func (this *userRepository) GetById(id int64) (entities.User, bool) {

	DB := xorm_mysql.GetInstance()
	_ = DB.Sync2(new(entities.User))

	user := &entities.User{}
	has, err := DB.Id(3).Get(user)

	if err != nil {
		fmt.Println("DATABASE ERROR!!!")
		panic(err)
	} else if has {
		return *user, true
	} else {
		return *user, false
	}

}
