package domain

type Interest_Topic_Users struct {
	Id      int64 `db:"ID" pk:"true"`
	Id_user int64 `db:"id_user" fk:"true"`
	Subcategory_id int64 `db:"subcategory_id" fk:"true"`
}
