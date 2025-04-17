package models

type FilePath struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Path string `json:"path"`
}
