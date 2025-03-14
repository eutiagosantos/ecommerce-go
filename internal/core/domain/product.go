package domain

import (
	"errors"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Price    uint64 `json:"price"`
	Category string `json:"category"`
}

func BuildProduct(name string, price uint64, category string) (*Product, error) {
	product := &Product{
		Name:     name,
		Price:    price,
		Category: category,
	}
	if err := product.validate(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) validate() error {

	if len(p.Name) < 10 {
		return errors.New("The name of product must contain at least 10 characters")
	}

	if p.Price <= 0 {
		return errors.New("The price of the product must be greater than zero")
	}

	return nil
}
