package entities

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	gorm.Model
	Products       []ProductQuantity
	TotalDiscounts float64
	TotalAmount    float64
	DeliveryCost   float64
	CouponID       int
}

type CartEntity struct {
	db *gorm.DB
}

func NewCartEntity(db *gorm.DB) CartEntity {
	return CartEntity{db}
}

func (h *CartEntity) DeleteCart(id int) (c Cart) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Delete(&c)
	}
	return
}

func (h *CartEntity) UpdateCart(ci Cart) (c Cart) {
	if err := h.db.Where("id = ?", ci.ID).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Save(&ci)
	}
	return ci
}

func (h *CartEntity) CreateCart(ci Cart) (co Cart) {
	if err := h.db.Create(&ci).Error; err != nil {
		log.Panic(err)
	}
	return ci
}

func (h *CartEntity) GetCartByID(id int) (c Cart) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}

func (h *CartEntity) GetCart() (c []Cart) {
	if err := h.db.Find(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}
