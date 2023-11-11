package domain

import (
	"time"
)


type Profile struct {
        
	UserID         int64     
    BirthDate      time.Time 
    Name           string    
    LastName       string    
    AboutMe        string    
    Genre          string    
    Address        string    
    Country        string    
    City           string    
    InsertionDate  time.Time 
    UpdateDate     time.Time 
    ProfilePicture string    
}

func NewProfile( userID int64, birthDate time.Time, name string, lastName string, aboutMe string, genre string, address string, country string, city string, profilePicture string) *Profile {
	return &Profile{
		
		UserID:         userID,
		BirthDate:      birthDate,
		Name:           name,
		LastName:       lastName,
		AboutMe:        aboutMe,
		Genre:          genre,
		Address:        address,
		Country:        country,
		City:           city,
		ProfilePicture: profilePicture,
	}
}

func NewProfileToImport(id_profile , name, lastname, username, profileImage, biography string) *Profile {
	return &Profile{
		Name:           name,
		LastName:       lastname,
		ProfilePicture: profileImage,
		AboutMe:        biography,
	}
}
