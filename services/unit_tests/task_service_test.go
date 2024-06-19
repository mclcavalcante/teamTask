package service_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	service "github.com/mclcavalcante/teamTask/services"
	"github.com/mclcavalcante/teamTask/services/unit_tests/mock"
	"go.uber.org/zap"
)

func NewTestService() service.Service {
	repo := mock.NewTestRepository()
	logger := zap.NewNop()
	return service.NewService(repo, logger)
}

func TestCreateTask(t *testing.T) {
	// Simulando um ambiente de teste, por exemplo,
	// criando uma instância de banco de dados em memória
	s := NewTestService()

	// Criando um mock de entrada para o serviço de criação de tarefa
	input := service.Task{
		Title:         "Test Task",
		Description:   "This is a test task",
		Priority:      "High",
		AssignedUsers: []int{0, 1, 2},
	}

	// Chamando o serviço de criação de tarefa
	taskID, err := s.CreateTask(input)

	// Verificando se houve algum erro na criação da tarefa
	if err != nil {
		t.Errorf("Erro ao criar tarefa: %v", err)
	}

	// Verificando se a tarefa foi atribuída com sucesso e possui um ID válido
	if taskID == 0 {
		t.Error("ID de tarefa inválido")
	}

	//TODO

	// Verificando se a tarefa criada existe no banco de dados
	// task, err := s.GetTaskByID(taskID)
	// if err != nil {
	// 	t.Errorf("Erro ao buscar tarefa do banco de dados: %v", err)
	// }

	// // Comparando os atributos da tarefa criada com os valores esperados
	// if task.Title != input.Title {
	// 	t.Errorf("Título da tarefa não corresponde: esperado %s, obtido %s", input.Title, task.Title)
	// }
	// if task.Description != input.Description {
	// 	t.Errorf("Descrição da tarefa não corresponde: esperado %s, obtido %s", input.Description, task.Description)
	// }
	// Verifique outras propriedades da tarefa, como prioridade e membros da equipe, da mesma maneira
}

func TestCreateTaskWithoutTitleOrDescription(t *testing.T) {
	s := NewTestService()

	input := service.Task{
		Title:         "",
		Description:   "",
		Priority:      "High",
		AssignedUsers: []int{0, 1},
	}

	_, err := s.CreateTask(input)
	if err == nil {
		t.Error("Esperava-se um erro ao criar uma tarefa sem título ou descrição")
	}
}

func TestCreateTaskWithInvalidPriority(t *testing.T) {
	s := NewTestService()

	input := service.Task{
		Title:         "Nova Tarefa",
		Description:   "Descrição da tarefa",
		Priority:      "Invalida",
		AssignedUsers: []int{0, 1},
	}

	_, err := s.CreateTask(input)
	if err == nil {
		t.Error("Esperava-se um erro ao criar uma tarefa com prioridade inválida")
	}
}

func TestCreateTaskWithoutTeamMembers(t *testing.T) {
	s := NewTestService()

	input := service.Task{
		Title:         "Nova Tarefa",
		Description:   "Descrição da tarefa",
		Priority:      "High",
		AssignedUsers: nil,
	}

	_, err := s.CreateTask(input)
	if err != nil {
		t.Error("Erro ao criar uma tarefa sem membros da equipe:", err)
	}
}

func TestGetVisibleTasksForUser(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar usuários de exemplo ao banco de dados
	user1 := service.User{ID: 1, Name: "User 1", Email: "user1@example.com", Password: "123"}
	user2 := service.User{ID: 2, Name: "User 2", Email: "user2@example.com", Password: "456"}
	s.RegisterNewUser(user1)
	s.RegisterNewUser(user2)

	// Adicionar tarefas de exemplo ao banco de dados
	task1 := service.Task{ID: 1, Title: "Task 1", Description: "Description for Task 1", Priority: "High", AssignedUsers: []int{1}}
	task2 := service.Task{ID: 2, Title: "Task 2", Description: "Description for Task 2", Priority: "High", AssignedUsers: []int{1}}
	s.CreateTask(task1)
	s.CreateTask(task2)

	// Chamar o serviço de Visualização de Tarefas para o usuário 1
	tasks, err := s.GetVisibleTasksForUser(user1.ID)
	if err != nil {
		t.Errorf("Erro inesperado ao obter tarefas para o usuário 1: %v", err)
	}

	// Verificar se as tarefas atribuídas ao usuário 1 foram retornadas
	if len(tasks) != 2 || tasks[0].ID != task1.ID {
		t.Errorf("Tarefas visíveis para o usuário 1 não correspondem às tarefas esperadas")
	}
}

func TestGetVisibleTasksForInvalidUser(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Tentar obter tarefas para um usuário inválido (ID não existente)
	_, err := s.GetVisibleTasksForUser(-1)
	if err == nil {
		t.Error("Esperava-se um erro ao obter tarefas para um usuário inválido")
	}
}

