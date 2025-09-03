package model

type HarborResponse struct {
	ID                    string   `json:"id"`
	HarborCode            string   `json:"harbor_code"`
	HarborName            string   `json:"harbor_name"`
	UNLocode              *string  `json:"un_locode"`
	Country               string   `json:"country"`
	Province              string   `json:"province"`
	City                  string   `json:"city"`
	Latitude              *float64 `json:"latitude"`
	Longitude             *float64 `json:"longitude"`
	TimeZone              *string  `json:"time_zone"`
	MaxShipLength         *float64 `json:"max_ship_length"`
	MaxShipBeam           *float64 `json:"max_ship_beam"`
	MaxShipDraft          *float64 `json:"max_ship_draft"`
	BerthCount            *int     `json:"berth_count"`
	AnchorageDepth        *float64 `json:"anchorage_depth"`
	ChannelDepth          *float64 `json:"channel_depth"`
	CargoHandlingCapacity *float64 `json:"cargo_handling_capacity"`
	StorageCapacity       *float64 `json:"storage_capacity"`
	ContactPerson         *string  `json:"contact_person"`
	ContactPhone          *string  `json:"contact_phone"`
	ContactEmail          *string  `json:"contact_email"`
	Website               *string  `json:"website"`
	OperatingHours        *string  `json:"operating_hours"`
	HasPilotage           bool     `json:"has_pilotage"`
	HasTugService         bool     `json:"has_tug_service"`
	HasQuarantine         bool     `json:"has_quarantine"`
	HasCustoms            bool     `json:"has_customs"`
	HasImmigration        bool     `json:"has_immigration"`
	HasSecurity           bool     `json:"has_security"`
	HasMedical            bool     `json:"has_medical"`
	HasRepair             bool     `json:"has_repair"`
	HasSupplies           bool     `json:"has_supplies"`
	HasFuel               bool     `json:"has_fuel"`
	HasWater              bool     `json:"has_water"`
	HasWaste              bool     `json:"has_waste"`
	HasCargo              bool     `json:"has_cargo"`
	HasPassenger          bool     `json:"has_passenger"`
	HasRoro               bool     `json:"has_roro"`
	HasContainer          bool     `json:"has_container"`
	HasBulk               bool     `json:"has_bulk"`
	HasLiquid             bool     `json:"has_liquid"`
	HasBreakbulk          bool     `json:"has_breakbulk"`
	IsActive              bool     `json:"is_active"`
	Notes                 *string  `json:"notes"`
	CreatedAt             int64    `json:"created_at"`
	UpdatedAt             int64    `json:"updated_at"`
}

type CreateHarborRequest struct {
	HarborCode            string   `json:"harbor_code" validate:"required,max=20"`
	HarborName            string   `json:"harbor_name" validate:"required,max=255"`
	UNLocode              *string  `json:"un_locode" validate:"omitempty,max=10"`
	Country               string   `json:"country" validate:"required,max=100"`
	Province              string   `json:"province" validate:"required,max=100"`
	City                  string   `json:"city" validate:"required,max=100"`
	Latitude              *float64 `json:"latitude" validate:"omitempty,min=-90,max=90"`
	Longitude             *float64 `json:"longitude" validate:"omitempty,min=-180,max=180"`
	TimeZone              *string  `json:"time_zone" validate:"omitempty,max=50"`
	MaxShipLength         *float64 `json:"max_ship_length" validate:"omitempty,min=0"`
	MaxShipBeam           *float64 `json:"max_ship_beam" validate:"omitempty,min=0"`
	MaxShipDraft          *float64 `json:"max_ship_draft" validate:"omitempty,min=0"`
	BerthCount            *int     `json:"berth_count" validate:"omitempty,min=0"`
	AnchorageDepth        *float64 `json:"anchorage_depth" validate:"omitempty,min=0"`
	ChannelDepth          *float64 `json:"channel_depth" validate:"omitempty,min=0"`
	CargoHandlingCapacity *float64 `json:"cargo_handling_capacity" validate:"omitempty,min=0"`
	StorageCapacity       *float64 `json:"storage_capacity" validate:"omitempty,min=0"`
	ContactPerson         *string  `json:"contact_person" validate:"omitempty,max=255"`
	ContactPhone          *string  `json:"contact_phone" validate:"omitempty,max=20"`
	ContactEmail          *string  `json:"contact_email" validate:"omitempty,email,max=255"`
	Website               *string  `json:"website" validate:"omitempty,url,max=500"`
	OperatingHours        *string  `json:"operating_hours" validate:"omitempty,max=255"`
	HasPilotage           *bool    `json:"has_pilotage"`
	HasTugService         *bool    `json:"has_tug_service"`
	HasQuarantine         *bool    `json:"has_quarantine"`
	HasCustoms            *bool    `json:"has_customs"`
	HasImmigration        *bool    `json:"has_immigration"`
	HasSecurity           *bool    `json:"has_security"`
	HasMedical            *bool    `json:"has_medical"`
	HasRepair             *bool    `json:"has_repair"`
	HasSupplies           *bool    `json:"has_supplies"`
	HasFuel               *bool    `json:"has_fuel"`
	HasWater              *bool    `json:"has_water"`
	HasWaste              *bool    `json:"has_waste"`
	HasCargo              *bool    `json:"has_cargo"`
	HasPassenger          *bool    `json:"has_passenger"`
	HasRoro               *bool    `json:"has_roro"`
	HasContainer          *bool    `json:"has_container"`
	HasBulk               *bool    `json:"has_bulk"`
	HasLiquid             *bool    `json:"has_liquid"`
	HasBreakbulk          *bool    `json:"has_breakbulk"`
	Notes                 *string  `json:"notes" validate:"omitempty,max=1000"`
}

