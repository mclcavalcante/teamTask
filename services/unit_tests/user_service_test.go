package service_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	service "github.com/mclcavalcante/teamTask/services"
)

func TestRegisterNewUserWithValidData(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Dados de exemplo para um novo usuário
	newUser := service.User{
		Name:     "John Doe",
		Email:    "john@example.com",
		Password: "password123",
	}

	// Chamar a função de registro de novo usuário
	_, err := s.RegisterNewUser(newUser)
	if err != nil {
		t.Errorf("Erro ao registrar novo usuário com dados válidos: %v", err)
	}

	//TODO
	// // Verificar se o usuário foi adicionado ao banco de dados
	// _, err = s.GetUserByEmail(newUser.Email)
	// if err != nil {
	// 	t.Errorf("Erro ao buscar usuário recém-registrado no banco de dados: %v", err)
	// }
}

func TestRegisterNewUserWithExistingEmail(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Dados de exemplo para um novo usuário
	existingUser := service.User{
		Name:     "Jane Smith",
		Email:    "jane@example.com",
		Password: "password456",
	}

	// Adicionar usuário existente ao banco de dados
	_, err := s.RegisterNewUser(existingUser)
	if err != nil {
		t.Fatalf("Erro ao adicionar usuário existente ao banco de dados: %v", err)
	}

	// Tentar registrar um novo usuário com o mesmo e-mail
	newUser := service.User{
		Name:     "John Doe",
		Email:    "jane@example.com",
		Password: "password789",
	}

	_, err = s.RegisterNewUser(newUser)
	if err == nil {
		t.Error("Esperava-se um erro ao tentar registrar um novo usuário com e-mail já existente")
	}
}

func TestRegisterNewUserWithInvalidData(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Tentar registrar um novo usuário com dados inválidos (sem e-mail)
	newUser := service.User{
		Name:     "John Doe",
		Password: "password123",
		Email:    "",
	}

	_, err := s.RegisterNewUser(newUser)
	if err == nil {
		t.Error("Esperava-se um erro ao tentar registrar um novo usuário com dados inválidos")
	}
}

// func TestDeleteUser(t *testing.T) {
// 	// Mock do banco de dados
// 	s := NewTestService()

// 	// Adicionar um usuário de exemplo ao banco de dados
// 	user := service.User{ID: 1, Name: "John Doe", Email: "john@example.com"}
// 	s.RegisterNewUser(user)

// 	// Deletar o usuário
// 	err := s.DeleteUser(1)
// 	if err != nil {
// 		t.Errorf("Erro inesperado ao deletar o usuário: %v", err)
// 	}

// 	// Verificar se o usuário foi deletado corretamente
// 	_, err = s.GetUserByID(1)
// 	if err == nil {
// 		t.Error("Esperava-se um erro ao tentar obter um usuário deletado")
// 	}
// }

func TestDeleteNonExistentUser(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Tentar deletar um usuário inexistente
	err := s.DeleteUser(999)
	if err == nil {
		t.Error("Esperava-se um erro ao tentar deletar um usuário inexistente")
	}
}

func TestGetUserByID(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Adicionar um usuário de exemplo ao banco de dados
	user := service.User{ID: 1, Name: "John Doe", Email: "john@example.com", Password: "123"}
	s.RegisterNewUser(user)

	// Buscar o usuário pelo ID
	returnedUser, err := s.GetUserByID(1)
	if err != nil {
		t.Errorf("Erro inesperado ao buscar o usuário pelo ID: %v", err)
	}

	// Verificar se o usuário retornado é o esperado
	if returnedUser.ID != user.ID || returnedUser.Name != user.Name || returnedUser.Email != user.Email {
		t.Error("O usuário retornado não corresponde ao usuário esperado")
	}
}

func TestGetNonExistentUserByID(t *testing.T) {
	// Mock do banco de dados
	s := NewTestService()

	// Tentar buscar um usuário inexistente pelo ID
	_, err := s.GetUserByID(-1)
	if err == nil {
		t.Error("Esperava-se um erro ao tentar buscar um usuário inexistente pelo ID")
	}
}
