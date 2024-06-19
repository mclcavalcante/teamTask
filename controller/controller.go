package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mclcavalcante/teamTask/services"
	"go.uber.org/zap"
)

type Controller interface {
	CreateTaskData(ctx *gin.Context)
	RegisterNewUser(ctx *gin.Context)
	GetVisibleTasksForUser(ctx *gin.Context)
	FilterTasksByStatusAndPriority(ctx *gin.Context)
	AssignMemberToTask(ctx *gin.Context)
	DeleteTask(ctx *gin.Context)
	EditTask(ctx *gin.Context)
	GetAllTasks(ctx *gin.Context)
	GetUserByID(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	GetTaskByID(ctx *gin.Context)
}

type TaskController struct {
	svc service.Service
	log *zap.Logger
}

func (c TaskController) CreateTaskData(ctx *gin.Context) {
	var request service.Task
	if err := ctx.ShouldBindJSON(&request); err != nil {
		c.log.Error(err.Error())
	}

	input := service.Task{
		ID:            request.ID,
		Title:         request.Title,
		Description:   request.Description,
		Priority:      request.Priority,
		Status:        request.Status,
		AssignedUsers: request.AssignedUsers,
	}

	task_id, err := c.svc.CreateTask(input)
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, task_id)
}

func (c TaskController) RegisterNewUser(ctx *gin.Context) {
	var request service.User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		//TODO
		c.log.Error(err.Error())
		return
	}

	input := service.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	c.log.Info("CONTROLLER: " + input.Name)

	user_id, err := c.svc.RegisterNewUser(input)
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, user_id)
}

func (c TaskController) GetVisibleTasksForUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("userID"))

	tasks, err := c.svc.GetVisibleTasksForUser(userId)

	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c TaskController) FilterTasksByStatusAndPriority(ctx *gin.Context) {
	status := ctx.Param("status")
	priority := ctx.Param("priority")

	tasks, err := c.svc.FilterTasksByStatusAndPriority(status, priority)
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c TaskController) AssignMemberToTask(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Param("userID"))
	taskID, _ := strconv.Atoi(ctx.Param("taskID"))

	err := c.svc.AssignMemberToTask(taskID, userID)
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}
}

func (c TaskController) DeleteTask(ctx *gin.Context) {
	taskID, _ := strconv.Atoi(ctx.Param("taskID"))

	err := c.svc.DeleteTask(taskID)
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}
}

func (c TaskController) EditTask(ctx *gin.Context) {
	taskID, _ := strconv.Atoi(ctx.Param("taskID"))

	var request service.Task
	if err := ctx.ShouldBindJSON(&request); err != nil {
		//TODO
		c.log.Error(err.Error())
	}

	input := service.Task{
		Title:         request.Title,
		Description:   request.Description,
		Priority:      request.Priority,
		Status:        request.Status,
		AssignedUsers: request.AssignedUsers,
	}

	err := c.svc.EditTask(taskID, input)
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}

}

func (c TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.svc.GetAllTasks()
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c TaskController) GetUserByID(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("userID"))

	user, err := c.svc.GetUserByID(userId)
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)

	}

	ctx.JSON(http.StatusOK, user)
}

func (c TaskController) DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("userID"))

	err := c.svc.DeleteUser(userId)
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)

	}
}

func (c TaskController) GetTaskByID(ctx *gin.Context) {
	taskID, _ := strconv.Atoi(ctx.Param("taskID"))

	task, err := c.svc.GetTaskByID(taskID)
	if err != nil {
		c.log.Error(err.Error())
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, task)
}

func ControllerInit(service service.Service, logger *zap.Logger) *TaskController {
	return &TaskController{
		svc: service,
		log: logger,
	}
}
