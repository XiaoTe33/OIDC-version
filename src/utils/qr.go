package utils

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
	"winter-examination/src/conf"
)

func GenerateQR(content string, name string) (link string) {
	qrCode, _ := qr.Encode(content, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	file, _ := os.Create(conf.LocalSavePathOfQR + name + ".jpg")
	defer file.Close()

	png.Encode(file, qrCode)

	return conf.WebLinkPathOfQR + name + ".jpg"
}
