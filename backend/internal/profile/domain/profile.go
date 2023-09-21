package domain

type Profile struct {
	Id_profile   string
	Username     string
	ProfileImage string
	Biography    string
}

func NewProfile(id_profile string, username string, profileImage string, biography string) *Profile {
    return &Profile{
        Id_profile:   id_profile,
        Username:     username,
        ProfileImage: profileImage,
        Biography:    biography,
    }
}
