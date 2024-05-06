package repository

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) port.IProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Find(id int) (*domain.Product, error) {
	var product *model.Product

	if err := r.db.Where("id=?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return product.ToDomain(), nil
}

func (r *ProductRepository) CreateProduct(product *domain.Product) (*model.Product, error) {
	productModel := new(model.Product)
	productModel.FromDomain(product)

	if err := r.db.Create(productModel).Error; err != nil {
		return nil, err
	}

	return productModel, nil
}
