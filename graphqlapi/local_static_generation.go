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
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/descriptions"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/utils"
	"github.com/graphql-go/graphql"
)

// generate the static parts of the schema
func genThingsAndActionsFieldsForWeaviateLocalGetObj(localGetActions *graphql.Object, localGetThings *graphql.Object) *graphql.Object {
	getThingsAndActionFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        "WeaviateLocalGetActions",
			Description: descriptions.LocalGetActionsDesc,
			Type:        localGetActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        "WeaviateLocalGetThings",
			Description: descriptions.LocalGetThingsDesc,
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
		Description: descriptions.LocalGetObjDesc,
	}

	return graphql.NewObject(getThingsAndActionFieldsObject)
}

func genThingsAndActionsFieldsForWeaviateLocalGetMetaObj(localGetMetaActions *graphql.Object, localGetMetaThings *graphql.Object) *graphql.Object {
	getMetaThingsAndActionFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        "WeaviateLocalGetMetaActions",
			Description: descriptions.LocalGetMetaActionsDesc,
			Type:        localGetMetaActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        "WeaviateLocalGetMetaThings",
			Description: descriptions.LocalGetMetaThingsDesc,
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
		Description: descriptions.LocalGetMetaObjDesc,
	}

	return graphql.NewObject(getMetaThingsAndActionFieldsObject)
}

func genGetAndGetMetaFields(localGetObject *graphql.Object, localGetMetaObject *graphql.Object, filterContainer *utils.FilterContainer) *graphql.Object {
	filterFields := genFilterFields(filterContainer)
	getAndGetMetaFields := graphql.Fields{

		"Get": &graphql.Field{
			Name:        "WeaviateLocalGet",
			Type:        localGetObject,
			Description: descriptions.LocalGetDesc,
			Args: graphql.FieldConfigArgument{
				"where": &graphql.ArgumentConfig{
					Description: descriptions.LocalGetWhereDesc,
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:        "WeaviateLocalGetWhereInpObj",
							Fields:      filterFields,
							Description: descriptions.LocalGetWhereInpObjDesc,
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
			Description: descriptions.LocalGetMetaDesc,
			Args: graphql.FieldConfigArgument{
				"where": &graphql.ArgumentConfig{
					Description: descriptions.LocalGetMetaWhereDesc,
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:        "WeaviateLocalGetMetaWhereInpObj",
							Fields:      filterFields,
							Description: descriptions.LocalGetMetaWhereInpObjDesc,
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
		Description: descriptions.LocalObjDesc,
	}

	return graphql.NewObject(*weaviateLocalObject)
}
