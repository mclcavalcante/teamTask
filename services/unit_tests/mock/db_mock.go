package mock

import (
	"errors"

	service "github.com/mclcavalcante/teamTask/services"
)

// MockDatabase representa uma implementação de banco de dados simulada para testes.
type MockDatabase struct {
	tasks       map[int]service.Task
	taskCounter int
	tasksByUser map[int][]service.Task // Mapeamento de IDs de usuário para tarefas atribuídas a esse usuário

	userCounter int
	usersByID   map[int]service.User // Mapeamento de IDs de usuário para usuários
}

// CreateTask cria uma nova tarefa simulada no banco de dados e retorna o ID da tarefa criada.
func (d *MockDatabase) CreateTask(title, description, status, priority string, assignedUsers []int) (int, error) {
	d.taskCounter++
	taskID := d.taskCounter

	task := service.Task{
		Title:         title,
		Description:   description,
		Status:        status,
		Priority:      priority,
		ID:            taskID,
		AssignedUsers: assignedUsers,
	}

	d.tasks[taskID] = task

	return taskID, nil
}

// GetTaskByID retorna os detalhes de uma tarefa simulada com base no ID da tarefa fornecido.
func (d *MockDatabase) GetTaskByID(taskID int) (service.Task, error) {
	task, ok := d.tasks[taskID]
	if !ok {
		return service.Task{}, errors.New("tarefa não encontrada")
	}
	return task, nil
}

// AssignTaskToUser atribui uma tarefa a um usuário no banco de dados simulado.
func (d *MockDatabase) AssignTaskToUser(taskID int, userID int) error {
	// Simula a atribuição de tarefa a um usuário no banco de dados
	// Verificar se a tarefa existe
	newMemberTask, taskExists := d.tasks[taskID]
	if !taskExists {
		return errors.New("tarefa não encontrada")
	}

	// Verificar se o membro da equipe existe
	_, memberExists := d.usersByID[userID]
	if !memberExists {
		return errors.New("membro da equipe não encontrado")
	}
	newMemberTask.AssignedUsers = append(newMemberTask.AssignedUsers, userID)

	// Associar o membro da equipe à tarefa
	d.tasks[taskID] = newMemberTask

	return nil
}

// GetUserByEmail simula a busca de um usuário no banco de dados pelo seu e-mail.
func (d *MockDatabase) GetUserByEmail(email string) (service.User, error) {

	for _, user := range d.usersByID {
		if user.Email == email {
			return user, nil
		}
	}

	return service.User{}, nil // Usuário não encontrado
}

// AddUser simula a adição de um novo usuário ao banco de dados.
func (d *MockDatabase) AddUser(newUser service.User) (int, error) {

	for _, user := range d.usersByID {
		if user.Email == newUser.Email {
			return 0, errors.New("e-mail já está em uso")
		}
	}

	d.userCounter++
	newUser.ID = d.userCounter

	d.usersByID[newUser.ID] = newUser
	return newUser.ID, nil
}

// GetUserByID simula a busca de um usuário no banco de dados pelo seu ID.
func (d *MockDatabase) GetUserByID(id int) (service.User, error) {
	user, ok := d.usersByID[id]
	if !ok {
		return service.User{}, errors.New("usuário não existe") // Usuário não encontrado
	}
	return user, nil
}

// GetTasksForUser simula a obtenção de tarefas atribuídas ao usuário especificado.
func (d *MockDatabase) GetTasksForUser(userID int) ([]service.Task, error) {
	var tasksForUser []service.Task
	for _, task := range d.tasks {
		for _, member := range task.AssignedUsers {
			if member == userID {
				tasksForUser = append(tasksForUser, task)
			}
		}

	}

	return tasksForUser, nil
}

func (d *MockDatabase) GetAllTasks() ([]service.Task, error) {
	tasks := make([]service.Task, 0, len(d.tasks))
	for _, task := range d.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// DeleteTask simula a exclusão de uma tarefa do banco de dados.
func (d *MockDatabase) DeleteTask(taskID int) error {
	// Verificar se a tarefa existe
	_, taskExists := d.tasks[taskID]
	if !taskExists {
		return errors.New("tarefa não encontrada")
	}

	// Excluir a tarefa do banco de dados mockado
	delete(d.tasks, taskID)

	return nil
}

// UpdateTask é um método para atualizar uma tarefa existente no banco de dados.
func (d *MockDatabase) UpdateTask(taskID int, updatedTask service.Task) error {
	// Verificar se a tarefa existe
	_, ok := d.tasks[taskID]
	if !ok {
		return errors.New("tarefa não encontrada")
	}

	// Atualizar a tarefa
	d.tasks[taskID] = updatedTask

	return nil
}

// RemoveUser é um método para remover um usuário existente do banco de dados.
func (d *MockDatabase) RemoveUser(userID int) error {
	// Verificar se o usuário existe
	_, ok := d.usersByID[userID]
	if !ok {
		return errors.New("usuário não encontrado")
	}

	// Remover o usuário
	delete(d.usersByID, userID)

	return nil
}

// NewTestDatabase cria uma nova instância do banco de dados simulado para testes.
func NewTestRepository() *MockDatabase {
	return &MockDatabase{
		tasks:       make(map[int]service.Task),
		taskCounter: 0,
		tasksByUser: make(map[int][]service.Task),

		userCounter: 0,
		usersByID:   make(map[int]service.User),
	}
}
