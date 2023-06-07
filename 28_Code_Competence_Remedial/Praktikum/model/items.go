package model

type Item struct {
	ID       string `gorm:"primaryKey;type:varchar(36);uniqueindex;not null"`
	Name     string `json:"name" form:"name"`
	Category int    `json:"category" form:"category"`
}
