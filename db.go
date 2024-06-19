package main

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"

	service "github.com/mclcavalcante/teamTask/services"
	"go.uber.org/zap"
)

// Database representa a camada de acesso ao banco de dados.
type Database struct {
	db  *sql.DB
	log *zap.Logger
}

// NewDatabase cria uma nova instância da camada de acesso ao banco de dados.
func NewDatabase(db *sql.DB) *Database {
	return &Database{db: db}
}

// GetUserByEmail busca um usuário no banco de dados pelo seu e-mail.
func (d *Database) GetUserByEmail(email string) (service.User, error) {
	query := "SELECT id, name, email, password FROM users WHERE email = ?"
	row := d.db.QueryRow(query, email)

	var user service.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return service.User{}, nil // Usuário não encontrado
		}
		return service.User{}, err
	}

	return user, nil
}

// AddUser adiciona um novo usuário ao banco de dados.
func (d *Database) AddUser(user service.User) (int, error) {
	query := "INSERT INTO User (name, email, password) VALUES (?, ?, ?)"
	result, err := d.db.Exec(query, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	d.log.Info(strconv.Itoa(int(id)))
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// GetUserById busca um usuário no banco de dados pelo seu ID.
func (d *Database) GetUserById(id int) (service.User, error) {
	query := "SELECT name, email, password FROM User WHERE id = ?"
	row := d.db.QueryRow(query, id)

	var user service.User
	err := row.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return service.User{}, errors.New("Usuário não existe") // Usuário não encontrado
		}
		return service.User{}, err
	}

	user.ID = id
	return user, nil
}

// CreateTask cria uma nova tarefa no banco de dados e retorna o ID da tarefa criada.
func (d *Database) CreateTask(title, description, status, priority string, assignedUsers []int) (int, error) {
	// Implementação para inserir uma nova tarefa no banco de dados e retornar o ID da tarefa criada
	// Exemplo simplificado:
	result, err := d.db.Exec("INSERT INTO Tasks (title, description, status, priority) VALUES (?, ?, ?, ?)", title, description, status, priority)
	if err != nil {
		d.log.Error(err.Error())
		return 0, err
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		d.log.Error("DB 2")
		return 0, err
	}

	// d.log.Info("assigned: " + strconv.Itoa(assignedUsers[0]))
	// Atribuir a tarefa aos usuários associados
	for _, userID := range assignedUsers {
		d.log.Info("assigned: " + strconv.Itoa(userID))
		err := d.AssignTaskToUser(int(taskID), userID)
		if err != nil {
			d.log.Error(err.Error())
			// Se ocorrer um erro ao atribuir a tarefa a um usuário, podemos optar por reverter a criação da tarefa
			// Neste exemplo, vamos apenas retornar um erro, indicando que a tarefa foi criada, mas não foi atribuída a todos os usuários associados com sucesso
			return int(taskID), err
		}
	}

	return int(taskID), nil
}

// AssignTaskToUser atribui uma tarefa a um usuário no banco de dados.
func (d *Database) AssignTaskToUser(taskID int, userID int) error {
	// Implementação para associar uma tarefa a um usuário no banco de dados

	user, err := d.GetUserById(userID)
	if err != nil {
		return err
	}

	emptyUsr := service.User{}

	if user != emptyUsr {
		d.log.Info("usr: " + user.Name)
		_, err := d.db.Exec("INSERT INTO Task_user_associations (task_id, user_id) VALUES (?, ?)", taskID, userID)
		if !strings.Contains(err.Error(), "Duplicate") {
			if err != nil {
				return err
			}
		} else {
			return errors.New("membro ja está associado")
		}
	}

	return nil
}

// GetTaskByID retorna os detalhes de uma tarefa com base no ID da tarefa fornecido.
func (d *Database) GetTaskByID(taskID int) (service.Task, error) {
	// Implementação para buscar os detalhes de uma tarefa no banco de dados com base no ID da tarefa
	// Exemplo simplificado:
	var task service.Task
	err := d.db.QueryRow("SELECT id, title, description, priority, status FROM Tasks WHERE id = ?", taskID).Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status)
	if err != nil {
		return service.Task{}, err
	}

	return task, nil
}

// GetUserByID busca um usuário no banco de dados pelo seu ID.
func (d *Database) GetUserByID(id int) (service.User, error) {
	query := "SELECT name, email, password FROM User WHERE id = ?"
	row := d.db.QueryRow(query, id)

	var user service.User
	err := row.Scan(&user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return service.User{}, errors.New("usuário inexistente") // Usuário não encontrado
		}
		return service.User{}, err
	}

	user.ID = id
	return user, nil
}

// GetTasksForUser retorna todas as tarefas atribuídas ao usuário especificado.
func (d *Database) GetTasksForUser(userID int) ([]service.Task, error) {
	query := "SELECT id, title, description FROM Tasks WHERE id IN (SELECT task_id FROM Task_user_associations WHERE user_id = ?)"
	rows, err := d.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []service.Task
	for rows.Next() {
		var task service.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetAllTasks retorna todas as tarefas armazenadas no banco de dados.
func (d *Database) GetAllTasks() ([]service.Task, error) {
	query := "SELECT id, title, description, status, priority FROM Tasks"
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []service.Task
	for rows.Next() {
		var task service.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Priority)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// DeleteTaskByID exclui uma tarefa do banco de dados com o ID especificado.
func (d *Database) DeleteTask(taskID int) error {
	// Preparar a declaração SQL para excluir a tarefa
	query := "DELETE FROM Task_user_associations WHERE task_id = ?"
	rows, err := d.db.Query(query, taskID)
	if err != nil {
		d.log.Info(err.Error())
		return err
	}
	defer rows.Close()

	// Preparar a declaração SQL para excluir a tarefa
	query = "DELETE FROM Tasks WHERE id = ?"
	rows, err = d.db.Query(query, taskID)
	if err != nil {
		d.log.Info(err.Error())
		return err
	}
	defer rows.Close()

	return nil
}

// UpdateTask atualiza uma tarefa existente no banco de dados.
func (d *Database) UpdateTask(taskID int, updatedTask service.Task) error {
	// Preparar a declaração SQL para atualizar a tarefa
	query := "UPDATE Tasks SET title = ?, description = ?, status = ?, priority = ? WHERE id = ?"
	// Executar a declaração SQL para atualizar a tarefa
	rows, err := d.db.Query(query, updatedTask.Title, updatedTask.Description, updatedTask.Status, updatedTask.Priority, taskID)
	if err != nil {
		d.log.Info(err.Error())
		return err
	}
	defer rows.Close()

	return nil
}

// RemoveUser remove um usuário existente do banco de dados.
func (d *Database) RemoveUser(userID int) error {
	// Preparar a declaração SQL para deletar o usuário
	query := "DELETE FROM User WHERE id = ?"
	// Executar a declaração SQL para deletar o usuário
	rows, err := d.db.Query(query, userID)
	if err != nil {
		d.log.Error(err.Error())
		return err
	}
	defer rows.Close()

	return nil
}

func NewRepository(db *sql.DB, logger *zap.Logger) service.Repository {
	return &Database{
		db:  db,
		log: logger,
	}
}
