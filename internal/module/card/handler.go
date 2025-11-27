package card

import (
	"net/http"
	"strconv"

	"app/internal/entity"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) {
	var card entity.Card

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateCard(&card); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save card"})
		return
	}

	c.JSON(http.StatusCreated, card)
}

func (h *Handler) GetAll(c *gin.Context) {
	cards, err := h.service.GetCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cards"})
		return
	}

	c.JSON(http.StatusOK, cards)
}

func (h *Handler) Modify(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var input entity.Card
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCard, err := h.service.ModifyCard(id, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCard)
}
