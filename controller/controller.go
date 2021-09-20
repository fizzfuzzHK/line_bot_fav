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

type LineBotController struct {
	u              *domain.User
	userRepository *database.UserRepository
	bot            *linebot.Client
}

func NewLineBotController(userRepository *database.UserRepository) *LineBotController {
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}

	user := &domain.User{}

	return &LineBotController{
		u:              user,
		userRepository: userRepository,
		bot:            bot,
	}
}

func (controller *LineBotController) HandleEvents() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する

		events, err := controller.bot.ParseRequest(c.Request())
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
					if _, err = controller.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
						fmt.Println(err)
					}
				case *linebot.StickerMessage:
					{
						replyMessage := fmt.Sprintf(
							"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
						if _, err = controller.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
							log.Print(err)
							fmt.Println(err)
						}
					}
				}
			} else if event.Type == linebot.EventTypeFollow {
				controller.u.UserId = event.Source.UserID
				controller.userRepository.AddUser(controller.u.UserId)

			}
		}
		return c.String(http.StatusOK, "")
	}
}
