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
		UserID:    source.UserID,
	}
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func toFeedFollow(source database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        source.ID,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
		UserID:    source.UserID,
		FeedID:    source.FeedID,
	}
}

func toFeedFollows(source []database.FeedFollow) []FeedFollow {
	result := []FeedFollow{}

	for _, feedFollow := range source {
		result = append(result, toFeedFollow(feedFollow))
	}

	return result
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func toPost(source database.Post) Post {
	var description *string

	if source.Description.Valid {
		description = &source.Description.String
	}

	return Post{
		ID:          source.ID,
		CreatedAt:   source.CreatedAt,
		UpdatedAt:   source.UpdatedAt,
		Title:       source.Title,
		Description: description,
		PublishedAt: source.PublishedAt,
		Url:         source.Url,
		FeedID:      source.FeedID,
	}
}

func toPosts(source []database.Post) []Post {
	result := []Post{}

	for _, post := range source {
		result = append(result, toPost(post))
	}

	return result
}
