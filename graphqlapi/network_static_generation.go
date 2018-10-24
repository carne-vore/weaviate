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
	"fmt"
	"github.com/graphql-go/graphql"
)

// temporary function that does nothing but display a Weaviate instance // TODO: delete this once p2p functionality is up
func insertDummyNetworkWeaviateField(weaviatesWithGetFields map[string]*graphql.Object, weaviatesWithMetaGetFields map[string]*graphql.Object) (*graphql.Object, *graphql.Object) {

	getWeaviates := graphql.Fields{}
	metaGetWeaviates := graphql.Fields{}

	for weaviate, weaviateFields := range weaviatesWithGetFields {
		getWeaviates["weaviateB"] = &graphql.Field{ // TODO first char to lower case, hardcoded for now as this is a placeholder func anyway
			Name:        weaviate,
			Description: "Object field for weaviate weaviateB in the network.",
			Type:        weaviateFields,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		}
		metaGetWeaviates["weaviateB"] = &graphql.Field{ // TODO first char to lower case, hardcoded for now as this is a placeholder func anyway
			Name:        fmt.Sprintf("%s%s", "Meta", weaviate),
			Description: "Object field for weaviate weaviateB in the network.",
			Type:        weaviatesWithMetaGetFields[weaviate],
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		}
	}

	dummyWeaviateGetObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkGetObject",
		Fields:      getWeaviates,
		Description: "Type of Get function to get Things or Actions from the Network",
	}
	dummyWeaviateGetMetaObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkGetMetaObject",
		Fields:      metaGetWeaviates,
		Description: "Type of Get function to get Things or Actions from the Network",
	}

	return graphql.NewObject(dummyWeaviateGetObject), graphql.NewObject(dummyWeaviateGetMetaObject)
}

// generate the static parts of the schema for network queries
func genThingsAndActionsFieldsForWeaviateNetworkGetObj(networkGetActions *graphql.Object, networkGetThings *graphql.Object, weaviate string) *graphql.Object {
	getThingsAndActionFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGet", weaviate, "Actions"),
			Description: "Get Actions on the Local Weaviate",
			Type:        networkGetActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGet", weaviate, "Things"),
			Description: "Get Things on the Local Weaviate",
			Type:        networkGetThings,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getNetworkThingsAndActionFieldsObject := graphql.ObjectConfig{
		Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGet", weaviate, "Obj"), // TODO: "WeaviateNetworkGetObj",
		Fields:      getThingsAndActionFields,
		Description: "Type of Get function to get Things or Actions on the Local Weaviate",
	}

	return graphql.NewObject(getNetworkThingsAndActionFieldsObject)
}

func genThingsAndActionsFieldsForWeaviateNetworkGetMetaObj(networkGetMetaActions *graphql.Object, networkGetMetaThings *graphql.Object, weaviate string) *graphql.Object {
	getNetworkMetaThingsAndActionFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGetMeta", weaviate, "Actions"),
			Description: "Get Meta information about Actions on the Local Weaviate",
			Type:        networkGetMetaActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGetMeta", weaviate, "Things"),
			Description: "Get Meta information about Things on the Local Weaviate",
			Type:        networkGetMetaThings,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getNetworkMetaThingsAndActionFieldsObject := graphql.ObjectConfig{
		Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGetMeta", weaviate, "Obj"), // TODO: fix this one p2p functionality is up
		Fields:      getNetworkMetaThingsAndActionFields,
		Description: "Type of Get function to get meta information about Things or Actions on the Local Weaviate",
	}

	return graphql.NewObject(getNetworkMetaThingsAndActionFieldsObject)
}

func genNetworkGetAndGetMetaFields(networkGetObject *graphql.Object, networkGetMetaObject *graphql.Object, networkFilterOptions map[string]*graphql.InputObject) *graphql.Object {
	filterFields := genNetworkFilterFields(networkFilterOptions)
	networkGetAndGetMetaFields := graphql.Fields{

		"Get": &graphql.Field{
			Name:        "WeaviateNetworkGet",
			Type:        networkGetObject,
			Description: "Get Things or Actions on the local weaviate",
			Args: graphql.FieldConfigArgument{
				"where": &graphql.ArgumentConfig{
					Description: "Filter options for the Get search, to convert the data to the filter input",
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:        "WeaviateNetworkGetWhereInpObj",
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
			Name:        "WeaviateNetworkGetMeta",
			Type:        networkGetMetaObject,
			Description: "Query to Get Meta information about the data in the local Weaviate instance",
			Args: graphql.FieldConfigArgument{
				"where": &graphql.ArgumentConfig{
					Description: "Filter options for the GetMeta search, to convert the data to the filter input",
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:        "WeaviateNetworkGetMetaWhereInpObj",
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

	weaviateNetworkObject := &graphql.ObjectConfig{
		Name:        "WeaviateNetworkObj",
		Fields:      networkGetAndGetMetaFields,
		Description: "Type of query on the local Weaviate",
	}

	return graphql.NewObject(*weaviateNetworkObject)
}
