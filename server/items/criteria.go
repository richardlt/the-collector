package items

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"

	"github.com/richardlt/the-collector/server/database"
)

type criteria struct {
	database.BasicCriteria
}

func newCriteria() *criteria { return &criteria{} }

func (c criteria) Build() database.Query { return c.BasicCriteria.Build() }

func (c *criteria) CollectionID(ids ...objectid.ObjectID) *criteria {
	c.AppendQuery(database.In("_collection_id", ids))
	return c
}
