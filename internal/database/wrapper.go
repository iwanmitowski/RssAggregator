package database

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database interface {
	// Define the methods you need for your application
	CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error)
	GetNextFeedsToFetch(ctx context.Context, limit int32) ([]Feed, error)
	MarkFeedFetched(ctx context.Context, id uuid.UUID) (Feed, error)
	GetNotFollowedFeeds(ctx context.Context, userID uuid.UUID) ([]Feed, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUserByAPIKey(ctx context.Context, apiKey string) (User, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]Post, error)
	CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error)
	UnfollowFeed(ctx context.Context, arg UnfollowFeedParams) error
	GetFeedFollows(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error)
}

type PostGresDBClient struct {
	*Queries
}

type MongoDBClient struct {
	*mongo.Client
}
