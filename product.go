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

func New(db *gorm.DB, pg *paginate.Pagination) *service {
	return &service{
		db: db,
		pg: pg,
	}
}

func (s *service) New(product *models.Product, newProduct models.Product) (*models.Product, error) {
	if product != nil {

		err1 := fmt.Errorf("no error ocured")
		return nil, err1
	}

	// if product.Name != "" {

	// 	return c.Status(fiber.StatusConflict).JSON(fiber.Map{
	// 		"error": "Product already exists",
	// 	})
	// }

	// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 	"error": "Internal server error",
	// })

	// return nil, nil

	return nil, nil
}

func (s *service) Get(product *models.Product, id uint) (*models.Product, error) {
	if product == nil {
		return nil, errors.New("product not found")
	}

	return nil, nil
}

func (s *service) Query(product *models.Product, ctx *fiber.Ctx) (paginate.Page, error) {
	stmt := s.db.Table("products").
		Select("product.id, product.name, product.description, product.stock, product.price, product.discount, product.discountEndDate, product.image")
	page := s.pg.With(stmt).Request(ctx.Request()).Response(&[]models.Product{})

	return page, nil
}

func (s *service) Update(product *models.Product, updatedProduct *models.Product) (*models.Product, error) {
	if product == nil {
		return nil, errors.New("product not found")
	}

	var dbProduct models.Product
	var res *gorm.DB

	if res = s.db.First(&dbProduct, updatedProduct.ID); res.Error != nil {
		return nil, res.Error
	}

	if res = s.db.Model(&dbProduct).Updates(updatedProduct); res.Error != nil {
		return nil, res.Error
	}

	return &dbProduct, nil
}
