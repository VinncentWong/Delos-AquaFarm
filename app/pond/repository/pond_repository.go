package repository

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/infrastructure"
	"gorm.io/gorm"
)

type IPondRepository interface {
	CreatePond(id string, pond *domain.Pond) (domain.Pond, error)
	GetPondByName(pondName string) (domain.Pond, error)
	GetPondById(id string) (domain.Pond, error)
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
	err := r.db.Where("name = ?", pondName).Take(&container).Limit(1).Error
	if err != nil {
		return domain.Pond{}, err
	}
	return container, nil
}

func (r *PondRepository) GetPondById(id string) (domain.Pond, error) {
	var container domain.Pond
	err := r.db.Where("id = ?", id).Take(&container).Error
	if err != nil {
		return domain.Pond{}, err
	}
	return container, nil
}
