package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Battle struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	MonsterAID uint           `gorm:"column:monster_a"`
	MonsterA   Monster        `gorm:"not null;foreignKey:MonsterAID" json:"monsterA"`
	MonsterBID uint           `gorm:"column:monster_b"`
	MonsterB   Monster        `gorm:"not null;foreignKey:MonsterBID" json:"monsterB"`
	WinnerID   uint           `gorm:"column:winner"`
	Winner     Monster        `gorm:"not null;foreignKey:WinnerID" json:"winner"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
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
