package service

// TaskInput representa a entrada para o serviço de criação de tarefa.
type Task struct {
	ID          int
	Title       string
	Description string
	Priority    string
	Status      string
	TeamMembers []int
}

type Service interface {
	CreateTask(input Task) (int, error)
}

type Repository interface {
	CreateTask(title, description, status, priority string, assignedUsers []int) (int, error)
	AssignTaskToUser(taskID, userID int) error
	GetTaskByID(taskID int) (Task, error)
}

type teamTaskService struct {
	db Repository
}

func NewService(db Repository) Service {
	return &teamTaskService{
		db: db,
	}
}
