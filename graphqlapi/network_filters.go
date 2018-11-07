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
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/utils"
	"github.com/graphql-go/graphql"
)

func genNetworkFilterFields(filterContainer *utils.FilterContainer) graphql.InputObjectConfigFieldMap {
	staticFilterElements := genNetworkStaticWhereFilterElements(filterContainer)

	filterFields := graphql.InputObjectConfigFieldMap{
		"operands": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(filterContainer.Operands),
			Description: "Operands in the 'where' filter field, is a list of objects",
		},
	}

	for key, value := range staticFilterElements {
		filterFields[key] = value
	}

	return filterFields
}

// generate these elements once
func genNetworkStaticWhereFilterElements(filterContainer *utils.FilterContainer) graphql.InputObjectConfigFieldMap {
	staticFilterElements := graphql.InputObjectConfigFieldMap{
		"operator": &graphql.InputObjectFieldConfig{
			Type:        filterContainer.WhereOperatorEnum,
			Description: "Operator in the 'where' filter field, value is one of the 'WhereOperatorEnum' object",
		},
		"path": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(graphql.String),
			Description: "Path of from 'Things' or 'Actions' to the property name through the classes",
		},
		"valueInt": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: "Integer value that the property at the provided path will be compared to by an operator",
		},
		"valueNumber": &graphql.InputObjectFieldConfig{
			Type:        graphql.Float,
			Description: "Number value that the property at the provided path will be compared to by an operator",
		},
		"valueBoolean": &graphql.InputObjectFieldConfig{
			Type:        graphql.Boolean,
			Description: "Boolean value that the property at the provided path will be compared to by an operator",
		},
		"valueString": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "String value that the property at the provided path will be compared to by an operator",
		},
	}

	return staticFilterElements
}

func genNetworkOperandsObjectFields(filterContainer *utils.FilterContainer, staticFilterElements graphql.InputObjectConfigFieldMap) graphql.InputObjectConfigFieldMap {
	outputFieldConfigMap := staticFilterElements

	outputFieldConfigMap["operands"] = &graphql.InputObjectFieldConfig{
		Type:        graphql.NewList(filterContainer.Operands),
		Description: "Operands in the 'where' filter field, is a list of objects",
	}

	return outputFieldConfigMap
}

// This is an exact translation of the Prototype from JS to Go. In the prototype some filter elements are declared as global variables, this is recreated here.
func genGlobalNetworkFilterElements(filterContainer *utils.FilterContainer) {
	filterContainer.WeaviateNetworkWhereNameKeywordsInpObj = genWeaviateNetworkWhereNameKeywordsInpObj()
	filterContainer.WeaviateNetworkIntrospectPropertiesObjField = genWeaviateNetworkIntrospectPropertiesObjField()
}

func genWeaviateNetworkWhereNameKeywordsInpObj() *graphql.InputObject {
	weaviateNetworkWhereNameKeywordsInpObj := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "WeaviateNetworkWhereNameKeywordsInpObj",
			Fields: graphql.InputObjectConfigFieldMap{
				"value": &graphql.InputObjectFieldConfig{
					Type:        graphql.String,
					Description: "", // TODO this has no description in the prototype
				},
				"weight": &graphql.InputObjectFieldConfig{
					Type:        graphql.Float,
					Description: "", // TODO this has no description in the prototype
				},
			},
			Description: "", // TODO this has no description in the prototype
		},
	)
	return weaviateNetworkWhereNameKeywordsInpObj
}

func genWeaviateNetworkIntrospectPropertiesObjField() *graphql.Field {
	weaviateNetworkIntrospectPropertiesObject := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "WeaviateNetworkIntrospectPropertiesObj",
			Fields: graphql.Fields{
				"propertyName": &graphql.Field{
					Type:        graphql.String,
					Description: "Which property name to filter properties on",
				},
				"certainty": &graphql.Field{
					Type:        graphql.Float,
					Description: "With which certainty 0-1 to filter properties on",
				},
			},
			Description: "which properties to filter on",
		},
	)

	weaviateNetworkIntrospectPropertiesObjField := &graphql.Field{
		Name:        "WeaviateNetworkIntrospectPropertiesObj",
		Description: "Which properties to filter on",
		Type:        graphql.NewList(weaviateNetworkIntrospectPropertiesObject),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			result, err := dbConnector.GetGraph(p)
			return result, err
		},
	}

	return weaviateNetworkIntrospectPropertiesObjField

}

func genNetworkFetchThingsAndActionsFilterFields(filterContainer *utils.FilterContainer) graphql.InputObjectConfigFieldMap {
	networkFetchWhereInpObjPropertiesObj := genNetworkFetchWhereInpObjPropertiesObj(filterContainer)
	networkFetchWhereInpObjClassInpObj := genNetworkFetchWhereInpObjClassInpObj(filterContainer)

	networkFetchThingsAndActionsFilterFields := graphql.InputObjectConfigFieldMap{
		"class": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(networkFetchWhereInpObjClassInpObj),
			Description: "", // TODO no description in prototype
		},
		"properties": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(networkFetchWhereInpObjPropertiesObj),
			Description: "", // TODO no description in prototype
		},
		"first": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: "", // TODO no description in prototype
		},
	}

	return networkFetchThingsAndActionsFilterFields
}

