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
		getWeaviates[weaviate] = &graphql.Field{ // TODO first char to lower case
			Name:        weaviate,
			Description: fmt.Sprintf("%s%s%s", "Object field for weaviate ", weaviate, " in the network."),
			Type:        weaviateFields,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		}
		metaGetWeaviates[weaviate] = &graphql.Field{ // TODO first char to lower case
			Name:        fmt.Sprintf("%s%s", "Meta", weaviate),
			Description: fmt.Sprintf("%s%s%s", "Object field for weaviate ", weaviate, " in the network."),
			Type:        weaviatesWithMetaGetFields[weaviate],
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		}
	}

	dummyWeaviateGetObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkGetObj",
		Fields:      getWeaviates,
		Description: "Type of Get function to get Things or Actions from the Network",
	}
	dummyWeaviateGetMetaObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkGetMetaObj",
		Fields:      metaGetWeaviates,
		Description: "Type of Get function to get meta information about Things or Actions on a Weaviate in the Network",
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
		Description: fmt.Sprintf("%s%s%s", "Objects for the what to Get from the weaviate ", weaviate, " in the network."), // TODO: edit this string. Possibly make weaviate reference dynamic?
	}
	return graphql.NewObject(getNetworkThingsAndActionFieldsObject)
}

func genThingsAndActionsFieldsForWeaviateNetworkGetMetaObj(networkGetMetaActions *graphql.Object, networkGetMetaThings *graphql.Object, weaviate string) *graphql.Object {
	getNetworkMetaThingsAndActionFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGetMeta", weaviate, "Actions"),
			Description: "Get Meta information about Actions on a Weaviate in the Network",
			Type:        networkGetMetaActions,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        fmt.Sprintf("%s%s%s", "WeaviateNetworkGetMeta", weaviate, "Things"),
			Description: "Get Meta information about Things on a Weaviate in a Network",
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
		Description: fmt.Sprintf("%s%s%s", "Objects for the what to Get Meta from the weaviate ", weaviate, " in the network."), // TODO
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
			Description: "Beacon as an Action result found in the Network",
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
			Description: "Beacon as a Thing result found in the Network",
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
			Description: "The certainty the node has a value matching the value searched on in the fuzzy network fetch",
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
				Name:        "WeaviateNetworkFetchWhereInpObj",
				Fields:      genNetworkFetchThingsAndActionsFilterFields(filterContainer),
				Description: "", //TODO no desc in prototype
			},
		)),
	}

	return whereFilterFields
}

