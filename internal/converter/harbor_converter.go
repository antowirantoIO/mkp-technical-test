package converter

import (
	"mkp-boarding-test/internal/entity"
	"mkp-boarding-test/internal/model"
)

func HarborToResponse(harbor *entity.Harbor) *model.HarborResponse {
	var unLocode *string
	if harbor.UNLocode != "" {
		unLocode = &harbor.UNLocode
	}

	var latitude *float64
	if harbor.Latitude != 0 {
		latitude = &harbor.Latitude
	}

	var longitude *float64
	if harbor.Longitude != 0 {
		longitude = &harbor.Longitude
	}

	var timezone *string
	if harbor.Timezone != "" {
		timezone = &harbor.Timezone
	}

	var berthCount *int
	if harbor.BerthCount > 0 {
		berthCount = &harbor.BerthCount
	}

	var contactPerson *string
	if harbor.ContactPerson != "" {
		contactPerson = &harbor.ContactPerson
	}

	var contactPhone *string
	if harbor.ContactPhone != "" {
		contactPhone = &harbor.ContactPhone
	}

	var contactEmail *string
	if harbor.ContactEmail != "" {
		contactEmail = &harbor.ContactEmail
	}

	var operatingHours *string
	if harbor.WorkingHours != "" {
		operatingHours = &harbor.WorkingHours
	}

	return &model.HarborResponse{
		ID:              harbor.ID,
		HarborCode:      harbor.HarborCode,
		HarborName:      harbor.HarborName,
		UNLocode:        unLocode,
		Country:         harbor.Country,
		Province:        harbor.Province,
		City:            harbor.City,
		Latitude:        latitude,
		Longitude:       longitude,
		TimeZone:        timezone,
		MaxShipLength:   harbor.MaxShipLength,
		MaxShipBeam:     harbor.MaxShipBeam,
		MaxShipDraft:    harbor.MaxShipDraft,
		BerthCount:      berthCount,
		StorageCapacity: harbor.StorageCapacity,
		ContactPerson:   contactPerson,
		ContactPhone:    contactPhone,
		ContactEmail:    contactEmail,
		Website:         harbor.Website,
		OperatingHours:  operatingHours,
		HasPilotage:     harbor.HasPilotage,
		HasTugService:   harbor.HasTugService,
		HasQuarantine:   harbor.HasQuarantine,
		HasCustoms:      harbor.HasCustoms,
		HasRepair:       harbor.HasRepairService,
		HasWaste:        harbor.HasWaste,
		IsActive:        harbor.IsActive,
		Notes:           harbor.Notes,
		CreatedAt:       harbor.CreatedAt,
		UpdatedAt:       harbor.UpdatedAt,
	}
}
