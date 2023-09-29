package searchpost

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func handleError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

// Function to manage the HTTP request
func SearchPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		handleError(w, errors.New("Disallowed method"), http.StatusMethodNotAllowed)
		return
	}

	requestData, err := extractDataFromRequest(r)
	if err != nil {
		handleError(w, err, http.StatusBadRequest)
		return
	}

	db, err := openDB()
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}
	defer closeDB(db)

	result, err := performSearch(db, requestData)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	if len(result) == 0 {
		writeResponse(w, []Post{})
		return
	}

	writeResponse(w, result)
}

// Function to extract the keyword from the request
func extractDataFromRequest(r *http.Request) (RequestData, error) {
	var requestData RequestData

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return requestData, err
	}

	err = json.Unmarshal(body, &requestData)
	if err != nil {
		return requestData, err
	}

	return requestData, nil
}

// Function to search the database
func performSearch(db *sql.DB, requestData RequestData) ([]Post, error) {
	query := `SELECT id_post, id_user, content, "postHour", "postDate", location, labels FROM Posts WHERE `

	args := make([]interface{}, len(requestData.Labels)+1)

	if len(requestData.Labels) > 0 {
		query += "("
		for i, label := range requestData.Labels {
			if i > 0 {
				query += " OR "
			}
			query += "labels ILIKE '%' || $" + strconv.Itoa(i+1) + " || '%'"
			args[i] = label
		}
		query += ") AND "
	}

	query += "content ILIKE '%' || $" + strconv.Itoa(len(requestData.Labels)+1) + " || '%'"
	args[len(args)-1] = requestData.Keyword

	rows, err := db.Query(query, args...)
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
		handleError(w, err, http.StatusInternalServerError)
		return
	}
}
