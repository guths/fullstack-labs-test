package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Monster struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Name      string         `gorm:"not null;size:255" json:"name"`
	Attack    uint           `gorm:"not null" json:"attack"`
	Defense   uint           `gorm:"not null" json:"defense"`
	Hp        uint           `gorm:"not null" json:"hp"`
	Speed     uint           `gorm:"not null" json:"speed"`
	ImageURL  string         `gorm:"not null" json:"imageUrl"`
	Battles   []Battle       `gorm:"foreignKey:ID" json:"battles"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
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
