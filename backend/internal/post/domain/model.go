package domain

import (
	"time"
	"errors"
	"regexp"
)

type PostBase struct {
	UserId        int64
	Title         string
	Description   string
	HasMultimedia bool
	Public        bool
	Multimedia    string
}

type Post struct {
	Id            int64
	InsertionDate time.Time
	UpdateDate    time.Time
	PostBase
}

type PostCreate struct {
	PostBase
}

type PostPaginationParams struct {
	Page  int
	Limit int
}

type PostPagination struct {
	Posts       []Post
	TotalCount  int
	CurrentPage int
}

func PostCreateToPost(p PostCreate) Post {
	post := Post{
		PostBase:      p.PostBase,
		InsertionDate: time.Now(),
		UpdateDate:    time.Now(),
	}
	return post
}

func (c *PostCreate) Validate() error {

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

type Location struct {
    Latitude float64
    Longitude float64
}

