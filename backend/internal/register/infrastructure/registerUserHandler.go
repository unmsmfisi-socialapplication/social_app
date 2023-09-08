package infrastructure


import (
	"encoding/json"
	"net/http"

	
)
type RegisterUserHandler struct {
	repo UserDBRepository
}
func NewRegisterUserHandler(r UserDBRepository) *RegisterUserHandler {
	return &RegisterUserHandler{repo: r}
}

func (rh *RegisterUserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

    var data struct{
		Username string `json:"username"`
		Email string `json:"email"`
		Password string `json:"password"`
		Name string `json:"name"`
		LastName string `json:"lastName"`
		Birthday string `json:"birthday"`
	}
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	rh.repo.RegisterUser(data.Username, data.Email, data.Password, data.Name, data.LastName, data.Birthday)
    json.NewEncoder(w).Encode(data)
}