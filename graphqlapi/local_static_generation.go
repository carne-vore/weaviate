/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2018 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * AUTHOR: Bob van Luijt (bob@kub.design)
 * See www.creativesoftwarefdn.org for details
 * Contact: @CreativeSofwFdn / bob@kub.design
 */

// Package graphqlapi provides the graphql endpoint for Weaviate
package graphqlapi

import (
	"github.com/graphql-go/graphql"
)

// generate the static parts of the schema
func genThingsAndActionsFieldsForWeaviateLocalGetObj(localGetActions *graphql.Object, localGetThings *graphql.Object) *graphql.Object {
	getThingsAndActionFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        "WeaviateLocalGetActions",
			Description: "Get Actions on the Local Weaviate",
			Type:        localGetActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        "WeaviateLocalGetThings",
			Description: "Get Things on the Local Weaviate",
			Type:        localGetThings,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getThingsAndActionFieldsObject := graphql.ObjectConfig{
		Name:        "WeaviateLocalGetObj",
		Fields:      getThingsAndActionFields,
		Description: "Type of Get function to get Things or Actions on the Local Weaviate",
	}

	return graphql.NewObject(getThingsAndActionFieldsObject)
}

func genThingsAndActionsFieldsForWeaviateLocalGetMetaObj(localGetMetaActions *graphql.Object, localGetMetaThings *graphql.Object) *graphql.Object {
	getMetaThingsAndActionFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        "WeaviateLocalGetMetaActions",
			Description: "Get Meta information about Actions on the Local Weaviate",
			Type:        localGetMetaActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        "WeaviateLocalGetMetaThings",
			Description: "Get Meta information about Things on the Local Weaviate",
			Type:        localGetMetaThings,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getMetaThingsAndActionFieldsObject := graphql.ObjectConfig{
		Name:        "WeaviateLocalGetMetaObj",
		Fields:      getMetaThingsAndActionFields,
		Description: "Type of Get function to get meta information about Things or Actions on the Local Weaviate",
	}

	return graphql.NewObject(getMetaThingsAndActionFieldsObject)
}

func genGetAndGetMetaFields(localGetObject *graphql.Object, localGetMetaObject *graphql.Object, filterOptions map[string]*graphql.InputObject) *graphql.Object {
	filterFields := genFilterFields(filterOptions)
	getAndGetMetaFields := graphql.Fields{

		"Get": &graphql.Field{
			Name:        "WeaviateLocalGet",
			Type:        localGetObject,
			Description: "Get Things or Actions on the local weaviate",
			Args: graphql.FieldConfigArgument{
				"where": &graphql.ArgumentConfig{
					Description: "Filter options for the Get search, to convert the data to the filter input",
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:        "WeaviateLocalGetWhereInpObj",
							Fields:      filterFields,
							Description: "Filter options for the Get search, to convert the data to the filter input",
						},
					),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"GetMeta": &graphql.Field{
			Name:        "WeaviateLocalGetMeta",
			Type:        localGetMetaObject,
			Description: "Query to Get Meta information about the data in the local Weaviate instance",
			Args: graphql.FieldConfigArgument{
				"where": &graphql.ArgumentConfig{
					Description: "Filter options for the GetMeta search, to convert the data to the filter input",
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:        "WeaviateLocalGetMetaWhereInpObj",
							Fields:      filterFields,
							Description: "Filter options for the GetMeta search, to convert the data to the filter input",
						},
					),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	weaviateLocalObject := &graphql.ObjectConfig{
		Name:        "WeaviateLocalObj",
		Fields:      getAndGetMetaFields,
		Description: "Type of query on the local Weaviate",
	}

	return graphql.NewObject(*weaviateLocalObject)
}
