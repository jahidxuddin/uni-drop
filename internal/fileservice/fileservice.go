package fileservice

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Server struct {
	UnimplementedFileServiceServer
}

func (s *Server) SendFile(ctx context.Context, req *FileRequest) (*FileResponse, error) {
	downloadsDir, err := getDownloadsFolderPath()
	if err != nil {
		return &FileResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to determine Downloads directory: %v", err),
		}, err
	}

	filePath := filepath.Join(downloadsDir, req.GetFileName())

	err = os.WriteFile(filePath, req.GetFileContent(), 0644)
	if err != nil {
		return &FileResponse{
			Success: false,
			Message: fmt.Sprintf("Failed to save the file: %v", err),
		}, err
	}

	return &FileResponse{
		Success: true,
		Message: "File uploaded and saved to Downloads folder successfully",
	}, nil
}

func getDownloadsFolderPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get home directory: %v", err)
	}

	switch runtime.GOOS {
	case "windows":
		downloadsPath := filepath.Join(homeDir, "Downloads")
		return downloadsPath, nil
	case "linux", "darwin":
		downloadsPath := filepath.Join(homeDir, "Downloads")
		return downloadsPath, nil
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
