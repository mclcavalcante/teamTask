package service

import (
	"errors"
)

// CreateTask cria uma nova tarefa com base nos dados de entrada fornecidos.
// Retorna o ID da tarefa criada e um erro, se houver.
func (service teamTaskService) CreateTask(input Task) (int, error) {
	// Validar entrada
	if input.Title == "" || input.Description == "" {
		return 0, errors.New("título e descrição são obrigatórios")
	}

	// Validar prioridade (por exemplo, garantir que seja um valor válido como "Alta", "Média" ou "Baixa")
	if input.Priority != "Alta" && input.Priority != "Média" && input.Priority != "Baixa" {
		return 0, errors.New("prioridade inválida")
	}

	// Criar a tarefa no banco de dados
	taskID, err := service.db.CreateTask(input.Title, input.Description, input.Status, input.Priority, input.TeamMembers)
	if err != nil {
		return 0, err
	}

	// Se tudo correu bem, retornamos o ID da tarefa criada
	return taskID, nil
}
