package main

import (
	"fmt"
	"net/http"

	"github.com/davifrjose/BOILERPLATE/internal/auth"
	"github.com/davifrjose/BOILERPLATE/internal/database"
)

type authedHandler func (http.ResponseWriter,	*http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func ( w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user , err := cfg.DB.GetUserByAPIKEY(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn´t get user: %v", err))
		return
	}
	handler(w, r, user)
	}
}