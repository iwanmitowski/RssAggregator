package database

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *MongoDBClient) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	db := c.Database("rssagg")
	collection := db.Collection("users")

	randomValue := make([]byte, 32) // 32 bytes for SHA-256
	rand.Read(randomValue)

	// Calculate the SHA-256 hash
	hash := sha256.Sum256(randomValue)

	// Encode the hash as hexadecimal
	hashHex := hex.EncodeToString(hash[:])

	user := &User{
		ID:        arg.ID,
		CreatedAt: arg.CreatedAt,
		UpdatedAt: arg.UpdatedAt,
		Name:      arg.Name,
		ApiKey:    hashHex,
	}

	_, err := collection.InsertOne(ctx, user)
	var result User
	if err != nil {
		return result, err
	}

	// Kak shte go vurne?
	collection.FindOne(ctx, bson.M{"id": arg.ID}).Decode(&result)

	return result, nil
}

func (c *MongoDBClient) GetUserByAPIKey(ctx context.Context, apiKey string) (User, error) {
	db := c.Database("rssagg")
	collection := db.Collection("users")

	var user User
	err := collection.FindOne(ctx, bson.M{"apikey": apiKey}).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}
