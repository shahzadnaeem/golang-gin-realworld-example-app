package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/gothinkster/golang-gin-realworld-example-app/articles"
	"github.com/gothinkster/golang-gin-realworld-example-app/common"
	"github.com/gothinkster/golang-gin-realworld-example-app/users"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	users.AutoMigrate()
	db.AutoMigrate(&articles.ArticleModel{})
	db.AutoMigrate(&articles.TagModel{})
	db.AutoMigrate(&articles.FavoriteModel{})
	db.AutoMigrate(&articles.ArticleUserModel{})
	db.AutoMigrate(&articles.CommentModel{})
}

func main() {

	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	articles.TagsAnonymousRegister(v1.Group("/tags"))

	v1.Use(users.AuthMiddleware(true))
	users.UserRegister(v1.Group("/user"))
	users.ProfileRegister(v1.Group("/profiles"))

	articles.ArticlesRegister(v1.Group("/articles"))

	// Simple /api/ping -> "pong"
	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Check if wew have already created user below
	var user users.UserModel
	chk1 := db.Begin()
	res := chk1.First(&user)
	chk1.Commit()
	if res.Error != nil {
		// test 1 to 1
		tx1 := db.Begin()
		userA := users.UserModel{
			Username: "AAAAAAAAAAAAAAAA",
			Email:    "aaaa@g.cn",
			Bio:      "hehddeda",
			Image:    nil,
		}
		tx1.Save(&userA)
		tx1.Commit()
		fmt.Println(userA)
	} else {
		fmt.Println("Initial user exists")
	}

	// Old initialisation code - ran always :(
	// // test 1 to 1
	// tx1 := db.Begin()
	// userA := users.UserModel{
	// 	Username: "AAAAAAAAAAAAAAAA",
	// 	Email:    "aaaa@g.cn",
	// 	Bio:      "hehddeda",
	// 	Image:    nil,
	// }
	// tx1.Save(&userA)
	// tx1.Commit()
	// fmt.Println(userA)

	// TODO: What is this - seems broken!
	//db.Save(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//})
	//var userAA ArticleUserModel
	//db.Where(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//}).First(&userAA)
	//fmt.Println(userAA)

	r.Run() // listen and serve on 0.0.0.0:8080
}
