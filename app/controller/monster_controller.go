package controller

import (
	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	"encoding/csv"
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FetchMonster(context *gin.Context) {
	monsterID := context.Param("monsterID")

	var monster models.Monster
	if result := db.CONN.First(&monster, monsterID); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		} else {
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}

		return
	}

	context.JSON(http.StatusOK, &monster)
}

func CreateMonster(context *gin.Context) {
	var monsterRequest struct {
		Name     string `json:"name" binding:"required,gte=1,lte=255"`
		Attack   uint   `json:"attack" binding:"required"`
		Defense  uint   `json:"defense" binding:"required"`
		Hp       uint   `json:"hp" binding:"required"`
		Speed    uint   `json:"speed" binding:"required"`
		ImageURL string `json:"imageUrl" binding:"required,gte=1,lte=255"`
	}

	if err := context.BindJSON(&monsterRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	monster := models.Monster{
		Name:     monsterRequest.Name,
		Attack:   monsterRequest.Attack,
		Defense:  monsterRequest.Defense,
		Hp:       monsterRequest.Hp,
		Speed:    monsterRequest.Speed,
		ImageURL: monsterRequest.ImageURL,
	}

	if result := db.CONN.Create(&monster); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})

		return
	}

	log.Printf("monster %v has been created", monster.Name)

	context.JSON(http.StatusCreated, &monster)
}

func UpdateMonster(context *gin.Context) {
	monsterID := context.Param("monsterID")

	var monsterRequest struct {
		Name     string
		Attack   uint
		Defense  uint
		Hp       uint
		Speed    uint
		ImageURL string
	}

	if err := context.BindJSON(&monsterRequest); err != nil {
		log.Fatalln("error parsing monster request")

		return
	}

	var monster models.Monster

	db.CONN.First(&monster, monsterID)

	monster.Name = monsterRequest.Name
	monster.Attack = monsterRequest.Attack
	monster.Defense = monsterRequest.Defense
	monster.Hp = monsterRequest.Hp
	monster.Speed = monsterRequest.Speed
	monster.ImageURL = monsterRequest.ImageURL

	if result := db.CONN.Save(&monster); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})

		return
	}

	context.JSON(http.StatusOK, &monster)
}

func DeleteMonster(context *gin.Context) {
	monsterID := context.Param("monsterID")

	if result := db.CONN.Delete(&models.Monster{}, monsterID); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})

		return
	}

	context.Status(http.StatusNoContent)
}

type csvUploadInput struct {
	CsvFile *multipart.FileHeader `form:"file" binding:"required"`
}

func ImportCSV(context *gin.Context) {
	log.Println("importing monsters CSV")

	var input csvUploadInput
	if err := context.ShouldBind(&input); err != nil {
		context.String(http.StatusBadRequest, "get form error: %s", err.Error())
	}

	file, err := input.CsvFile.Open()
	if err != nil {
		context.String(http.StatusBadRequest, "got error oppening csv: %s", err.Error())
	}

	reader := csv.NewReader(file)

	records, csvErr := reader.ReadAll()
	if csvErr != nil {
		context.String(http.StatusBadRequest, "got error reading csv: %s", err.Error())
	}

	for i := 1; i < len(records); i++ {
		var line = records[i]
		attack, _ := strconv.Atoi(line[1])
		defense, _ := strconv.Atoi(line[2])
		hp, _ := strconv.Atoi(line[3])
		speed, _ := strconv.Atoi(line[4])
		monster := models.Monster{
			Name:     line[0],
			Attack:   uint(attack),
			Defense:  uint(defense),
			Hp:       uint(hp),
			Speed:    uint(speed),
			ImageURL: line[5],
		}

		db.CONN.Create(&monster)
	}

	log.Printf("%v records imported", len(records))
	log.Printf("file %s imported successfully", input.CsvFile.Filename)

	defer file.Close()

	context.Status(http.StatusAccepted)
}
