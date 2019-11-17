package gql

import (
	"personal-website-graphql/postgres"

	"github.com/graphql-go/graphql"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot(db *postgres.Db) *Root {
	// Create a resolver
	resolver := Resolver{db: db}

	// Create root for query
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"projects": &graphql.Field{
						// Slice of Project type
						Type:    graphql.NewList(Project),
						Resolve: resolver.ProjectResolver,
					},
				},
			},
		),
	}
	return &root
}
