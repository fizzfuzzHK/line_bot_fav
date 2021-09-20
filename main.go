package main

import (
	"os"

	infrastructure "github.com/fizzfuzzHK/line_bot_fav/infrastrcture"
	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := infrastructure.Connect()
	if err != nil {
		logrus.Infof("Error connecting DB: %v", err)
		db, _ = infrastructure.Connect()
	}
	defer db.Close()

	e := echo.New()
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

	Initialize(e, db)
}
