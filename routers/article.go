package routers

import (
	"gin-boilerplate/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArticleRoutes(route *gin.Engine, db *gorm.DB) {
	ctrl := controllers.Controller{DB: db}
	v1 := route.Group("article")
	v1.GET("/", ctrl.ArticleGetList)
	v1.GET("/:id", ctrl.ArticleGetById)
	v1.POST("/", ctrl.ArticleAdd)
	v1.PUT("/:id", ctrl.ArticleEdit)
	v1.DELETE("/:id", ctrl.ArticleDelete)
}
