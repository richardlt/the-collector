package database

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

// EnsureIndexes create custom and default indexes.
func EnsureIndexes(ctx context.Context, c *mongo.Collection, ks ...string) error {
	var is []mongo.IndexModel

	for _, k := range append([]string{
		"uuid",
		"creation_date",
		"update_date"},
		ks...,
	) {
		is = append(is, mongo.IndexModel{
			Keys: bson.NewDocument(bson.EC.Int32(k, -1)),
		})
	}

	_, err := c.Indexes().CreateMany(ctx, is)
	return errors.WithStack(err)
}
