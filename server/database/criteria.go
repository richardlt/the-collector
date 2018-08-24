package database

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// Criteria is used to build database query.
type Criteria interface {
	Build() Query
}

// NewBasicCriteria .
func NewBasicCriteria() *BasicCriteria { return &BasicCriteria{} }

// BasicCriteria for documents.
type BasicCriteria struct{ queries []Query }

// AppendQuery add a new query to the list.
func (b *BasicCriteria) AppendQuery(q Query) { b.queries = append(b.queries, q) }

// Build the query from criteria.
func (b *BasicCriteria) Build() Query { return And(b.queries...) }

// ID appender for basic criteria.
func (b *BasicCriteria) ID(ids ...objectid.ObjectID) *BasicCriteria {
	b.AppendQuery(In("_id", ids))
	return b
}

// UUID appender for basic criteria.
func (b *BasicCriteria) UUID(uuids ...string) *BasicCriteria {
	b.AppendQuery(In("uuid", uuids))
	return b
}
