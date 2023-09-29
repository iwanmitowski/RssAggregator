package database

import (
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *MongoDBClient) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	db := c.Database("rssagg")
	collection := db.Collection("posts")

	_, err := collection.InsertOne(ctx, arg)
	var result Post
	if err != nil {
		return result, err
	}

	collection.FindOne(ctx, bson.M{"id": arg.ID}).Decode(&result)

	return result, nil
}

func (c *MongoDBClient) GetPostsForUser(ctx context.Context, arg GetPostsForUserParams) ([]Post, error) {
	db := c.Database("rssagg")

	feedFollowsCollection := db.Collection("feed_follows")
	feedFollowsCursor, err := feedFollowsCollection.Find(ctx, bson.M{"user_id": arg.UserID})

	var result []Post
	var feedFollows []FeedFollow

	if err != nil {
		return result, err
	}

	for feedFollowsCursor.Next(ctx) {
		var feedFollow FeedFollow
		if err := feedFollowsCursor.Decode(&feedFollows); err != nil {
			continue
		}
		feedFollows = append(feedFollows, feedFollow)
	}

	if len(feedFollows) == 0 {
		return result, nil
	}

	followedFeedIds := make([]uuid.UUID, len(feedFollows))
	for i, feedFollow := range feedFollows {
		followedFeedIds[i] = feedFollow.FeedID
	}

	postsCollection := db.Collection("posts")
	postsCursor, err := postsCollection.Find(ctx, bson.M{"feed_id": bson.M{"$in": followedFeedIds}})

	if err != nil {
		return result, err
	}

	for postsCursor.Next(ctx) {
		var post Post
		if err := postsCursor.Decode(&feedFollows); err != nil {
			continue
		}
		result = append(result, post)
	}

	return result, nil
}
