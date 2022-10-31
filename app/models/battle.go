package models

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Battle struct {
	gorm.Model
	MonsterA Monster `gorm:"not null;foreignKey:ID"`
	MonsterB Monster `gorm:"not null;foreignKey:ID"`
	Winner   Monster `gorm:"not null;foreignKey:ID"`
}

func (b *Battle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		ID       uint    `json:"id"`
		MonsterA Monster `json:"monsterA"`
		MonsterB Monster `json:"monsterB"`
		Winner   Monster `json:"winner"`
	}{
		b.ID, b.MonsterA, b.MonsterB, b.Winner,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Battle. %w", err)
	}

	return j, nil
}
