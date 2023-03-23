package repository

import (
	"errors"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/infrastructure"
	"gorm.io/gorm"
)

type IPondRepository interface {
	CreatePond(id string, pond *domain.Pond) (domain.Pond, error)
	GetPondByName(pondName string) (domain.Pond, error)
	GetPondById(id string) (domain.Pond, error)
	UpdatePond(pond *domain.Pond) error
	DeletePond(id string) error
	GetAll() ([]domain.Pond, error)
}

type PondRepository struct {
	db *gorm.DB
}

func NewPondRepository() IPondRepository {
	return &PondRepository{
		db: infrastructure.GetDb(),
	}
}

func (r *PondRepository) CreatePond(idFarm string, pond *domain.Pond) (domain.Pond, error) {
	pond.FarmID = idFarm
	err := r.db.Save(pond).Error
	if err != nil {
		return domain.Pond{}, err
	}
	return *pond, nil
}

func (r *PondRepository) GetPondByName(pondName string) (domain.Pond, error) {
	var container domain.Pond
	err := r.db.Preload("Farm").Where("name = ?", pondName).Take(&container).Limit(1).Error
	if err != nil {
		return domain.Pond{}, err
	}
	return container, nil
}

func (r *PondRepository) GetPondById(id string) (domain.Pond, error) {
	var container domain.Pond
	err := r.db.Preload("Farm").Where("id = ?", id).Take(&container).Error
	if err != nil {
		return domain.Pond{}, err
	}
	if container.Farm.ID == 0 && container.Farm.FarmName == "" {
		return domain.Pond{}, errors.New("farm already removed, pond doens't exist")
	}
	return container, nil
}

func (r *PondRepository) UpdatePond(pond *domain.Pond) error {
	err := r.db.Where("id = ?", pond.ID).Updates(pond).Error
	return err
}

func (r *PondRepository) DeletePond(id string) error {
	err := r.db.Where("id = ?", id).Delete(&domain.Pond{})
	return err.Error
}

func (r *PondRepository) GetAll() ([]domain.Pond, error) {
	var container []domain.Pond
	err := r.db.Preload("Farm").Find(&container).Error
	if err != nil {
		return []domain.Pond{}, err
	}
	var ponds []domain.Pond
	// exclude ponds with farm that has been removed
	for _, data := range container {
		if data.Farm.ID != 0 && data.Farm.FarmName != "" {
			ponds = append(ponds, data)
		}
	}
	return ponds, nil
}
