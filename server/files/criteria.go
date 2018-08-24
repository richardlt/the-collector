package files

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/richardlt/the-collector/server/database"
)

// Criteria .
type Criteria struct {
	database.BasicCriteria
}

// NewCriteria .
func NewCriteria() *Criteria { return &Criteria{} }

// Build .
func (c Criteria) Build() database.Query { return c.BasicCriteria.Build() }

// Name .
func (c *Criteria) Name(fs ...string) *Criteria {
	c.AppendQuery(database.In("name", fs))
	return c
}

// ResourceType .
func (c *Criteria) ResourceType(ts ...string) *Criteria {
	c.AppendQuery(database.In("resource_type", ts))
	return c
}

// ResourceID .
func (c *Criteria) ResourceID(ids ...objectid.ObjectID) *Criteria {
	c.AppendQuery(database.In("_resource_id", ids))
	return c
}
