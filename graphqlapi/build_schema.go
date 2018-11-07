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

	dbconnector "github.com/creativesoftwarefdn/weaviate/connectors"
	utils "github.com/creativesoftwarefdn/weaviate/graphqlapi/utils"
	"github.com/graphql-go/graphql"
)

var dbConnector dbconnector.DatabaseConnector

// Build the GraphQL schema based on
// 1) the static query structure (e.g. Get)
// 2) the (dynamic) database schema from Weaviate

func (g *GraphQL) buildGraphqlSchema() error {
	rootFieldsObject, err := assembleFullSchema(g)

	if err != nil {
		return fmt.Errorf("could not build GraphQL schema, because: %v", err)
	}

	schemaObject := graphql.ObjectConfig{
		Name:        "WeaviateObj",
		Fields:      rootFieldsObject,
		Description: "Location of the root query",
	}

	// Run grahpql.NewSchema in a sub-closure, so that we can recover from panics.
	// We need to use panics to return errors deep inside the dynamic generation of the GraphQL schema,
	// inside the FieldThunks. There is _no_ way to bubble up an error besides panicking.
	func() {
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				err, ok = r.(error) // can't shadow err here; we need the err from outside the function closure.

				if !ok {
					err = fmt.Errorf("%v", err)
				}
			}
		}()

		g.weaviateGraphQLSchema, err = graphql.NewSchema(graphql.SchemaConfig{
			Query: graphql.NewObject(schemaObject),
		})
	}()

	if err != nil {
		return fmt.Errorf("could not build GraphQL schema, because: %v", err)
	}
	return nil
}

func assembleFullSchema(g *GraphQL) (graphql.Fields, error) {
	filters := utils.FilterContainer{
		WhereOperatorEnum: genOperatorObject(),
	}

	localField, localErr := assembleLocalSchema(g, &filters)
	if localErr != nil {
		return nil, localErr
	}

	networkField, networkErr := assembleNetworkSchema(g, &filters)
	if networkErr != nil {
		return nil, networkErr
	}

	rootFields := graphql.Fields{
		"Local":   localField,
		"Network": networkField,
	}

	return rootFields, nil
}

func assembleLocalSchema(g *GraphQL, filterContainer *utils.FilterContainer) (*graphql.Field, error) {
	// This map is used to store all the Thing and Action Objects, so that we can use them in references.
	getActionsAndThings := make(map[string]*graphql.Object)
	// this map is used to store all the Filter InputObjects, so that we can use them in references.
	filterContainer.LocalFilterOptions = make(map[string]*graphql.InputObject)

	localGetActions, err := genActionClassFieldsFromSchema(g, &getActionsAndThings)

	if err != nil {
		return nil, fmt.Errorf("failed to generate action fields from schema for local Get because: %v", err)
	}

	localGetThings, err := genThingClassFieldsFromSchema(g, &getActionsAndThings)

	if err != nil {
		return nil, fmt.Errorf("failed to generate thing fields from schema for local Get because: %v", err)
	}

	classParentTypeIsAction := true
	localGetMetaActions, err := genMetaClassFieldsFromSchema(g.databaseSchema.ActionSchema.Schema.Classes, classParentTypeIsAction)

	if err != nil {
		return nil, fmt.Errorf("failed to generate action fields from schema for local MetaGet because: %v", err)
	}

	classParentTypeIsAction = false
	localGetMetaThings, err := genMetaClassFieldsFromSchema(g.databaseSchema.ThingSchema.Schema.Classes, classParentTypeIsAction)

	if err != nil {
		return nil, fmt.Errorf("failed to generate thing fields from schema for local MetaGet because: %v", err)
	}

	localGetObject := genThingsAndActionsFieldsForWeaviateLocalGetObj(localGetActions, localGetThings)

	localGetMetaObject := genThingsAndActionsFieldsForWeaviateLocalGetMetaObj(localGetMetaActions, localGetMetaThings)

	localGetAndGetMetaObject := genGetAndGetMetaFields(localGetObject, localGetMetaObject, filterContainer)

	localField := &graphql.Field{
		Type:        localGetAndGetMetaObject,
		Description: "Query a local Weaviate instance",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			result, err := dbConnector.GetGraph(p)
			return result, err
		},
	}

	return localField, nil
}

