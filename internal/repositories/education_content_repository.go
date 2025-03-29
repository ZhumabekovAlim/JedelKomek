package repositories

import (
	"your-project/internal/models"

	"gorm.io/gorm"
)

type EducationContentRepository interface {
	Create(content *models.EducationContent) error
	GetAll() ([]models.EducationContent, error)
	GetByID(id uint) (*models.EducationContent, error)
	Update(content *models.EducationContent) error
	Delete(id uint) error
}

type educationContentRepository struct {
	db *gorm.DB
}

func NewEducationContentRepository(db *gorm.DB) EducationContentRepository {
	return &educationContentRepository{db: db}
}

func (r *educationContentRepository) Create(content *models.EducationContent) error {
	return r.db.Create(content).Error
}

func (r *educationContentRepository) GetAll() ([]models.EducationContent, error) {
	var contents []models.EducationContent
	err := r.db.Preload("Author").Find(&contents).Error
	return contents, err
}

func (r *educationContentRepository) GetByID(id uint) (*models.EducationContent, error) {
	var content models.EducationContent
	err := r.db.Preload("Author").First(&content, id).Error
	return &content, err
}

func (r *educationContentRepository) Update(content *models.EducationContent) error {
	return r.db.Save(content).Error
}

func (r *educationContentRepository) Delete(id uint) error {
	return r.db.Delete(&models.EducationContent{}, id).Error
}
