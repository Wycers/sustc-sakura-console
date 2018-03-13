package model

type Report struct {
	Model
	Content    string   `gorm:"size:128" json:"content"`
	Contact    string   `gorm:"type:mediumtext" json:"contact"`
}
