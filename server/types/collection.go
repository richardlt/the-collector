package types

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/richardlt/the-collector/server/database"
)

// Collection .
type Collection struct {
	database.BasicEntity `bson:",inline"`
	UserID               objectid.ObjectID `bson:"_user_id" json:"-"`
	Name                 string            `bson:"name" json:"name"`
	Slug                 string            `bson:"slug" json:"slug"`
}
