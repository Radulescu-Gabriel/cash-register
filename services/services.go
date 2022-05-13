package services

import (
	"errors"
	"fmt"

	"github.com/Radulescu-Gabriel/cash-register/models"
	"github.com/gofiber/fiber/v2"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
	pg *paginate.Pagination
}

var dbProduct models.Product
resutl := models.Product{}

func (s *service) New(product *models.Product, newProduct models.Product) (*models.Product, error) {
	if product != nil {

		err1 := fmt.Errorf("No error ocured")
		return nil, err1
	}

	if result != "" {

		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Product already exists",
		})
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"error": "Internal server error",
	})

	return nil, nil
}

func (s *service) Get(product *models.Product, id uint) (*models.Product, error) {
	if product == nil {
		return nil, errors.New("Product not found")
	}
}

func (s *service) Query(product *models.Product, ctx *fiber.Ctx) (paginate.Page, error) {
	stmt := s.db.Table("products").
		Select("product.id, product.name, product.description, product.stock, product.price, product.discount, product.discountEndDate, product.image")
	page := s.pg.With(stmt).Request(ctx.Request()).Response(&[]models.Product{})

	return page, nil
}

func (s *service) Update(user *model.User, updatedUser *model.User) (*model.User, error) {
	// Guests not allowed
	if user == nil {
		return nil, errors.New(services.ErrNotAllowed)
	}

	// Basic users should only be able to update themselves
	if user.IsBasic() && user.ID != updatedUser.ID {
		return nil, errors.New(services.ErrNotAllowed)
	}

	if errs := services.ValidateStruct(*updatedUser); errs != nil {
		// todo: return all errors
		return nil, errs[0]
	}

	var dbUser model.User
	var res *gorm.DB

	if res = s.db.First(&dbUser, updatedUser.ID); res.Error != nil {
		return nil, res.Error
	}

	if res = s.db.Model(&dbUser).Updates(updatedUser); res.Error != nil {
		return nil, res.Error
	}

	// todo: send notification
	dbUser.Password = ""

	return &dbUser, nil
}
