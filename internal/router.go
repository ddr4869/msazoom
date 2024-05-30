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

	api.GET("/board", s.GetBoardList)
	api.GET("/board/:board_id", s.GetBoardWithIDValid, s.GetBoardWithID)

	api.POST("/board", utils.ParseJWT(), s.CreateBoardValid, s.CreateBoard)
	api.POST("/board/recommend", s.RecommendBoardValid, s.RecommendBoard)
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
