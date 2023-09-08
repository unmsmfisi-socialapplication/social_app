package domain

type Topic_Category struct {

	Id int64 `db:"ID" pk:"true"`
	Name string `db:"name"`
}