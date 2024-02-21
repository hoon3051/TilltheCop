package service

import (
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