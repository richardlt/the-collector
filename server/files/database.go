package files

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/richardlt/the-collector/server/database"
	"github.com/richardlt/the-collector/server/types"
)

var collection *mongo.Collection

func newFiles() files { return files{data: []*types.File{}} }

type files struct{ data []*types.File }

func (f *files) New() interface{} {
	fi := &types.File{}
	f.data = append(f.data, fi)
	return fi
}

// InitDatabase .
func InitDatabase(ctx context.Context, db *mongo.Database) error {
	collection = db.Collection("files")
	return database.EnsureIndexes(ctx, collection)
}

// Create .
func Create(ctx context.Context, f *types.File) error {
	return database.Create(ctx, collection, f)
}

// GetAll .
func GetAll(ctx context.Context, c database.Criteria) ([]*types.File, error) {
	return GetAllCustom(ctx, database.Pipeline{database.Match(c.Build())})
}

// GetAllCustom .
func GetAllCustom(ctx context.Context, p database.Pipeline) ([]*types.File, error) {
	f := newFiles()
	return f.data, database.GetAll(ctx, collection, p, &f)
}

// Get .
func Get(ctx context.Context, c database.Criteria) (*types.File, error) {
	var f *types.File
	return f, database.Get(ctx, collection, c.Build(), &f)
}
