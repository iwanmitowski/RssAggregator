package database

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *MongoDBClient) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	db := c.Database("rssagg")
	collection := db.Collection("feeds")

	_, err := collection.InsertOne(ctx, arg)
	var result Feed
	if err != nil {
		return result, err
	}

	collection.FindOne(ctx, bson.M{"id": arg.ID}).Decode(&result)

	return result, nil
}

func (c *MongoDBClient) GetNextFeedsToFetch(ctx context.Context, limit int32) ([]Feed, error) {
	db := c.Database("rssagg")
	collection := db.Collection("feeds")

	findOptions := options.Find().SetSort(map[string]int{"last_fetched_at": 1}).SetLimit(int64(limit))

	feedsCursor, err := collection.Find(ctx, nil, findOptions)

	var result []Feed
	if err != nil {
		return result, err
	}

	for feedsCursor.Next(ctx) {
		var feed Feed
		if err := feedsCursor.Decode(&feed); err != nil {
			return result, err
		}
		result = append(result, feed)
	}

	return result, nil
}

func (c *MongoDBClient) MarkFeedFetched(ctx context.Context, id uuid.UUID) (Feed, error) {
	db := c.Database("rssagg")
	collection := db.Collection("feeds")

	update := bson.M{
		"$set": bson.M{
			"last_fetched_at": time.Now().UTC(),
			"updated_at":      time.Now().UTC(),
		},
	}

	var feed Feed
	err := collection.FindOneAndUpdate(ctx, bson.M{"id": id}, update).Decode(&feed)

	if err != nil {
		return feed, err
	}

	return feed, nil
}

func (c *MongoDBClient) GetNotFollowedFeeds(ctx context.Context, userID uuid.UUID) ([]Feed, error) {
	db := c.Database("rssagg")
	
	feedFollowsCollection := db.Collection("feed_follows")
	feedIdsFollowedByUserCursor, err := feedFollowsCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer feedIdsFollowedByUserCursor.Close(ctx)

	var feedIdsFollowedByUser []uuid.UUID
	for feedIdsFollowedByUserCursor.Next(ctx) {
		var feedFollow FeedFollow
		if err := feedIdsFollowedByUserCursor.Decode(&feedFollow); err != nil {
			continue
		}
		feedIdsFollowedByUser = append(feedIdsFollowedByUser, feedFollow.FeedID)
	}

	feedsCollection := db.Collection("feeds")
	feedsNotFollowedByUserCursor, err := feedsCollection.Find(ctx, bson.M{"_id": bson.M{"$nin": feedIdsFollowedByUser}})
	if err != nil {
		return nil, err
	}
	defer feedsNotFollowedByUserCursor.Close(ctx)

	var result []Feed
	for feedsNotFollowedByUserCursor.Next(ctx) {
		var feed Feed
		if err := feedsNotFollowedByUserCursor.Decode(&feed); err != nil {
			continue
		}
		result = append(result, feed)
	}

	return result, nil
}
