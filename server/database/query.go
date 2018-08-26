package database

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/pkg/errors"
)

// Query type for database request.
type Query map[string]interface{}

// And returns new $and query for given list.
func And(qs ...Query) Query {
	if len(qs) > 1 {
		return Query{"$and": qs}
	} else if len(qs) == 1 {
		return qs[0]
	}
	return Query{}
}

// Or returns new $or query for given list.
func Or(qs ...Query) Query {
	if len(qs) > 1 {
		return Query{"$or": qs}
	} else if len(qs) == 1 {
		return qs[0]
	}
	return Query{}
}

// In returns new $in query for f key.
func In(f string, i interface{}) Query { return Query{f: Query{"$in": i}} }

// Match returns new $match query from query.
func Match(q Query) Query { return Query{"$match": q} }

// Pipeline type for database request.
type Pipeline []Query

func (p Pipeline) convertToDocumentSlice() ([]*bson.Document, error) {
	var res []*bson.Document
	for _, q := range p {
		d, err := mongo.TransformDocument(q)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		res = append(res, d)
	}
	return res, nil
}
