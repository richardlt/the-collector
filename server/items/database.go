package items

import (
	"github.com/richardlt/the-collector/server/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Collection : database collection for items
var Collection *mgo.Collection

// Create : create item in database
func Create(i *types.Item) error {
	return i.Create(Collection, i)
}

// GetByUUID : get item by uuid from database
func GetByUUID(UUID string) (*types.Item, error) {
	var i *types.Item
	err := Collection.Find(bson.M{"uuid": UUID}).One(&i)
	return i, err
}

// GetAllByCollectionID : get all items by collection id from database
func GetAllByCollectionID(collectionID bson.ObjectId) ([]*types.Item, error) {
	is := []*types.Item{}
	err := Collection.Find(bson.M{"_collection_id": collectionID}).All(&is)
	return is, err
}
