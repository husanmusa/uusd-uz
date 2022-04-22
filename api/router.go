package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/husanmusa/uusd-uz/api/docs" // swag
	"github.com/husanmusa/uusd-uz/api/handlers/v1"
	"github.com/husanmusa/uusd-uz/config"
	"github.com/husanmusa/uusd-uz/pkg/logger"
)

type Option struct {
	Db     *sqlx.DB
	Conf   config.Config
	Logger logger.Logger
}

//func CORSMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
//		if c.Request.Method == "OPTIONS" {
//			c.AbortWithStatus(204)
//			return
//		}
//
//		c.Next()
//	}
//}

func Routers(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//router.Use(CORSMiddleware())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Db:     option.Db,
		Logger: option.Logger,
		Cfg:    option.Conf,
	})

	api := router.Group("/v1")
	//api.Static("/media", "./media")

	// Company
	api.POST("/companies", handlerV1.CreateCompany)
	api.GET("/companies/:id", handlerV1.GetCompany)
	api.GET("/companies", handlerV1.GetListCompanies)
	api.PUT("/companies/:id", handlerV1.UpdateCompany)
	api.DELETE("/companies/:id", handlerV1.DeleteCompany)

	// Service
	api.POST("/services", handlerV1.CreateService)
	api.GET("/services/:id", handlerV1.GetService)
	api.GET("/services", handlerV1.GetListServices)
	api.PUT("/services/:id", handlerV1.UpdateService)
	api.DELETE("/services/:id", handlerV1.DeleteService)

	// Set
	api.POST("/sets", handlerV1.CreateSet)
	api.GET("/sets/:id", handlerV1.GetSet)
	api.GET("/sets", handlerV1.GetListSets)
	api.PUT("/sets/:id", handlerV1.UpdateSet)
	api.DELETE("/sets/:id", handlerV1.DeleteSet)

	// Package
	api.POST("/packages", handlerV1.CreatePackage)
	api.GET("/packages/:id", handlerV1.GetPackage)
	api.GET("/packages", handlerV1.GetListPackages)
	api.PUT("/packages/:id", handlerV1.UpdatePackage)
	api.DELETE("/packages/:id", handlerV1.DeletePackage)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
