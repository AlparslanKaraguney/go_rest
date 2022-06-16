package user

// DB logic for user

import (
	"gorm.io/gorm"
)

type Data interface {
	GetAll() ([]Model, error)
	GetByID(id uint) (*Model, error)
	Create(model Model) (uint, error)
	// Update(id uint, model Model) (Model, error)
	// Delete(id uint) error
	Migrate() error
}

type data struct {
	db *gorm.DB
}

// Compiled time check for interface implementation
var _ Data = data{}

func NewData(db *gorm.DB) Data {
	return data{db: db}
}

func (d data) GetAll() ([]Model, error) {
	var models []Model
	err := d.db.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (d data) GetByID(id uint) (*Model, error) {
	model := &Model{ID: id}
	err := d.db.First(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (d data) Create(model Model) (uint, error) {
	err := d.db.Create(&model).Error
	if err != nil {
		return 0, err
	}
	return model.ID, nil
}

func (d data) Migrate() error {
	return d.db.AutoMigrate(&Model{})
}
