package entity

// Harbor is a struct that represents a harbor entity
type Harbor struct {
	ID               string   `gorm:"column:id;primaryKey"`
	HarborCode       string   `gorm:"column:harbor_code;uniqueIndex"`
	HarborName       string   `gorm:"column:harbor_name"`
	UNLocode         string   `gorm:"column:un_locode;uniqueIndex"`
	Country          string   `gorm:"column:country"`
	Province         string   `gorm:"column:province"`
	City             string   `gorm:"column:city"`
	Address          string   `gorm:"column:address"`
	PostalCode       string   `gorm:"column:postal_code"`
	Latitude         float64  `gorm:"column:latitude"`
	Longitude        float64  `gorm:"column:longitude"`
	HarborType       string   `gorm:"column:harbor_type"`
	HarborCategory   string   `gorm:"column:harbor_category"`
	Status           string   `gorm:"column:status;default:active"`
	MaxShipLength    *float64 `gorm:"column:max_ship_length"`
	MaxShipBeam      *float64 `gorm:"column:max_ship_beam"`
	MaxShipDraft     *float64 `gorm:"column:max_ship_draft"`
	MaxShipDWT       *float64 `gorm:"column:max_ship_dwt"`
	BerthCount       int      `gorm:"column:berth_count;default:0"`
	CraneCount       int      `gorm:"column:crane_count;default:0"`
	StorageCapacity  *float64 `gorm:"column:storage_capacity"`
	WaterDepth       float64  `gorm:"column:water_depth"`
	TidalRange       *float64 `gorm:"column:tidal_range"`
	WorkingHours     string   `gorm:"column:working_hours"`
	Timezone         string   `gorm:"column:timezone"`
	ContactPerson    string   `gorm:"column:contact_person"`
	ContactPhone     string   `gorm:"column:contact_phone"`
	ContactEmail     string   `gorm:"column:contact_email"`
	Website          *string  `gorm:"column:website"`
	HasCustoms       bool     `gorm:"column:has_customs;default:false"`
	HasQuarantine    bool     `gorm:"column:has_quarantine;default:false"`
	HasPilotage      bool     `gorm:"column:has_pilotage;default:false"`
	HasTugService    bool     `gorm:"column:has_tug_service;default:false"`
	HasBunkering     bool     `gorm:"column:has_bunkering;default:false"`
	HasRepairService bool     `gorm:"column:has_repair_service;default:false"`
	HasWaste         bool     `gorm:"column:has_waste;default:false"`
	IsActive         bool     `gorm:"column:is_active;default:true"`
	EstablishedAt    *int64   `gorm:"column:established_at"`
	Notes            *string  `gorm:"column:notes"`
	CreatedAt        int64    `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt        int64    `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt        *int64   `gorm:"column:deleted_at"`
}

func (h *Harbor) TableName() string {
	return "harbors"
}
