package models

import "time"

type Recipes struct {
	ID          uint          `json:"id" gorm:"primary_key"`
	Title       string        `json:"title"`
	Author      string        `json:"author"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	UserId      int           `json:"userId"`
	User        Users         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Ingredients []Ingredients `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category    []Categories  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
