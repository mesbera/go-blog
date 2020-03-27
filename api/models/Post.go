package models

import (
	"time"
)

type Post struct {
	PostId    uint64    `gorm:"primary_key;auto_increment" json:"post_id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"type:text;not null;" json:"content"`
	Author    User      `json:"author"`
	AuthorId  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
