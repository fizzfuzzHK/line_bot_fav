package main

import (
	"os"

	controller "github.com/fizzfuzzHK/line_bot_fav/controller"
	domain "github.com/fizzfuzzHK/line_bot_fav/domain"
	infrastructure "github.com/fizzfuzzHK/line_bot_fav/infrastrcture"
	database "github.com/fizzfuzzHK/line_bot_fav/infrastrcture/database"
	echo "github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	db, err := infrastructure.Connect()
	if err != nil {
		logrus.Infof("Error connecting DB: %v", err)
		// Heroku用 アプリの起動に合わせてDBが起動できないことがあるので再接続を試みる
		db, _ = infrastructure.Connect()
	}
	defer db.Close()

	userRepo := database.NewUserRepository(db)
	user := new(domain.User)

	e := echo.New()

	e.POST("/callback", controller.HandlerMainPage(userRepo, user))

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	// res := weather.GetWeather()
	// テキストメッセージを生成する
	// テキストメッセージを友達登録しているユーザー全員に配信する

}
