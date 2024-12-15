package filetransfer

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Server struct{
	UnimplementedFileTransferServer
}

func (s *Server) UploadFile(stream FileTransfer_UploadFileServer) error {
	var fileName string
	fileData := make([][]byte, 0)

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to receive chunk: %v", err)
		}

		fileName = chunk.Filename
		fileData = append(fileData, chunk.Data)
	}

	// Write the file to disk
	outputFile, err := os.Create(fileName)
	if err != nil {
		return stream.SendAndClose(&UploadStatus{
			Success: false,
			Message: fmt.Sprintf("failed to create file: %v", err),
		})
	}
	defer outputFile.Close()

	for _, chunk := range fileData {
		_, err := outputFile.Write(chunk)
		if err != nil {
			return stream.SendAndClose(&UploadStatus{
				Success: false,
				Message: fmt.Sprintf("failed to write chunk: %v", err),
			})
		}
	}

	return stream.SendAndClose(&UploadStatus{
		Success: true,
		Message: "File uploaded successfully",
	})
}

func (s *Server) DownloadFile(req *FileRequest, stream FileTransfer_DownloadFileServer) error {
	file, err := os.Open(req.Filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	chunkIndex := int64(0)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read file: %v", err)
		}

		err = stream.Send(&FileChunk{
			Filename:   req.Filename,
			Data:       buffer[:n],
			ChunkIndex: chunkIndex,
		})
		if err != nil {
			return fmt.Errorf("failed to send chunk: %v", err)
		}

		chunkIndex++
	}

	return nil
}

func (s *Server) ListFiles(ctx context.Context, req *DirectoryRequest) (*FileList, error) {
	files := []string{}

	err := filepath.Walk(req.DirectoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %v", err)
	}

	return &FileList{
		Files: files,
	}, nil
}
