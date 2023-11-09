package domain

import (
	"time"
	"errors"
	"regexp"
)

type Post struct {
	Id            int64
	UserId        int64
	Title         string
	Description   string
	HasMultimedia bool
	Public        bool
	Multimedia    string
	InsertionDate time.Time
	UpdateDate    time.Time
}

type CreatePost struct {
	UserId        int64  `json:"userId"`
	Title         string `json:"title" db:"title" validate:"max=100"`
	Description   string `json:"description" db:"description" validate:"max=1000"`
	HasMultimedia bool   `json:"hasMultimedia"`
	Public        bool   `json:"public"`
	Multimedia    string `json:"multimedia" db:"multimedia" validate:"max=1000"`
}

func (c *CreatePost) Validate() error {

	titleRegex := regexp.MustCompile(`^.{0,100}$`)
	descriptionRegex := regexp.MustCompile(`^.{0,1000}$`)
	
    if !titleRegex.MatchString(c.Title) {
        return errors.New("el título debe tener un máximo de 100 caracteres")
    }
	if !descriptionRegex.MatchString(c.Title) {
        return errors.New("la descripcion debe tener un máximo de 1000 caracteres")
    }
    return nil
}
