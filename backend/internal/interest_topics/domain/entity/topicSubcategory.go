package entity

type Topic_Subcategory struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Category_id int64  `json:"category_id"`
}
