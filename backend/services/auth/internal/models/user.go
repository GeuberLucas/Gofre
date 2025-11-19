package models

import (
	"time"
)

type User struct {
	ID        int64
	Username  string
	Name      string
	LastName  string
	Cellphone string
	Email     string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}


func (u User) Validate() bool{
	if (u.Username == ""){
		return false
	}
	if(u.Email == ""){
		return false
	}
	if (len(u.Password)==0){
		return false
	}
	return true
}