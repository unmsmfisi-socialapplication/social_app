package application


import "github.com/unmsmfisi-socialapplication/social_app/internal/profile/domain"
//crea un caso de uso para la creacion de un perfil
type CreateProfileUseCase struct {
	profileRepository IProfileRepository
}

//crea un nuevo caso de uso para la creacion de un perfil
func NewCreateProfileUseCase(profileRepository IProfileRepository) *CreateProfileUseCase {
	return &CreateProfileUseCase{profileRepository}
}

//ejecuta el caso de uso para la creacion de un perfil
func (cpuc *CreateProfileUseCase) CreateProfile(p *domain.Profile) error {
	//crea un nuevo perfil
	profile := domain.NewProfile( p.UserID, p.BirthDate, p.Name, p.LastName, p.AboutMe, p.Genre, p.Address, p.Country, p.City, p.InsertionDate, p.UpdateDate, p.ProfilePicture)

	//guarda el perfil
	err := cpuc.profileRepository.CreateProfile(profile)
	if err != nil {
		return err
	}
	return nil
}

