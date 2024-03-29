package models

import (
	"time"

	"gorm.io/gorm"
)

type Battle struct {
	ID         uint           `gorm:"primarykey"                     json:"id"`
	MonsterAID uint           `gorm:"column:monster_a"               json:"-"`
	MonsterA   Monster        `gorm:"not null;foreignKey:MonsterAID" json:"monsterA" validate:"required"`
	MonsterBID uint           `gorm:"column:monster_b"               json:"-"`
	MonsterB   Monster        `gorm:"not null;foreignKey:MonsterBID" json:"monsterB" validate:"required"`
	WinnerID   uint           `gorm:"column:winner"                  json:"-"`
	Winner     Monster        `gorm:"not null;foreignKey:WinnerID"   json:"winner"   validate:"required"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (b *Battle) SetWinner() {
	var winner Monster

	fAttack, sAttack := getFirstAndSecondAttacker(b.MonsterA, b.MonsterB)
	stopBattle := false
	fDamage := fAttack.CalculateDamage()
	sDamage := sAttack.CalculateDamage()

	for !stopBattle {
		sAttack.BeAttacked(fDamage)
		fAttack.BeAttacked(sDamage)

		if fAttack.Hp == 0 {
			winner = b.MonsterB
			stopBattle = true
		}

		if sAttack.Hp == 0 {
			winner = b.MonsterA
			stopBattle = true
		}
	}

	b.Winner = winner
	b.WinnerID = winner.ID
}

func getFirstAndSecondAttacker(monsterA, monsterB Monster) (Monster, Monster) {
	if monsterA.Speed == monsterB.Speed {
		return getHigherAndLowerAttacker(monsterA, monsterB)
	}

	if monsterA.Speed > monsterB.Speed {
		return monsterA, monsterB
	}

	return monsterB, monsterA
}

func getHigherAndLowerAttacker(monsterA, monsterB Monster) (Monster, Monster) {
	if monsterA.Attack > monsterB.Attack {
		return monsterA, monsterB
	}

	return monsterB, monsterA
}
