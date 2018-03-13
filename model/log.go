package model

type Log struct {
	Model
	StudentID    string    `gorm:"size:32" json:"sid"`
	
}
