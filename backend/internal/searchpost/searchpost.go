// En searchpost.go dentro de la carpeta searchpost
package searchpost

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func SearchPost(w http.ResponseWriter, r *http.Request) {

	connStr := "host=localhost user=postgres password=admin dbname=app_social sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	keyword := r.URL.Query().Get("keyword")

	query := `SELECT * FROM Posts WHERE content ILIKE '%' || $1 || '%' OR labels ILIKE '%' || $1 || '%'`
	rows, err := db.Query(query, keyword)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	resultados := []Post{}

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.id_post, &post.id_user, &post.content, &post.postHour, &post.postDate, &post.location, &post.labels)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resultados = append(resultados, post)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resultados)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
