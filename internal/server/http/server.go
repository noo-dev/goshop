package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/quangdangfit/gocommon/logger"
	"gorm.io/gorm"
	"goshop/pkg/config"
	"log"
)

type Server struct {
	engine    *gin.Engine
	cfg       *config.Schema
	db        *gorm.DB
	validator *validator.Validate
}

func NewServer(db *gorm.DB, validator *validator.Validate) *Server {
	return &Server{
		engine:    gin.Default(),
		cfg:       config.GetConfig(),
		db:        db,
		validator: validator,
	}
}

func (s Server) Run() error {
	_ = s.engine.SetTrustedProxies(nil)

	// Map routes
	if err := s.MapRoutes(); err != nil {
		log.Fatalf("Map routes error: %v", err)
	}

	// Start http server
	logger.Info("HTTP server is listening on PORT: ", s.cfg.HttpPort)
	if err := s.engine.Run(fmt.Sprintf(":%d", s.cfg.HttpPort)); err != nil {
		log.Fatalf("Running HTTP server: %v", err)
	}

	return nil
}

func (s Server) GetEngine() *gin.Engine {
	return s.engine
}