func genFieldsObjForNetworkIntrospect(filterContainer *utils.FilterContainer) *graphql.Object {
	networkIntrospectActionsFields := genNetworkIntrospectActionsFieldsObj(filterContainer)
	networkIntrospectThingsFields := genNetworkIntrospectThingsFieldsObj(filterContainer)
	networkIntrospectBeaconFields := genNetworkIntrospectBeaconFieldsObj()
	networkIntrospectWhereFilterFields := genNetworkIntrospectThingsActionsWhereFilterFields(filterContainer)

	networkIntrospectFields := graphql.Fields{

		"Actions": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectActions",
			Description: "Actions to introspect in the network",
			Type:        graphql.NewList(networkIntrospectActionsFields),
			Args: graphql.FieldConfigArgument{
				"where": networkIntrospectWhereFilterFields,
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Things": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectThings",
			Description: "Things to introspect in the network",
			Type:        graphql.NewList(networkIntrospectThingsFields),
			Args: graphql.FieldConfigArgument{
				"where": networkIntrospectWhereFilterFields,
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Beacon": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectBeacon",
			Description: "Beacon to introspect in the network",
			Type:        networkIntrospectBeaconFields,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Description: "",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	networkIntrospectFieldsObj := graphql.ObjectConfig{
		Name:        "WeaviateNetworkIntrospectObj",
		Fields:      networkIntrospectFields,
		Description: "type of object to introspect in the network",
	}

	return graphql.NewObject(networkIntrospectFieldsObj)
}

func genNetworkIntrospectActionsFieldsObj(filterContainer *utils.FilterContainer) *graphql.Object {
	getNetworkIntrospectActionsFields := graphql.Fields{

		"weaviate": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectActionsWeaviate",
			Description: "Weaviate node the found class in the Network Introspection search is in.",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"className": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectActionsClassName",
			Description: "To filter on which class name",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"certainty": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectActionsCertainty",
			Description: "To filter with which certainty from 0-1",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"properties": filterContainer.WeaviateNetworkIntrospectPropertiesObjField,
	}

	getNetworkIntrospectActionsFieldsObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkIntrospectActionsObj",
		Fields:      getNetworkIntrospectActionsFields,
		Description: "Object for which Actions to introspect in the network",
	}

	return graphql.NewObject(getNetworkIntrospectActionsFieldsObject)
}

func genNetworkIntrospectThingsFieldsObj(filterContainer *utils.FilterContainer) *graphql.Object {
	getNetworkIntrospectThingsFields := graphql.Fields{

		"weaviate": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectThingsWeaviate",
			Description: "Weaviate node the found class in the Network Introspection search is in.",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"className": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectThingsClassName",
			Description: "To filter on which class name",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"certainty": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectThingsCertainty",
			Description: "To filter with which certainty from 0-1",
			Type:        graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"properties": filterContainer.WeaviateNetworkIntrospectPropertiesObjField,
	}

	getNetworkIntrospectThingsFieldsObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkIntrospectThingsObj",
		Fields:      getNetworkIntrospectThingsFields,
		Description: "Object for which Things to introspect in the network",
	}

	return graphql.NewObject(getNetworkIntrospectThingsFieldsObject)
}

func genNetworkIntrospectBeaconFieldsObj() *graphql.Object {
	beaconPropertiesObj := genWeaviateNetworkIntrospectBeaconPropertiesObj()

	introspectBeaconFields := graphql.Fields{

		"weaviate": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectBeaconWeaviate",
			Description: "Weaviate node the found class in the Network Introspection search is in.",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"className": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectBeaconclassName",
			Description: "To filter on which class name",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"properties": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectBeaconPropertiesObj",
			Description: "Which properties to filter on",
			Type:        graphql.NewList(beaconPropertiesObj),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	getNetworkFetchFuzzyFieldsObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkIntrospectBeaconObj",
		Fields:      introspectBeaconFields,
		Description: "Object for which beacon to introspect in the network",
	}

	return graphql.NewObject(getNetworkFetchFuzzyFieldsObject)
}

func genWeaviateNetworkIntrospectBeaconPropertiesObj() *graphql.Object {
	beaconPropertiesFields := graphql.Fields{

		"propertyName": &graphql.Field{
			Name:        "WeaviateNetworkIntrospectBeaconPropertiesObjPropertyName",
			Description: "Which property name to filter properties on",
			Type:        graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},
	}

	beaconPropertiesObject := graphql.ObjectConfig{
		Name:        "WeaviateNetworkIntrospectBeaconPropertiesObj",
		Fields:      beaconPropertiesFields,
		Description: "Which properties to filter on",
	}

	return graphql.NewObject(beaconPropertiesObject)
}

func genNetworkIntrospectThingsActionsWhereFilterFields(filterContainer *utils.FilterContainer) *graphql.ArgumentConfig {
	whereFilterFields := &graphql.ArgumentConfig{
		Description: "", // TODO no desc
		Type: graphql.NewNonNull(graphql.NewList(graphql.NewInputObject(
			graphql.InputObjectConfig{
				Name:        "WeaviateNetworkIntrospectWhereInpObj",
				Fields:      genNetworkIntrospectThingsAndActionsFilterFields(filterContainer),
				Description: "", // TODO
			},
		))),
	}

	return whereFilterFields
}

func genNetworkFields(graphQLNetworkFieldContents *utils.GraphQLNetworkFieldContents, filterContainer *utils.FilterContainer) *graphql.Object {
	getGetMetaFilterFields := genNetworkFilterFields(filterContainer)
	networkGetAndGetMetaFields := graphql.Fields{

		"Get": &graphql.Field{
			Name:        "WeaviateNetworkGet",
			Type:        graphQLNetworkFieldContents.NetworkGetObject,
			Description: "Get Things or Actions on the network of weaviate",
			Args: graphql.FieldConfigArgument{
				"where": &graphql.ArgumentConfig{
					Description: "Filter options for the Network Get search, to convert the data to the filter input",
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
			Description: "Query to Get Meta information about the data in the Network Weaviate",
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
			Description: "Do a fuzzy search fetch to search Things or Actions on the network weaviate",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				result, err := dbConnector.GetGraph(p)
				return result, err
			},
		},

		"Introspect": &graphql.Field{
			Name:        "WeaviateNetworkIntrospection",
			Type:        graphQLNetworkFieldContents.NetworkIntrospectObject,
			Description: "To fetch meta information about the ontology of Things or Actions on the network weaviate",
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
