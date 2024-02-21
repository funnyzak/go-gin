package image

import (
	"bytes"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"os"
)

func GetImageSize(imgpath string) (int, int, error) {
	file, err := os.Open(imgpath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()
	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}
	return img.Width, img.Height, nil
}

func IsImage(data []byte) bool {
	_, _, err := image.Decode(bytes.NewReader(data))
	return err == nil
}
