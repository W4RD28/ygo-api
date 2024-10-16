package handlers

import (
	"net/http"
	"strconv"

	"github.com/W4RD28/ygo-api/internal/inputs"
	"github.com/W4RD28/ygo-api/internal/models"
	"github.com/W4RD28/ygo-api/internal/services"
	"github.com/gin-gonic/gin"
)

func AddCard(c *gin.Context) {
	var input inputs.CardAddInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card := models.Card{
		Name:        input.Name,
		Race:        input.Race,
		Type:        input.Type,
		Description: input.Description,
		Level:       input.Level,
		Attack:      input.Attack,
		Defense:     input.Defense,
	}

	savedCard, err := services.SaveCard(&card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create card"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedCard})
}

func EditCard(c *gin.Context) {
	var input inputs.CardUpdateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card, err := services.FindCardById(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	card.Name = input.Name
	card.Race = input.Race
	card.Type = input.Type
	card.Description = input.Description
	card.Level = input.Level
	card.Attack = input.Attack
	card.Defense = input.Defense

	updatedCard, err := services.EditCard(card)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update card"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": updatedCard})
}

func DeleteCard(c *gin.Context) {
	var input inputs.CardDeleteInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	card, err := services.FindCardById(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	err = services.DeleteCard(card)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete card"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": card})
}

func FindCards(c *gin.Context) {
	query := c.Query("query")
	cards, err := services.FindCards(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cards"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cards})
}

func FindCardById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}
	card, err := services.FindCardById(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": card})
}
