package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
	"testApi/internal/api/DB"
	"testApi/internal/api/Handlers"
	"testApi/internal/api/Repositories"
)

type Server struct {
	host   string
	port   int
	router *gin.Engine
	Handlers.Handler
}

func New(host string, port int) *Server {
	return &Server{host: host, port: port}
}

func (s *Server) Init() {
	db := DB.New()

	s.Handler = Handlers.Handler{
		Repo: Repositories.NewAlbumRepository(db),
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
	{
		apiGroup.GET("/albums", s.GetAlbums)
		apiGroup.GET("/albums/:id", s.GetAlbumByID)
		apiGroup.POST("/albums", s.PostAlbums)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
