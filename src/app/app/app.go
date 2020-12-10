package app

import (
	"microservice_tut/bookstore_oauth_api/src/domain/access_token"
	"microservice_tut/bookstore_oauth_api/src/http"
	"microservice_tut/bookstore_oauth_api/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	dbRepo := db.NewRepo()
	atService := access_token.NewService(dbRepo)
	atHandler := http.NewHandler(atService)

	router.GET("/oauth/access_token/:token_id", atHandler.GetByID)

	router.Run(":8080")
}
