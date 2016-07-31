package types

import "gopkg.in/mgo.v2/bson"

// Item : item struct
type Item struct {
	Entity       `bson:",inline"`
	CollectionID bson.ObjectId `bson:"_collection_id" json:"-" valid:"-"`
	Name         string        `bson:"name" json:"name" valid:"string,len(0|100)"`
	// generated
	Picture *File `bson:"picture,omitempty" json:"picture,omitempty" valid:"-"`
}