func genNetworkFetchWhereInpObjPropertiesObj(filterContainer *utils.FilterContainer) *graphql.InputObject {
	filterPropertiesElements := genNetworkStaticWhereFilterElements(filterContainer)
	// delete "path" key/value set
	delete(filterPropertiesElements, "path")

	filterPropertiesElements["certainty"] = &graphql.InputObjectFieldConfig{
		Type:        graphql.Float,
		Description: "", // TODO this has no description in the prototype
	}
	filterPropertiesElements["name"] = &graphql.InputObjectFieldConfig{
		Type:        graphql.String,
		Description: "", // TODO this has no description in the prototype
	}
	filterPropertiesElements["keywords"] = &graphql.InputObjectFieldConfig{
		Type:        graphql.NewList(genNetworkFetchWherePropertyWhereKeywordsInpObj()),
		Description: "", // TODO this has no description in the prototype
	}

	networkFetchWhereInpObjPropertiesObj := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:        "WeaviateNetworkFetchWhereInpObjProperties",
			Fields:      filterPropertiesElements,
			Description: "", // TODO no description in prototype
		},
	)

	return networkFetchWhereInpObjPropertiesObj
}

func genNetworkFetchWherePropertyWhereKeywordsInpObj() *graphql.InputObject {
	outputObject := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "NetworkFetchWherePropertyWhereKeywordsInpObj",
			Fields: graphql.InputObjectConfigFieldMap{
				"value": &graphql.InputObjectFieldConfig{
					Type:        graphql.String,
					Description: "", // TODO this has no description in the prototype
				},
				"weight": &graphql.InputObjectFieldConfig{
					Type:        graphql.Float,
					Description: "", // TODO this has no description in the prototype
				},
			},
			Description: "", // TODO this has no description in the prototype
		},
	)
	return outputObject
}

func genNetworkFetchWhereInpObjClassInpObj(filterContainer *utils.FilterContainer) *graphql.InputObject {
	filterClassElements := graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "", // TODO: no description in prototype
		},
		"certainty": &graphql.InputObjectFieldConfig{
			Type:        graphql.Float,
			Description: "", // TODO: no description in prototype
		},
		"keywords": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(filterContainer.WeaviateNetworkWhereNameKeywordsInpObj),
			Description: "", // TODO: no description in prototype
		},
		"first": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: "", // TODO: no description in prototype
		},
	}

	networkFetchWhereInpObjClassInpObj := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:        "WeaviateNetworkFetchWhereInpObjClassInpObj",
			Fields:      filterClassElements,
			Description: "", // TODO no description in prototype
		},
	)
	return networkFetchWhereInpObjClassInpObj
}

func genNetworkIntrospectThingsAndActionsFilterFields(filterContainer *utils.FilterContainer) graphql.InputObjectConfigFieldMap {
	weaviateNetworkIntrospectWherePropertiesObj := genWeaviateNetworkIntrospectWherePropertiesObj(filterContainer)
	weaviateNetworkIntrospectWhereClassObj := genWeaviateNetworkIntrospectWhereClassObj(filterContainer)

	networkIntrospectThingsAndActionsFilterFields := graphql.InputObjectConfigFieldMap{
		"class": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(weaviateNetworkIntrospectWhereClassObj),
			Description: "", // TODO no description in prototype
		},
		"properties": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(weaviateNetworkIntrospectWherePropertiesObj),
			Description: "", // TODO no description in prototype
		},
		//		"first": &graphql.InputObjectFieldConfig{
		//			Type:        graphql.Int,
		//			Description: "", // TODO no description in prototype
		//		},
	}

	return networkIntrospectThingsAndActionsFilterFields
}

func genWeaviateNetworkIntrospectWherePropertiesObj(filterContainer *utils.FilterContainer) *graphql.InputObject {
	filterPropertiesElements := graphql.InputObjectConfigFieldMap{
		"first": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: "", // TODO this has no description in the prototype
		},
		"certainty": &graphql.InputObjectFieldConfig{
			Type:        graphql.Float,
			Description: "", // TODO this has no description in the prototype
		},
		"name": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "", // TODO this has no description in the prototype
		},
		"keywords": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(filterContainer.WeaviateNetworkWhereNameKeywordsInpObj),
			Description: "", // TODO this has no description in the prototype
		},
	}

	weaviateNetworkIntrospectWherePropertiesObj := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:        "WeaviateNetworkIntrospectWherePropertiesObj",
			Fields:      filterPropertiesElements,
			Description: "", // TODO no description in prototype
		},
	)

	return weaviateNetworkIntrospectWherePropertiesObj
}

func genWeaviateNetworkIntrospectWhereClassObj(filterContainer *utils.FilterContainer) *graphql.InputObject {
	filterClassElements := graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "", // TODO: no description in prototype
		},
		"certainty": &graphql.InputObjectFieldConfig{
			Type:        graphql.Float,
			Description: "", // TODO: no description in prototype
		},
		"keywords": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(filterContainer.WeaviateNetworkWhereNameKeywordsInpObj),
			Description: "", // TODO: no description in prototype
		},
		"first": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: "", // TODO: no description in prototype
		},
	}

	weaviateNetworkIntrospectWhereClassObj := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name:        "WeaviateNetworkIntrospectWhereClassObj",
			Fields:      filterClassElements,
			Description: "", // TODO no description in prototype
		},
	)
	return weaviateNetworkIntrospectWhereClassObj
}