func TestFilterTasksByStatusAndPriority(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar usuários de exemplo ao banco de dados
	user := service.User{ID: 1, Name: "User 1", Email: "user1@example.com", Password: "1234"}
	s.RegisterNewUser(user)

	// Adicionar tarefas de exemplo ao banco de dados
	task1 := service.Task{ID: 1, Title: "Task 1", Description: "Description for Task 1", Status: "Open", Priority: "High", AssignedUsers: []int{1}}
	task2 := service.Task{ID: 2, Title: "Task 2", Description: "Description for Task 2", Status: "In Progress", Priority: "Medium", AssignedUsers: []int{1}}
	task3 := service.Task{ID: 3, Title: "Task 3", Description: "Description for Task 3", Status: "Closed", Priority: "Low", AssignedUsers: []int{1}}
	s.CreateTask(task1)
	s.CreateTask(task2)
	s.CreateTask(task3)

	// Chamar o serviço de Filtragem de Tarefas para o usuário 1 com status "Open" e prioridade "High"
	filteredTasks, err := s.FilterTasksByStatusAndPriority("Open", "High")
	if err != nil {
		t.Errorf("Erro inesperado ao filtrar tarefas: %v", err)
	}

	// Verificar se apenas a tarefa 1 foi retornada
	if len(filteredTasks) != 1 || filteredTasks[0].ID != task1.ID {
		t.Errorf("Tarefas filtradas não correspondem às tarefas esperadas")
	}
}

func TestFilterTasksForInvalidUser(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Tentar filtrar tarefas para um usuário inválido (ID não existente)
	_, err := s.FilterTasksByStatusAndPriority("", "")
	if err == nil {
		t.Error("Esperava-se um erro ao filtrar tarefas para um usuário inválido")
	}
}

func TestFilterTasksWithDatabaseError(t *testing.T) {
	// Mock do banco de dados que sempre retorna um erro
	s := NewTestService()

	// Tentar filtrar tarefas com um banco de dados que sempre retorna um erro
	_, err := s.FilterTasksByStatusAndPriority("", "")
	if err == nil {
		t.Error("Esperava-se um erro ao filtrar tarefas com um banco de dados com erro")
	}
}

func TestAssignMemberToTask(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar um membro da equipe de exemplo ao banco de dados
	teamMember := service.User{Name: "User", Email: "user2@example.com", Password: "123"}
	teamMember.ID, _ = s.RegisterNewUser(teamMember)

	// Adicionar uma tarefa de exemplo ao banco de dados
	task := service.Task{Title: "Task 1", Description: "Description for Task 1"}
	task.ID, _ = s.CreateTask(task)

	// Marcar o membro da equipe na tarefa
	err := s.AssignMemberToTask(task.ID, teamMember.ID)
	if err != nil {
		t.Errorf("Erro inesperado ao marcar membro da equipe na tarefa: %v", err)
	}

	// Verificar se o membro da equipe foi marcado corretamente na tarefa
	taskForUser, err := s.GetVisibleTasksForUser(teamMember.ID)
	if err != nil {
		t.Errorf("Erro ao obter tarefa com membros da equipe: %v", err)
	}

	found := false
	for _, t := range taskForUser {
		if t.ID == task.ID {
			found = true
		}
	}

	if !found {
		t.Error("Membro da equipe não foi marcado corretamente na tarefa")
	}
}

func TestAssignMemberToNonExistentTask(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Tentar marcar um membro da equipe em uma tarefa inexistente
	err := s.AssignMemberToTask(999, 1)
	if err == nil {
		t.Error("Esperava-se um erro ao marcar membro da equipe em uma tarefa inexistente")
	}
}

// TestAssignNonExistentMemberToTask verifica se há erro ao tentar marcar um membro da equipe que não existe.
func TestAssignNonExistentMemberToTask(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Tentar marcar um membro da equipe que não existe em uma tarefa
	err := s.AssignMemberToTask(1, 999)
	if err == nil {
		t.Error("Esperava-se um erro ao marcar membro da equipe que não existe")
	}
}

// TestAssignMemberToTaskWithNoPermission verifica se há erro ao tentar marcar um membro da equipe em uma tarefa onde o usuário atual não tem permissão.
func TestAssignMemberToTaskWithNoPermission(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar um usuário de exemplo ao banco de dados
	user := service.User{ID: 1, Name: "User 1", Email: "user1@example.com"}
	s.RegisterNewUser(user)

	// Adicionar um membro da equipe de exemplo ao banco de dados
	teamMember := service.User{ID: 2, Name: "User 2", Email: "user2@example.com"}
	s.RegisterNewUser(teamMember)

	// Adicionar uma tarefa de exemplo ao banco de dados
	task := service.Task{ID: 1, Title: "Task 1", Description: "Description for Task 1"}
	s.CreateTask(task)

	// Tentar marcar o membro da equipe na tarefa sem permissão
	err := s.AssignMemberToTask(task.ID, teamMember.ID)
	if err == nil {
		t.Error("Esperava-se um erro ao marcar membro da equipe em uma tarefa sem permissão")
	}
}

