package dao

import (
	"github.com/xavier268/go-demo-gin-gorm/internal/models"
)

// CountProducts return the number of products in db.
func (d *DAO) CountProducts() int {
	var count int
	err := GetDAO().Model(&models.Product{}).Count(&count).Error
	if err != nil {
		d.Close()
		panic(err)
	}
	return count
}

// CreateProduct register new product, return ID.
// Duplicates if already exists.
func (d *DAO) CreateProduct(price uint, code string) uint {
	p := new(models.Product)
	p.Price, p.Code = price, code
	GetDAO().Create(&p)
	return p.ID
}

// DeleteProduct using primary key.
func (d *DAO) DeleteProduct(id uint) {
	p := new(models.Product)
	p.ID = id
	GetDAO().Delete(&p)
}

// DeleteProducts delete all products
func (d *DAO) DeleteProducts() {
	GetDAO().Delete(&models.Product{})
}

// AllProducts dumps table content.
func (d *DAO) AllProducts() models.Products {
	var pp models.Products
	GetDAO().Order("price desc, code").Find(&pp)
	return pp
}

// GetProduct by id
func (d *DAO) GetProduct(id uint) *models.Product {
	p := new(models.Product)
	GetDAO().First(p, id)
	return p
}
