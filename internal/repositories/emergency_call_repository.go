package repositories

import (
	"your-project/internal/models"

	"gorm.io/gorm"
)

type EmergencyCallRepository interface {
	Create(call *models.EmergencyCall) error
	GetAll() ([]models.EmergencyCall, error)
	GetByID(id uint) (*models.EmergencyCall, error)
	Update(call *models.EmergencyCall) error
	Delete(id uint) error
}

type emergencyCallRepository struct {
	db *gorm.DB
}

func NewEmergencyCallRepository(db *gorm.DB) EmergencyCallRepository {
	return &emergencyCallRepository{db: db}
}

func (r *emergencyCallRepository) Create(call *models.EmergencyCall) error {
	return r.db.Create(call).Error
}

func (r *emergencyCallRepository) GetAll() ([]models.EmergencyCall, error) {
	var calls []models.EmergencyCall
	err := r.db.Preload("User").Find(&calls).Error
	return calls, err
}

func (r *emergencyCallRepository) GetByID(id uint) (*models.EmergencyCall, error) {
	var call models.EmergencyCall
	err := r.db.Preload("User").First(&call, id).Error
	return &call, err
}

func (r *emergencyCallRepository) Update(call *models.EmergencyCall) error {
	return r.db.Save(call).Error
}

func (r *emergencyCallRepository) Delete(id uint) error {
	return r.db.Delete(&models.EmergencyCall{}, id).Error
}
