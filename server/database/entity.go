package database

import (
	"context"
	"reflect"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/updateopt"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Entity interface.
type Entity interface {
	GetID() objectid.ObjectID
	GetUUID() string
	GenerateIDs()
}

// EntityArray struct.
type EntityArray interface {
	New() interface{}
}

// BasicEntity struct contains fields for documents.
type BasicEntity struct {
	ID           objectid.ObjectID `bson:"_id" json:"-"`
	UUID         string            `bson:"uuid" json:"uuid"`
	CreationDate *time.Time        `bson:"creation_date,omitempty" json:"creation_date"`
	UpdateDate   *time.Time        `bson:"update_date,omitempty" json:"update_date"`
}

// GetID returns ID for basic entity.
func (b BasicEntity) GetID() objectid.ObjectID { return b.ID }

// GetUUID returns UUID for basic entity.
func (b BasicEntity) GetUUID() string { return b.UUID }

// GenerateIDs create values for ID and UUID.
func (b *BasicEntity) GenerateIDs() {
	b.ID = objectid.New()
	b.UUID = uuid.NewV5(uuid.NamespaceOID, b.ID.Hex()).String()
}

// Create new entity in database.
func Create(ctx context.Context, c *mongo.Collection, e Entity) error {
	e.GenerateIDs()

	_, err := c.UpdateOne(ctx,
		bson.NewDocument(bson.EC.ObjectID("_id", e.GetID())),
		bson.NewDocument(
			bson.EC.Interface("$set", e),
			bson.EC.SubDocument("$currentDate", bson.NewDocument(
				bson.EC.SubDocument("creation_date", bson.NewDocument(
					bson.EC.String("$type", "date"),
				)),
				bson.EC.SubDocument("update_date", bson.NewDocument(
					bson.EC.String("$type", "date"),
				)),
			)),
		),
		updateopt.Upsert(true),
	)
	if err != nil {
		return errors.WithStack(err)
	}

	res := c.FindOne(ctx, bson.NewDocument(bson.EC.ObjectID("_id", e.GetID())))
	return errors.WithStack(res.Decode(e))
}

// Delete existing entity in database.
func Delete(ctx context.Context, c *mongo.Collection, e Entity) error {
	_, err := c.DeleteOne(ctx, bson.NewDocument(bson.EC.ObjectID("_id", e.GetID())))
	return errors.WithStack(err)
}

// GetAll populates entity array with results of given pipeline.
func GetAll(ctx context.Context, c *mongo.Collection, p Pipeline, e EntityArray) error {
	res, err := p.convertToDocumentSlice()
	if err != nil {
		return err
	}

	cu, err := c.Aggregate(ctx, res)
	if err != nil {
		return errors.WithStack(err)
	}

	for cu.Next(ctx) {
		if err := cu.Err(); err != nil {
			return errors.WithStack(err)
		}
		if err := cu.Decode(e.New()); err != nil {
			return errors.WithStack(err)
		}
	}

	return errors.WithStack(cu.Close(ctx))
}

// Get result of given pipeline, take a ptr to value.
// If no result found reset the value.
func Get(ctx context.Context, c *mongo.Collection, q Query, i interface{}) error {
	v := reflect.ValueOf(i)
	t := v.Elem().Type()
	k := t.Kind()

	// if given value is an ptr of ptr
	if k == reflect.Ptr {
		v = v.Elem()
	}

	// init nil ptr with value to prevent error in decode
	if k == reflect.Ptr && v.IsNil() {
		v.Set(reflect.New(t.Elem()))
	}

	err := c.FindOne(ctx, q).Decode(v.Interface())
	if err == mongo.ErrNoDocuments {
		// reset value if ptr of ptr
		if k == reflect.Ptr {
			v.Set(reflect.Zero(t))
		}
		return nil
	}

	return errors.WithStack(err)
}

// DeleteAll documents in collection that match query.
func DeleteAll(ctx context.Context, c *mongo.Collection, q Query) (int64, error) {
	res, err := c.DeleteMany(ctx, q)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return res.DeletedCount, nil
}
