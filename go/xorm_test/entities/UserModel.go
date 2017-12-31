package entities

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int64     `json:"id" form:"id"`
	Firstname string    `json:"firstname" form:"firstname" xorm:"varchar(90)"`
	Lastname  string    `json:"lastname" form:"lastname" xorm:"varchar(90)"`
	Email     string    `json:"email" form:"email" xorm:"varchar(60)"`
	Username  string    `json:"username" form:"username" xorm:"varchar(60)"`
	Password  string    `json:"password" form:"password" xorm:"password" xorm:"varchar(200)"`
	CreatedAt time.Time `json:"created_at" form:"created_at" xorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at" xorm:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
func (this *User) GeneratePassword() {
	hashed, err := bcrypt.GenerateFromPassword([]byte(this.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Printf("\n ERROR ON USER MODEL \n")
		panic(err)
	}

	this.Password = string(hashed)
}

// ValidatePassword will check if passwords are matched.
func (this *User) ValidatePassword(userPassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(this.Password), []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

func (this *User) ValidateHash(password string, hash []byte) bool {
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		return false
	}
	return true
}

func (this *User) GetHash(userPassword string) []byte {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	return hashed
}

// IsValid can do some very very simple "low-level" data validations.
func (u *User) IsValid() bool {
	return u.Id > 0
}
func (this User) CreatedAtFormated() string {
	return this.CreatedAt.Format("2006-01-02 15:04:05")
}
func (this User) UpdatedAtFormated() string {
	return this.UpdatedAt.Format("2006-01-02 15:04:05")
}
