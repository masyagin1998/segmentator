package segmentator

import "errors"

// FGPCIterative is a simple binarization method.
func FGPCIterative(img Image) (threshold int) {
	duplicate := make([][]Pixel, len(img.Pixels))
	for i := range img.Pixels {
		duplicate[i] = make([]Pixel, len(img.Pixels[i]))
		copy(duplicate[i], img.Pixels[i])
	}

	// Find maximum and minimum.
	min := 255
	max := 0

	for x := 0; x < img.Height; x++ {
		for y := 0; y < img.Width; y++ {
			color := duplicate[x][y].R
			if color > max {
				max = color
			}
			if color < min {
				min = color
			}
		}
	}

	prevT := -1
	currT := (min + max) / 2
	eps := 1

	// Iterations.
	for {
		numBackground := int64(0)
		sumBackground := int64(0)

		numForeground := int64(0)
		sumForeground := int64(0)

		for x := 0; x < img.Height; x++ {
			for y := 0; y < img.Width; y++ {
				if duplicate[x][y].R > currT {
					numForeground++
					sumForeground += int64(duplicate[x][y].R)
				} else {
					numBackground++
					sumBackground += int64(duplicate[x][y].R)
				}
			}
		}

		if (prevT != -1) && (abs(prevT-currT) < eps) {
			for x := 0; x < img.Height; x++ {
				for y := 0; y < img.Width; y++ {
					color := 0
					if duplicate[x][y].R > currT {
						color = 255
					}
					duplicate[x][y].R = color
					duplicate[x][y].G = color
					duplicate[x][y].B = color
				}
			}
			break
		}
		prevT = currT
		currT = int(sumForeground/numForeground+sumBackground/numBackground) / 2
	}

	threshold = currT

	return
}

// FGPCOtsuThresholding2 uses Otsu method for image binarization.
func FGPCOtsuThresholding2(img Image) (threshold int) {
	duplicate := make([][]Pixel, len(img.Pixels))
	for i := range img.Pixels {
		duplicate[i] = make([]Pixel, len(img.Pixels[i]))
		copy(duplicate[i], img.Pixels[i])
	}

	// Построение нормированной гистограммы - p[i] = n[i] / (M * N).
	normalizedHist := make([]float64, 256)

	for x := 0; x < img.Height; x++ {
		for y := 0; y < img.Width; y++ {
			normalizedHist[duplicate[x][y].R]++
		}
	}

	for i := 0; i < len(normalizedHist); i++ {
		normalizedHist[i] /= float64(img.Width * img.Height)
	}

	k := 0
	dispMax := float64(0)

	for i := 0; i < len(normalizedHist); i++ {
		// Вычисление накопленных сумм - P1(k) = S[i = 0...k] p[i].
		P1 := float64(0)
		for j := 0; j <= i; j++ {
			P1 += normalizedHist[j]
		}
		// Вычисление накопленных сумм - m(k) = S[i = 0...k] i*p[i].
		m := float64(0)
		for j := 0; j <= i; j++ {
			m += float64(j) * normalizedHist[j]
		}
		// Вычисление глобальной средней яркости - mG = S[i = 0...L-1] i*p[i]
		mG := float64(0)
		for j := 0; j < len(normalizedHist); j++ {
			mG += float64(j) * normalizedHist[j]
		}
		// Вычисление межклассовой дисперсии - disp = (mG * P1[k] - m(k))^2 / (P1(k) * (1 - P1(k)))
		disp := (mG*P1 - m) * (mG*P1 - m) / (P1 * (1 - P1))
		if disp > dispMax {
			dispMax = disp
			k = i
		}
	}

	threshold = k

	return
}

// FGPCCustomThreshold uses custom thresholds to segmentate image.
func FGPCThreshold(img Image, thresholds []int, colors []Pixel) (err error) {
	if len(thresholds) == 0 {
		err = errors.New("Wrong thresholds array")
		return
	}
	if len(colors) != len(thresholds)+1 {
		err = errors.New("Wrong colors array")
		return
	}
	if len(thresholds) == 1 {
		t := thresholds[0]
		for x := 0; x < img.Height; x++ {
			for y := 0; y < img.Width; y++ {
				index := -1
				if img.Pixels[x][y].R > t {
					index = 1
				} else {
					index = 0
				}
				img.Pixels[x][y].R = colors[index].R
				img.Pixels[x][y].G = colors[index].G
				img.Pixels[x][y].B = colors[index].B
				img.Pixels[x][y].A = colors[index].A
			}
		}
	} else {
		for x := 0; x < img.Height; x++ {
			for y := 0; y < img.Width; y++ {
				index := -1
				if img.Pixels[x][y].R <= thresholds[0] {
					index = 0
				}
				if img.Pixels[x][y].R > thresholds[len(thresholds)-1] {
					index = len(colors) - 1
				}
				if index == -1 {
					for i := 0; i < len(thresholds)-1; i++ {
						if (img.Pixels[x][y].R > thresholds[i]) && (img.Pixels[x][y].R <= thresholds[i+1]) {
							index = i + 1
							break
						}
					}
				}
				img.Pixels[x][y].R = colors[index].R
				img.Pixels[x][y].G = colors[index].G
				img.Pixels[x][y].B = colors[index].B
				img.Pixels[x][y].A = colors[index].A
			}
		}
	}
	return
}
