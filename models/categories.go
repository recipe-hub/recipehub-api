package models

import "time"

type Categories struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
