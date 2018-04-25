package main

import (
	"fmt"
	"log"
	"segmentator/segmentator"
)

func main() {
	// Information.
	fmt.Println("____________________________________________________________________")
	fmt.Println("| Segmentator tutorial. V-0.1 by Mikhail Masyagin. BMSTU. IU9-42B. |")
	fmt.Println("‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾")
	fmt.Println()

	// Loading image.
	path := ""
	name := ""
	fmt.Println("Choose path and name to load image:")
	fmt.Printf("1) ")
	fmt.Scanf("%s\n", &path)
	fmt.Printf("2) ")
	fmt.Scanf("%s\n", &name)
	fmt.Println()
	path = "/home/mikhail/"
	name = "igor.jpg"
	img, err := segmentator.LoadImage(path, name)
	if err != nil {
		log.Fatalf("Error occured, while loading image: %s", err)
	}

	segmentator.GSLuma(img)
	segmentator.FGEDSobel(img, segmentator.SQRTGXGY)

	// Saving Image.
	path = ""
	name = ""
	fmt.Println("Choose path and name to save image:")
	fmt.Printf("1) ")
	fmt.Scanf("%s\n", &path)
	fmt.Printf("2) ")
	fmt.Scanf("%s\n", &name)
	path = "/home/mikhail/"
	name = "igor1.jpg"
	err = segmentator.SaveImage(path, name, img)
	if err != nil {
		log.Fatalf("Error occured, while saving image: %s", err)
	}
	fmt.Println("Done")
}
