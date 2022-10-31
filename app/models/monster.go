package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

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

func (m *Monster) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ID       uint     `json:"id"`
		Name     string   `json:"name"`
		Attack   uint     `json:"attack"`
		Defense  uint     `json:"defense"`
		Hp       uint     `json:"hp"`
		Speed    uint     `json:"speed"`
		ImageURL string   `json:"imageUrl"`
		Battles  []Battle `json:"battles"`
	}{
		m.ID, m.Name, m.Attack, m.Defense, m.Hp, m.Speed, m.ImageURL, m.Battles,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Monster. %w", err)
	}

	return j, nil
}
