package entities

import (
	"log"

	"github.com/jinzhu/gorm"
)

type ProductQuantity struct {
	gorm.Model
	Product   Product
	ProductID int
	Quantity  int
	CartID    int
}

type ProductQuantityEntity struct {
	db *gorm.DB
}

func NewProductQuantityEntity(db *gorm.DB) ProductQuantityEntity {
	return ProductQuantityEntity{db}
}

func (h *ProductQuantityEntity) DeleteProductQuantity(id int) (c ProductQuantity) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Delete(&c)
	}
	return
}

func (h *ProductQuantityEntity) UpdateProductQuantity(ci ProductQuantity) (c ProductQuantity) {
	if err := h.db.Where("id = ?", ci.ID).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Save(&ci)
	}
	return ci
}

func (h *ProductQuantityEntity) CreateProductQuantity(ci ProductQuantity) (co ProductQuantity) {
	if err := h.db.Create(&ci).Error; err != nil {
		log.Panic(err)
	}
	return ci
}

func (h *ProductQuantityEntity) GetProductQuantityByID(id int) (c ProductQuantity) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}

func (h *ProductQuantityEntity) GetProductQuantity() (c []ProductQuantity) {
	if err := h.db.Find(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}

func (h *ProductQuantityEntity) GetProductQuantityByCartID(id int) (c []ProductQuantity) {
	if err := h.db.Where("cart_id = ?", id).Find(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}
