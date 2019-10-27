package handler

import (
	"log"

	"github.com/mytheta/line-bot-template/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Handler struct {
	lineBotClient *linebot.Client
}

func NewHandler(client *linebot.Client) Handler {
	return Handler{lineBotClient: client}
}

func (h *Handler) PostMessage(c *gin.Context) {
	received, err := h.lineBotClient.ParseRequest(c.Request)
	event, message, ok := validate(received)
	if ok {
		postMessage := linebot.NewTextMessage(service.Create(message))
		if _, err = h.lineBotClient.ReplyMessage(event.ReplyToken, postMessage).Do(); err != nil {
			log.Print(err)
		}
	}
}

//lineのメッセージからtextのものを抽出し、メッセージとtrueを返す。メッセージがない場合は、(nil,false)
func validate(received []*linebot.Event) (*linebot.Event, string, bool) {
	for _, event := range received {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				source := event.Source
				if source.Type == linebot.EventSourceTypeUser {
					return event, message.Text, true
				}
			}
		}
	}
	return nil, "", false
}
