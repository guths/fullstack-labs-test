package models

import (
	"errors"
	"reflect"
	"time"

	"gorm.io/gorm"
)

var err = errors.New("passed columns do not match the model")

type Monster struct {
	ID        uint           `gorm:"primarykey"        json:"id"`
	Name      string         `gorm:"not null;size:255" json:"name"     validate:"required"`
	Attack    uint           `gorm:"not null"          json:"attack"   validate:"required"`
	Defense   uint           `gorm:"not null"          json:"defense"  validate:"required"`
	Hp        uint           `gorm:"not null"          json:"hp"       validate:"gte=0"`
	Speed     uint           `gorm:"not null"          json:"speed"    validate:"required"`
	ImageURL  string         `gorm:"not null"          json:"imageUrl"`
	Battles   []Battle       `gorm:"foreignKey:ID"     json:"battles"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"             json:"-"`
}

func (m Monster) CalculateDamage() int {
	if m.Attack <= m.Defense {
		return 1
	}

	return int(m.Attack) - int(m.Defense)
}

func (m *Monster) BeAttacked(damage int) {
	m.Hp = uint(int(m.Hp) - damage)
}

func (m Monster) VerifyColumnsInModel(column []string) error {
	structType := reflect.TypeOf(m)
	sum := 0

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		name := field.Tag.Get("json")

		for _, v := range column {
			if name == v {
				sum++
			}
		}
	}

	if sum != len(column) {
		return err
	}

	return nil
}
