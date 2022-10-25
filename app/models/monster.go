package models

import "gorm.io/gorm"

type Monster struct {
	gorm.Model
	Name     string   `gorm:"not null;size:255"`
	Attack   uint     `gorm:"not null"`
	Defense  uint     `gorm:"not null"`
	Hp       uint     `gorm:"not null"`
	Speed    uint     `gorm:"not null"`
	ImageURL string   `gorm:"not null"`
	Battles  []Battle `gorm:"foreignKey:ID"`
}
