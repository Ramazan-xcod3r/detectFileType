package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Dosya adını al
	fileName := os.Args[1]

	// Dosyayı aç
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Dosya açılamadı:", err)
		return
	}
	defer file.Close()

	// Dosya içeriğini oku
	reader := bufio.NewReader(file)
	magicNumber, err := reader.Peek(4)
	if err != nil {
		fmt.Println("Dosya okunamadı:", err)
		return
	}

	// Dosya tipini belirle
	var fileType string
	switch {
	case magicNumber[0] == 0xFF && magicNumber[1] == 0xD8 && magicNumber[2] == 0xFF:
		fileType = "image/jpeg"
	case magicNumber[0] == 0x47 && magicNumber[1] == 0x49 && magicNumber[2] == 0x46 && magicNumber[3] == 0x38:
		fileType = "image/gif"
	case magicNumber[0] == 0x89 && magicNumber[1] == 0x50 && magicNumber[2] == 0x4E && magicNumber[3] == 0x47:
		fileType = "image/png"
	default:
		fileType = "bilinmiyor"
	}
	fmt.Println("Dosya tipi:", fileType)
}
