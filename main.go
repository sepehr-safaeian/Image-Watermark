package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func main() {
	fmt.Print("Please enter the path of the folder the images:")
	var inputPath string
	fmt.Scan(&inputPath)

	outputPath := filepath.Join(inputPath, "watermarked")
	err := os.Mkdir(outputPath, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Error creating watermarked folder:", err)
		return
	}

	err = filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		img, _, err := loadImage(path)
		if err != nil {
			fmt.Printf("Error opening image %s: %v\n", path, err)
			return nil
		}

		// Convert image to *image.RGBA type
		rgba := image.NewRGBA(img.Bounds())
		draw.Draw(rgba, rgba.Bounds(), img, image.Point{}, draw.Over)

		watermarkPath := "logo.png"
		watermark, _, err := loadImage(watermarkPath)
		if err != nil {
			fmt.Printf("Error opening watermark %s: %v\n", watermarkPath, err)
			return nil
		}

		bounds := img.Bounds()
		offset := image.Pt(bounds.Dx()/2-watermark.Bounds().Dx()/2, bounds.Dy()/2-watermark.Bounds().Dy()/2)
		draw.Draw(rgba, watermark.Bounds().Add(offset), watermark, image.Point{}, draw.Over)

		outputFilePath := filepath.Join(outputPath, info.Name())
		err = saveImage(outputFilePath, rgba)
		if err != nil {
			fmt.Printf("Error saving watermarked image %s: %v\n", outputFilePath, err)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error in folder navigation:", err)
	}
}

func loadImage(path string) (image.Image, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

func saveImage(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	switch filepath.Ext(path) {
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, img, nil)
	case ".png":
		return png.Encode(file, img)
	default:
		return fmt.Errorf("Image format not supported: %s", filepath.Ext(path))
	}
}
