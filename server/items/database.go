package items

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/richardlt/the-collector/server/database"
	"github.com/richardlt/the-collector/server/types"
)

var collection *mongo.Collection

func newItems() items { return items{data: []*types.Item{}} }

type items struct{ data []*types.Item }

func (i *items) New() interface{} {
	it := &types.Item{}
	i.data = append(i.data, it)
	return it
}

// InitDatabase .
func InitDatabase(ctx context.Context, db *mongo.Database) error {
	collection = db.Collection("items")
	return database.EnsureIndexes(ctx, collection, "_collection_id")
}

// Create .
func Create(ctx context.Context, i *types.Item) error {
	return database.Create(ctx, collection, i)
}

// Delete .
func Delete(ctx context.Context, i *types.Item) error {
	return database.Delete(ctx, collection, i)
}

// GetAll .
func GetAll(ctx context.Context, c database.Criteria) ([]*types.Item, error) {
	return GetAllCustom(ctx, database.Pipeline{database.Match(c.Build())})
}

// GetAllCustom .
func GetAllCustom(ctx context.Context, p database.Pipeline) ([]*types.Item, error) {
	i := newItems()
	return i.data, database.GetAll(ctx, collection, p, &i)
}

// Get .
func Get(ctx context.Context, c database.Criteria) (*types.Item, error) {
	var i *types.Item
	return i, database.Get(ctx, collection, c.Build(), &i)
}
