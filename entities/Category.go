package entities

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Title            string
	ParentCategory   *Category
	ParentCategoryID int
}

type CategoryEntity struct {
	db *gorm.DB
}

func NewCategoryEntity(db *gorm.DB) CategoryEntity {
	return CategoryEntity{db}
}

func (h *CategoryEntity) DeleteCategory(id int) (c Category) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Delete(&c)
	}
	return
}

func (h *CategoryEntity) UpdateCategory(ci Category) (c Category) {
	if err := h.db.Where("id = ?", ci.ID).First(&c).Error; err != nil {
		log.Panic(err)
	} else {
		h.db.Save(&ci)
	}
	return ci
}

func (h *CategoryEntity) CreateCategory(ci Category) (co Category) {
	if err := h.db.Create(&ci).Error; err != nil {
		log.Panic(err)
	}
	return ci
}

func (h *CategoryEntity) GetCategoryByID(id int) (c Category) {
	if err := h.db.Where("id = ?", id).First(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}

func (h *CategoryEntity) GetCategory() (c []Category) {
	if err := h.db.Find(&c).Error; err != nil {
		log.Panic(err)
	}
	return
}
