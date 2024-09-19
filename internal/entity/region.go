package entity

// Province mysql region entity.
type Province struct {
	ID           uint    `gorm:"primarykey"`
	ProvinceName string  `gorm:"size:100;not null"`
	Longitude    float64 `gorm:"type:float;not null"`
	Latitude     float64 `gorm:"type:float;not null"`
	Geocode      string  `gorm:"size:6"`

	Regencies *[]Regency `gorm:"foreignKey:ProvinceID;references:ID"`
}

// Regency mysql region entity.
type Regency struct {
	ID          uint    `gorm:"primarykey"`
	RegencyName string  `gorm:"size:100;not null"`
	ProvinceID  uint    `gorm:"not null"`
	Longitude   float64 `gorm:"type:float;not null"`
	Latitude    float64 `gorm:"type:float;not null"`

	Districts *[]District `gorm:"foreignKey:RegencyID;references:ID"`
}

// District mysql region entity.
type District struct {
	ID           uint    `gorm:"primarykey"`
	DistrictName string  `gorm:"size:100;not null"`
	RegencyID    uint    `gorm:"not null"`
	Longitude    float64 `gorm:"type:float;not null"`
	Latitude     float64 `gorm:"type:float;not null"`

	Villages *[]Village `gorm:"foreignKey:DistrictID;references:ID"`
}

// Village mysql region entity.
type Village struct {
	ID          uint    `gorm:"primarykey"`
	VillageName string  `gorm:"size:100;not null"`
	DistrictID  uint    `gorm:"not null"`
	Longitude   float64 `gorm:"type:float;not null"`
	Latitude    float64 `gorm:"type:float;not null"`
}
