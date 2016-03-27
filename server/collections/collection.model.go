package collections

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DatabaseCollection : database collection for collections
var DatabaseCollection *mgo.Collection

// ExistSlug : check if slug exist in database
func ExistSlug(slug string) (bool, error) {
	c, err := DatabaseCollection.Find(bson.M{"slug": slug}).Count()
	if err != nil {
		return true, err
	}
	return c > 0, nil
}

// GetBySlug : get collection by slug from database
func GetBySlug(slug string) (*Collection, error) {
	c := new(Collection)
	err := DatabaseCollection.Find(bson.M{"slug": slug}).One(&c)
	return c, err
}

// GetAll : get all collections from database
func GetAll() ([]Collection, error) {
	c := make([]Collection, 0)
	err := DatabaseCollection.Find(nil).All(&c)
	return c, err
}

// Collection : collection struct
type Collection struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Slug string        `bson:"slug"`
}

// Restify : return rest collection
func (c Collection) Restify() CollectionRest {
	return CollectionRest{
		Name: c.Name,
		Slug: c.Slug,
	}
}

// Create : create collection in database
func (c *Collection) Create() error {
	c.ID = bson.NewObjectId()
	err := DatabaseCollection.Insert(c)
	return err
}

// CollectionRest : collection rest struct
type CollectionRest struct {
	Name string `json:"name"`
	Slug string `json:"name"`
}

// ToCollection : return collection
func (c CollectionRest) ToCollection() Collection {
	return Collection{Name: c.Name, Slug: c.Slug}
}
