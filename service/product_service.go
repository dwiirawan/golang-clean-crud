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

type productRepository struct {
	repo repository.ProductRepository
}

func NewProductService(p repository.ProductRepository) ProductService {
	return &productRepository{p}
}

// Create implements ProductService.
func (p *productRepository) Create(product models.Product) (models.Product, error) {
	if product.Name == "" {
		return product, errors.New("name required")
	}

	return p.repo.Create(product)
}

// Delete implements ProductService.
func (p *productRepository) Delete(id uint) error {
	return p.repo.Delete(id)
}

// GetAll implements ProductService.
func (p *productRepository) GetAll() ([]models.Product, error) {
	return p.repo.FindAll()
}

// GetbyID implements ProductService.
func (p *productRepository) GetbyID(id uint) (models.Product, error) {
	return p.repo.FindByID(id)
}

// Update implements ProductService.
func (p *productRepository) Update(id uint, input models.Product) (models.Product, error) {
	product, err := p.repo.FindByID(id)
	if err != nil {
		return product, err
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock

	return p.repo.Update(product)
}
