package items

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// DatabaseCollection : database collection for items
var DatabaseCollection *mgo.Collection

// ExistSlug : check if slug exist in database
func ExistSlug(slug string) (bool, error) {
	c, err := DatabaseCollection.Find(bson.M{"slug": slug}).Count()
	if err != nil {
		return true, err
	}
	return c > 0, nil
}

// GetBySlug : get item by slug from database
func GetBySlug(slug string) (*Item, error) {
	i := new(Item)
	err := DatabaseCollection.Find(bson.M{"slug": slug}).One(&i)
	return i, err
}

// GetAll : get all items from database
func GetAll() ([]Item, error) {
	i := make([]Item, 0)
	err := DatabaseCollection.Find(nil).All(&i)
	return i, err
}

// Item : item struct
type Item struct {
	ID    bson.ObjectId `bson:"_id"`
	Name  string        `bson:"name"`
	Slug  string        `bson:"slug"`
	Image []byte        `bson:"image"`
}

// Restify : return rest item
func (i Item) Restify() ItemRest {
	return ItemRest{
		Name: i.Name,
		Slug: i.Slug,
	}
}

// Create : create item in database
func (i *Item) Create() error {
	i.ID = bson.NewObjectId()
	err := DatabaseCollection.Insert(i)
	return err
}

// ItemRest : item rest struct
type ItemRest struct {
	Name string `json:"name"`
	Slug string `json:"name"`
}

// ToItem : return item
func (c ItemRest) ToItem() Item {
	return Item{Name: c.Name, Slug: c.Slug}
}