type UpdateHarborRequest struct {
	ID                    string   `json:"-" validate:"required,max=100,uuid"`
	HarborCode            *string  `json:"harbor_code" validate:"omitempty,max=20"`
	HarborName            *string  `json:"harbor_name" validate:"omitempty,max=255"`
	UNLocode              *string  `json:"un_locode" validate:"omitempty,max=10"`
	Country               *string  `json:"country" validate:"omitempty,max=100"`
	Province              *string  `json:"province" validate:"omitempty,max=100"`
	City                  *string  `json:"city" validate:"omitempty,max=100"`
	Latitude              *float64 `json:"latitude" validate:"omitempty,min=-90,max=90"`
	Longitude             *float64 `json:"longitude" validate:"omitempty,min=-180,max=180"`
	TimeZone              *string  `json:"time_zone" validate:"omitempty,max=50"`
	MaxShipLength         *float64 `json:"max_ship_length" validate:"omitempty,min=0"`
	MaxShipBeam           *float64 `json:"max_ship_beam" validate:"omitempty,min=0"`
	MaxShipDraft          *float64 `json:"max_ship_draft" validate:"omitempty,min=0"`
	BerthCount            *int     `json:"berth_count" validate:"omitempty,min=0"`
	AnchorageDepth        *float64 `json:"anchorage_depth" validate:"omitempty,min=0"`
	ChannelDepth          *float64 `json:"channel_depth" validate:"omitempty,min=0"`
	CargoHandlingCapacity *float64 `json:"cargo_handling_capacity" validate:"omitempty,min=0"`
	StorageCapacity       *float64 `json:"storage_capacity" validate:"omitempty,min=0"`
	ContactPerson         *string  `json:"contact_person" validate:"omitempty,max=255"`
	ContactPhone          *string  `json:"contact_phone" validate:"omitempty,max=20"`
	ContactEmail          *string  `json:"contact_email" validate:"omitempty,email,max=255"`
	Website               *string  `json:"website" validate:"omitempty,url,max=500"`
	OperatingHours        *string  `json:"operating_hours" validate:"omitempty,max=255"`
	HasPilotage           *bool    `json:"has_pilotage"`
	HasTugService         *bool    `json:"has_tug_service"`
	HasQuarantine         *bool    `json:"has_quarantine"`
	HasCustoms            *bool    `json:"has_customs"`
	HasImmigration        *bool    `json:"has_immigration"`
	HasSecurity           *bool    `json:"has_security"`
	HasMedical            *bool    `json:"has_medical"`
	HasRepair             *bool    `json:"has_repair"`
	HasSupplies           *bool    `json:"has_supplies"`
	HasFuel               *bool    `json:"has_fuel"`
	HasWater              *bool    `json:"has_water"`
	HasWaste              *bool    `json:"has_waste"`
	HasCargo              *bool    `json:"has_cargo"`
	HasPassenger          *bool    `json:"has_passenger"`
	HasRoro               *bool    `json:"has_roro"`
	HasContainer          *bool    `json:"has_container"`
	HasBulk               *bool    `json:"has_bulk"`
	HasLiquid             *bool    `json:"has_liquid"`
	HasBreakbulk          *bool    `json:"has_breakbulk"`
	IsActive              *bool    `json:"is_active"`
	Notes                 *string  `json:"notes" validate:"omitempty,max=1000"`
}

type GetHarborRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type DeleteHarborRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type ListHarborRequest struct {
	Page     int     `json:"page" validate:"min=1"`
	Size     int     `json:"size" validate:"min=1,max=100"`
	IsActive *bool   `json:"is_active"`
	Country  *string `json:"country"`
	Province *string `json:"province"`
	City     *string `json:"city"`
	Name     *string `json:"name"`
}

type SearchHarborRequest struct {
	Latitude  float64 `json:"latitude" validate:"required,min=-90,max=90"`
	Longitude float64 `json:"longitude" validate:"required,min=-180,max=180"`
	Radius    float64 `json:"radius" validate:"required,min=0,max=1000"`
	Page      int     `json:"page" validate:"min=1"`
	Size      int     `json:"size" validate:"min=1,max=100"`
}