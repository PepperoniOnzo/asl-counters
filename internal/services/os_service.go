package services

import (
	"errors"
	"io"
	"os"

	"github.com/PepperoniOnzo/asl-counters/internal/models"
)

const assets = "/assets"

func GetContentFromFolder(path string) (*models.ContentResponse, error) {

	entries, err := os.ReadDir(assets + path)
	if err != nil {
		return nil, err
	}

	var res []*models.FileStructure

	for i := range entries {
		res = append(res, &models.FileStructure{
			Id:          entries[i].Name(),
			IsDirectory: entries[i].IsDir(),
		})
	}

	return &models.ContentResponse{Content: res}, nil
}

func GetImage(path string) ([]byte, error) {
	file, err := os.Open(assets + path)
	if err != nil {
		return nil, errors.New("image not found")
	}
	defer file.Close()

	imgData, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("failed to read image")
	}

	return imgData, nil
}
