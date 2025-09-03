package entity

// Ship is a struct that represents a ship entity
type Ship struct {
	ID                     string   `gorm:"column:id;primaryKey"`
	OperatorID             string   `gorm:"column:operator_id"`
	ShipName               string   `gorm:"column:ship_name"`
	IMONumber              string   `gorm:"column:imo_number;uniqueIndex"`
	CallSign               string   `gorm:"column:call_sign;uniqueIndex"`
	MMSI                   string   `gorm:"column:mmsi;uniqueIndex"`
	ShipType               string   `gorm:"column:ship_type"`
	FlagState              string   `gorm:"column:flag_state"`
	PortOfRegistry         string   `gorm:"column:port_of_registry"`
	BuildYear              *int     `gorm:"column:build_year"`
	Builder                *string  `gorm:"column:builder"`
	Length                 *float64 `gorm:"column:length"`
	Beam                   *float64 `gorm:"column:beam"`
	Draft                  *float64 `gorm:"column:draft"`
	GrossTonnage           *float64 `gorm:"column:gross_tonnage"`
	NetTonnage             *float64 `gorm:"column:net_tonnage"`
	DeadweightTonnage      *float64 `gorm:"column:deadweight_tonnage"`
	MaxSpeed               *float64 `gorm:"column:max_speed"`
	PassengerCapacity      *int     `gorm:"column:passenger_capacity"`
	CrewCapacity           *int     `gorm:"column:crew_capacity"`
	ClassificationSociety  *string  `gorm:"column:classification_society"`
	Status                 string   `gorm:"column:status;default:active"`
	IsActive               bool     `gorm:"column:is_active;default:true"`
	LastInspection         *int64   `gorm:"column:last_inspection"`
	NextInspection         *int64   `gorm:"column:next_inspection"`
	InsuranceExpiry        *int64   `gorm:"column:insurance_expiry"`
	CertificateExpiry      *int64   `gorm:"column:certificate_expiry"`
	CurrentLatitude        *float64 `gorm:"column:current_latitude"`
	CurrentLongitude       *float64 `gorm:"column:current_longitude"`
	LastPosition           *int64   `gorm:"column:last_position"`
	Notes                  *string  `gorm:"column:notes"`
	CreatedAt              int64    `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt              int64    `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	DeletedAt              *int64   `gorm:"column:deleted_at"`

	// Relations
	Operator Operator `gorm:"foreignKey:operator_id;references:id"`
}

func (s *Ship) TableName() string {
	return "ships"
}