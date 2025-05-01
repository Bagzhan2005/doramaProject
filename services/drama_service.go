package services

import (
	"doramaPro/models"
	"gorm.io/gorm"
)

type DramaService struct {
	db *gorm.DB
}

func NewDramaService(db *gorm.DB) *DramaService {
	return &DramaService{db: db}
}

func (s *DramaService) GetGenresByIDs(ids []int, genres *[]models.Genre) error {
	return s.db.Where("id IN ?", ids).Find(genres).Error
}

func (s *DramaService) CreateDrama(title, originalTitle, country string, releaseYear, episodes int, status, description string, rating float64, genreIDs []int) (models.Drama, error) {
	var genres []models.Genre
	s.db.Where("id IN ?", genreIDs).Find(&genres)

	drama := models.Drama{
		Title:         title,
		OriginalTitle: originalTitle,
		Country:       country,
		ReleaseYear:   releaseYear,
		Episodes:      episodes,
		Status:        status,
		Description:   description,
		Rating:        rating,
		Genres:        genres,
	}

	err := s.db.Create(&drama).Error
	return drama, err
}

func (s *DramaService) GetAllDramas() ([]models.Drama, error) {
	var dramas []models.Drama
	err := s.db.Preload("Genres").Find(&dramas).Error
	return dramas, err
}

func (s *DramaService) UpdateDrama(id uint, updatedDrama *models.Drama) (models.Drama, error) {
	var drama models.Drama
	if err := s.db.First(&drama, id).Error; err != nil {
		return models.Drama{}, err
	}
	drama.Title = updatedDrama.Title
	drama.OriginalTitle = updatedDrama.OriginalTitle
	drama.Country = updatedDrama.Country
	drama.ReleaseYear = updatedDrama.ReleaseYear
	drama.Episodes = updatedDrama.Episodes
	drama.Status = updatedDrama.Status
	drama.Description = updatedDrama.Description
	drama.Rating = updatedDrama.Rating
	drama.Genres = updatedDrama.Genres

	s.db.Save(&drama)
	return drama, nil
}

func (s *DramaService) DeleteDrama(id uint) error {
	return s.db.Delete(&models.Drama{}, id).Error
}
