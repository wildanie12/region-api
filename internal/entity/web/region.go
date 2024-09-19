package web

// Province mysql region entity.
type Province struct {
	ID           uint    `json:"id"`
	ProvinceName string  `json:"province_name"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	Geocode      string  `json:"geocode"`
}

// Regency mysql region entity.
type Regency struct {
	ID          uint    `json:"id"`
	RegencyName string  `json:"regency_name"`
	ProvinceID  uint    `json:"province_id"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

// District mysql region entity.
type District struct {
	ID           uint    `json:"id"`
	DistrictName string  `json:"district_name"`
	RegencyID    uint    `json:"regency_id"`
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
}

// Village mysql region entity.
type Village struct {
	ID          uint    `json:"id"`
	VillageName string  `json:"village_name"`
	DistrictID  uint    `json:"district_id"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}
