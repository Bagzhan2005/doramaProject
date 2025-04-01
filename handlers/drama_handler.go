package handlers

import (
	"doramaPro/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type DramaHandler struct {
	db *gorm.DB
}

func NewDramaHandler(db *gorm.DB) *DramaHandler {
	return &DramaHandler{db: db}
}

// Барлық драмаларды алу
func (h *DramaHandler) GetDramas(c *gin.Context) {
	var dramas []models.Drama
	if err := h.db.Preload("Genres").Find(&dramas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dramas)
}

// Жаңа драма қосу
func (h *DramaHandler) CreateDrama(c *gin.Context) {
	var input struct {
		Title  string `json:"title"`
		Genres []int  `json:"genres"` // Принимаем ID жанров
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1️⃣ Найдём жанры по их ID
	var genres []models.Genre
	if len(input.Genres) > 0 { // Проверяем, есть ли жанры в запросе
		if err := h.db.Where("id IN ?", input.Genres).Find(&genres).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Genres not found"})
			return
		}
	}

	// 2️⃣ Создаём драму с найденными жанрами
	drama := models.Drama{
		Title:  input.Title,
		Genres: genres, // Присваиваем найденные жанры
	}

	// 3️⃣ Сохраняем драму в базе
	if err := h.db.Create(&drama).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, drama)
}

// Драманы жаңарту
func (h *DramaHandler) UpdateDrama(c *gin.Context) {
	id := c.Param("id")
	var drama models.Drama

	if err := h.db.First(&drama, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Drama not found"})
		return
	}

	if err := c.ShouldBindJSON(&drama); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.Save(&drama)
	c.JSON(http.StatusOK, drama)
}

// Драманы жою
func (h *DramaHandler) DeleteDrama(c *gin.Context) {
	id := c.Param("id")
	if err := h.db.Delete(&models.Drama{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
