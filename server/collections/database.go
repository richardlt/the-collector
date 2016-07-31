package collections

import (
	"github.com/richardlt/the-collector/server/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Collection : database collection for collections
var Collection *mgo.Collection

// Create : create collection in database
func Create(c *types.Collection) error {
	return c.Create(Collection, c)
}

// GetByUUID : get collection by uuid from database
func GetByUUID(UUID string) (*types.Collection, error) {
	var c *types.Collection
	err := Collection.Find(bson.M{"uuid": UUID}).One(&c)
	return c, err
}

// GetAll : get all collections from database
func GetAll() ([]*types.Collection, error) {
	cs := []*types.Collection{}
	err := Collection.Find(nil).All(&cs)
	return cs, err
}
