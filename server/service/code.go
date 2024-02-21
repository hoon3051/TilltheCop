package service

import (
	"strconv"

	"github.com/hoon3051/TilltheCop/server/initializer"
	"github.com/hoon3051/TilltheCop/server/model"
	"github.com/skip2/go-qrcode"
)

type CodeService struct{}

func (s CodeService) GenerateQRCode(userID uint, reportID string) ([]byte, error) {
	// Generate QR code
	qrCode, err := qrcode.Encode(reportID, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return qrCode, nil
}

func (s CodeService) CreateRecord(userID uint, reportID string) (record model.Record, err error) {
	// Parse reportID
	reportIDuint, err := strconv.ParseUint(reportID, 10, 64)
	if err != nil {
		return record, err
	}


	// Create record
	record = model.Record{
		User_id:   userID,
		Report_id: uint(reportIDuint),
	}
	err = initializer.DB.Create(&record).Error
	if err != nil {
		return record, err
	}

	return record, nil

}
