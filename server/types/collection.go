package types

// Collection : collection struct
type Collection struct {
	Entity `bson:",inline"`
	Name   string `bson:"name" json:"name" valid:"string,len(0|100)"`
	Slug   string `bson:"slug" json:"slug" valid:"-"`
}
