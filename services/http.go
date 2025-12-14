package services

import (
	ctx2 "context"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/alphabatem/common/context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hxuan190/dex_aggregator/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HttpService struct {
	context.DefaultService
	BaseURL string
	Port    int
	FePort  int
	apiKey  string

	server *http.Server
}

func (svc HttpService) Id() string {
	return "http"
}

func (svc *HttpService) Configure(ctx *context.Context) error {
	// get Port if not fallback
	if port := os.Getenv("HTTP_PORT"); port != "" {
		var err error
		if svc.Port, err = strconv.Atoi(port); err != nil {
			return err
		}
	} else {
		svc.Port = 8000
	}

	// get FePort if not fallback
	if fePort := os.Getenv("FE_PORT"); fePort != "" {
		var err error
		if svc.FePort, err = strconv.Atoi(fePort); err != nil {
			return err
		}
	} else {
		svc.FePort = 5173
	}

	if baseURL := os.Getenv("BASE_URL"); baseURL != "" {
		svc.BaseURL = baseURL
	} else {
		svc.BaseURL = "http://localhost"
	}

	svc.apiKey = os.Getenv("API_KEY")
	if svc.apiKey == "" {
		return errors.New("API_KEY is not set")
	}

	return svc.DefaultService.Configure(ctx)
}

func (svc *HttpService) Shutdown() {
	err := svc.server.Shutdown(ctx2.TODO())
	if err != nil {
		return
	}
}

func (svc *HttpService) Start() error {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))

	// Validation endpoints
	r.GET("/ping", svc.ping)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/v1")

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(svc.Port),
		Handler: r,
	}

	svc.server = server
	listenErr := svc.server.ListenAndServe()
	if listenErr != nil {
		return listenErr
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return nil
}

func (svc *HttpService) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
