package collections

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/richardlt/the-collector/server/database"
	"github.com/richardlt/the-collector/server/types"
)

var collection *mongo.Collection

func newCollections() collections {
	return collections{data: []*types.Collection{}}
}

type collections struct{ data []*types.Collection }

func (c *collections) New() interface{} {
	co := &types.Collection{}
	c.data = append(c.data, co)
	return co
}

// InitDatabase .
func InitDatabase(ctx context.Context, db *mongo.Database) error {
	collection = db.Collection("collections")
	return database.EnsureIndexes(ctx, collection, "slug")
}

// Create .
func Create(ctx context.Context, c *types.Collection) error {
	return database.Create(ctx, collection, c)
}

// Delete .
func Delete(ctx context.Context, c *types.Collection) error {
	return database.Delete(ctx, collection, c)
}

// GetAll .
func GetAll(ctx context.Context, c database.Criteria) ([]*types.Collection, error) {
	return GetAllCustom(ctx, database.Pipeline{database.Match(c.Build())})
}

// GetAllCustom .
func GetAllCustom(ctx context.Context, p database.Pipeline) ([]*types.Collection, error) {
	cs := newCollections()
	return cs.data, database.GetAll(ctx, collection, p, &cs)
}

// Get .
func Get(ctx context.Context, c database.Criteria) (*types.Collection, error) {
	var co *types.Collection
	return co, database.Get(ctx, collection, c.Build(), &co)
}
