package gql

import "github.com/graphql-go/graphql"

// About describes a graphql object containing About page information
var About = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "About",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "identifier",
			},
			"header": &graphql.Field{
				Type:        graphql.String,
				Description: "header of about page",
			},
			"paragraph": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "paragraph of about page",
			},
		},
	},
)

// Experience describes a graphql object containing Experience page information
var Experience = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Experience",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "identifier",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of position",
			},
			"organization": &graphql.Field{
				Type:        graphql.String,
				Description: "name of organization",
			},
			"location": &graphql.Field{
				Type:        graphql.String,
				Description: "location of organization",
			},
			"dateRange": &graphql.Field{
				Type:        graphql.String,
				Description: "date range of position",
			},
			"summary": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "summary of experience",
			},
		},
	},
)

// Project describes a graphql object containing Project information
var Project = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Project",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "identifier",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of project",
			},
			"subtitle": &graphql.Field{
				Type:        graphql.String,
				Description: "subtitle of project",
			},
			"workInProgress": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "flag if the project is a work in progress",
			},
			"date": &graphql.Field{
				Type:        graphql.String,
				Description: "date of project",
			},
			"summary": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "summary of project",
			},
			"moreInfo": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "more info about project",
			},
			"source": &graphql.Field{
				Type:        graphql.String,
				Description: "source code for project",
			},
			"videos": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "videos for project",
			},
			"images": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "images for project",
			},
		},
	},
)

// Education describes a graphql object containing Education page information
var Education = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Education",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "identifier",
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "title of program",
			},
			"institution": &graphql.Field{
				Type:        graphql.String,
				Description: "name of institution",
			},
			"location": &graphql.Field{
				Type:        graphql.String,
				Description: "location of institution",
			},
			"date": &graphql.Field{
				Type:        graphql.String,
				Description: "date attended",
			},
			"relevantProjects": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "relevant projects completed",
			},
			"relevantCourses": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "relevant courses completed",
			},
		},
	},
)
