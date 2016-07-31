package types

import (
	"time"

	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Entity : database entity struct
type Entity struct {
	ID           bson.ObjectId `bson:"_id" json:"-" valid:"-"`
	UUID         string        `bson:"uuid" json:"uuid" valid:"-"`
	CreationDate time.Time     `bson:"creation_date,omitempty" json:"creation_date" valid:"-"`
	UpdateDate   time.Time     `bson:"update_date,omitempty" json:"update_date" valid:"-"`
}

// Create : create new entity in database
func (e *Entity) Create(c *mgo.Collection, i interface{}) error {
	e.ID = bson.NewObjectId()
	e.UUID = uuid.NewV5(uuid.NamespaceOID, e.ID.Hex()).String()
	_, err := c.Upsert(bson.M{"_id": e.ID}, bson.M{
		"$set": i,
		"$currentDate": bson.M{
			"creation_date": bson.M{"$type": "date"},
			"update_date":   bson.M{"$type": "date"},
		},
	})
	if err != nil {
		return err
	}
	return c.Find(bson.M{"_id": e.ID}).One(i)
}

// Save : save entity in database
func (e *Entity) Save(c *mgo.Collection, i interface{}) error {
	e.UpdateDate = time.Time{}
	_, err := c.Upsert(bson.M{"_id": e.ID}, bson.M{
		"$set": i,
		"$currentDate": bson.M{
			"update_date": bson.M{"$type": "date"},
		},
	})
	if err != nil {
		return err
	}
	return c.Find(bson.M{"_id": e.ID}).One(i)
}

// Delete : delete entity in database
func (e *Entity) Delete(c *mgo.Collection, i interface{}) error {
	return c.RemoveId(e.ID)
}
