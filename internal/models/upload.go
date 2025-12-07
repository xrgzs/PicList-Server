package models

import (
	"github.com/google/uuid"
)

type UploadRequest struct {
	List []string `json:"list"`
}

type UploadResponse struct {
	Success    bool         `json:"success"`
	Result     []string     `json:"result,omitempty"`
	Message    string       `json:"message,omitempty"`
	FullResult []FullResult `json:"fullResult,omitempty"`
}

type FullResult struct {
	FileName  string    `json:"fileName"`
	ImgURL    string    `json:"imgUrl"`
	Extname   string    `json:"extname"`
	Type      string    `json:"type"`
	ID        uuid.UUID `json:"id"`
	CreatedAt int64     `json:"createdAt"`
	UpdatedAt int64     `json:"updatedAt"`
	// Width    int    `json:"width"`
	// Height   int    `json:"height"`
	// Hash        string `json:"hash"`
	// GalleryPath string `json:"galleryPath"`
	// ShortURL  string    `json:"shortUrl"`
}

type DeleteRequest struct {
	List []FullResult `json:"list"`
}
