package main

import (
	controller "github.com/fizzfuzzHK/line_bot_fav/controller"
	infrastructure "github.com/fizzfuzzHK/line_bot_fav/infrastrcture"
	database "github.com/fizzfuzzHK/line_bot_fav/infrastrcture/database"
	echo "github.com/labstack/echo/v4"

	"github.com/jmoiron/sqlx"
)

func Initialize(e *echo.Echo, db *sqlx.DB) *infrastructure.Router {
	userRepository := database.NewUserRepository(db)
	lineBotController := controller.NewLineBotController(userRepository)
	r := infrastructure.NewRouter(e, lineBotController)
	return r
}
