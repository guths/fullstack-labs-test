package models_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	utilstests "battle-of-monsters/app/tests/utils"
)

var _ = Describe("Monster", func() {
	utilstests.LoadEnv()
	db.Connect()

	BeforeEach(func() {
		if err := db.CONN.Exec("DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete monsters. %w", err))
		}
	})

	AfterEach(func() {
		if err := db.CONN.Exec("DELETE FROM monsters;").Error; err != nil {
			panic(fmt.Errorf("failed to delete monsters. %w", err))
		}
	})

	Describe("Marshal", func() {

		var darkSnake *models.Monster
		var m []byte
		var expected []byte

		JustBeforeEach(func() {
			darkSnake = &models.Monster{
				Name:     "Dark Snake",
				Attack:   10,
				Defense:  15,
				Hp:       8,
				Speed:    18,
				ImageURL: "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/dark-snake.png",
			}

			db.CONN.Create(darkSnake)

			m, _ = darkSnake.MarshalJSON()

			expected = []byte(`{
				"id": 1,
				"name": "Dark Snake",
				"attack": 10,
				"defense": 15,
				"hp": 8,
				"speed": 18,
				"imageUrl": "https://fsl-assessment-public-files.s3.amazonaws.com/assessment-cc-01/dark-snake.png",
				"battles": null
			}`)

		})

		Context("should parse moster correctly", func() {

			It("monster should match with the expected json", func() {
				Expect(m).Should(MatchJSON(expected))
			})

		})

	})

})
