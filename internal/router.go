package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/ddr4869/msazoom/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetUp(s *Server) {
	r := s.router
	api := r.Group("/api")
	api.GET("/ping", s.Ping)

	api.POST("/user", s.UserCreateValid, s.UserCreate)
	api.POST("/user/login", s.UserLoginValid, s.UserLogin)
	api.POST("/user/friend", utils.ParseJWT(), s.AddFriendValid, s.AddFriend)
	api.GET("/user/friend", utils.ParseJWT(), s.GetFriendList)
	api.GET("/user/friend/follower", utils.ParseJWT(), s.GetFriendRequestList)
	api.GET("/user/friend/check", utils.ParseJWT(), s.CheckFriendValid, s.CheckFriend)
	api.DELETE("/user/friend", utils.ParseJWT(), s.RemoveFriendValid, s.RemoveFriend)

	api.GET("/board", s.GetBoardList)
	api.GET("/board/:board_id", s.GetBoardWithIDValid, s.GetBoardWithID)

	api.POST("/board", utils.ParseJWT(), s.CreateBoardValid, s.CreateBoard)
	api.POST("/board/recommend", s.RecommendBoardValid, s.RecommendBoard)
	api.POST("/board/remove", utils.ParseJWT(), s.DeleteBoardValid, s.DeleteBoard)

	//api.GET("/message/connect", s.ConnectMessageValid, s.ConnectMessage)
	api.GET("/message", utils.ParseJWT(), s.GetFriendMessageValid, s.GetFriendMessage)
	api.GET("/message/connect", s.ConnectMessageValid, s.ConnectMessage)
	api.GET("/message/condition", s.MessageConditionCheck)

	api.GET("/chat/:chat_id", utils.ParseJWT(), s.GetChatValid, s.GetChat)
	api.GET("/chat/create", utils.ParseJWT(), s.CreateChatValid, s.CreateChat)
	api.GET("/chat", utils.ParseJWT(), s.GetChatList)
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
