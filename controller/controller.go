package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	domain "github.com/fizzfuzzHK/line_bot_fav/domain"
	"github.com/fizzfuzzHK/line_bot_fav/infrastrcture/database"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
)

func HandlerMainPage(i database.IUserRepository, u *domain.User) echo.HandlerFunc {
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
						replyMessage = "ぱおん"
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
				i.AddUser(u.UserId)

			}
		}
		return c.String(http.StatusOK, "")
	}
}
