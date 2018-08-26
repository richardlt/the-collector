package users

import (
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/richardlt/the-collector/server/database"
)

func newCriteria() *criteria { return &criteria{} }

type criteria struct {
	database.BasicCriteria
}

func (c *criteria) FacebookID(ids ...string) *criteria {
	c.AppendQuery(database.In("facebook_id", ids))
	return c
}

func (c *criteria) ID(ids objectid.ObjectID) *criteria {
	c.BasicCriteria.ID(ids)
	return c
}

func newCriteriaAuth() *criteriaAuth { return &criteriaAuth{} }

type criteriaAuth struct {
	database.BasicCriteria
}

func (c *criteriaAuth) Type(ts ...string) *criteriaAuth {
	c.AppendQuery(database.In("type", ts))
	return c
}

func (c *criteriaAuth) UserID(ids ...objectid.ObjectID) *criteriaAuth {
	c.AppendQuery(database.In("_user_id", ids))
	return c
}

func (c *criteriaAuth) AccessToken(ts ...string) *criteriaAuth {
	c.AppendQuery(database.In("access_token", ts))
	return c
}
