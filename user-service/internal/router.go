package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/ddr4869/msazoom/user-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func SetUp(s *Server) {
	r := s.router
	api := r.Group("/msazoom.userservice/api")
	api.GET("/ping", s.Ping)

	api.POST("/user", s.UserCreateValid, s.UserCreate)
	api.POST("/user/login", s.UserLoginValid, s.UserLogin)
	api.POST("/user/friend", utils.ParseJWT(), s.AddFriendValid, s.AddFriend)
	api.GET("/user/friend", utils.ParseJWT(), s.GetFriendList)
	api.GET("/user/friend/follower", utils.ParseJWT(), s.GetFriendRequestList)
	api.GET("/user/friend/check", utils.ParseJWT(), s.CheckFriendValid, s.CheckFriend)
	api.DELETE("/user/friend", utils.ParseJWT(), s.RemoveFriendValid, s.RemoveFriend)
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
