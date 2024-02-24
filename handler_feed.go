package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/davifrjose/BOILERPLATE/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct { 
	Name      string `json:"name"`
	Url       string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
	}
	feed , err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams {
		ID: uuid.New() ,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.Url,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}
	// creation of feedFollow
	feedFollows , err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil  {
		responseWithError(w, http.StatusInternalServerError, "Couldn't create feed follow ")
		return
	}
	
	respondWithJSON(w, http.StatusCreated,  struct {
		feed       Feed
		feedFollow FeedFollows
	}{
		feed:       databaseFeedToFeed(feed),
		feedFollow: databaseFeedFollowsToFeedFollows(feedFollows),
	})
}

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
	feeds , err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}
	respondWithJSON(w, http.StatusOK, databasesFeedToFeed(feeds))
}