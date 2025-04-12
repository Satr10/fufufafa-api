package model

type Post struct {
	ID       int    `gorm:"primaryKey;column:id" json:"id"`
	Content  string `gorm:"type:text" json:"content"`
	Datetime string `gorm:"index" json:"datetime"`
	Doksli   string `json:"doksli"`
	ImageURL string `gorm:"column:image_url" json:"image_url"`
}