func TestDeleteTask(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar uma tarefa de exemplo ao banco de dados
	task := service.Task{ID: 1, Title: "Task 1", Description: "Description for Task 1"}
	s.CreateTask(task)

	// Excluir a tarefa
	err := s.DeleteTask(1)
	if err != nil {
		t.Errorf("Erro inesperado ao excluir a tarefa: %v", err)
	}

	// Verificar se a tarefa foi excluída corretamente
	_, err = s.GetTaskByID(1)
	if err == nil {
		t.Error("A tarefa não foi excluída corretamente")
	}
}

func TestDeleteNonExistentTask(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Tentar excluir uma tarefa inexistente
	err := s.DeleteTask(999)
	if err == nil {
		t.Error("Esperava-se um erro ao tentar excluir uma tarefa inexistente")
	}
}

func TestEditTask(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar uma tarefa de exemplo ao banco de dados
	task := service.Task{ID: 1, Title: "Task 1", Description: "Description for Task 1"}
	s.CreateTask(task)

	// Nova descrição para a tarefa
	newDescription := "New description for Task 1"

	// Editar a tarefa
	err := s.EditTask(1, service.Task{Description: newDescription})
	if err != nil {
		t.Errorf("Erro inesperado ao editar a tarefa: %v", err)
	}

	// Verificar se a tarefa foi editada corretamente
	editedTask, err := s.GetTaskByID(1)
	if err != nil {
		t.Errorf("Erro ao obter a tarefa editada: %v", err)
	}
	if editedTask.Description != newDescription {
		t.Errorf("A descrição da tarefa não foi editada corretamente. Esperado: %s, Obtido: %s", newDescription, editedTask.Description)
	}
}

func TestEditNonExistentTask(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Tentar editar uma tarefa inexistente
	err := s.EditTask(999, service.Task{})
	if err == nil {
		t.Error("Esperava-se um erro ao tentar editar uma tarefa inexistente")
	}
}

func TestEditTaskTitle(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar uma tarefa de exemplo ao banco de dados
	task := service.Task{ID: 1, Title: "Task 1", Description: "Description for Task 1"}
	s.CreateTask(task)

	// Novo título para a tarefa
	newTitle := "New Task Title"

	// Editar o título da tarefa
	err := s.EditTask(1, service.Task{Title: newTitle})
	if err != nil {
		t.Errorf("Erro inesperado ao editar o título da tarefa: %v", err)
	}

	// Verificar se o título da tarefa foi editado corretamente
	editedTask, err := s.GetTaskByID(1)
	if err != nil {
		t.Errorf("Erro ao obter a tarefa editada: %v", err)
	}
	if editedTask.Title != newTitle {
		t.Errorf("O título da tarefa não foi editado corretamente. Esperado: %s, Obtido: %s", newTitle, editedTask.Title)
	}
}

func TestEditTaskPriority(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar uma tarefa de exemplo ao banco de dados
	task := service.Task{ID: 1, Title: "Task 1", Description: "Description for Task 1", Priority: "High"}
	s.CreateTask(task)

	// Nova prioridade para a tarefa
	newPriority := "Low"

	// Editar a prioridade da tarefa
	err := s.EditTask(1, service.Task{Priority: newPriority})
	if err != nil {
		t.Errorf("Erro inesperado ao editar a prioridade da tarefa: %v", err)
	}

	// Verificar se a prioridade da tarefa foi editada corretamente
	editedTask, err := s.GetTaskByID(1)
	if err != nil {
		t.Errorf("Erro ao obter a tarefa editada: %v", err)
	}
	if editedTask.Priority != newPriority {
		t.Errorf("A prioridade da tarefa não foi editada corretamente. Esperado: %s, Obtido: %s", newPriority, editedTask.Priority)
	}
}

func TestEditTaskStatus(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar uma tarefa de exemplo ao banco de dados
	task := service.Task{ID: 1, Title: "Task 1", Description: "Description for Task 1", Status: "Open"}
	s.CreateTask(task)

	// Novo status para a tarefa
	newStatus := "Closed"

	// Editar o status da tarefa
	err := s.EditTask(1, service.Task{Status: newStatus})
	if err != nil {
		t.Errorf("Erro inesperado ao editar o status da tarefa: %v", err)
	}

	// Verificar se o status da tarefa foi editado corretamente
	editedTask, err := s.GetTaskByID(1)
	if err != nil {
		t.Errorf("Erro ao obter a tarefa editada: %v", err)
	}
	if editedTask.Status != newStatus {
		t.Errorf("O status da tarefa não foi editado corretamente. Esperado: %s, Obtido: %s", newStatus, editedTask.Status)
	}
}
