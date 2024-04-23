package service

import (
	"errors"
	"fmt"
)

// MockDatabase representa uma implementação de banco de dados simulada para testes.
type MockDatabase struct {
	tasks       map[int]Task
	taskCounter int
}

// CreateTask cria uma nova tarefa simulada no banco de dados e retorna o ID da tarefa criada.
func (d *MockDatabase) CreateTask(title, description, status, priority string, assignedUsers []int) (int, error) {
	d.taskCounter++
	taskID := d.taskCounter

	task := Task{
		Title:       title,
		Description: description,
		Status:      status,
		Priority:    priority,
	}

	d.tasks[taskID] = task

	// Simula a atribuição de tarefa a usuários (neste exemplo, não faz nada)

	return taskID, nil
}

// GetTaskByID retorna os detalhes de uma tarefa simulada com base no ID da tarefa fornecido.
func (d *MockDatabase) GetTaskByID(taskID int) (Task, error) {
	task, ok := d.tasks[taskID]
	if !ok {
		return Task{}, errors.New("tarefa não encontrada")
	}
	return task, nil
}

// AssignTaskToUser atribui uma tarefa a um usuário no banco de dados simulado.
func (d *MockDatabase) AssignTaskToUser(taskID, userID int) error {
	// Simula a atribuição de tarefa a um usuário no banco de dados
	// Aqui você pode adicionar a lógica necessária para simular a atribuição de tarefa a um usuário
	fmt.Printf("Tarefa %d atribuída ao usuário %d\n", taskID, userID)
	return nil
}
