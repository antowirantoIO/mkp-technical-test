package request

type CreateShipRequest struct {
	OperatorID            string   `json:"operator_id" validate:"required,uuid"`
	ShipName              string   `json:"ship_name" validate:"required,max=255"`
	IMONumber             *string  `json:"imo_number" validate:"omitempty,max=20"`
	CallSign              *string  `json:"call_sign" validate:"omitempty,max=20"`
	MMSI                  *string  `json:"mmsi" validate:"omitempty,max=20"`
	ShipType              string   `json:"ship_type" validate:"required,max=100"`
	FlagState             string   `json:"flag_state" validate:"required,max=100"`
	PortOfRegistry        string   `json:"port_of_registry" validate:"required,max=255"`
	BuildYear             *int     `json:"build_year" validate:"omitempty,min=1800,max=2100"`
	Builder               *string  `json:"builder" validate:"omitempty,max=255"`
	Length                *float64 `json:"length" validate:"omitempty,min=0"`
	Beam                  *float64 `json:"beam" validate:"omitempty,min=0"`
	Draft                 *float64 `json:"draft" validate:"omitempty,min=0"`
	GrossTonnage          *float64 `json:"gross_tonnage" validate:"omitempty,min=0"`
	NetTonnage            *float64 `json:"net_tonnage" validate:"omitempty,min=0"`
	DeadweightTonnage     *float64 `json:"deadweight_tonnage" validate:"omitempty,min=0"`
	MaxSpeed              *float64 `json:"max_speed" validate:"omitempty,min=0"`
	PassengerCapacity     *int     `json:"passenger_capacity" validate:"omitempty,min=0"`
	CrewCapacity          *int     `json:"crew_capacity" validate:"omitempty,min=0"`
	ClassificationSociety *string  `json:"classification_society" validate:"omitempty,max=255"`
	LastInspection        *int64   `json:"last_inspection"`
	NextInspection        *int64   `json:"next_inspection"`
	InsuranceExpiry       *int64   `json:"insurance_expiry"`
	CertificateExpiry     *int64   `json:"certificate_expiry"`
	CurrentLatitude       *float64 `json:"current_latitude" validate:"omitempty,min=-90,max=90"`
	CurrentLongitude      *float64 `json:"current_longitude" validate:"omitempty,min=-180,max=180"`
	Notes                 *string  `json:"notes" validate:"omitempty,max=1000"`
}

type UpdateShipRequest struct {
	ID                    string   `json:"-" validate:"required,max=100,uuid"`
	ShipName              *string  `json:"ship_name" validate:"omitempty,max=255"`
	IMONumber             *string  `json:"imo_number" validate:"omitempty,max=20"`
	CallSign              *string  `json:"call_sign" validate:"omitempty,max=20"`
	MMSI                  *string  `json:"mmsi" validate:"omitempty,max=20"`
	ShipType              *string  `json:"ship_type" validate:"omitempty,max=100"`
	FlagState             *string  `json:"flag_state" validate:"omitempty,max=100"`
	PortOfRegistry        *string  `json:"port_of_registry" validate:"omitempty,max=255"`
	BuildYear             *int     `json:"build_year" validate:"omitempty,min=1800,max=2100"`
	Builder               *string  `json:"builder" validate:"omitempty,max=255"`
	Length                *float64 `json:"length" validate:"omitempty,min=0"`
	Beam                  *float64 `json:"beam" validate:"omitempty,min=0"`
	Draft                 *float64 `json:"draft" validate:"omitempty,min=0"`
	GrossTonnage          *float64 `json:"gross_tonnage" validate:"omitempty,min=0"`
	NetTonnage            *float64 `json:"net_tonnage" validate:"omitempty,min=0"`
	DeadweightTonnage     *float64 `json:"deadweight_tonnage" validate:"omitempty,min=0"`
	MaxSpeed              *float64 `json:"max_speed" validate:"omitempty,min=0"`
	PassengerCapacity     *int     `json:"passenger_capacity" validate:"omitempty,min=0"`
	CrewCapacity          *int     `json:"crew_capacity" validate:"omitempty,min=0"`
	ClassificationSociety *string  `json:"classification_society" validate:"omitempty,max=255"`
	Status                *string  `json:"status" validate:"omitempty,oneof=active inactive maintenance"`
	IsActive              *bool    `json:"is_active"`
	LastInspection        *int64   `json:"last_inspection"`
	NextInspection        *int64   `json:"next_inspection"`
	InsuranceExpiry       *int64   `json:"insurance_expiry"`
	CertificateExpiry     *int64   `json:"certificate_expiry"`
	CurrentLatitude       *float64 `json:"current_latitude" validate:"omitempty,min=-90,max=90"`
	CurrentLongitude      *float64 `json:"current_longitude" validate:"omitempty,min=-180,max=180"`
	Notes                 *string  `json:"notes" validate:"omitempty,max=1000"`
}

type GetShipRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteShipRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type ListShipRequest struct {
	Page       int     `json:"page" validate:"min=1"`
	Size       int     `json:"size" validate:"min=1,max=100"`
	OperatorID *string `json:"operator_id"`
	IsActive   *bool   `json:"is_active"`
	Status     *string `json:"status"`
	ShipType   *string `json:"ship_type"`
	FlagState  *string `json:"flag_state"`
	ShipName   *string `json:"ship_name"`
}

type GetShipsByOperatorRequest struct {
	OperatorID string `json:"-" validate:"required,max=100,uuid"`
	Page       int    `json:"page" validate:"min=1"`
	Size       int    `json:"size" validate:"min=1,max=100"`
}

type UpdateShipPositionRequest struct {
	ID               string  `json:"-" validate:"required,max=100,uuid"`
	CurrentLatitude  float64 `json:"current_latitude" validate:"required,min=-90,max=90"`
	CurrentLongitude float64 `json:"current_longitude" validate:"required,min=-180,max=180"`
}