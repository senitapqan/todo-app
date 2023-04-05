package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webapp/model"
)

func (h *Handler) createItem(wr http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		newErrorResponse(wr, http.StatusBadRequest, "invalid id param")
		return
	}

	listId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newErrorResponse(wr, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input model.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(wr, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)

	if err != nil {
		newErrorResponse(wr, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(wr).Encode(map[string] interface{} {
		"id": id,
	})
}

func (h *Handler) getAllItems(wr http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		newErrorResponse(wr, http.StatusBadRequest, "invalid id param")
		return
	}

	listId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newErrorResponse(wr, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)

	if err != nil {
		newErrorResponse(wr, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(wr).Encode(items)
}

func (h *Handler) getItemById(wr http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		newErrorResponse(wr, http.StatusBadRequest, "invalid id param")
		return
	}

	itemId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		newErrorResponse(wr, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)

	if err != nil {
		newErrorResponse(wr, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(wr).Encode(item)
}

func (h *Handler) deleteItem(wr http.ResponseWriter, r *http.Request) {

}