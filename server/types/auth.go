package types

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/richardlt/the-collector/server/database"
)

// AuthType values.
const (
	Local    string = "local"
	Facebook string = "facebook"
)

// Auth .
type Auth struct {
	database.BasicEntity `bson:",inline"`
	UserID               objectid.ObjectID `bson:"_user_id"`
	Type                 string            `bson:"type"`
	AccessToken          string            `bson:"access_token"`
	Expires              time.Time         `bson:"expires"`
}
