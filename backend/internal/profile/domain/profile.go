package domain

type Profile struct {
	Id_profile   string
	Name         string
	Lastname     string
	Username     string
	ProfileImage string
	Biography    string
}

func NewProfile(id_profile, name, lastname, username, profileImage, biography string) *Profile {
	return &Profile{
		Id_profile:   id_profile,
		Name:         name,
		Lastname:     lastname,
		Username:     username,
		ProfileImage: profileImage,
		Biography:    biography,
	}
}
