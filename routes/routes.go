package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/heyrmi/go-webapi/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	//All CRUD
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
	authenticated.GET("/events/registrations", getRegistrations)
	authenticated.GET("/user", getUsers)

	//user management service
	server.POST("/signup", signup)
	server.POST("/login", login)

}
