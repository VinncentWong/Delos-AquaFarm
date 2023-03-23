package repository

import (
	"fmt"

	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/domain/response"
	"github.com/VinncentWong/Delos-AquaFarm/infrastructure"
	"gorm.io/gorm"
)

type IRecordRepository interface {
	GetRecord() ([]response.RecordResponse, error)
	GetUniqueAgent(endpoint string) (int, error)
}

type RecordRepository struct {
	db *gorm.DB
}

func NewRecordRepository() IRecordRepository {
	return &RecordRepository{
		db: infrastructure.GetDb(),
	}
}

func (r *RecordRepository) GetRecord() ([]response.RecordResponse, error) {
	var container []response.RecordResponse
	err := r.db.
		Model(&response.RecordResponse{}).
		Raw("SELECT method_name, endpoint, SUM(count) AS count FROM record_apis GROUP BY method_name,endpoint").
		Scan(&container).
		Error
	if err != nil {
		return []response.RecordResponse{}, err
	}
	return container, nil
}

func (r *RecordRepository) GetUniqueAgent(endpoint string) (int, error) {
	var container int
	err := r.db.Model(&domain.RecordApi{}).
		Raw(fmt.Sprintf("SELECT DISTINCT COUNT(ip_address) FROM record_apis WHERE endpoint='%s'", endpoint)).
		Scan(&container).
		Error
	if err != nil {
		return -1, err
	}
	return container, nil
}
