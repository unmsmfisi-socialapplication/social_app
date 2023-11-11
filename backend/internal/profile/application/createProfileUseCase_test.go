package application

import (
	"testing"
	"time"

	"github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
)

// MockProfileRepository2 es una implementación de IProfileRepository para pruebas.
type MockProfileRepository2 struct {
	createdProfile2 *domain.Profile
	updatedProfile2 *domain.Profile
	err            error
}

func (mpr *MockProfileRepository2) CreateProfile(profile *domain.Profile) error {
	mpr.createdProfile2 = profile
	return mpr.err
}

func (mpr *MockProfileRepository2) UpdateProfile(profile *domain.Profile) error {
	mpr.updatedProfile2 = profile
	return mpr.err
}

func TestCreateProfileUseCase_CreateProfile(t *testing.T) {
	// Configuración
	mockRepo2 := &MockProfileRepository2{}
	useCase := NewCreateProfileUseCase(mockRepo2)

	// Datos de prueba
	testProfile := &domain.Profile{
		UserID:         1,
		BirthDate:      time.Now(),
		Name:           "John",
		LastName:       "Doe",
		AboutMe:        "Test about me",
		Genre:          "Male",
		Address:        "123 Main St",
		Country:        "Test Country",
		City:           "Test City",
		ProfilePicture: "test.jpg",
	}

	// Ejecutar el caso de uso
	err := useCase.CreateProfile(testProfile)

	// Verificar resultados
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if testProfile.Name != mockRepo2.createdProfile2.Name {
		t.Errorf("Error: los nombres no coinciden. Esperado: %s, Obtenido: %s", testProfile.Name, mockRepo2.createdProfile2.Name)
	}
	//valida todos los campos
	if testProfile.LastName != mockRepo2.createdProfile2.LastName {
		t.Errorf("Error: los apellidos no coinciden. Esperado: %s, Obtenido: %s", testProfile.LastName, mockRepo2.createdProfile2.LastName)
	}
	if testProfile.AboutMe != mockRepo2.createdProfile2.AboutMe {
		t.Errorf("Error: los about me no coinciden. Esperado: %s, Obtenido: %s", testProfile.AboutMe, mockRepo2.createdProfile2.AboutMe)
	}
	if testProfile.Genre != mockRepo2.createdProfile2.Genre {
		t.Errorf("Error: los generos no coinciden. Esperado: %s, Obtenido: %s", testProfile.Genre, mockRepo2.createdProfile2.Genre)
	}
	if testProfile.Address != mockRepo2.createdProfile2.Address {
		t.Errorf("Error: las direcciones no coinciden. Esperado: %s, Obtenido: %s", testProfile.Address, mockRepo2.createdProfile2.Address)
	}
	if testProfile.Country != mockRepo2.createdProfile2.Country {
		t.Errorf("Error: los paises no coinciden. Esperado: %s, Obtenido: %s", testProfile.Country, mockRepo2.createdProfile2.Country)
	}
	if testProfile.City != mockRepo2.createdProfile2.City {
		t.Errorf("Error: las ciudades no coinciden. Esperado: %s, Obtenido: %s", testProfile.City, mockRepo2.createdProfile2.City)
	}
	if testProfile.ProfilePicture != mockRepo2.createdProfile2.ProfilePicture {
		t.Errorf("Error: las fotos de perfil no coinciden. Esperado: %s, Obtenido: %s", testProfile.ProfilePicture, mockRepo2.createdProfile2.ProfilePicture)
	}

	
}
