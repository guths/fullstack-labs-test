package models

import "gorm.io/gorm"

type Battle struct {
	gorm.Model
	MonsterA Monster `gorm:"not null;foreignKey:ID"`
	MonsterB Monster `gorm:"not null;foreignKey:ID"`
	Winner   Monster `gorm:"not null;foreignKey:ID"`
}
