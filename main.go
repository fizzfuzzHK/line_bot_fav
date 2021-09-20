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
		// Heroku用 アプリの起動に合わせてDBが起動できないことがあるので再接続を試みる
		db, _ = infrastructure.Connect()
	}
	defer db.Close()

	e := echo.New()
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

	Initialize(e, db)
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	// res := weather.GetWeather()
	// テキストメッセージを生成する
	// テキストメッセージを友達登録しているユーザー全員に配信する

}
