package response

type ShipResponse struct {
	ID                    string   `json:"id"`
	OperatorID            string   `json:"operator_id"`
	ShipName              string   `json:"ship_name"`
	IMONumber             *string  `json:"imo_number"`
	CallSign              *string  `json:"call_sign"`
	MMSI                  *string  `json:"mmsi"`
	ShipType              string   `json:"ship_type"`
	FlagState             string   `json:"flag_state"`
	PortOfRegistry        string   `json:"port_of_registry"`
	BuildYear             *int     `json:"build_year"`
	Builder               *string  `json:"builder"`
	Length                *float64 `json:"length"`
	Beam                  *float64 `json:"beam"`
	Draft                 *float64 `json:"draft"`
	GrossTonnage          *float64 `json:"gross_tonnage"`
	NetTonnage            *float64 `json:"net_tonnage"`
	DeadweightTonnage     *float64 `json:"deadweight_tonnage"`
	MaxSpeed              *float64 `json:"max_speed"`
	PassengerCapacity     *int     `json:"passenger_capacity"`
	CrewCapacity          *int     `json:"crew_capacity"`
	ClassificationSociety *string  `json:"classification_society"`
	Status                string   `json:"status"`
	IsActive              bool     `json:"is_active"`
	LastInspection        *int64   `json:"last_inspection"`
	NextInspection        *int64   `json:"next_inspection"`
	InsuranceExpiry       *int64   `json:"insurance_expiry"`
	CertificateExpiry     *int64   `json:"certificate_expiry"`
	CurrentLatitude       *float64 `json:"current_latitude"`
	CurrentLongitude      *float64 `json:"current_longitude"`
	LastPosition          *int64   `json:"last_position"`
	Notes                 *string  `json:"notes"`
	CreatedAt             int64    `json:"created_at"`
	UpdatedAt             int64    `json:"updated_at"`
}