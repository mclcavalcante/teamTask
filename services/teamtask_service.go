package service

import (
	"errors"
)

// RegisterNewUser registra um novo usuário no sistema.
func (service teamTaskService) RegisterNewUser(user User) (int, error) {
	// Verificar se o e-mail do usuário já existe no banco de dados
	emptyUsr := User{}
	existingUser, err := service.db.GetUserByEmail(user.Email)
	if err == nil && existingUser != emptyUsr {
		service.log.Info("E-mail já está em uso")

		return 0, errors.New("e-mail já está em uso")
	}

	// Validar os dados do usuário
	if user.Name == "" || user.Email == "" || user.Password == "" {
		service.log.Info("Dados do usuário incompletos")

		return 0, errors.New("dados do usuário incompletos")
	}

	// Adicionar o novo usuário ao banco de dados
	user_id, err := service.db.AddUser(user)
	if err != nil {
		service.log.Info("Erro ao registrar novo usuário")
		return 0, errors.New("erro ao registrar novo usuário")
	}

	return user_id, nil
}

// CreateTask cria uma nova tarefa com base nos dados de entrada fornecidos.
// Retorna o ID da tarefa criada e um erro, se houver.
func (service teamTaskService) CreateTask(input Task) (int, error) {
	// Validar entrada
	if input.Title == "" || input.Description == "" {
		service.log.Error("título e descrição são obrigatórios")
		return 0, errors.New("título e descrição são obrigatórios")
	}

	// Validar prioridade (por exemplo, garantir que seja um valor válido como "High", "Medium" ou "Low")
	if (input.Priority != "Alta" && input.Priority != "Média" && input.Priority != "Baixa") && input.Priority != "" {
		service.log.Error("prioridade inválida")
		return 0, errors.New("prioridade inválida")
	}

	// Criar a tarefa no banco de dados
	taskID, err := service.db.CreateTask(input.Title, input.Description, input.Status, input.Priority, input.AssignedUsers)
	if err != nil {
		service.log.Error("Error salvado a task")
		return 0, err
	}

	// Se tudo correu bem, retornamos o ID da tarefa criada
	return taskID, nil
}

func (service teamTaskService) GetVisibleTasksForUser(userID int) ([]Task, error) {
	// Verificar se o usuário existe
	_, err := service.db.GetUserByID(userID)
	if err != nil {
		return nil, errors.Join(err, errors.New("usuário não encontrado"))
	}

	// Obter todas as tarefas visíveis para o usuário
	tasks, err := service.db.GetTasksForUser(userID)
	if err != nil {
		return nil, errors.Join(err, errors.New("erro ao obter tarefas para o usuário"))
	}

	return tasks, nil
}

// FilterTasksByStatusAndPriority retorna todas as tarefas  com o status e a prioridade especificados.
func (service teamTaskService) FilterTasksByStatusAndPriority(status, priority string) ([]Task, error) {

	// Filtrar tarefas com base no status e na prioridade
	tasks, err := service.db.GetAllTasks()
	if err != nil {
		return nil, errors.Join(err, errors.New("erro ao obter tarefas"))
	}

	var filteredTasks []Task
	for _, task := range tasks {
		if (status == "" || task.Status == status) && (priority == "" || task.Priority == priority) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks, nil
}

// AssignMemberToTask associa um membro da equipe a uma tarefa específica.
func (service teamTaskService) AssignMemberToTask(taskID, memberID int) error {

	// Verificar se a tarefa existe
	_, err := service.db.GetTaskByID(taskID)
	if err != nil {
		return errors.Join(err, errors.New("tarefa não encontrada"))
	}

	// Verificar se o membro da equipe existe
	_, err = service.db.GetUserByID(memberID)
	if err != nil {
		return errors.Join(err, errors.New("membro da equipe não encontrado"))
	}

	// Associar o membro da equipe à tarefa
	err = service.db.AssignTaskToUser(taskID, memberID)
	if err != nil {
		return errors.Join(err, errors.New("erro ao associar membro da equipe à tarefa"))
	}

	return nil
}

// DeleteTask exclui uma tarefa específica do banco de dados.
func (service teamTaskService) DeleteTask(taskID int) error {
	// Verificar se a tarefa existe
	_, err := service.db.GetTaskByID(taskID)
	if err != nil {
		return errors.Join(errors.New("tarefa não encontrada"))
	}

	// Excluir a tarefa do banco de dados
	err = service.db.DeleteTask(taskID)
	if err != nil {
		return errors.Join(errors.New("erro ao excluir a tarefa"))
	}

	return nil
}

func (service teamTaskService) GetTaskByID(taskID int) (Task, error) {
	// Verificar se a tarefa existe
	task, err := service.db.GetTaskByID(taskID)
	if err != nil {
		return Task{}, errors.Join(err, errors.New("tarefa não encontrada"))
	}

	return task, nil
}

// EditTask edita uma tarefa existente no banco de dados.
func (service teamTaskService) EditTask(taskID int, updatedTask Task) error {
	// Verificar se a tarefa existe
	_, err := service.GetTaskByID(taskID)
	if err != nil {
		return errors.Join(err, errors.New("tarefa não encontrada"))
	}

	// Executar a edição da tarefa no banco de dados
	err = service.db.UpdateTask(taskID, updatedTask)
	if err != nil {
		return errors.Join(err, errors.New("erro ao editar a tarefa"))
	}

	return nil
}

func (service teamTaskService) GetAllTasks() ([]Task, error) {
	tasks, err := service.db.GetAllTasks()
	if err != nil {
		return []Task{}, errors.Join(err, errors.New("erro ao recuperar as tarefas"))
	}

	return tasks, nil
}

// DeleteUser deleta um usuário existente do banco de dados.
func (service teamTaskService) DeleteUser(userID int) error {
	// Verificar se o usuário existe
	_, err := service.db.GetUserByID(userID)
	if err != nil {
		return errors.Join(err, errors.New("usuário não encontrado"))
	}

	// Deletar o usuário do banco de dados
	err = service.db.RemoveUser(userID)
	if err != nil {
		return errors.Join(err, errors.New("erro ao deletar o usuário"))
	}

	return nil
}

// GetUserByID busca um usuário no banco de dados pelo ID.
func (service teamTaskService) GetUserByID(userID int) (User, error) {
	// Verificar se a tarefa existe
	task, err := service.db.GetUserByID(userID)
	if err != nil {
		return User{}, err
	}

	return task, nil
}
