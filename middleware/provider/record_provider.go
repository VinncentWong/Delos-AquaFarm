package provider

import (
	"github.com/VinncentWong/Delos-AquaFarm/domain"
	"github.com/VinncentWong/Delos-AquaFarm/infrastructure"
	"gorm.io/gorm"
)

type IRecordProvider interface {
	SaveRecord(record *domain.RecordApi) error
	GetRecord(ipAddress string, endpoint string) (domain.RecordApi, error)
}

type RecordProvider struct {
	db *gorm.DB
}

func NewRecordProvider() IRecordProvider {
	return &RecordProvider{
		db: infrastructure.GetDb(),
	}
}

func (p *RecordProvider) SaveRecord(record *domain.RecordApi) error {
	result, err := p.GetRecord(record.IpAddress, record.Endpoint)
	if err != nil {
		record.Count++
		err = p.db.Save(record).Error
		return err
	} else {
		result.Count++
		err = p.db.Model(&domain.RecordApi{}).Where("ip_address = ? AND endpoint = ?", result.IpAddress, result.Endpoint).Updates(&result).Error
		return err
	}
}

func (p *RecordProvider) GetRecord(ipAddress string, endpoint string) (domain.RecordApi, error) {
	var container domain.RecordApi
	err := p.db.Where("ip_address = ? AND endpoint = ?", ipAddress, endpoint).Take(&container).Limit(1).Error
	if err != nil {
		return domain.RecordApi{}, err
	}
	return container, nil
}
