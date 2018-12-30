package users

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/richardlt/the-collector/server/database"
	"github.com/richardlt/the-collector/server/types"
)

var collection *mongo.Collection
var collectionAuth *mongo.Collection

// InitDatabase .
func InitDatabase(ctx context.Context, db *mongo.Database) error {
	collection = db.Collection("users")
	collectionAuth = db.Collection("auths")
	if err := database.EnsureIndexes(ctx, collection); err != nil {
		return err
	}
	return database.EnsureIndexes(ctx, collectionAuth)
}

func create(ctx context.Context, u *types.User) error {
	return database.Create(ctx, collection, u)
}

func get(ctx context.Context, c *criteria) (*types.User, error) {
	var u *types.User
	return u, database.Get(ctx, collection, c.Build(), &u)
}

func createAuth(ctx context.Context, a *types.Auth) error {
	return database.Create(ctx, collectionAuth, a)
}

func getAuth(ctx context.Context, c *criteriaAuth) (*types.Auth, error) {
	var a *types.Auth
	return a, database.Get(ctx, collectionAuth, c.Build(), &a)
}

// deleteAllAuth .
func deleteAllAuth(ctx context.Context, c *criteriaAuth) (int64, error) {
	return database.DeleteAll(ctx, collectionAuth, c.Build())
}
