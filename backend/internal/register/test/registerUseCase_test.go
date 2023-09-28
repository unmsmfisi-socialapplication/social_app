package test

import (
	"errors"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/application"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
)

// Mock UserRepository para pruebas
type MockUserRepository struct {
	UserByEmail *domain.User
	Err        error
}

func (m *MockUserRepository) GetUserByEmail(email string) (*domain.User, error) {
	return m.UserByEmail, m.Err
}

func TestRegistrationUseCase_Register_Success(t *testing.T) {
	// Crear un UserRepository de prueba
	mockRepo := &MockUserRepository{}

	// Crear una instancia de RegistrationUseCase con el UserRepository de prueba
	useCase := application.NewRegistrationUseCase(mockRepo)

	// Definir los datos de entrada
	phone := "123456789"
	email := "correo@example.com"
	username := "usuario"
	password := "contraseña"

	// Simular un usuario que no existe en el UserRepository de prueba
	mockRepo.UserByEmail = nil
	mockRepo.Err = nil

	// Llamar a la función Register
	newUser, err := useCase.Register(phone, email, username, password)

	// Afirmar que no haya errores
	if err != nil {
		t.Fatal(err)
	}

	// Afirmar que el nuevo usuario no sea nulo
	if newUser == nil {
		
		t.Fatal("el usuario no debe ser nulo")
	}
}

func TestRegistrationUseCase_Register_UserAlreadyExists(t *testing.T) {
	// Crear un UserRepository de prueba
	mockRepo := &MockUserRepository{}

	// Crear una instancia de RegistrationUseCase con el UserRepository de prueba
	useCase :=application.NewRegistrationUseCase(mockRepo)

	// Definir los datos de entrada
	phone := "123456789"
	email := "correo@example.com"
	username := "usuario"
	password := "contraseña"

	// Simular un usuario que ya existe en el UserRepository de prueba
	mockRepo.UserByEmail = &domain.User{}
	mockRepo.Err = nil

	// Llamar a la función Register
	newUser, err := useCase.Register(phone, email, username, password)

	// Afirmar que se devuelve un error
	
	if err == nil {
		t.Fatal("se debe devolver un error")
	}
	
	if newUser != nil {
		t.Fatal("el usuario debe ser nulo")
	}
	if err.Error() != "el nombre de usuario ya está en uso" {
		t.Fatalf("se debe devolver el error %s, se devolvió %s", "el nombre de usuario ya está en uso", err.Error())
	}
	
}

func TestRegistrationUseCase_Register_UserRepositoryError(t *testing.T) {
	// Crear un UserRepository de prueba
	mockRepo := &MockUserRepository{}

	// Crear una instancia de RegistrationUseCase con el UserRepository de prueba
	useCase := application.NewRegistrationUseCase(mockRepo)

	// Definir los datos de entrada
	phone := "123456789"
	email := "correo@example.com"
	username := "usuario"
	password := "contraseña"

	// Simular un error en el UserRepository de prueba
	mockRepo.UserByEmail = nil
	mockRepo.Err = errors.New("error en el repositorio")

	// Llamar a la función Register
	newUser, err := useCase.Register(phone, email, username, password)

	// Afirmar que se devuelve un error
	
	if err == nil {
		t.Fatal("se debe devolver un error")
	}

	
	if newUser != nil {
		t.Fatal("el usuario debe ser nulo")
	}

	if err.Error() != "error en el repositorio" {
		t.Fatalf("se debe devolver el error %s, se devolvió %s", "error en el repositorio", err.Error())
	}

}
