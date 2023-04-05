package handler

import (
	"encoding/json"
	"net/http"
	"webapp/model"
)

func (h *Handler) signUp(wr http.ResponseWriter, r *http.Request) {
	var input model.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}

	err = json.NewEncoder(wr).Encode(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		http.Error(wr, err.Error(), 500)
	}
	return
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(wr http.ResponseWriter, r *http.Request) {
	var input signInInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}

	err = json.NewEncoder(wr).Encode(map[string]interface{}{
		"token": token,
	})

	if err != nil {
		http.Error(wr, err.Error(), 500)
	}
	return

}
