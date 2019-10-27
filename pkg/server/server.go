package server

import (
	"net/http"
	"time"

	"github.com/mytheta/line-bot-template/conf"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/mytheta/line-bot-template/pkg/handler"
)

func Init() {
	i := injector()
	r := router(i)
	r.Run(":8080")
}

func router(h handler.Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.POST("/hook", h.PostMessage)

	return router
}

//依存関係
func injector() handler.Handler {
	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	bot, err := linebot.New(conf.GetBotConfig().ChannelSecret, conf.GetBotConfig().ChannelToken, linebot.WithHTTPClient(client))
	if err != nil {
		panic(err)
	}
	return handler.NewHandler(bot)
}
