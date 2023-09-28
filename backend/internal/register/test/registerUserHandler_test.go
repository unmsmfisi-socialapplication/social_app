package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unmsmfisi-socialapplication/social_app/internal/register/domain"
	"github.com/unmsmfisi-socialapplication/social_app/internal/register/infrastructure"
)

// Mock UserDBRepository para pruebas
type MockUserDBRepository struct {
	RegisterUserFn func(phone, email, username, password string) error
}

func (m *MockUserDBRepository) RegisterUser(phone, email, username, password string) error {
	return m.RegisterUserFn(phone, email, username, password)
}
func (m *MockUserDBRepository) GetUserByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func TestRegisterUserHandler_RegisterUser_Success(t *testing.T) {
	// Crear un MockUserDBRepository de prueba
	mockRepo := &MockUserDBRepository{
		RegisterUserFn: func(phone, email, username, password string) error {
			return nil // Simular un registro exitoso
		},
	}

	// Crear una instancia de RegisterUserHandler con el MockUserDBRepository de prueba
	handler := infrastructure.NewRegisterUserHandler(mockRepo)

	// Crear una solicitud HTTP de prueba con datos JSON
	requestBody := `{"phone": "123456789", "email": "correo@example.com", "user_name": "usuario", "password": "contraseña"}`
	req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Llamar a la función RegisterUser del manejador
	handler.RegisterUser(w, req)

	// Afirmar que se recibió una respuesta HTTP exitosa (200 OK)
	
	if w.Code != http.StatusOK {
		t.Fatal("el codigo de estado no es 200")
	}

	// Leer la respuesta JSON y verificar que los datos sean correctos
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	
	if err != nil {
		t.Fatal(err)
	}
	if response["phone"] != "123456789" {
		t.Errorf("expected %s, got %s", "123456789", response["phone"])
	}
	if response["email"] != "correo@example.com" {
		t.Errorf("expected %s, got %s", "correo@example.com", response["email"])
	}
	if response["user_name"] != "usuario" {
		t.Errorf("expected %s, got %s", "usuario", response["user_name"])
	}
	if response["password"] != "contraseña" {
		t.Errorf("expected %s, got %s", "contraseña", response["password"])
	}
	
}

func TestRegisterUserHandler_RegisterUser_RepositoryError(t *testing.T) {
	// Crear un MockUserDBRepository de prueba
	mockRepo := &MockUserDBRepository{
		RegisterUserFn: func(phone, email, username, password string) error {
			return errors.New("error en el repositorio") // Simular un error en el registro
		},
	}

	// Crear una instancia de RegisterUserHandler con el MockUserDBRepository de prueba
	handler := NewRegisterUserHandler(mockRepo)

	// Crear una solicitud HTTP de prueba con datos JSON
	requestBody := `{"phone": "123456789", "email": "correo@example.com", "user_name": "usuario", "password": "contraseña"}`
	req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Llamar a la función RegisterUser del manejador
	handler.RegisterUser(w, req)

	// Afirmar que se recibió una respuesta HTTP de error (400 Bad Request)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Leer el mensaje de error de la respuesta y verificar que sea correcto
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "error en el repositorio", response["error"])
}

func TestRegisterUserHandler_RegisterUser_InvalidJSON(t *testing.T) {
	// Crear una instancia de RegisterUserHandler sin un MockUserDBRepository (no se usará en esta prueba)

	handler := NewRegisterUserHandler(nil)

	// Crear una solicitud HTTP de prueba con un cuerpo JSON no válido
	requestBody := `{"phone": "123456789", "email": "correo@example.com", "user_name": "usuario", "password": "contraseña"`
	req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Llamar a la función RegisterUser del manejador
	handler.RegisterUser(w, req)

	// Afirmar que se recibió una respuesta HTTP de error (400 Bad Request) debido a JSON no válido
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
