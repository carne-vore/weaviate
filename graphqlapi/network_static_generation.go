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
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/utils"
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
			Description: "Get Actions from the Network",
			Type:        networkGetActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGet", weaviate, "Things"),
			Description: "Get Things from the Network",
			Type:        networkGetThings,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getNetworkThingsAndActionFieldsObject := graphql.ObjectConfig{
		Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGet", weaviate, "Obj"),
		Fields:      getThingsAndActionFields,
		Description: "Objects for the what to Get from the weaviate weaviateB in the network.", // TODO: edit this string. Possibly make weaviate reference dynamic?
	}
	return graphql.NewObject(getNetworkThingsAndActionFieldsObject)
}

func genThingsAndActionsFieldsForWeaviateNetworkGetMetaObj(networkGetMetaActions *graphql.Object, networkGetMetaThings *graphql.Object, weaviate string) *graphql.Object {
	getNetworkMetaThingsAndActionFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGetMeta", weaviate, "Actions"),
			Description: "Get Meta information about Actions from the Network",
			Type:        networkGetMetaActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGetMeta", weaviate, "Things"),
			Description: "Get Meta information about Things from the Network",
			Type:        networkGetMetaThings,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getNetworkMetaThingsAndActionFieldsObject := graphql.ObjectConfig{
		Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGetMeta", weaviate, "Obj"),
		Fields:      getNetworkMetaThingsAndActionFields,
		Description: "Type of Get function to get meta information about Things or Actions from the Network",
	}

	return graphql.NewObject(getNetworkMetaThingsAndActionFieldsObject)
}

