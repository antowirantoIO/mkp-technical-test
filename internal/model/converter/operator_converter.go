package converter

import (
	"mkp-boarding-test/internal/entity"
	"mkp-boarding-test/internal/model"
)

func OperatorToResponse(operator *entity.Operator) *model.OperatorResponse {
	return &model.OperatorResponse{
		ID:            operator.ID,
		UserID:        operator.UserID,
		OperatorCode:  operator.OperatorCode,
		CompanyName:   operator.CompanyName,
		LicenseNumber: operator.LicenseNumber,
		ContactPerson: operator.ContactPerson,
		ContactPhone:  operator.ContactPhone,
		ContactEmail:  operator.ContactEmail,
		Address:       operator.Address,
		City:          operator.City,
		Province:      operator.Province,
		Country:       operator.Country,
		PostalCode:    operator.PostalCode,
		Website:       operator.Website,
		OperatorType:  operator.OperatorType,
		Status:        operator.Status,
		EstablishedAt: operator.EstablishedAt,
		LicenseExpiry: operator.LicenseExpiry,
		IsActive:      operator.IsActive,
		Notes:         operator.Notes,
		CreatedAt:     operator.CreatedAt,
		UpdatedAt:     operator.UpdatedAt,
	}
}
