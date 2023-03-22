package repository

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/infrastructure"
	"gorm.io/gorm"
)

type IFarmRepository interface {
	CreateFarm(farm *domain.Farm) (domain.Farm, error)
	GetFarmByName(name string) (domain.Farm, error)
	GetFarmById(id string) (domain.Farm, error)
	UpdateFarm(farm *domain.Farm) error
}

type FarmRepository struct {
	db *gorm.DB
}

func NewFarmRepository() IFarmRepository {
	return &FarmRepository{
		db: infrastructure.GetDb(),
	}
}

func (r *FarmRepository) CreateFarm(data *domain.Farm) (domain.Farm, error) {
	err := r.db.Save(data).Error
	if err != nil {
		return domain.Farm{}, err
	}
	return *data, nil
}

func (r *FarmRepository) GetFarmByName(name string) (domain.Farm, error) {
	var container domain.Farm
	result := r.db.Where("farm_name = ?", name).Take(&container).Limit(1)
	if result.Error != nil {
		return domain.Farm{}, result.Error
	}
	return container, nil
}

func (r *FarmRepository) GetFarmById(id string) (domain.Farm, error) {
	var container domain.Farm
	result := r.db.Where("id = ?", id).Take(&container)
	if result.Error != nil {
		return domain.Farm{}, result.Error
	}
	return container, result.Error
}

func (r *FarmRepository) UpdateFarm(farm *domain.Farm) error {
	err := r.db.Save(farm).Error
	return err
}
