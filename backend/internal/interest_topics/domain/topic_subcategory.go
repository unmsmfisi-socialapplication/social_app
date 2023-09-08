package domain

type Topic_Subcategory struct {

	Id int64 `db:"ID" pk:"true"`
	Name string `db:"name"`
	Category_id int64 `db:"category_id" fk:"true"`
}