// En models.go dentro de la carpeta searchpost
package searchpost

import (
	"time"
)

// Estructura para representar una publicación (Post)
type Post struct {
	id_post  int64     `json:"id_post"`
	id_user  int64     `json:"id_user"`
	content  string    `json:"content"`
	postHour time.Time `json:"postHour"`
	postDate time.Time `json:"postDate"`
	location string    `json:"location"`
	labels   string    `json:"labels"`
}
