package gql

import (
	"personal-website-graphql/postgres"

	"github.com/graphql-go/graphql"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *postgres.Db
}

// ProjectResolver resolves our user query through a db call to GetProjects
func (r *Resolver) ProjectResolver(p graphql.ResolveParams) (interface{}, error) {
	projects := r.db.GetAllProjects()
	return projects, nil
}
