package searchpost

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Function to manage the HTTP request
func SearchPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Disallowed method", http.StatusMethodNotAllowed)
		return
	}

	keyword, err := extractKeywordFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db, err := openDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer closeDB(db)

	result, err := performSearch(db, keyword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, result)
}

// Function to extract the keyword from the request
func extractKeywordFromRequest(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	var requestData struct {
		Keyword string `json:"keyword"`
	}

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		return "", err
	}

	return requestData.Keyword, nil
}

// Function to search the database
func performSearch(db *sql.DB, keyword string) ([]Post, error) {

	query := `SELECT id_post, id_user, content,"postHour","postDate", location, labels FROM Posts WHERE content ILIKE '%' || $1 || '%' OR labels ILIKE '%' || $1 || '%'`
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

// Function to write the HTTP response
func writeResponse(w http.ResponseWriter, result []Post) {
	w.Header().Set("Content-Type", "application/json")

	if len(result) == 0 {
		responseMessage := "No available publications found."
		responseJSON := []byte(`[{"message": "` + responseMessage + `"}]`)
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
		return
	}

	err := json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
