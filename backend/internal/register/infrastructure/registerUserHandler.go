package infrastructure

import (
	"encoding/json"
	"net/http"

	
	
)
type RegisterUserHandler struct {
	repo UserDBRepository
}
func NewRegisterUserHandler(r UserDBRepository ) *RegisterUserHandler {
	return &RegisterUserHandler{repo: r}
}

func (rh *RegisterUserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

    var data struct{
		
		Phone string `json:"phone"`
		Email string `json:"email"`
		User_name string `json:"user_name"`
		Password string `json:"password"`
	}
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	_,er:= rh.repo.RegisterUser(data.Phone,data.Email,data.User_name,data.Password)
	if er!=nil{
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}
    json.NewEncoder(w).Encode(data)
}