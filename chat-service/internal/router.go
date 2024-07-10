package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/ddr4869/msazoom/chat-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetUp(s *Server) {
	r := s.router
	api := r.Group("/msazoom.chatservice/api")
	api.GET("/ping", s.Ping)

	api.GET("/chat/:chat_id", utils.ParseJWT(), s.GetChatValid, s.GetChat)
	api.GET("/chat/create", utils.ParseJWT(), s.CreateChatValid, s.CreateChat)
	api.GET("/chat", utils.ParseJWT(), s.GetChatList)
	api.POST("/chat/check_password", utils.ParseJWT(), s.CheckPasswordValid, s.CheckPassword)
	// ws
	api.GET("/chat/join", s.JoinChatValid, s.JoinChat)
	api.GET("/chat/random", utils.ParseJWT(), s.RandomChating)
	api.GET("/chat/room_condition", s.RoomConditionCheck)
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
