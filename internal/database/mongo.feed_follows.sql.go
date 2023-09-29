package database

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *MongoDBClient) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (FeedFollow, error) {
	db := c.Database("rssagg")
	collection := db.Collection("feed_follows")

	_, err := collection.InsertOne(ctx, arg)
	var result FeedFollow
	if err != nil {
		return result, err
	}

	collection.FindOne(ctx, bson.M{"id": arg.ID}).Decode(&result)

	return result, nil
}

func (c *MongoDBClient) UnfollowFeed(ctx context.Context, arg UnfollowFeedParams) error {
	db := c.Database("rssagg")

	feedFollowsCollection := db.Collection("feed_follows")

	_, err := feedFollowsCollection.DeleteOne(ctx,  bson.M{
		"id":     arg.ID, 
		"user_id": arg.UserID,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *MongoDBClient) GetFeedFollows(ctx context.Context, userID uuid.UUID) ([]FeedFollow, error) {
	db := c.Database("rssagg")
	collection := db.Collection("feeds")

	feedsCursor, err := collection.Find(ctx, bson.M{"user_id": userID})

	var result []FeedFollow
	if err != nil {
		return result, err
	}

	for feedsCursor.Next(ctx) {
		var feedFollow FeedFollow
		if err := feedsCursor.Decode(&feedFollow); err != nil {
			return result, err
		}
		result = append(result, feedFollow)
	}

	return result, nil
}