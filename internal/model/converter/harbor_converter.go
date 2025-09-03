package converter

import (
	"mkp-boarding-test/internal/domain/entity"
	"mkp-boarding-test/internal/model"
)

func HarborToResponse(harbor *entity.Harbor) *model.HarborResponse {
	return &model.HarborResponse{
		ID:              harbor.ID,
		HarborCode:      harbor.HarborCode,
		HarborName:      harbor.HarborName,
		UNLocode:        &harbor.UNLocode,
		Country:         harbor.Country,
		Province:        harbor.Province,
		City:            harbor.City,
		Latitude:        &harbor.Latitude,
		Longitude:       &harbor.Longitude,
		TimeZone:        &harbor.Timezone,
		MaxShipLength:   harbor.MaxShipLength,
		MaxShipBeam:     harbor.MaxShipBeam,
		MaxShipDraft:    harbor.MaxShipDraft,
		BerthCount:      &harbor.BerthCount,
		StorageCapacity: harbor.StorageCapacity,
		ContactPerson:   &harbor.ContactPerson,
		ContactPhone:    &harbor.ContactPhone,
		ContactEmail:    &harbor.ContactEmail,
		Website:         harbor.Website,
		OperatingHours:  &harbor.WorkingHours,
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
