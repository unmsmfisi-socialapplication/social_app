package searchpost

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Database configuration will then go in a separate file
const (
	dbHost     = "localhost"
	dbUser     = "postgres"
	dbPassword = "admin"
	dbName     = "app_social"
)

// connection that will then go in another file
func openDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Db Close
func closeDB(db *sql.DB) {
	if db != nil {
		db.Close()
	}
}

// Function to search publications in the database
func searchPosts(db *sql.DB, keyword string) ([]Post, error) {
	query := `SELECT * FROM Posts WHERE content ILIKE '%' || $1 || '%' OR labels ILIKE '%' || $1 || '%'`
	rows, err := db.Query(query, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []Post{}

	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.PostHour, &post.PostDate, &post.Location, &post.Labels)
		if err != nil {
			return nil, err
		}
		result = append(result, post)
	}

	return result, nil
}

// Function to handle HTTP request
func SearchPost(w http.ResponseWriter, r *http.Request) {
	db, err := openDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer closeDB(db)

	keyword := r.URL.Query().Get("keyword")

	result, err := searchPosts(db, keyword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
