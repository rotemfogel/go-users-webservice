package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users = make(map[int]*User)
	nextId = 1
)

func GetUsers() []*User {
	arr := make([]*User, len(users))
	i := 0
	for _, v := range users {
		arr[i] = v
		i++
	}
	return arr
}

func GetUser(id int) (*User, error) {
	user := users[id]
	var err error = nil
	if user == nil {
		err = fmt.Errorf("User with id [%v] does not exist", id)
	}
	return user, err
}

func AddUser(u User) (User, error) {
	if u.Id != 0 {
		return u, errors.New("User object cannot have and Id")
	}
	u.Id = nextId
	nextId++
	users[u.Id] = &u
	return u, nil
}

func UpdateUser(u User) (User, error) {
	_, err := GetUser(u.Id)
	if err == nil {
		users[u.Id] = &u
	}
	return u, err
}

func RemoveUser(id int) error {
	u, err := GetUser(id)
	if u != nil {
		delete(users, id)
	}
	return err
}