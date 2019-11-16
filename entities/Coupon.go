package entities

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Coupon struct {
	gorm.Model
	Value          string
	Discount       float64
	IsActive       bool `sql:"type:integer;"`
	ExpirationDate time.Time
	MinAmount      float64
	DiscountType   DiscountType `sql:"type:varchar(50);"`
}

type CouponEntity struct {
	db *gorm.DB
}

func NewCouponEntity(db *gorm.DB) CouponEntity {
	return CouponEntity{db}
}

func (h *CouponEntity) GetCouponByValue(value string) Coupon {
	var couponObj Coupon
	if err := h.db.Where("value = ?", value).First(&couponObj).Error; err != nil {
		log.Panic(err)
	}
	return couponObj
}

func (h *CouponEntity) DeleteCoupon(id int) (c Coupon) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Delete(&c)
	}
	return
}

func (h *CouponEntity) UpdateCoupon(ci Coupon) (c Coupon) {
	if err := h.db.Where("id = ?", ci.ID).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Save(&ci)
	}
	return ci
}

func (h *CouponEntity) CreateCoupon(ci Coupon) (co Coupon) {
	if err := h.db.Create(&ci).Error; err != nil {
		log.Panic(err)
	}
	return ci
}

func (h *CouponEntity) GetCouponByID(id int) (c Coupon) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}

func (h *CouponEntity) GetCoupon() (c []Coupon) {
	if err := h.db.Find(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}
