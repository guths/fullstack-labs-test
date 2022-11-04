package models_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	utilstests "battle-of-monsters/app/tests/utils"
)

var _ = Describe("Battle", func() {
	utilstests.LoadEnv()
	db.Connect()

	BeforeEach(func() {
		if err := db.CONN.Exec("DELETE FROM battles; DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete battle and monsters. %w", err))
		}
	})

	AfterEach(func() {
		if err := db.CONN.Exec("DELETE FROM battles; DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete battle and monsters. %w", err))
		}
	})

	Describe("Marshal", func() {

		var b []byte
		var expected []byte

		JustBeforeEach(func() {
			blueSnake := models.Monster{
				Name:     "Blue Snake",
				Attack:   10,
				Defense:  15,
				Hp:       8,
				Speed:    18,
				ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png",
			}

			redUnicorn := models.Monster{
				Name:     "Red Unicorn",
				Attack:   12,
				Defense:  14,
				Hp:       10,
				Speed:    9,
				ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-unicorn.png",
			}

			db.CONN.Debug().Create(&blueSnake)
			db.CONN.Debug().Create(&redUnicorn)

			battle := models.Battle{
				MonsterA: models.Monster{ID: blueSnake.ID},
				MonsterB: models.Monster{ID: redUnicorn.ID},
				Winner:   models.Monster{ID: blueSnake.ID},
			}

			db.CONN.Debug().Create(&battle)

			db.CONN.Preload("Winner").Preload("MonsterA").Preload("MonsterB").First(&battle, battle.ID)

			b, _ = battle.MarshalJSON()

			expected = []byte(`{
				"id": 1,
				"monsterA": {
					"id": 1,
          "name": "Blue Snake",
          "attack": 10,
          "defense": 15,
          "hp": 8,
          "speed": 18,
          "imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png",
          "battles": null
				},
				"monsterB": {
					"id": 2,
          "name": "Red Unicorn",
          "attack": 12,
          "defense": 14,
          "hp": 10,
          "speed": 9,
          "imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/red-unicorn.png",
          "battles": null
				},
				"winner": {
					"id": 1,
          "name": "Blue Snake",
          "attack": 10,
          "defense": 15,
          "hp": 8,
          "speed": 18,
          "imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/blue-snake.png",
          "battles": null
				}
			}`)

		})

		Context("should parse battle correctly", func() {

			It("battle should match with the expected json", func() {
				Expect(b).Should(MatchJSON(expected))
			})

		})

	})

})
