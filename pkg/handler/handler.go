package handler

import (
	"github.com/ProxS/todo-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		list := api.Group("/lists")
		{
			list.POST("/", h.createList)
			list.GET("/", h.getAllLists)
			list.GET("/:id", h.getListById)
			list.PUT("/:id", h.updateList)
			list.DELETE("/:id", h.deleteList)
		}
	}

	return router
}
