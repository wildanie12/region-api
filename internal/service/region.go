package service

import (
	"context"
	"fmt"

	"github.com/wildanie12/region-api/internal/entity"
	"github.com/wildanie12/region-api/internal/repository"
)

// RegionService main implementation object.
type RegionService struct {
	regionRepository repository.RegionInterface
}

// NewRegion contructs instance of region service.
func NewRegion(rr repository.RegionInterface) *RegionService {
	return &RegionService{
		regionRepository: rr,
	}
}

// ListProvince return list of provinces.
func (s *RegionService) ListProvince(ctx context.Context) ([]entity.Province, error) {
	// get list of provinces ordered by name
	data, err := s.regionRepository.ListProvince(
		ctx,
		[]entity.ListFilter{},
		[]entity.ListSort{{Column: "province_name", Order: entity.SortAsc}},
		entity.Preload{},
	)
	if err != nil {
		return data, s.makeError(err)
	}
	return data, nil
}

// ListRegency return list of regencies.
func (s *RegionService) ListRegency(ctx context.Context, provinceID uint) ([]entity.Regency, error) {
	// get list of regencies ordered by name
	data, err := s.regionRepository.ListRegency(
		ctx,
		[]entity.ListFilter{{Column: "province_id", Operator: "=", Value: provinceID}},
		[]entity.ListSort{{Column: "regency_name", Order: entity.SortAsc}},
		entity.Preload{},
	)
	if err != nil {
		return data, s.makeError(err)
	}
	return data, nil
}

func (s *RegionService) makeError(err error) error {
	return fmt.Errorf("(x) region service err: %v", err)
}
