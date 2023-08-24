package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/iwanmitowski/RssAggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func toUser(source database.User) User {
	return User{
		ID:        source.ID,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
		Name:      source.Name,
		APIKey:    source.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func toFeed(source database.Feed) Feed {
	return Feed{
		ID:        source.ID,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
		Name:      source.Name,
		Url:       source.Url,
		UserID:    source.ID,
	}
}
