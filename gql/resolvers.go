package gql

import (
	"personal-website-graphql/postgres"

	"github.com/graphql-go/graphql"
)

// Resolver struct holds a connection to our database
type Resolver struct {
	db *postgres.Db
}

// AboutResolver resolves our user query through a db call to GetAbout
func (r *Resolver) AboutResolver(p graphql.ResolveParams) (interface{}, error) {
	about := r.db.GetAbout()
	return about, nil
}

// EducationResolver resolves our user query through a db call to GetEducation
func (r *Resolver) EducationResolver(p graphql.ResolveParams) (interface{}, error) {
	education := r.db.GetEducation()
	return education, nil
}

// ExperienceResolver resolves our user query through a db call to GetExperience
func (r *Resolver) ExperienceResolver(p graphql.ResolveParams) (interface{}, error) {
	experience := r.db.GetExperience()
	return experience, nil
}

// ProjectResolver resolves our user query through a db call to GetProjects
func (r *Resolver) ProjectResolver(p graphql.ResolveParams) (interface{}, error) {
	projects := r.db.GetAllProjects()
	return projects, nil
}
