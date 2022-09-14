package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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

func AddRoutes(r *gin.Engine) {

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

	testAuth.GET("/", logger.SetLogger(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func InitDb(db *gorm.DB) {
	// Check if wew have already created user below
	var user users.UserModel
	chk1 := db.Begin()
	res := chk1.First(&user)
	chk1.Commit()
	if res.Error != nil {
		log.Info().Msg("INIT: Adding users")

		tx1 := db.Begin()
		userA := users.UserModel{
			Username: "AAAAAAAAAAAAAAAA",
			Email:    "aaaa@g.cn",
			Bio:      "hehddeda",
			Image:    nil,
		}
		tx1.Save(&userA)

		userB := users.UserModel{
			Username: "BBBBBBBBBBBBBBBB",
			Email:    "bbbb@g.cn",
			Bio:      "Bee bop",
			Image:    nil,
		}
		tx1.Save(&userB)
		tx1.Commit()

		fmt.Println(userA)
		fmt.Println(userB)
		fmt.Println()
	} else {
		log.Info().Msg("INIT: Initial user exists")
	}
}

func main() {

	// Setup logging
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	db := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()
	AddRoutes(r)

	InitDb(db)

	r.Run() // listen and serve on 0.0.0.0:8080
}
