package segmentator

import (
	"errors"
	"math"
)

// do-parameters for FGEDSobel.
const (
	// Gx - Brightness = |Gx|.
	GX = iota
	// Gy - Brightness = |Gy|.
	GY = iota
	// GxGy - Brightness = |Gx| + |Gy|
	GXGY = iota
	// SqrtGxGy - Brightness = sqrt(Gx^2 + Gy^2)
	SQRTGXGY = iota
)

// FGEDRoberts uses Roberts operators for fingding Gx and Gy.
func FGEDRoberts(img Image, do int) (err error) {
	// Gx Roberts operator.
	GxOp := [][]int{
		{-1, 0},
		{0, 1},
	}

	// Gy Roberts operator.
	GyOp := [][]int{
		{0, -1},
		{1, 0},
	}

	return FGEDCustomOperators(img, GxOp, GyOp, do)
}

// FGEDPrevitt uses Previtt operators for finding Gx and Gy.
func FGEDPrevitt(img Image, do int) (err error) {
	// Gx Previtt operator.
	GxOp := [][]int{
		{-1, -1, -1},
		{0, 0, 0},
		{1, 1, 1},
	}

	// Gy Previtt operator.
	GyOp := [][]int{
		{-1, 0, 1},
		{-1, 0, 1},
		{-1, 0, 1},
	}

	return FGEDCustomOperators(img, GxOp, GyOp, do)
}

// FGEDSobel uses Sobel operators for finding Gx and Gy.
func FGEDSobel(img Image, do int) (err error) {
	// Gx Sobel operator.
	GxOp := [][]int{
		{-1, -2, -1},
		{0, 0, 0},
		{1, 2, 1},
	}

	// Gy Sobel operator.
	GyOp := [][]int{
		{-1, 0, 1},
		{-2, 0, 2},
		{-1, 0, 1},
	}

	return FGEDCustomOperators(img, GxOp, GyOp, do)
}

// FGEDScharr uses Scharr operators for finding Gx and Gy.
func FGEDScharr(img Image, do int) (err error) {
	// Gx Scharr operator.
	GxOp := [][]int{
		{-3, -10, -3},
		{0, 0, 0},
		{3, 10, 3},
	}

	// Gy Scharr operator.
	GyOp := [][]int{
		{-3, 0, 3},
		{-10, 0, 10},
		{-3, 0, 3},
	}

	return FGEDCustomOperators(img, GxOp, GyOp, do)
}

// FGEDCustomOperators uses programmers operators for finding Gx and Gy.
func FGEDCustomOperators(img Image, GxOp, GyOp [][]int, do int) (err error) {
	// Check do-parameter.
	if (do != GX) && (do != GY) && (do != GXGY) && (do != SQRTGXGY) {
		err = errors.New("Unknown do-parameter")
		return
	}

	// Check matricies.
	if len(GxOp) != len(GyOp) {
		err = errors.New("Wrong matricies")
		return
	}

	// Gx and Gy matricies
	var GxMat [][]int
	var GyMat [][]int

	// length of matrix
	l := len(GxOp)

	// Cycles.
	for x := 0; x < img.Height; x++ {
		var GxRow []int
		var GyRow []int
		for y := 0; y < img.Width; y++ {
			Gx := 0
			Gy := 0
			for i := -l / 2; i < (l+1)/2; i++ {
				if (len(GxOp) != len(GxOp[i+l/2])) || (len(GxOp) != len(GyOp[i+l/2])) {
					err = errors.New("Wrong matricies")
					return
				}
				for j := -l / 2; j < (l+1)/2; j++ {
					if (x+i < 0) || (x+i >= img.Height) || (y+j < 0) || (y+j >= img.Width) {
						continue
					}
					Gx += GxOp[i+l/2][j+l/2] * img.Pixels[x+i][y+j].R
					Gy += GyOp[i+l/2][j+l/2] * img.Pixels[x+i][y+j].R
				}
			}
			GxRow = append(GxRow, Gx)
			GyRow = append(GyRow, Gy)
		}
		GxMat = append(GxMat, GxRow)
		GyMat = append(GyMat, GyRow)
	}

	switch do {
	case GX:
		for x := 0; x < img.Height; x++ {
			for y := 0; y < img.Width; y++ {
				color := abs(GxMat[x][y])
				if color > 255 {
					color = 255
				}
				img.Pixels[x][y] = Pixel{color, color, color, img.Pixels[x][y].A}
			}
		}
	case GY:
		for x := 0; x < img.Height; x++ {
			for y := 0; y < img.Width; y++ {
				color := abs(GyMat[x][y])
				if color > 255 {
					color = 255
				}
				img.Pixels[x][y] = Pixel{color, color, color, img.Pixels[x][y].A}
			}
		}
	case GXGY:
		for x := 0; x < img.Height; x++ {
			for y := 0; y < img.Width; y++ {
				color := abs(GxMat[x][y]) + abs(GyMat[x][y])
				if color > 255 {
					color = 255
				}
				img.Pixels[x][y] = Pixel{color, color, color, img.Pixels[x][y].A}
			}
		}
	case SQRTGXGY:
		for x := 0; x < img.Height; x++ {
			for y := 0; y < img.Width; y++ {
				color := int(math.Sqrt(float64(GxMat[x][y]*GxMat[x][y] + GyMat[x][y]*GyMat[x][y])))
				if color > 255 {
					color = 255
				}
				img.Pixels[x][y] = Pixel{color, color, color, img.Pixels[x][y].A}
			}
		}
	}
	return
}

// FGEDMarrHildreth uses Marr-Hildreth algorithm for finding Gx and Gy.
func FGEDMarrHildreth(img Image) {

}
