package app

import (
	"github.com/bm1905/bookstore_oauth_api/src/domain/access_token"
	"github.com/bm1905/bookstore_oauth_api/src/http"
	"github.com/bm1905/bookstore_oauth_api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	repository := db.NewRepository()
	atService := access_token.NewService(repository)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
