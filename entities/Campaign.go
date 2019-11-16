package entities

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Campaign struct {
	gorm.Model
	Discount          float64
	ProductCategory   Category
	ProductCategoryID int
	ProductCount      int
	DiscountType      DiscountType `sql:"type:varchar(50);"`
}

type CampaignEntity struct {
	db *gorm.DB
}

func NewCampaignEntity(db *gorm.DB) CampaignEntity {
	return CampaignEntity{db}
}

func (h *CampaignEntity) GetCampaignByCategory(id int) []Campaign {
	var Campaigns []Campaign
	if err := h.db.Where("product_category_id = ?", id).Find(&Campaigns).Error; err != nil {
		log.Panic(err)
	}

	return Campaigns
}

func (h *CampaignEntity) DeleteCampaign(id int) (c Campaign) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Delete(&c)
	}
	return
}

func (h *CampaignEntity) UpdateCampaign(ci Campaign) (c Campaign) {
	if err := h.db.Where("id = ?", ci.ID).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Save(&ci)
	}
	return ci
}

func (h *CampaignEntity) CreateCampaign(ci Campaign) (co Campaign) {
	if err := h.db.Create(&ci).Error; err != nil {
		log.Panic(err)
	}
	return ci
}

func (h *CampaignEntity) GetCampaignByID(id int) (c Campaign) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}

func (h *CampaignEntity) GetCampaign() (c []Campaign) {
	if err := h.db.Find(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}
