package service

import (
	"errors"
	"golang-clean-crud/models"
	"golang-clean-crud/repository"
)

type ProductService interface {
	GetAll() ([]models.Product, error)
	GetbyID(id uint) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(id uint, input models.Product) (models.Product, error)
	Delete(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

// Function NewProductService receive ProductRepository interface
// Return_Type: ProductService interface
// Return_Value: productService, struct implementasi service
// (struct automatically becomes an interface )
func NewProductService(p repository.ProductRepository) ProductService {
	return &productService{p}
}

// Create implements ProductService.
func (s *productService) Create(product models.Product) (models.Product, error) {
	if product.Name == "" {
		return product, errors.New("name required")
	}

	return s.repo.Create(product)
}

// Delete implements ProductService.
func (s *productService) Delete(id uint) error {
	return s.repo.Delete(id)
}

// GetAll implements ProductService.
func (s *productService) GetAll() ([]models.Product, error) {
	return s.repo.FindAll()
}

// GetbyID implements ProductService.
func (s *productService) GetbyID(id uint) (models.Product, error) {
	return s.repo.FindByID(id)
}

// Update implements ProductService.
func (s *productService) Update(id uint, input models.Product) (models.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return product, err
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock

	return s.repo.Update(product)
}
