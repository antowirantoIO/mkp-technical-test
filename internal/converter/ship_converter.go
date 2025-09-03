package converter

import (
	"golang-clean-architecture/internal/entity"
	"golang-clean-architecture/internal/model"
)

func ShipToResponse(ship *entity.Ship) *model.ShipResponse {
	return &model.ShipResponse{
		ID:                    ship.ID,
		OperatorID:            ship.OperatorID,
		ShipName:              ship.ShipName,
		IMONumber:             &ship.IMONumber,
		CallSign:              &ship.CallSign,
		MMSI:                  &ship.MMSI,
		ShipType:              ship.ShipType,
		FlagState:             ship.FlagState,
		PortOfRegistry:        ship.PortOfRegistry,
		BuildYear:             ship.BuildYear,
		Builder:               ship.Builder,
		Length:                ship.Length,
		Beam:                  ship.Beam,
		Draft:                 ship.Draft,
		GrossTonnage:          ship.GrossTonnage,
		NetTonnage:            ship.NetTonnage,
		DeadweightTonnage:     ship.DeadweightTonnage,
		MaxSpeed:              ship.MaxSpeed,
		PassengerCapacity:     ship.PassengerCapacity,
		CrewCapacity:          ship.CrewCapacity,
		ClassificationSociety: ship.ClassificationSociety,
		Status:                ship.Status,
		IsActive:              ship.IsActive,
		LastInspection:        ship.LastInspection,
		NextInspection:        ship.NextInspection,
		InsuranceExpiry:       ship.InsuranceExpiry,
		CertificateExpiry:     ship.CertificateExpiry,
		CurrentLatitude:       ship.CurrentLatitude,
		CurrentLongitude:      ship.CurrentLongitude,
		LastPosition:          ship.LastPosition,
		Notes:                 ship.Notes,
		CreatedAt:             ship.CreatedAt,
		UpdatedAt:             ship.UpdatedAt,
	}
}