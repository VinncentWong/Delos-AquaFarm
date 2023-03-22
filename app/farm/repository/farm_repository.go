package repository

import (
	"errors"
	"fmt"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/infrastructure"
	"gorm.io/gorm"
)

type IFarmRepository interface {
	CreateFarm(farm *domain.Farm) (domain.Farm, error)
	GetFarmByName(name string) (domain.Farm, error)
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
	fmt.Println("data.Farmname = ", data.FarmName)
	farm, err := r.GetFarmByName(data.FarmName)
	if err != nil {
		return domain.Farm{}, nil
	}
	if len(farm.FarmName) != 0 {
		return domain.Farm{}, errors.New("farm already exist in database, a duplicate should be denied")
	}
	err = r.db.Save(data).Error
	if err != nil {
		return domain.Farm{}, err
	}
	return *data, nil
}

func (r *FarmRepository) GetFarmByName(name string) (domain.Farm, error) {
	var container domain.Farm
	result := r.db.Where("farm_name = ?", name).Take(&container)
	if result.Error != nil {
		return domain.Farm{}, nil
	}
	return container, nil
}
