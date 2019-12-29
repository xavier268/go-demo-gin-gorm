package dao

import (
	"github.com/xavier268/go-demo-gin-gorm/internal/models"
)

// CountProducts return the number of products in db.
func (s *Source) CountProducts() int {
	var count int
	err := s.GetDAO().Model(&models.Product{}).Count(&count).Error
	if err != nil {
		s.Close()
		panic(err)
	}
	return count
}

// CreateProduct register new product, return ID.
// Duplicates if already exists.
func (s *Source) CreateProduct(price uint, code string) uint {
	p := new(models.Product)
	p.Price, p.Code = price, code
	s.GetDAO().Create(&p)
	return p.ID
}

// DeleteProduct using primary key.
func (s *Source) DeleteProduct(id uint) {
	p := new(models.Product)
	p.ID = id
	s.GetDAO().Delete(&p)
}

// DeleteProducts delete all products
func (s *Source) DeleteProducts() {
	s.GetDAO().Delete(&models.Product{})
}

// AllProducts dumps table content.
func (s *Source) AllProducts() models.Products {
	var pp models.Products
	s.GetDAO().Order("price desc, code").Find(&pp)
	return pp
}

// GetProduct by id
func (s *Source) GetProduct(id uint) *models.Product {
	p := new(models.Product)
	s.GetDAO().First(p, id)
	return p
}
