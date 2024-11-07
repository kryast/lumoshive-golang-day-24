package library

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Fungsi untuk meng-upload file ke folder tertentu
func UploadFile(r *http.Request, fileField string, folder string, fileExt string) (string, error) {
	// Cek apakah file ada di form
	file, _, err := r.FormFile(fileField)
	if err != nil {
		// Jika file tidak ada, kembalikan string kosong
		return "", nil
	}
	defer file.Close()

	// Tentukan nama file dengan timestamp
	fileName := fmt.Sprintf("%s_%s.%s", fileField, time.Now().Format("20060102150405"), fileExt)
	filePath := folder + "/" + fileName

	// Buat folder jika belum ada
	err = os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory for file: %v", err)
	}

	// Simpan file ke server
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}
	defer outFile.Close()

	// Salin isi file ke file yang baru dibuat
	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file: %v", err)
	}

	// Kembalikan path file yang telah disimpan
	return filePath, nil
}
