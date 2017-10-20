package models

import (
	"errors"
)

var (
	UserList map[string]*User1
)

func init() {
	UserList = make(map[string]*User1)
	u1 := User1{
		1,
		"yanghuolong1",
		"yhl27ml@163.com",
		"123456",
	}
	u2 := User1{
		2,
		"yanghuolong2",
		"yhl27ml@126.com",
		"123456",
	}

	UserList["1"] = &u1
	UserList["2"] = &u2
}

type User1 struct {
	Id       int
	Username string
	Email    string
	Password string
}

func GetUser(id string) (u *User1, err error) {
	if u, ok := UserList[id]; ok {
		return u, nil
	}

	return nil, errors.New("没有此数据")
}

func GetAllUser() map[string]*User1 {
	return UserList
}

func GetUserLogin(email, password string) (u *User1, err error) {
	for _, u := range UserList {
		if u.Email == email && u.Password == password {
			return u, nil
		}
	}

	return nil, errors.New("帐号或密码有错误!")
}
