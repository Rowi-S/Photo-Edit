package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	greetingScreen()
	file, errFile := openingFile()
	if errFile != nil {
		fmt.Println("Error: File could not be opened", errFile)
		return
	}

	defer file.Close()
	img, errImg := decodingFile(file)
	if errImg != nil {
		fmt.Println("Error: Could not decode image", errImg)
		return
	}

	fmt.Println("Image dimensions:", img.Bounds())

	bwImg := convertToBlackAndWhite(img)

	if err := saveImage(bwImg, "bw_go_mascot.png"); err != nil {
		fmt.Println("Failed to save image", err)
		return
	}

}

func greetingScreen() {
	fmt.Println("Welcome to the photo edit cli: ")
}

func openingFile() (*os.File, error) {
	file, err := os.Open("go_mascot.png")
	fmt.Println("File is closed.")
	return file, err
}

func decodingFile(file *os.File) (image.Image, error) {
	img, err := png.Decode(file)
	return img, err
}

func convertToBlackAndWhite(img image.Image) image.Image {
	bounds := img.Bounds()
	bwImg := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			originalColor := img.At(x, y)
			grayColor := color.GrayModel.Convert(originalColor)
			bwImg.Set(x, y, grayColor)
		}
	}
	return bwImg
}

func saveImage(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %w ", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("could not encode image: %w", err)
	}

	fmt.Println("Image saved successfully as:", filename)
	return nil
}
