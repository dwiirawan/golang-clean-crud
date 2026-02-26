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
func (r *productRepository) Create(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

// Delete implements ProductRepository.
func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error

}

// FindAll implements ProductRepository.
func (r *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error
	return products, err
}

// FindByID implements ProductRepository.
func (r *productRepository) FindByID(id uint) (models.Product, error) {
	var products models.Product
	err := r.db.First(&products, id).Error
	return products, err
}

// Update implements ProductRepository.
func (r *productRepository) Update(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}
