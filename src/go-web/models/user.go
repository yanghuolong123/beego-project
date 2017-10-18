package models

import (
	"errors"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u1 := User{
		1,
		"yanghuolong1",
		"yhl27ml@163.com",
		"123456",
	}
	u2 := User{
		2,
		"yanghuolong2",
		"yhl27ml@126.com",
		"123456",
	}

	UserList["1"] = &u1
	UserList["2"] = &u2
}

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

func GetUser(id string) (u *User, err error) {
	if u, ok := UserList[id]; ok {
		return u, nil
	}

	return nil, errors.New("没有此数据")
}

func GetAllUser() map[string]*User {
	return UserList
}

func Login(email, password string) bool {
	return true
}
