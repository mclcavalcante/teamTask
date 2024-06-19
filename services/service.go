package service

import "go.uber.org/zap"

// TaskInput representa a entrada para o serviço de criação de tarefa.
type Task struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Priority      string `json:"priority"`
	Status        string `json:"status"`
	AssignedUsers []int  `json:"assignedUsers"`
}

// Definição da estrutura de dados do usuário
type User struct {
	ID       int
	Name     string
	Role     string
	Email    string
	Password string
}

type Service interface {
	CreateTask(input Task) (int, error)
	GetVisibleTasksForUser(userID int) ([]Task, error)
	FilterTasksByStatusAndPriority(status, priority string) ([]Task, error)
	AssignMemberToTask(taskID, memberID int) error
	DeleteTask(taskID int) error
	GetTaskByID(taskID int) (Task, error)
	EditTask(taskID int, updatedTask Task) error
	GetAllTasks() ([]Task, error)

	RegisterNewUser(user User) (int, error)
	DeleteUser(userID int) error
	GetUserByID(userID int) (User, error)
}

type Repository interface {
	CreateTask(title, description, status, priority string, assignedUsers []int) (int, error)
	AssignTaskToUser(taskID int, userID int) error
	GetTaskByID(taskID int) (Task, error)
	GetTasksForUser(userID int) ([]Task, error)
	GetAllTasks() ([]Task, error)
	DeleteTask(taskID int) error
	UpdateTask(taskID int, updatedTask Task) error

	GetUserByEmail(email string) (User, error)
	GetUserByID(id int) (User, error)
	AddUser(user User) (int, error)
	RemoveUser(userID int) error
}

type teamTaskService struct {
	db  Repository
	log *zap.Logger
}

func NewService(db Repository, logger *zap.Logger) Service {
	return &teamTaskService{
		db:  db,
		log: logger,
	}
}
