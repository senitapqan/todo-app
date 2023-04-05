package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentify(ht http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			newErrorResponse(w, http.StatusUnauthorized, "invalid is empty")
			return
		}

		if len(headerParts[1]) == 0 {
			newErrorResponse(w, http.StatusUnauthorized, "token is empty")
			return
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			newErrorResponse(w, http.StatusUnauthorized, "invalid is empty")
			return
		}

		r.Header.Set(userCtx, strconv.Itoa(userId))
		ht.ServeHTTP(w, r)
	})
}

func getUserId(r *http.Request) (int, error) {
	id := r.Header.Get(userCtx)
	if id == "" {
		return 0, errors.New("user id not found")
	}

	var userId int
	userId, err := strconv.Atoi(id)

	if err != nil {
		return 0, errors.New("user id not found")
	}

	return userId, nil
}
