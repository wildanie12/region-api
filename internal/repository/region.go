package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/wildanie12/region-api/internal/entity"
	"gorm.io/gorm"
)

// RegionInterface holds method must be implemented for it to be RegionRepository.
type RegionInterface interface {
	// ListProvince return list of data province.
	ListProvince(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.Province, error)
	// ListRegency return list of data province.
	ListRegency(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.Regency, error)
	// ListDistrict return list of data province.
	ListDistrict(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.District, error)
	// ListVillage return list of data province.
	ListVillage(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.Village, error)
	// JoinedSearch returned list of villages that can be searched by province, region, district and villages.
	JoinedSearch(ctx context.Context, search string, limit int) ([]entity.Village, error)
}

// regionRepository is a concrete implementation of RegionRepository.
type regionRepository struct {
	db *gorm.DB
}

// NewRegion constructs RegionRepository.
func NewRegion(db *gorm.DB) RegionInterface {
	return &regionRepository{
		db: db,
	}
}

// ListProvince return list of data province.
func (r *regionRepository) ListProvince(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.Province, error) {
	data := []entity.Province{}
	tx := r.db.Model(entity.Province{})

	// apply filter
	for _, filter := range filters {
		tx = tx.Where(fmt.Sprintf("%s %s ?", filter.Column, filter.Operator), filter.Value)
	}
	// apply sort
	for _, sort := range sorts {
		tx = tx.Order(fmt.Sprintf("%s %s", sort.Column, sort.Order))
	}
	// apply preload
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// execute query
	tx = tx.Find(&data)
	if tx.Error != nil {
		return data, tx.Error
	}
	return data, nil
}

// ListRegency return list of data province.
func (r *regionRepository) ListRegency(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.Regency, error) {
	data := []entity.Regency{}
	tx := r.db.Model(entity.Regency{})

	// apply filter
	for _, filter := range filters {
		tx = tx.Where(fmt.Sprintf("%s %s ?", filter.Column, filter.Operator), filter.Value)
	}
	// apply sort
	for _, sort := range sorts {
		tx = tx.Order(fmt.Sprintf("%s %s", sort.Column, sort.Order))
	}
	// apply preload
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// execute query
	tx = tx.Find(&data)
	if tx.Error != nil {
		return data, tx.Error
	}
	return data, nil
}

// ListDistrict return list of data province.
func (r *regionRepository) ListDistrict(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.District, error) {
	data := []entity.District{}
	tx := r.db.Model(entity.District{})

	// apply filter
	for _, filter := range filters {
		tx = tx.Where(fmt.Sprintf("%s %s ?", filter.Column, filter.Operator), filter.Value)
	}
	// apply sort
	for _, sort := range sorts {
		tx = tx.Order(fmt.Sprintf("%s %s", sort.Column, sort.Order))
	}
	// apply preload
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// execute query
	tx = tx.Find(&data)
	if tx.Error != nil {
		return data, tx.Error
	}
	return data, nil
}

// ListVillage return list of data province.
func (r *regionRepository) ListVillage(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.Village, error) {
	data := []entity.Village{}
	tx := r.db.Model(entity.Village{})

	// apply filter
	for _, filter := range filters {
		tx = tx.Where(fmt.Sprintf("%s %s ?", filter.Column, filter.Operator), filter.Value)
	}
	// apply sort
	for _, sort := range sorts {
		tx = tx.Order(fmt.Sprintf("%s %s", sort.Column, sort.Order))
	}
	// apply preload
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// execute query
	tx = tx.Find(&data)
	if tx.Error != nil {
		return data, tx.Error
	}
	return data, nil
}

// JoinedSearch returned list of villages that can be searched by province, region, district and villages.
// the returned data is already preloaded up to province entity.
func (r *regionRepository) JoinedSearch(ctx context.Context, search string, limit int) ([]entity.Village, error) {
	data := []entity.Village{}

	inputArr := strings.Split(search, ",")

	// build query
	tx := r.db.Model(&entity.Village{}).
		Joins("JOIN districts ON district_id = districts.id").
		Joins("JOIN regencies ON regency_id = regencies.id").
		Joins("JOIN provinces ON province_id = provinces.id")

	switch len(inputArr) {
	case 1:
		fmt.Println("case 1")
		tx = tx.Or("villages.village_name LIKE ?", "%"+search+"%").
			Or("districts.district_name LIKE ?", "%"+search+"%").
			Or("regencies.regency_name LIKE ?", "%"+search+"%").
			Or("provinces.province_name LIKE ?", "%"+search+"%")
	case 2:
		fmt.Println("case 2")
		province := strings.TrimSpace(inputArr[0])
		regency := strings.TrimSpace(inputArr[1])
		tx = tx.Where("provinces.province_name LIKE ?", "%"+province+"%").
			Where("regencies.regency_name LIKE ?", "%"+regency+"%")
		tx = tx.Group("regency_name")
	case 3:
		fmt.Println("case 3")
		province := strings.TrimSpace(inputArr[0])
		regency := strings.TrimSpace(inputArr[1])
		district := strings.TrimSpace(inputArr[2])
		tx = tx.Where("provinces.province_name LIKE ?", "%"+province+"%").
			Where("regencies.regency_name LIKE ?", "%"+regency+"%").
			Where("districts.district_name LIKE ?", "%"+district+"%")
	case 4:
		fmt.Println("case 4")
		province := strings.TrimSpace(inputArr[0])
		regency := strings.TrimSpace(inputArr[1])
		district := strings.TrimSpace(inputArr[2])
		village := strings.TrimSpace(inputArr[3])
		tx = tx.Where("provinces.province_name LIKE ?", "%"+province+"%").
			Where("regencies.regency_name LIKE ?", "%"+regency+"%").
			Where("districts.district_name LIKE ?", "%"+district+"%").
			Where("villages.village_name LIKE ?", "%"+village+"%")
	}

	tx = tx.Preload("District").Preload("District.Regency").Preload("District.Regency.Province")
	// find data
	err := tx.Limit(limit).Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
