package handlers

import (
	"doramaPro/models"
	"doramaPro/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DramaHandler struct {
	service *services.DramaService
}

func NewDramaHandler(service *services.DramaService) *DramaHandler {
	return &DramaHandler{service: service}
}

func (h *DramaHandler) GetDramas(c *gin.Context) {
	dramas, err := h.service.GetAllDramas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dramas)
}

func (h *DramaHandler) CreateDrama(c *gin.Context) {
	var input struct {
		Title         string  `json:"title" binding:"required"`
		OriginalTitle string  `json:"original_title"`
		Country       string  `json:"country"`
		ReleaseYear   int     `json:"release_year"`
		Episodes      int     `json:"episodes"`
		Status        string  `json:"status"`
		Description   string  `json:"description"`
		Rating        float64 `json:"rating"`
		Genres        []int   `json:"genres"` // Массив ID жанров
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	drama, err := h.service.CreateDrama(
		input.Title,
		input.OriginalTitle,
		input.Country,
		input.ReleaseYear,
		input.Episodes,
		input.Status,
		input.Description,
		input.Rating,
		input.Genres,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, drama)
}

func (h *DramaHandler) UpdateDrama(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID драмы"})
		return
	}

	var updatedDrama models.Drama
	if err := c.ShouldBindJSON(&updatedDrama); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Преобразуем жанры в массив ID
	genreIDs := updatedDrama.Genres // Теперь передаем массив ID жанров, а не самих жанров

	var genres []models.Genre
	// Получаем жанры по ID
	err = h.service.GetGenresByIDs(genreIDs, &genres)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения жанров"})
		return
	}

	updatedDrama.Genres = genres

	drama, err := h.service.UpdateDrama(uint(parsedID), &updatedDrama)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Драма не найдена"})
		return
	}

	c.JSON(http.StatusOK, drama)
}

func (h *DramaHandler) DeleteDrama(c *gin.Context) {
	id := c.Param("id")

	parsedID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID драмы"})
		return
	}

	err = h.service.DeleteDrama(uint(parsedID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
