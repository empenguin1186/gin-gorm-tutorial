package repository

import (
	"errors"

	. "github.com/empenguin1186/gin-gorm-tutorial/entity"
)

var (
	users []User
)

type UserRepoitoryOnMemory struct {
	users []User
}

func NewUserRepositoryOnMemory() UserRepoitoryOnMemory {
	return UserRepoitoryOnMemory{users: []User{}}
}

func (u UserRepoitoryOnMemory) FindAll() ([]User, error) {
	return users, nil
}

func (u UserRepoitoryOnMemory) FindOne(id string) (User, error) {
	for _, e := range users {
		if e.ID == id {
			return e, nil
		}
	}

	return User{}, errors.New("Could not find user")
}

func (u UserRepoitoryOnMemory) Create(user User) error {
	users = append(u.users, user)
	return nil
}

func (u UserRepoitoryOnMemory) Update(id string, user User) error {
	for i, e := range users {
		if e.ID == id {
			users = append(users[:i], user)
			for _, u := range users[i+1:] {
				users = append(users, u)
			}
			return nil
		}
	}
	return errors.New("Could not find user.")
}

func (u UserRepoitoryOnMemory) Delete(id string) error {
	for i, e := range users {
		if e.ID == id {
			for _, u := range users[i+1:] {
				users = append(users[:i], u)
			}
			return nil
		}
	}
	return errors.New("Could not find user.")
}
