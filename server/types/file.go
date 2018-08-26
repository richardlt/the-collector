package types

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
	"github.com/richardlt/the-collector/server/database"
)

// Content types
const (
	ImageGif  string = "image/gif"
	ImageJpeg string = "image/jpeg"
	ImagePng  string = "image/png"
)

// IsImageContentType .
func IsImageContentType(contentType string) bool {
	switch contentType {
	case ImageGif, ImageJpeg, ImagePng:
		return true
	}
	return false
}

// Image sizes
const (
	Small    string = "small"
	Medium   string = "medium"
	Large    string = "large"
	Original string = ""
)

// IsValidSize .
func IsValidSize(size string) bool {
	switch size {
	case Small, Medium, Large, Original:
		return true
	}
	return false
}

// SizeToPixels .
func SizeToPixels(size string) int {
	switch size {
	case Small:
		return 100
	case Medium:
		return 200
	case Large:
		return 400
	}
	return 600
}

// ResourceType values.
const (
	ItemResource string = "item"
)

// File .
type File struct {
	database.BasicEntity `bson:",inline"`
	ResourceType         string            `bson:"resource_type"`
	ResourceID           objectid.ObjectID `bson:"_resource_id"`
	Name                 string            `bson:"name"`
	Path                 string            `bson:"path"`
	Size                 int64             `bson:"size"`
	ContentType          string            `bson:"content_type"`
}

// GenerateURI returns new URI to access a file four 24 hours.
func (f File) GenerateURI(jwtSecret string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type": "file",
		"uuid": f.UUID,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})

	token, err := t.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.WithStack(err)
	}

	return fmt.Sprintf("/api/files/%s/%s", token, f.Name), nil
}