func assembleNetworkSchema(g *GraphQL, filterContainer *utils.FilterContainer) (*graphql.Field, error) {
	// TODO: placeholder loop, remove this once p2p functionality is up
	weaviateInstances := []string{"WeaviateB", "WeaviateC"}
	weaviateNetworkGetResults := make(map[string]*graphql.Object)
	weaviateNetworkGetMetaResults := make(map[string]*graphql.Object)

	// this map is used to store all the Filter InputObjects, so that we can use them in references.
	filterContainer.NetworkFilterOptions = make(map[string]*graphql.InputObject)

	// TODO implement function that capitalizes all Weaviate names

	for _, weaviate := range weaviateInstances {

		// This map is used to store all the Thing and Action Objects, so that we can use them in references.
		getNetworkActionsAndThings := make(map[string]*graphql.Object)

		networkGetActions, err := genNetworkActionClassFieldsFromSchema(g, &getNetworkActionsAndThings, weaviate)

		if err != nil {
			return nil, fmt.Errorf("failed to generate action fields from schema for network Get because: %v", err)
		}

		networkGetThings, err := genNetworkThingClassFieldsFromSchema(g, &getNetworkActionsAndThings, weaviate)

		if err != nil {
			return nil, fmt.Errorf("failed to generate thing fields from schema for network Get because: %v", err)
		}

		classParentTypeIsAction := true
		networkGetMetaActions, err := genNetworkMetaClassFieldsFromSchema(g.databaseSchema.ActionSchema.Schema.Classes, classParentTypeIsAction, weaviate)

		if err != nil {
			return nil, fmt.Errorf("failed to generate action fields from schema for network MetaGet because: %v", err)
		}

		classParentTypeIsAction = false
		networkGetMetaThings, err := genNetworkMetaClassFieldsFromSchema(g.databaseSchema.ThingSchema.Schema.Classes, classParentTypeIsAction, weaviate)

		if err != nil {
			return nil, fmt.Errorf("failed to generate thing fields from schema for network MetaGet because: %v", err)
		}

		networkGetObject := genThingsAndActionsFieldsForWeaviateNetworkGetObj(networkGetActions, networkGetThings, weaviate)
		networkGetMetaObject := genThingsAndActionsFieldsForWeaviateNetworkGetMetaObj(networkGetMetaActions, networkGetMetaThings, weaviate)
		weaviateNetworkGetResults[weaviate] = networkGetObject
		weaviateNetworkGetMetaResults[weaviate] = networkGetMetaObject

	}
	// TODO this is a temp function, inserts a temp weaviate obj in between Get and Things/Actions
	networkGetObject, networkGetMetaObject := insertDummyNetworkWeaviateField(weaviateNetworkGetResults, weaviateNetworkGetMetaResults)

	genGlobalNetworkFilterElements(filterContainer)

	networkFetchObj := genFieldsObjForNetworkFetch(filterContainer)

	networkIntrospectObj := genFieldsObjForNetworkIntrospect(filterContainer)

	graphQLNetworkFieldContents := utils.GraphQLNetworkFieldContents{
		networkGetObject,
		networkGetMetaObject,
		networkFetchObj,
		networkIntrospectObj,
	}

	networkGetAndGetMetaObject := genNetworkFields(&graphQLNetworkFieldContents, filterContainer)

	networkField := &graphql.Field{
		Type:        networkGetAndGetMetaObject,
		Description: "Query a Weaviate network",
		Args: graphql.FieldConfigArgument{
			"networkTimeout": &graphql.ArgumentConfig{
				Description: "The max time to request an answer from the network. It is the time in seconds",
				Type:        graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			result, err := dbConnector.GetGraph(p)
			return result, err
		},
	}

	return networkField, nil
}
