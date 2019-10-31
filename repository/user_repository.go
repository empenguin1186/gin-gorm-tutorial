package repository

import (
	"github.com/empenguin1186/gin-gorm-tutorial/db"
	"github.com/empenguin1186/gin-gorm-tutorial/entity"
)

type User entity.User

type UserRepository interface {
	FindAll() ([]User, error)
	FindOne(string) (User, error)
	Create(User) error
	Update(string, User) error
	Delete(string) error
}

type UserRepoitoryPostgre struct{}

func (repository *UserRepoitoryPostgre) FindAll() ([]User, error) {
	db := db.GetDB()
	var u []User

	if err := db.Find(&u).Error; err != nil {
		return nil, err
	}

	return u, err
}

func (repository *UserRepoitoryPostgre) FindOne(id string) (User, error) {
	db := db.GetDB()
	var u User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (repository *UserRepoitoryPostgre) Create(u User) error {
	db := db.GetDB()
	if db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (repository *UserRepoitoryPostgre) Update(id string, u User) error {
	db := db.GetDB()

	var us User
	if err := db.Where("id = ?", id).First(&us).Error; err != nil {
		return err
	}
	us = u
	db.Save(&us)
	return nil
}

func (repository *UserRepoitoryPostgre) Delete(id string) error {
	db := db.GetDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}
