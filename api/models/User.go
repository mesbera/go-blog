package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    uint32    `gorm:"primary_key;auto_increment" json:"user_id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) Prepare() {
	u.UserId = 0
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Username == "" {
			return errors.New("Required username")
		}
		if u.Password == "" {
			return errors.New("Required password")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required password")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil

	default:
		if u.Username == "" {
			return errors.New("Required username")
		}
		if u.Password == "" {
			return errors.New("Required password")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid email")
		}
		return nil
	}
}

//CountPostsAndCommentsByUserId counts all the posts and comments added by the user in the past X days
func (u *User) CountPostsAndCommentsByUserId(db *gorm.DB, uid uint32, days int) (int32, error) {

	var err error
	var countedposts int32
	var countedcomments int32

	var startTime time.Time
	startTime = time.Now().AddDate(0, 0, -days)
	err = db.Debug().Model(Post{}).Where("author_id = ? AND created_at > ? ", uid, startTime).Count(&countedposts).Error
	if err != nil {
		return -1, err
	}
	err = db.Debug().Model(Comment{}).Where("author_id = ? AND created_at > ? ", uid, startTime).Count(&countedcomments).Error
	if err != nil {
		return -1, err
	}

	return countedposts + countedcomments, err
}
