package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sebasromero/api/authorization"
	"github.com/sebasromero/api/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Method not allowed.", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp := newResponse(Error, "Invalid struct", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}
	if !isLoginValid(&data) {
		resp := newResponse(Error, "User or password invalid", nil)
		responseJSON(w, http.StatusBadRequest, resp)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		resp := newResponse(Error, "Could not generate token", nil)
		responseJSON(w, http.StatusInternalServerError, resp)
		return
	}

	dataToken := map[string]string{"token": token}
	resp := newResponse(Message, "Ok", dataToken)
	responseJSON(w, http.StatusOK, resp)
}

func isLoginValid(data *model.Login) bool {
	return data.Email == "sebas@hotmail.com" && data.Password == "123456"
}
