package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/f4nt0md3v/bookstore_oauth-api/src/domain/access_token"
	"github.com/f4nt0md3v/bookstore_oauth-api/src/http"
	"github.com/f4nt0md3v/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	tokenHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", tokenHandler.GetById)
	router.POST("/oauth/access_token", tokenHandler.Create)

	err := router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		panic(err)
	}
}
