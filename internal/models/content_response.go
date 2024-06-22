package models

type FileStructure struct {
	Id          string `json:"id"`
	IsDirectory bool   `json:"is_directory"`
}

type ContentResponse struct {
	Content []*FileStructure `json:"content"`
}
