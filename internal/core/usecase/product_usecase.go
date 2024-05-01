package usecase

import (
	"github.com/br4tech/go-with-gemini/internal/core/domain"
	"github.com/br4tech/go-with-gemini/internal/core/port"
	"github.com/br4tech/go-with-gemini/internal/dto"
	validator "github.com/br4tech/go-with-gemini/pkg"
)

type ProductUsecase struct {
	productRepository port.IProductRepository
}

func NewProductUseCase(productRepository port.IProductRepository) port.IProductUseCase {
	return &ProductUsecase{productRepository: productRepository}
}

func (uc ProductUsecase) Find(id int) (*domain.Product, error) {
	product, err := uc.productRepository.Find(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (uc *ProductUsecase) CreateProduct(productDTO dto.ProductDTO) (*domain.Product, error) {
	product := domain.NewProduct(
		productDTO.Name,
		productDTO.Code,
		productDTO.Image,
	)

	if err := validator.ValidateStruct(product); err != nil {
		return nil, err
	}

	createdProduct, err := uc.productRepository.CreateProduct(product)

	if err != nil {
		return nil, err
	}

	return createdProduct.ToDomain(), nil
}
