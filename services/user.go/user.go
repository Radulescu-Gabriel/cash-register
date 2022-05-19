package services

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/Radulescu-Gabriel/cash-register/models"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
	pg *paginate.Pagination
}

var dbUser models.User

func New(db *gorm.DB, pg *paginate.Pagination) *service {
	return &service{
		db: db,
		pg: pg,
	}
}

func (s *service) New(user *models.User, newUser models.User) (*models.User, error) {
	if user != nil {

		err1 := fmt.Errorf("no error ocured")
		return nil, err1
	}

	if len(newUser.Password) > 0 {
		// User with password
		encryptedPassword := sha256.Sum256([]byte(newUser.Password))
		newUser.Password = string(encryptedPassword[:])
	}

	if res := s.db.Create(&newUser); res.Error != nil {
		return nil, res.Error
	}

	newUser.Password = ""

	return &newUser, nil
}

func (s *service) Get(user *models.User, id uint) (*models.User, error) {
	// Guests not allowed
	if user == nil {
		return nil, errors.New("must log in first")
	}

	return &dbUser, nil
}
