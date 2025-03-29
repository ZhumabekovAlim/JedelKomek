package repositories

import (
	"your-project/internal/models"

	"gorm.io/gorm"
)

type PoliceDepartmentRepository interface {
	Create(dept *models.PoliceDepartment) error
	GetAll() ([]models.PoliceDepartment, error)
	GetByID(id uint) (*models.PoliceDepartment, error)
	Update(dept *models.PoliceDepartment) error
	Delete(id uint) error
}

type policeDepartmentRepository struct {
	db *gorm.DB
}

func NewPoliceDepartmentRepository(db *gorm.DB) PoliceDepartmentRepository {
	return &policeDepartmentRepository{db: db}
}

func (r *policeDepartmentRepository) Create(dept *models.PoliceDepartment) error {
	return r.db.Create(dept).Error
}

func (r *policeDepartmentRepository) GetAll() ([]models.PoliceDepartment, error) {
	var depts []models.PoliceDepartment
	err := r.db.Find(&depts).Error
	return depts, err
}

func (r *policeDepartmentRepository) GetByID(id uint) (*models.PoliceDepartment, error) {
	var dept models.PoliceDepartment
	err := r.db.First(&dept, id).Error
	return &dept, err
}

func (r *policeDepartmentRepository) Update(dept *models.PoliceDepartment) error {
	return r.db.Save(dept).Error
}

func (r *policeDepartmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.PoliceDepartment{}, id).Error
}
