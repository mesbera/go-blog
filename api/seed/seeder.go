package seed

import (
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	"github.com/mesbera/go-blog/api/models"
)

var users = []models.User{
	models.User{
		Username: "Steven victor",
		Email:    "steven@gmail.com",
		Password: "password",
	},
	models.User{
		Username: "Martin Luther",
		Email:    "luther@gmail.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "Title 1",
		Content: "Hello world 1",
	},
	models.Post{
		Title:   "Title 2",
		Content: "Hello world 2",
	},
}

var comments = []models.Comment{
	models.Comment{
		Comment: "Comment of comments 1",
	},
	models.Comment{
		Comment: "Comment of comments 2",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Comment{}, &models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(user_id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Comment{}).AddForeignKey("author_id", "users(user_id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	err = db.Debug().Model(&models.Comment{}).AddForeignKey("post_id", "posts(post_id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorId = users[i].UserId
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
		comments[i].AuthorId = users[i].UserId
		comments[i].PostId = posts[i].PostId
		err = db.Debug().Model(&models.Comment{}).Create(&comments[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
