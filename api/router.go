package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	//
	// Company
	api.POST("/companies", handlerV1.CreateCompany)
	api.GET("/companies/:id", handlerV1.GetCompany)
	api.GET("/companies", handlerV1.GetListCompanies)
	api.PUT("/companies/:id", handlerV1.UpdateCompany)
	api.DELETE("/companies/:id", handlerV1.DeleteCompany)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
