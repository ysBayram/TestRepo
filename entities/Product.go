package entities

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Title      string
	Price      float64
	Category   Category
	CategoryID int
}

type ProductEntity struct {
	db *gorm.DB
}

func NewProductEntity(db *gorm.DB) ProductEntity {
	return ProductEntity{db}
}

func (h *ProductEntity) DeleteProduct(id int) (c Product) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Delete(&c)
	}
	return
}

func (h *ProductEntity) UpdateProduct(ci Product) (c Product) {
	if err := h.db.Where("id = ?", ci.ID).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Save(&ci)
	}
	return ci
}

func (h *ProductEntity) CreateProduct(ci Product) (co Product) {
	if err := h.db.Create(&ci).Error; err != nil {
		log.Panic(err)
	}
	return ci
}

func (h *ProductEntity) GetProductByID(id int) (c Product) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}

func (h *ProductEntity) GetProduct() (c []Product) {
	if err := h.db.Find(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}
