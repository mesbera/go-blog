package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Comment struct {
	CommentId uint64    `gorm:"primary_key;auto_increment" json:"comment_id"`
	Author    User      `json: author`
	AuthorId  uint32    `gorm:"not null" json:"author_id"`
	Post      Post      `json: post`
	PostId    uint64    `gorm:"not null" json:"post_id"`
	Comment   string    `gorm:"type:text;not null;" json:"comment"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Comment) Prepare() {
	c.CommentId = 0
	c.Author = User{}
	c.Post = Post{}
	c.Comment = html.EscapeString(strings.TrimSpace(c.Comment))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *Comment) Validate() error {

	if c.Comment == "" {
		return errors.New("Required comment")
	}
	if c.PostId < 1 {
		return errors.New("Required post")
	}
	if c.AuthorId < 1 {
		return errors.New("Required author")
	}
	return nil
}

func (c *Comment) IsPostAuthor(db *gorm.DB) (bool, error) {
	var err error

	post := Post{}

	err = db.Debug().Model(&Post{}).Where("post_id = ? AND author_id = ?", c.PostId, c.AuthorId).Take(&post).Error
	if err != nil {
		return false, err
	}
	if post.PostId != 0 {
		return true, nil
	}
	return false, nil
}

func (c *Comment) SaveComment(db *gorm.DB) (*Comment, error) {
	var err error

	err = db.Debug().Model(&Comment{}).Create(&c).Error
	if err != nil {
		return &Comment{}, err
	}
	if c.CommentId != 0 {
		err = db.Debug().Model(&User{}).Where("user_id = ?", c.AuthorId).Take(&c.Author).Error
		if err != nil {
			return &Comment{}, err
		}
		err = db.Debug().Model(&Post{}).Where("post_id = ?", c.PostId).Take(&c.Post).Error
		if err != nil {
			return &Comment{}, err
		}
	}
	return c, nil
}
