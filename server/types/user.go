package types

import (
	"github.com/richardlt/the-collector/server/database"
)

// User .
type User struct {
	database.BasicEntity `bson:",inline"`
	Name                 string `bson:"name" json:"name"`
	Email                string `bson:"email" json:"email,omitempty"`
	FacebookID           string `bson:"facebook_id" json:"-"`
}
