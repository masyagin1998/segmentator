# segmentator
This is a simple library for image segmentation written in Golang.
I need it for my CG essay, but hope it'll be usefull for You too :)

Structure:

0) Loading and Saving Images.

1) Grayscale algorithms:
   - Averaging ("Quick and Dirty");
   - Correcting the human eye ("Luma");
   - Desaturation;
   - Decomposition (Maximal and Minimal);
   - Single Color Channel (Red, Green or Blue);

2) First generation algorithms:
   - Based on Edge-Detection:
      - Roberts operator: func FGEDRoberts(img Image, do int) (err error)
      - Previtt operator: func FGEDPrevitt(img Image, do int) (err error)
      - Sobel operator:   func FGEDSobel(img Image, do int) (err error)
      - Scharr operator:  func FGEDScharr(img Image, do int) (err error)
      - Custom operator:  func FGEDCustomOperators(img Image, GxOp, GyOp [][]int, do int) (err error)
      - Marr-Hildreth:
   - Based on Pixel Classification:
   - Based on Regions:
  
3) Second generation algorithms:
   - Based on Edge-Detection:
   - Based on Pixel Classification:
   - Based on Regions:
   
4) Third generation algorithms:
   - Based on Edge-Detection:
   - Based on Pixel Classification:
   - Based on Regions:

# structure of the library

# important notes

1) Disputes on which coordinate system to use when processing images continues to the day. I prefer Gonsales-Woods varinant:
![coordinate system](info/Coordinates.png)
