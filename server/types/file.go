package types

// File : file struct
type File struct {
	Entity `bson:",inline"`
	Path   string `bson:"path" json:"-" valid:"-"`
	// generated
	Href string `bson:",omitempty" json:"href,omitempty" valid:"-"`
}
