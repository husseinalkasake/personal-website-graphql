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
					"about": &graphql.Field{
						Type:    graphql.NewList(About),
						Resolve: resolver.AboutResolver,
					},
					"education": &graphql.Field{
						Type:    graphql.NewList(Education),
						Resolve: resolver.EducationResolver,
					},
					"experience": &graphql.Field{
						Type:    graphql.NewList(Experience),
						Resolve: resolver.ExperienceResolver,
					},
					"projects": &graphql.Field{
						Type:    graphql.NewList(Project),
						Resolve: resolver.ProjectResolver,
					},
				},
			},
		),
	}
	return &root
}
