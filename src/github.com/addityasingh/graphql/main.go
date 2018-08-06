package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

var blogpostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Blogpost",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				m := p.Source.(Blogpost)
				return m.id, nil
			},
		},
		"title": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				m := p.Source.(Blogpost)
				return m.title, nil
			},
		},
		"body": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				m := p.Source.(Blogpost)
				return m.body, nil
			},
		},
		"created": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				m := p.Source.(Blogpost)
				return m.created, nil
			},
		},
		"author": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				m := p.Source.(Blogpost)
				return m.author, nil
			},
		},
	},
})

type Blogpost struct {
	id      string
	title   string
	body    string
	created string
	author  string
}

func main() {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
		"blogpostsBetween": &graphql.Field{
			Type: graphql.NewList(blogpostType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// ... some more code ...
				// what I return here was placed to demonstrate the error.
				return []Blogpost{
					Blogpost{
						id:      "new-post",
						title:   "new post",
						body:    "Lorem ipsum",
						author:  "John Doe",
						created: "12.12.2012",
					},
				}, nil
			},
			Args: graphql.FieldConfigArgument{
				"from": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"to": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
				{
					blogpostsBetween {
						id,
						title,
						body,
						author
					}
				}
			`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	fmt.Printf("%s \n", rJSON)
}
