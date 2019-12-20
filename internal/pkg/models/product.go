package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Product model
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// ToString human readable format.
func (p *Product) ToString() string {
	return fmt.Sprintf("ID : %d\tPrice : %d\tCode : %s", p.ID, p.Price, p.Code)
}

// Products are an array of Product
type Products []Product

// ToString human readable format
func (pp Products) ToString() string {
	s := ""
	for _, ss := range pp {
		s += ss.ToString() + "\n"
	}
	return s
}
