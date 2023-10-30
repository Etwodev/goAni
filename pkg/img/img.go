package img

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func save(path string, image image.Image) error {
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("save: failed creating file path: %w", err)
	}
	defer out.Close()

	err = jpeg.Encode(out, image, nil)
	if err != nil {
		return fmt.Errorf("save: failed writing file: %w", err)
	}
	return nil
}

func Resize(file []byte, path string, size int) error {
	var newImage image.Image

	img, _, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return fmt.Errorf("Resize: failed decoding file: %w", err)
	}

	if bounds := img.Bounds(); bounds.Dy() > bounds.Dx() {
		newImage = resize.Resize(uint(size), 0, img, resize.Lanczos3)
	} else {
		newImage = resize.Resize(0, uint(size), img, resize.Lanczos3)
	}

	cropSize := image.Rect(0, 0, size, size)
	croppedImage := newImage.(SubImager).SubImage(cropSize)

	err = save(path, croppedImage)
	if err != nil {
		return fmt.Errorf("Resize: failed to save file: %w", err)
	}

	return nil
}

// [channel][width][height]
func ImageToArray(path string) ([]float64, error) {
	var full []float64
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ImageToArray: failed opening file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("ImageToArray: failed decoding file: %w", err)
	}

	bounds := img.Bounds()
	width := bounds.Dy()
	height := bounds.Dx()

	red := make([]float64, width*height)
	green := make([]float64, width*height)
	blue := make([]float64, width*height)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			red[i*j] = float64(r)
			green[i*j] = float64(g)
			blue[i*j] = float64(b)
		}
	}

	full = append(full, red...)
	full = append(full, green...)
	full = append(full, blue...)

	return full, nil
}
