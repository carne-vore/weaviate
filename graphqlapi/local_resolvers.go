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
	//"fmt"

	"github.com/graphql-go/graphql"
)

const WeaviateRequestType string = "weaviate_request_type"
const LocalGetThing string = "local_get_thing"
const LocalGetAction string = "local_get_action"
const LocalGetMetaThing string = "local_get_meta_thing"
const LocalGetMetaAction string = "local_get_meta_action"

/* unexported consts from dataloader/connector.go
const thingsDataLoader string = "thingsDataLoader"
const keysDataLoader string = "keysDataLoader"
const actionsDataLoader string = "actionsDataLoader"
*/

func resolveLocalGetThing(request graphql.ResolveParams) (interface{}, error) {
	// Init variables used by the dataloader
	//var result interface{}
	//var loader *dataloader.Loader
	//var ok bool

	//ctx := request.Context
	//request.Args[WeaviateRequestType] = LocalGetThing

	// hier moet je een Dataloader ref hebben uit je connector ding
	//result, err := &dataloader.DataLoader.GetLocalGraph(request) // TODO dit werkt nog niet, ik kan de koppeling nog niet maken

	//	// Load the dataloader from the context
	//	if loader, ok = ctx.Value("thingsDataLoader").(*dataloader.Loader); !ok {
	//		return nil, fmt.Errorf("graphql resolver: dataloader not found in context")
	//	}
	//	fmt.Printf(ctx)
	//	result, err := loader.databaseConnector.GetGraph(request) // TODO fix this, where to get the actual db connector?

	return nil, nil //result, err //result, err
}

//func resolveLocalGetAction(request graphql.ResolveParams) (interface{}, error) {
//    p.Args[WeaviateRequestType] = LocalGetAction
//    result, err := db_connector.GetGraph(p) // TODO fix this
//    return (result, err)
//}
//
//func resolveLocalGetMetaThing(request graphql.ResolveParams) (interface{}, error) {
//    p.Args[WeaviateRequestType] = LocalGetMetaThing
//    result, err := db_connector.GetGraph(p) // TODO fix this
//    return (result, err)
//}
//
//func resolveLocalGetMetaAction(request graphql.ResolveParams) (interface{}, error) {
//    p.Args[WeaviateRequestType] = LocalGetMetaAction
//    result, err := db_connector.GetGraph(p) // TODO fix this
//    return (result, err)
//}
