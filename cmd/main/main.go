package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/jahidxuddin/uni-drop/internal/fileservice"
	networkscanner "github.com/jahidxuddin/uni-drop/internal/network_scanner"
	"google.golang.org/grpc"
)

type PageData struct {
	Devices map[string]string
}

func main() {
	go startGRPCServer()

	http.Handle("/script.js", http.FileServer(http.Dir("static")))
	http.Handle("/style.css", http.FileServer(http.Dir("static")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handleFileUpload(w, r)
		} else {
			serveTemplate(w, r)
		}
	})

	if err := http.ListenAndServe(":50000", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	fileservice.RegisterFileServiceServer(s, &fileservice.Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // Limit upload size to 10 MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	conn, err := grpc.NewClient(r.FormValue("ip") + ":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := fileservice.NewFileServiceClient(conn)

	files := r.MultipartForm.File["file"]
	if len(files) == 0 {
		http.Error(w, "No files uploaded", http.StatusBadRequest)
		return
	}

	for _, file := range files {
		fileContent, err := file.Open()
		if err != nil {
			http.Error(w, "Unable to open uploaded file", http.StatusInternalServerError)
			return
		}
		defer fileContent.Close()

		contentBytes, err := io.ReadAll(fileContent)
		if err != nil {
			http.Error(w, "Unable to read file content", http.StatusInternalServerError)
			return
		}

		request := &fileservice.FileRequest{
			FileName:    file.Filename,
			FileContent: contentBytes,
		}

		response, err := client.SendFile(context.Background(), request)
		if err != nil {
			log.Fatalf("Failed to upload file: %v", err)
		}

		if !response.GetSuccess() {
			fmt.Printf("File upload failed: %s\n", response.GetMessage())
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
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
