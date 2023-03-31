package app

import (
	"eight-learn/config"
	"eight-learn/repository"
	"eight-learn/route"
	"eight-learn/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.PSQL.DB)
	app := service.NewService(repo)
	route.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
