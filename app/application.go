package app

import (
	"github.com/gin-gonic/gin"
	http "github.com/ravishen/bookstore_oauth-api/http/access_token/access_token_http"
	"github.com/ravishen/bookstore_oauth-api/repository/db"
	"github.com/ravishen/bookstore_oauth-api/src/domain/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")

}
