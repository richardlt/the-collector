package types

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/richardlt/the-collector/server/database"
)

// Item .
type Item struct {
	database.BasicEntity `bson:",inline"`
	CollectionID         objectid.ObjectID `bson:"_collection_id" json:"-"`
	Picture              string            `bson:"-" json:"picture"`
}
