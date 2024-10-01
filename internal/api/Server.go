package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
	"testApi/internal/api/db"
	"testApi/internal/api/handlers"
	"testApi/internal/api/middleware"
	"testApi/internal/api/repositories"
)

type Server struct {
	host   string
	port   int
	router *gin.Engine
	handlers.Handler
}

func New(host string, port int) *Server {
	return &Server{host: host, port: port}
}

func (s *Server) Init() {
	db := db.New()

	s.Handler = handlers.Handler{
		Repo: repositories.NewAlbumRepository(db),
	}

	s.router = s.setRouter()

	s.start()
}

func (s *Server) start() {
	err := s.router.Run(s.host + ":" + strconv.Itoa(s.port))
	if err != nil {
		return
	}
}

func (s *Server) setRouter() *gin.Engine {
	router := gin.Default()

	apiGroup := router.Group("/api/v1")
	apiGroup.Use(middleware.RequestIdMiddleware())
	{
		apiGroup.GET("/albums", s.GetAlbums)
		apiGroup.GET("/albums/:id", s.GetAlbumByID)
		apiGroup.POST("/albums", s.PostAlbums)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
