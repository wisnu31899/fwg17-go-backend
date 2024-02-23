package lib

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadFile(c *gin.Context, dest string) (string, error) {
	// Mendapatkan file dari form
	file, err := c.FormFile("picture")
	if err != nil {
		return "", err
	}

	// Memeriksa ukuran file
	if file.Size > 5<<20 { // 5 MB dalam byte
		return "", fmt.Errorf("file size exceeds the limit of 5MB")
	}

	// Menentukan ekstensi file
	ext := map[string]string{
		"image/jpeg": ".jpg",
		"image/png":  ".png",
	}
	fileType := file.Header["Content-Type"][0]

	// Memeriksa apakah ekstensi file didukung
	if _, ok := ext[fileType]; !ok {
		return "", fmt.Errorf("file format is not supported")
	}

	// Membuat nama file unik dengan ekstensi yang sesuai
	fileName := fmt.Sprintf("uploads/%v/%v%v", dest, uuid.NewString(), ext[fileType])

	// Menyimpan file
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		return "", err
	}

	// Mengembalikan nama file dalam bentuk string
	return fileName, nil
}

func UploadFileProduct(c *gin.Context, dest string) (string, error) {
	// Mendapatkan file dari form
	file, err := c.FormFile("image")
	if err != nil {
		return "", err
	}
	// Memeriksa ukuran file
	if file.Size > 5<<20 { // 5 MB dalam byte
		return "", fmt.Errorf("file size exceeds the limit of 5MB")
	}

	// Menentukan ekstensi file
	ext := map[string]string{
		"image/jpeg": ".jpg",
		"image/png":  ".png",
	}
	fileType := file.Header["Content-Type"][0]

	// Memeriksa apakah ekstensi file didukung
	if _, ok := ext[fileType]; !ok {
		return "", fmt.Errorf("file format is not supported")
	}

	// Membuat nama file unik dengan ekstensi yang sesuai
	fileName := fmt.Sprintf("uploads/%v/%v%v", dest, uuid.NewString(), ext[fileType])

	// Menyimpan file
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		return "", err
	}

	fmt.Println(fileName)
	// Mengembalikan nama file dalam bentuk string
	return fileName, nil
}
