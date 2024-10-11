package main

import (
	"fmt"
	"image"
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
