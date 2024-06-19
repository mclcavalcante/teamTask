package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mclcavalcante/teamTask/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Default())

	api := router.Group("/task")
	{
		api.POST("/", init.Controller.CreateTaskData)
		api.GET("/:taskID", init.Controller.GetTaskByID)
		api.DELETE("/:taskID", init.Controller.DeleteTask)
		api.PUT("/:taskID", init.Controller.EditTask)
		api.GET("/all/:userID", init.Controller.GetVisibleTasksForUser)
		api.GET("/all", init.Controller.GetAllTasks)
	}

	router.GET("/filter/:status/:priority", init.Controller.FilterTasksByStatusAndPriority)
	router.POST("/:userID/:taskID", init.Controller.AssignMemberToTask)

	router.GET("/user/:userID", init.Controller.GetUserByID)
	router.POST("/user", init.Controller.RegisterNewUser)
	router.DELETE("/user/:userID", init.Controller.DeleteUser)

	return router
}
