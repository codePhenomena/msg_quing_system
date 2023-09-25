package controller

import (
	"log"
	"image"
	"image/jpeg"
	"os"
	"strings"
	"path/filepath"
	"github.com/nfnt/resize"
)

func CompressAndStoreInLocal(imagePaths []string)([]string,error){

	// Destination directory for compressed images
	compressedDir := "./compressed_images/"
	err := os.MkdirAll(compressedDir, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	var compressedPaths []string

	for _, imagePath := range imagePaths {
		compressedPath, err := CompressImage(imagePath, compressedDir)
		if err != nil {
			log.Printf("Error compressing %s: %v", imagePath, err)
			continue
		}
		compressedPaths = append(compressedPaths, compressedPath)
	}
	return compressedPaths,err
}

func CompressImage(inputPath, outputPath string) (string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	FailOnError(err,"")

	// Resize the image to a smaller size using the resize package
	newImg := resize.Resize(800, 0, img, resize.Lanczos3)

	outputFileName := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath)) + "_compressed.jpg"
	outputFilePath := filepath.Join(outputPath, outputFileName)

	out, err := os.Create(outputFilePath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the compressed image to the output file
	err = jpeg.Encode(out, newImg, nil)
	if err != nil {
		return "", err
	}

	return outputFilePath, nil
}