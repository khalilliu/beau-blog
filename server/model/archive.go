package model

type Archive struct {
	Year      string `gorm:"size:4" json:"year"`
	Month     string `gorm:"size:2" json:"month"`
	PostCount int    `json:"postCount"`
}
