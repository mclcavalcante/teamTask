package service

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// NewTestDatabase cria uma nova instância do banco de dados simulado para testes.
func NewTestDatabase() *MockDatabase {
	return &MockDatabase{
		tasks:       make(map[int]Task),
		taskCounter: 0,
	}
}

func NewTestService() *teamTaskService {
	return &teamTaskService{
		db: NewTestDatabase(),
	}
}

func TestCreateTask(t *testing.T) {
	// Simulando um ambiente de teste, por exemplo, criando uma instância de banco de dados em memória
	s := NewTestService()

	// Criando um mock de entrada para o serviço de criação de tarefa
	input := Task{
		Title:       "Test Task",
		Description: "This is a test task",
		Priority:    "Alta",
		TeamMembers: []int{1, 2, 3},
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

	// Verificando se a tarefa criada existe no banco de dados
	task, err := s.db.GetTaskByID(taskID)
	if err != nil {
		t.Errorf("Erro ao buscar tarefa do banco de dados: %v", err)
	}

	// Comparando os atributos da tarefa criada com os valores esperados
	if task.Title != input.Title {
		t.Errorf("Título da tarefa não corresponde: esperado %s, obtido %s", input.Title, task.Title)
	}
	if task.Description != input.Description {
		t.Errorf("Descrição da tarefa não corresponde: esperado %s, obtido %s", input.Description, task.Description)
	}
	// Verifique outras propriedades da tarefa, como prioridade e membros da equipe, da mesma maneira
}

func TestCreateTaskWithoutTitleOrDescription(t *testing.T) {
	s := NewTestService()

	input := Task{
		Title:       "",
		Description: "",
		Priority:    "Alta",
		TeamMembers: []int{1, 2},
	}

	_, err := s.CreateTask(input)
	if err == nil {
		t.Error("Esperava-se um erro ao criar uma tarefa sem título ou descrição")
	}
}

func TestCreateTaskWithInvalidPriority(t *testing.T) {
	s := NewTestService()

	input := Task{
		Title:       "Nova Tarefa",
		Description: "Descrição da tarefa",
		Priority:    "Invalida",
		TeamMembers: []int{1, 2},
	}

	_, err := s.CreateTask(input)
	if err == nil {
		t.Error("Esperava-se um erro ao criar uma tarefa com prioridade inválida")
	}
}

func TestCreateTaskWithoutTeamMembers(t *testing.T) {
	s := NewTestService()

	input := Task{
		Title:       "Nova Tarefa",
		Description: "Descrição da tarefa",
		Priority:    "Alta",
		TeamMembers: nil,
	}

	_, err := s.CreateTask(input)
	if err != nil {
		t.Error("Erro ao criar uma tarefa sem membros da equipe:", err)
	}
}
