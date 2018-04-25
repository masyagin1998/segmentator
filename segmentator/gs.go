package segmentator

import "errors"

// GSAveraging - Brightness = (Red + Green + Blue) / 3.
func GSAveraging(img Image) {
	// Color.
	for x := 0; x < img.Height; x++ {
		for y := 0; y < img.Width; y++ {
			brightness := (img.Pixels[x][y].R + img.Pixels[x][y].G + img.Pixels[x][y].B) / 3
			img.Pixels[x][y].R = brightness
			img.Pixels[x][y].G = brightness
			img.Pixels[x][y].B = brightness
		}
	}
}

// GSLuma - Brightness = (Red * 0.2126 + Green * 0.7152 + Blue * 0.0722).
func GSLuma(img Image) {
	// Color.
	for x := 0; x < img.Height; x++ {
		for y := 0; y < img.Width; y++ {
			brightness := int((float64(img.Pixels[x][y].R)*0.2126 +
				float64(img.Pixels[x][y].G)*0.7152 +
				float64(img.Pixels[x][y].B)*0.0722) / 3.0)
			img.Pixels[x][y].R = brightness
			img.Pixels[x][y].G = brightness
			img.Pixels[x][y].B = brightness
		}
	}
}

// GSDesaturation - Brightness = ( Max(Red, Green, Blue) + Min(Red, Green, Blue) ) / 2.
func GSDesaturation(img Image) {
	// Color.
	for x := 0; x < img.Height; x++ {
		for y := 0; y < img.Width; y++ {
			brightness := (max(img.Pixels[x][y].R, img.Pixels[x][y].G, img.Pixels[x][y].B) -
				min(img.Pixels[x][y].R, img.Pixels[x][y].G, img.Pixels[x][y].B)) / 2
			img.Pixels[x][y].R = brightness
			img.Pixels[x][y].G = brightness
			img.Pixels[x][y].B = brightness
		}
	}
}

// do-parameters for GSDecomposition
const (
	// DOMIN - Brightness = Min(Red, Green, Blue).
	DOMIN = iota
	// DOMAX - Brightness = Max(Red, Green, Blue).
	DOMAX = iota
)

// GSDecomposition - Brightness = Min or Max(Red, Green, Blue).
func GSDecomposition(img Image, do int) (err error) {
	// Check do-parameter.
	if (do != DOMIN) && (do != DOMAX) {
		err = errors.New("Unknown do-parameter")
		return
	}
	// Color.
	for x := 0; x < img.Height; x++ {
		for y := 0; y < img.Width; y++ {
			brightness := 0
			if do == DOMIN {
				brightness = min(img.Pixels[x][y].R, img.Pixels[x][y].G, img.Pixels[x][y].B)
			} else if do == DOMAX {
				brightness = max(img.Pixels[x][y].R, img.Pixels[x][y].G, img.Pixels[x][y].B)
			}
			img.Pixels[x][y].R = brightness
			img.Pixels[x][y].G = brightness
			img.Pixels[x][y].B = brightness
		}
	}
	return
}

// color-parameters for GSSingleColor
const (
	// RED - Brightness = Red
	RED = iota
	// GREEN - Brightness = Green
	GREEN = iota
	// BLUE - Brightness = Blue
	BLUE = iota
)

// GSSingleColor - Brightness = Red or Green or Blue
func GSSingleColor(img Image, color int) (err error) {
	// Check color-parameter.
	if (color != RED) && (color != GREEN) && (color != BLUE) {
		err = errors.New("Unknown color-parameter")
		return
	}
	// Color.
	for x := 0; x < img.Height; x++ {
		for y := 0; y < img.Width; y++ {
			brightness := 0
			if color == RED {
				brightness = img.Pixels[x][y].R
			} else if color == GREEN {
				brightness = img.Pixels[x][y].G
			} else if color == BLUE {
				brightness = img.Pixels[x][y].B
			}
			img.Pixels[x][y].R = brightness
			img.Pixels[x][y].G = brightness
			img.Pixels[x][y].B = brightness
		}
	}
	return
}
