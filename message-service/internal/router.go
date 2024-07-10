package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/ddr4869/msazoom/message-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetUp(s *Server) {
	r := s.router
	api := r.Group("/msazoom.messageservice/api")
	api.GET("/ping", s.Ping)

	//api.GET("/message/connect", s.ConnectMessageValid, s.ConnectMessage)
	api.GET("/message", utils.ParseJWT(), s.GetFriendMessageValid, s.GetFriendMessage)
	api.GET("/message/unread/count", utils.ParseJWT(), s.GetFriendMessageValid, s.GetNumberOfUnreadMessage)
	api.GET("/message/connect", s.ConnectMessageValid, s.ConnectMessage)
	api.GET("/message/condition", s.MessageConditionCheck)
}

func (s *Server) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

func (s *Server) Start() error {

	srv := &http.Server{
		Addr:    ":" + os.Getenv("SERVER_PORT"),
		Handler: s.router,
	}

	log.Printf("Listening and serving HTTP on %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("listen: %s\n", err)
		return err
	}

	return nil
}
