package main

import (
	"database/sql"
)

// Task representa uma tarefa no sistema.
type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	Priority    string
}

// Database representa a camada de acesso ao banco de dados.
type Database struct {
	db *sql.DB
}

// NewDatabase cria uma nova instância da camada de acesso ao banco de dados.
func NewDatabase(db *sql.DB) *Database {
	return &Database{db: db}
}

// CreateTask cria uma nova tarefa no banco de dados e retorna o ID da tarefa criada.
func (d *Database) CreateTask(title, description, status, priority string, assignedUsers []int) (int, error) {
	// Implementação para inserir uma nova tarefa no banco de dados e retornar o ID da tarefa criada
	// Exemplo simplificado:
	result, err := d.db.Exec("INSERT INTO tasks (title, description, status, priority) VALUES (?, ?, ?, ?)", title, description, status, priority)
	if err != nil {
		return 0, err
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Atribuir a tarefa aos usuários associados
	for _, userID := range assignedUsers {
		err := d.AssignTaskToUser(int(taskID), userID)
		if err != nil {
			// Se ocorrer um erro ao atribuir a tarefa a um usuário, podemos optar por reverter a criação da tarefa
			// Neste exemplo, vamos apenas retornar um erro, indicando que a tarefa foi criada, mas não foi atribuída a todos os usuários associados com sucesso
			return int(taskID), err
		}
	}

	return int(taskID), nil
}

// AssignTaskToUser atribui uma tarefa a um usuário no banco de dados.
func (d *Database) AssignTaskToUser(taskID, userID int) error {
	// Implementação para associar uma tarefa a um usuário no banco de dados
	// Exemplo simplificado:
	_, err := d.db.Exec("INSERT INTO task_user_associations (task_id, user_id) VALUES (?, ?)", taskID, userID)
	if err != nil {
		return err
	}

	return nil
}

// GetTaskByID retorna os detalhes de uma tarefa com base no ID da tarefa fornecido.
func (d *Database) GetTaskByID(taskID int) (Task, error) {
	// Implementação para buscar os detalhes de uma tarefa no banco de dados com base no ID da tarefa
	// Exemplo simplificado:
	var task Task
	err := d.db.QueryRow("SELECT id, title, description, priority, status FROM tasks WHERE id = ?", taskID).Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}
