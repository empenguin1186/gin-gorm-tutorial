package user

import (
	"github.com/empenguin1186/gin-gorm-tutorial/db"
	"github.com/empenguin1186/gin-gorm-tutorial/entity"
	"github.com/empenguin1186/gin-gorm-tutorial/repository"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Repository repository.UserRepository
}

func NewService(r repository.UserRepository) Service {
	return Service{Repository: r}
}

// User is alias of entyty.User struct
type User entity.User

// GetAll is get all User
func (s Service) GetAll() ([]User, error) {
	// db := db.GetDB()
	// var u []User

	// // DB.Find() メソッドは全てのレコードを返す
	// if err := db.Find(&u).Error; err != nil {
	// 	return nil, err
	// }

	// return u, nil
	return s.Repository.FindAll()
}

// CreateModel is create User model
func (s Service) CreateModel(c *gin.Context) (User, error) {
	db := db.GetDB()
	var u User

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	// if err := db.Create(&u).Error; err != nil {
	// 	return u, err
	// }

	// return u, nil
	return u, s.Repository.Create(u)
}

// GetByID is get a User
func (s Service) GetByID(id string) (User, error) {
	// db := db.GetDB()
	// var u User

	// if err := db.Where("id = ?", id).First(&u).Error; err != nil {
	// 	return u, err
	// }

	// return u, nil
	return s.Repository.FindOne(id)
}

// UpdateByID is update a User
func (s Service) UpdateByID(id string, c *gin.Context) (User, error) {
	// db := db.GetDB()
	// var u User

	// if err := db.Where("id = ?", id).First(&u).Error; err != nil {
	// 	return u, err
	// }

	if err := c.BindJSON(&u); err != nil {
		return u, err
	}

	// db.Save(&u)
	return u, s.Repository.Update(id, u)
}

// DeleteByID is delete a User
func (s Service) DeleteByID(id string) error {
	// db := db.GetDB()
	// var u User

	// if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
	// 	return err
	// }

	return s.Repository.Delete(id)
}
