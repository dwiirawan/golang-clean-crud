package repository

import (
	"golang-clean-crud/models"

	"gorm.io/gorm"
)

// Contract
type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindByID(id uint) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(id uint) error
}

type productRepository struct {
	db *gorm.DB
}

// Create object reposiory
// Fill dependency database
// Return ProductRepository interface
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

// Create implements ProductRepository.
func (p *productRepository) Create(product models.Product) (models.Product, error) {
	err := p.db.Create(&product).Error
	return product, err
}

// Delete implements ProductRepository.
func (p *productRepository) Delete(id uint) error {
	return p.db.Delete(&models.Product{}, id).Error

}

// FindAll implements ProductRepository.
func (p *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := p.db.Find(&products).Error
	return products, err
}

// FindByID implements ProductRepository.
func (p *productRepository) FindByID(id uint) (models.Product, error) {
	var products models.Product
	err := p.db.First(&products, id).Error
	return products, err
}

// Update implements ProductRepository.
func (p *productRepository) Update(product models.Product) (models.Product, error) {
	err := p.db.Save(&product).Error
	return product, err
}
