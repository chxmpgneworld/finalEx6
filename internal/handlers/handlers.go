package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "non-root path requested", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusInternalServerError)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "parse error", http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "file receiving error", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	filedata, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "file read error", http.StatusInternalServerError)
		return
	}
	result, err := service.Determinant(string(filedata))
	if err != nil {
		http.Error(w, "converted error", http.StatusInternalServerError)
		return
	}
	newFileName := time.Now().UTC().Format("02_01_06_15_04_05")
	ext := filepath.Ext(handler.Filename)
	newFileName = newFileName + ext
	outputPath := filepath.Join("../", newFileName)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		http.Error(w, "create error", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()
	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "write error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
