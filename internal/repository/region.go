package repository

import (
	"context"
	"fmt"

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
	ListDistrict(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.Regency, error)
	// ListVillage return list of data province.
	ListVillage(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.Village, error)
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
func (r *regionRepository) ListDistrict(ctx context.Context, filters []entity.ListFilter, sorts []entity.ListSort, preloads entity.Preload) ([]entity.Regency, error) {
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
