package collections

import "github.com/richardlt/the-collector/server/database"

type criteria struct {
	database.BasicCriteria
}

func newCriteria() *criteria { return &criteria{} }

func (c criteria) Build() database.Query { return c.BasicCriteria.Build() }

func (c *criteria) SlugOrUUID(v ...string) *criteria {
	c.AppendQuery(database.Or(
		database.In("slug", v),
		database.In("uuid", v),
	))
	return c
}

func (c *criteria) Slug(v ...string) *criteria {
	c.AppendQuery(database.In("slug", v))
	return c
}
