package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/iwanmitowski/RssAggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: ", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create feed_follow: ", err))
		return
	}

	respondWithJSON(w, 201, toFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt get feed_follow: ", err))
		return
	}

	respondWithJSON(w, 201, toFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDParam := chi.URLParam(r, "feedFollowID")

	feedFollowID, err := uuid.Parse((feedFollowIDParam))

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt parse feed follow id: ", err))
		return
	}

	// Not returning error if feed is not followed - TODO
	err = apiCfg.DB.UnfollowFeed(r.Context(), database.UnfollowFeedParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt unfollow feed: ", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