func genFieldsObjForNetworkFetch(filterContainer *utils.FilterContainer) *graphql.Object {
	networkFetchActionsFields := genNetworkFetchActionsFieldsObj()
	networkFetchThingsFields := genNetworkFetchThingsFieldsObj()
	networkFetchFuzzyFields := genNetworkFetchFuzzyFieldsObj()
	networkFetchWhereFilterFields := genNetworkFetchThingsActionsWhereFilterFields(filterContainer)

	networkFetchFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        "WeaviateNetworkFetchActions",
			Description: "Actions to fetch on the network",
			Type:        graphql.NewList(networkFetchActionsFields),
			Args: graphql.FieldConfigArgument{
				"where": networkFetchWhereFilterFields,
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        "WeaviateNetworkFetchThings",
			Description: "Things to fetch on the network",
			Type:        graphql.NewList(networkFetchThingsFields),
			Args: graphql.FieldConfigArgument{
				"where": networkFetchWhereFilterFields,
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Fuzzy": &graphql.Field{
			Name:        "WeaviateNetworkFetchFuzzy",
			Description: "To do a fuzzy fetch, with only a ontology value, on the Network",
			Type:        graphql.NewList(networkFetchFuzzyFields),
			Args: graphql.FieldConfigArgument{
				"value": &graphql.ArgumentConfig{
					Description: "The ontology value to fetch Things or Actions on the Network on",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"certainty": &graphql.ArgumentConfig{
					Description: "The minimum certainty the nodes should match the given ontology value",
					Type:        graphql.NewNonNull(graphql.Float),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	networkFetchFieldsObj := graphql.ObjectConfig{
		Name:        "WeaviateNetworkFetchObj",
		Fields:      networkFetchFields,
		Description: "Type of network fetch: e.g. Things, Actions",
	}

	return graphql.NewObject(networkFetchFieldsObj)
}

func genNetworkFetchActionsFieldsObj() *graphql.Object {
	getNetworkFetchActionsFields := graphql.Fields{

		"beacon": &graphql.Field{
			Name:        "WeaviateNetworkFetchActionsBeacon",
			Description: "The beacon of the node that is a result from the fuzzy network fetch",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"certainty": &graphql.Field{
			Name:        "WeaviateNetworkFetchActionsCertainty",
			Description: "Certainty of beacon result found in the Network has expected ontology characterisics",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getNetworkFetchActionsFieldsObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkFetchActionsObj",
		Fields:      getNetworkFetchActionsFields,
		Description: "Type of Actions i.e. classes to fetch on the network",
	}

	return graphql.NewObject(getNetworkFetchActionsFieldsObject)
}

func genNetworkFetchThingsFieldsObj() *graphql.Object {
	getNetworkFetchThingsFields := graphql.Fields{

		"beacon": &graphql.Field{
			Name:        "WeaviateNetworkFetchThingsBeacon",
			Description: "The beacon of the node that is a result from the fuzzy network fetch",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"certainty": &graphql.Field{
			Name:        "WeaviateNetworkFetchThingsCertainty",
			Description: "Certainty of beacon result found in the Network has expected ontology characterisics", // TODO typo in original string
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getNetworkFetchThingsFieldsObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkFetchThingsObj",
		Fields:      getNetworkFetchThingsFields,
		Description: "Type of Things i.e. classes to fetch on the network",
	}

	return graphql.NewObject(getNetworkFetchThingsFieldsObject)
}

func genNetworkFetchFuzzyFieldsObj() *graphql.Object {
	getNetworkFetchFuzzyFields := graphql.Fields{

		"beacon": &graphql.Field{
			Name:        "WeaviateNetworkFetchFuzzyBeacon",
			Description: "The beacon of the node that is a result from the fuzzy network fetch",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"certainty": &graphql.Field{
			Name:        "WeaviateNetworkFetchFuzzyCertainty",
			Description: "Certainty of beacon result found in the Network has expected ontology characterisics", // TODO typo in original string
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getNetworkFetchFuzzyFieldsObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkFetchFuzzyObj",
		Fields:      getNetworkFetchFuzzyFields,
		Description: "The objects what to request from this network fuzzy fetch query",
	}

	return graphql.NewObject(getNetworkFetchFuzzyFieldsObject)
}

func genNetworkFetchThingsActionsWhereFilterFields(filterContainer *utils.FilterContainer) *graphql.ArgumentConfig {
	whereFilterFields := &graphql.ArgumentConfig{
		Description: "Filter options for the Network GetMeta search, to convert the data to the filter input",
		Type: graphql.NewNonNull(graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name:        "NetworkFetchWhereInpObj",
				Fields:      genNetworkFetchThingsAndActionsFilterFields(filterContainer),
				Description: "Filter options for the Network GetMeta search, to convert the data to the filter input",
			},
		)),
	}

	return whereFilterFields
}

func genNetworkFields(graphQLNetworkFieldContents *utils.GraphQLNetworkFieldContents, filterContainer *utils.FilterContainer) *graphql.Object {
	getGetMetaFilterFields := genNetworkFilterFields(filterContainer)
	networkGetAndGetMetaFields := graphql.Fields{

		"Get": &graphql.Field{
			Name:        "WeaviateNetworkGet",
			Type:        graphQLNetworkFieldContents.NetworkGetObject,
			Description: "Get Things or Actions on the local weaviate",
			Args: graphql.FieldConfigArgument{
				"where": &graphql.ArgumentConfig{
					Description: "Filter options for the Get search, to convert the data to the filter input",
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:        "WeaviateNetworkGetWhereInpObj",
							Fields:      getGetMetaFilterFields,
							Description: "Filter options for the Network Get search, to convert the data to the filter input",
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
			Type:        graphQLNetworkFieldContents.NetworkGetMetaObject,
			Description: "Query to Get Meta information about data on the Network",
			Args: graphql.FieldConfigArgument{
				"where": &graphql.ArgumentConfig{
					Description: "Filter options for the Network GetMeta search, to convert the data to the filter input",
					Type: graphql.NewInputObject(
						graphql.InputObjectConfig{
							Name:        "WeaviateNetworkGetMetaWhereInpObj",
							Fields:      getGetMetaFilterFields,
							Description: "Filter options for the Network GetMeta search, to convert the data to the filter input",
						},
					),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Fetch": &graphql.Field{
			Name:        "WeaviateNetworkFetch",
			Type:        graphQLNetworkFieldContents.NetworkFetchObject,
			Description: "Query to Get Meta information about data on the Network",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Introspection": &graphql.Field{
			Name:        "WeaviateNetworkIntrospection",
			Type:        graphql.String,
			Description: "Query to Get Meta information about data on the Network",
			//			Args: graphql.FieldConfigArgument{
			//				"where": &graphql.ArgumentConfig{
			//					Description: "Filter options for the Network GetMeta search, to convert the data to the filter input",
			//					Type: graphql.NewInputObject(
			//						graphql.InputObjectConfig{
			//							Name:        "WeaviateNetworkGetMetaWhereInpObj",
			//							Fields:      filterFields,
			//							Description: "Filter options for the Network GetMeta search, to convert the data to the filter input",
			//						},
			//					),
			//				},
			//			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	weaviateNetworkObject := &graphql.ObjectConfig{
		Name:        "WeaviateNetworkObj",
		Fields:      networkGetAndGetMetaFields,
		Description: "Type of query on the Weaviate network",
	}

	return graphql.NewObject(*weaviateNetworkObject)
}
