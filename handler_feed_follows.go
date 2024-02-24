package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/davifrjose/BOILERPLATE/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedsFolowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't decode parameter")
	}
	feedFollows , err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (cfg *apiConfig) handlerFeedsFolowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	idstr := chi.URLParam(r, "id")
	if idstr == "" {
		responseWithError(w, http.StatusInternalServerError, "ID is required")
		return
	}
	iduuid, err := uuid.Parse(idstr)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't parse from string to uuid")
		return
	}
	  err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
			ID:	iduuid ,
			UserID: user.ID,
		})

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}
	respondWithJSON(w, http.StatusOK, struct {}{})
}

func (cfg *apiConfig) handlerFeedsFollowFromUser(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows , err := cfg.DB.GetFeedFollowForUser(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't get feedFollows")
		return
	}
	respondWithJSON(w, http.StatusOK, databasesFeedFollowToFeedFollow(feedFollows))
}