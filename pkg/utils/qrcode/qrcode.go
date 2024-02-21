package qrcode

import (
	"image/color"
	"os"

	QREncode "github.com/skip2/go-qrcode"
	QRParse "github.com/tuotoo/qrcode"
)

type QRCode struct {
	Size            int
	Level           QREncode.RecoveryLevel
	BackgroundColor color.Color
	ForegroundColor color.Color
}

func (qr *QRCode) GenerateQRCodeToBytes(content string) ([]byte, error) {
	_qr, err := QREncode.New(content, qr.Level)
	if err != nil {
		return nil, err
	}
	if qr.BackgroundColor != nil {
		_qr.BackgroundColor = qr.BackgroundColor
	}
	if qr.ForegroundColor != nil {
		_qr.ForegroundColor = qr.ForegroundColor
	}
	png, err := _qr.PNG(qr.Size)

	if err != nil {
		return nil, err
	}
	return png, nil
}

func (qr *QRCode) GenerateQRCode(content string, filename string) error {
	png, err := qr.GenerateQRCodeToBytes(content)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, png, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (qr *QRCode) DecodeQRCode(filename string) (string, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer fi.Close()

	qrmatrix, err := QRParse.Decode(fi)
	if err != nil {
		return "", err
	}
	return qrmatrix.Content, nil
}

func NewQRCode(size int) *QRCode {
	return &QRCode{Size: size, Level: QREncode.Medium, BackgroundColor: color.White, ForegroundColor: color.Black}
}

func Example() {
	qr := NewQRCode(256)
	qr.GenerateQRCode("hello world", "hello.png")
	content, _ := qr.DecodeQRCode("hello.png")
	println(content)
}
