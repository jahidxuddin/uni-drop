package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	networkscanner "github.com/jahidxuddin/uni-drop/internal/network_scanner"
)

type PageData struct {
	Devices map[string]string
}

func main() {
	http.Handle("/script.js", http.FileServer(http.Dir("static")))
	http.Handle("/style.css", http.FileServer(http.Dir("static")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleFileUpload(w, r)
		} else {
			serveTemplate(w, r)
		}
	})

	log.Println("Starting web server on port 50000")
	if err := http.ListenAndServe(":50000", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
func serveTemplate(w http.ResponseWriter, _ *http.Request) {
	devices, err := networkscanner.RunNetworkScan()
	if err != nil {
		fmt.Printf("Fehler beim Netzwerkscan: %v\n", err)
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, PageData{Devices: devices})
	if err != nil {
		http.Error(w, "Failed to execute template", http.StatusInternalServerError)
	}
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Limit upload size to 10 MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["file"]
	if len(files) == 0 {
		http.Error(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	uploadDir := "uploads"
	if err := createUploadDirectory(uploadDir); err != nil {
		http.Error(w, "Unable to create upload directory", http.StatusInternalServerError)
		return
	}

	for _, fileHeader := range files {
		err := saveUploadedFile(fileHeader, uploadDir)
		if err != nil {
			http.Error(w, "Unable to save file", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createUploadDirectory(uploadDir string) error {
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		return os.Mkdir(uploadDir, os.ModePerm)
	}
	return nil
}

func saveUploadedFile(fileHeader *multipart.FileHeader, uploadDir string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	filePath := filepath.Join(uploadDir, fileHeader.Filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	return err
}
