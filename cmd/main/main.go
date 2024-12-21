package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/jahidxuddin/uni-drop/internal/filetransfer"
	"google.golang.org/grpc"
)

func main() {
	go startGRPCServer()

	waitForServer(":9000")

	go func() {
		err := openURL("http://localhost:50000")
		if err != nil {
			log.Println("Failed to open default web browser:", err)
		}
	}()

	http.Handle("/script.js", http.FileServer(http.Dir("static")))
	http.Handle("/style.css", http.FileServer(http.Dir("static")))
	http.Handle("/index.html", http.FileServer(http.Dir("static")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleFileUpload(w, r)
		} else {
			http.ServeFile(w, r, "static/index.html")
		}
	})

	log.Println("Starting web server on port 50000")
	if err := http.ListenAndServe(":50000", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	filetransfer.RegisterFileTransferServer(grpcServer, &filetransfer.Server{})

	log.Println("Starting grpc server on port 9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func waitForServer(address string) {
	for {
		conn, err := net.Dial("tcp", address)
		if err == nil {
			conn.Close()
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
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

func openURL(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return fmt.Errorf("unsupported platform")
	}

	return exec.Command(cmd, args...).Start()
}
