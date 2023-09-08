package infraestructure

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type EditUserHandler struct {
	repo UserDBRepository
}

func NewEditUserHandler(r UserDBRepository) *EditUserHandler {
	return &EditUserHandler{repo: r}
}

func (rh *EditUserHandler) EditUser(w http.ResponseWriter, r *http.Request) {
	// Get and return json
	id_user, err1 := strconv.Atoi(r.URL.Query().Get("id_user"))
	print(id_user)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	var data struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		LastName string `json:"lastName"`
		Birthday string `json:"birthday"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//username, email, password, name, lastName, birthday string
	rh.repo.UpdateUser(id_user, data.Username, data.Email, data.Password, data.Name, data.LastName, data.Birthday)
	json.NewEncoder(w).Encode(data)
}
