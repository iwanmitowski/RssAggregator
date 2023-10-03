package database

import (
	"time"

	"github.com/google/uuid"
)

type UserDto struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func ToUserDto(source User) UserDto {
	return UserDto{
		ID:        source.ID,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
		Name:      source.Name,
		APIKey:    source.ApiKey,
	}
}

type FeedDto struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func ToFeedDto(source Feed) FeedDto {
	return FeedDto{
		ID:        source.ID,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
		Name:      source.Name,
		Url:       source.Url,
		UserID:    source.UserID,
	}
}

func ToFeedDtos(source []Feed) []FeedDto {
	result := []FeedDto{}

	for _, feedDto := range source {
		result = append(result, ToFeedDto(feedDto))
	}

	return result
}

type FeedFollowDto struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id" bson:"feedid"`
}

func ToFeedFollowDto(source FeedFollow) FeedFollowDto {
	return FeedFollowDto{
		ID:        source.ID,
		CreatedAt: source.CreatedAt,
		UpdatedAt: source.UpdatedAt,
		UserID:    source.UserID,
		FeedID:    source.FeedID,
	}
}

func ToFeedFollowDtos(source []FeedFollow) []FeedFollowDto {
	result := []FeedFollowDto{}

	for _, feedFollowDto := range source {
		result = append(result, ToFeedFollowDto(feedFollowDto))
	}

	return result
}

type PostDto struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func ToPostDto(source Post) PostDto {
	var description *string

	if source.Description.Valid {
		description = &source.Description.String
	}

	return PostDto{
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

func ToPostDtos(source []Post) []PostDto {
	result := []PostDto{}

	for _, postDto := range source {
		result = append(result, ToPostDto(postDto))
	}

	return result
}
