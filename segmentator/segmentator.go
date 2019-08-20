package segmentator

import (
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

// Pixel contains RGBA codes.
type Pixel struct {
	R int
	G int
	B int
	A int
}

// Image contains pixel matrix, name and format
// of original image.
type Image struct {
	Pixels [][]Pixel
	Width  int
	Height int
	Path   string
	Name   string
}

// init register most common image formats: "jpeg", "jpg", "png", "gif"
func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jpeg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

// LoadImage loads an image with specified name.
func LoadImage(path, name string) (img Image, err error) {
	// Check if image name is correct.
	num := -1
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			num = i
			break
		}
	}
	if (num == -1) || (num == len(name)-1) {
		err = errors.New("Invalid image name")
		return
	} else if num == 0 {
		err = errors.New("Invalid image format")
		return
	}

	// Check if image exists.
	var file *os.File
	if _, err = os.Stat(path + name); os.IsNotExist(err) {
		return
	}

	// Open file.
	file, err = os.Open(path + name)
	if err != nil {
		return
	}
	defer file.Close()

	// Decode file.
	var decodedFile image.Image
	decodedFile, _, err = image.Decode(file)
	if err != nil {
		return
	}

	// Create Image in runtime.
	bounds := decodedFile.Bounds()
	img.Width = bounds.Max.X
	img.Height = bounds.Max.Y
	img.Path = path
	img.Name = name
	for x := 0; x < img.Height; x++ {
		var row []Pixel
		for y := 0; y < img.Width; y++ {
			r, g, b, a := decodedFile.At(x, y).RGBA()
			row = append(row, Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)})
		}
		img.Pixels = append(img.Pixels, row)
	}
	return
}

// SaveImage saves an image with specified name
func SaveImage(path, name string, img Image) (err error) {
	// Check if image name is correct.
	num := -1
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			num = i
			break
		}
	}
	if (num == -1) || (num == len(name)-1) {
		err = errors.New("Invalid image name")
		return
	} else if num == 0 {
		err = errors.New("Invalid image format")
		return
	}

	// Create new image.
	newImg := image.NewRGBA(image.Rect(0, 0, img.Width, img.Height))
	for x := 0; x < img.Height; x++ {
		for y := 0; y < img.Width; y++ {
			newImg.Set(y, x, color.RGBA{
				uint8(img.Pixels[x][y].R),
				uint8(img.Pixels[x][y].G),
				uint8(img.Pixels[x][y].B),
				uint8(img.Pixels[x][y].A)})
		}
	}

	// Save new image.
	file, err := os.OpenFile(path+name, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	switch name[num+1:] {
	case "jpeg", "jpg":
		err = jpeg.Encode(file, newImg, &jpeg.Options{Quality: 100})
	case "png":
		err = png.Encode(file, newImg)
	default:
		err = errors.New("Invalid image format")
	}
	return
}
