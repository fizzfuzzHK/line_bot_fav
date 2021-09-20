package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	domain "github.com/fizzfuzzHK/line_bot_fav/domain"
	infrastructure "github.com/fizzfuzzHK/line_bot_fav/infrastrcture"
	"github.com/jinzhu/gorm"
	echo "github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
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

	u := new(domain.User)

	e := echo.New()

	e.POST("/callback", handlerMainPage(db, u))

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	// res := weather.GetWeather()
	// テキストメッセージを生成する
	// テキストメッセージを友達登録しているユーザー全員に配信する

}

func handlerMainPage(db *gorm.DB, u *domain.User) echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		fmt.Println("callbacked")
		bot, err := linebot.New(
			os.Getenv("LINE_BOT_CHANNEL_SECRET"),
			os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
		)
		// エラーに値があればログに出力し終了する
		if err != nil {
			log.Fatal(err)
		}

		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			return nil
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				fmt.Println("message delivered")
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					replyMessage := message.Text
					if replyMessage == "ぴえん" {
						replyMessage = fmt.Sprintf("ぱおん")
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
						fmt.Println(err)
					}
				case *linebot.StickerMessage:
					{
						replyMessage := fmt.Sprintf(
							"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
							log.Print(err)
							fmt.Println(err)
						}
					}
				}
			} else if event.Type == linebot.EventTypeFollow {
				u.UserId = event.Source.UserID
				db.Create(&u)

			} else if event.Type == linebot.EventTypeUnfollow {
				user := event.Source.UserID
				fmt.Println(user)
			}
		}
		return c.String(http.StatusOK, "")
	}
}
