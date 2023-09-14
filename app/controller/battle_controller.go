package controller

import (
	"battle-of-monsters/app/db"
	"battle-of-monsters/app/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListBattles(context *gin.Context) {
	var battle []models.Battle

	var result *gorm.DB

	if result = db.CONN.Find(&battle); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}

	log.Printf("Found %v battles", result.RowsAffected)
	context.JSON(http.StatusOK, &battle)
}

func Battle(context *gin.Context) {
	type MonsterRequest struct {
		ID       uint   `binding:"required"               json:"id"`
		Name     string `binding:"required,gte=1,lte=255" json:"name"`
		Attack   uint   `binding:"required"               json:"attack"`
		Defense  uint   `binding:"required"               json:"defense"`
		Hp       uint   `binding:"required"               json:"hp"`
		Speed    uint   `binding:"required"               json:"speed"`
		ImageURL string `binding:"required,gte=1,lte=255" json:"imageUrl"`
	}

	var battleRequest struct {
		MonsterA MonsterRequest `binding:"required" json:"monsterA"`
		MonsterB MonsterRequest `binding:"required" json:"monsterB"`
	}

	var monsterA models.Monster

	var monsterB models.Monster

	if err := context.ShouldBindJSON(&battleRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if result := db.CONN.First(&monsterA, battleRequest.MonsterA.ID); result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "monster A not found"})
	}

	if result := db.CONN.First(&monsterB, battleRequest.MonsterB.ID); result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "monster B not found"})
	}

	battle := models.Battle{
		MonsterAID: monsterA.ID,
		MonsterBID: monsterB.ID,
		MonsterA:   monsterA,
		MonsterB:   monsterB,
	}

	battle.SetWinner()

	if result := db.CONN.Save(&battle); result.Error != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": result.Error})
	}

	context.JSON(http.StatusCreated, &battle)
}

func DeleteBattle(context *gin.Context) {
	battleID := context.Param("battleID")

	var battle models.Battle

	if result := db.CONN.First(&battle, battleID); result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
	}

	if result := db.CONN.Delete(&models.Battle{}, battleID); result.Error != nil {
		context.Status(http.StatusBadRequest)
	}

	context.Status(http.StatusNoContent)
}
