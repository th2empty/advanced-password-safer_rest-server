package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/th2empty/auth-server/pkg/service"
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
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users")
		{
			users.POST("/delete", h.DeleteUser)
		}

		passwordLists := api.Group("/password_lists")
		{
			passwordLists.POST("/", h.CreateList) // create list
			passwordLists.GET("/", h.GetAllLists) // get all lists
			passwordLists.GET("/:id", h.GetList)  // get list by id
			//passwordLists.PUT("/:id", h.UpdateList)    // update list by id
			passwordLists.DELETE("/:id", h.DeleteList) // delete list by id

			passwords := passwordLists.Group(":id/passwords")
			{
				passwords.POST("/", h.AddPassword)    // add pass
				passwords.GET("/", h.GetAllPasswords) // get all passwords
			}
		}

		passwords := api.Group("passwords")
		{
			passwords.GET("/:id", h.GetPassword)       // get password by id
			passwords.PUT("/:id", h.UpdatePassword)    // update password by id
			passwords.DELETE("/:id", h.DeletePassword) // delete password by id
		}

	}

	return router
}
