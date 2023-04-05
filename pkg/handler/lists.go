package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/model"
)

func (h *Handler) createList(wr http.ResponseWriter, r *http.Request) {
	fmt.Println("i was here")
	userId, err := getUserId(r)
	

	if err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}

	var input model.TodoList
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}

	json.NewEncoder(wr).Encode(id)

}

type getAllListsResponse struct {
	Data []model.TodoList `json:"data"`
}

func (h *Handler) getAllLists(wr http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		http.Error(wr, err.Error(), 500)
		return
	}

	json.NewEncoder(wr).Encode(getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(wr http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		http.Error(wr, err.Error(), 400)
		return
	}

	var id int
	id, err = strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		newErrorResponse(wr, http.StatusBadRequest, "Invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, id)

	if err != nil {
		newErrorResponse(wr, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(wr).Encode(list)
}

func (h *Handler) deleteList(wr http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		return
	}

	var id int
	id, err = strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		newErrorResponse(wr, http.StatusBadRequest, "Invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, id)

	if err != nil {
		newErrorResponse(wr, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(wr).Encode(statusResponse{Status: "ok"})
}
