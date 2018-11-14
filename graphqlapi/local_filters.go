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

func genFilterFields(filterContainer *utils.FilterContainer) graphql.InputObjectConfigFieldMap {
	staticFilterElements := genStaticWhereFilterElements(filterContainer)

	filterFields := graphql.InputObjectConfigFieldMap{
		"operands": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(genOperandsObject(filterContainer, staticFilterElements)),
			Description: descriptions.WhereOperandsDesc,
		},
	}

	for key, value := range staticFilterElements {
		filterFields[key] = value
	}

	return filterFields
}

// generate these elements once
func genStaticWhereFilterElements(filterContainer *utils.FilterContainer) graphql.InputObjectConfigFieldMap {
	staticFilterElements := graphql.InputObjectConfigFieldMap{
		"operator": &graphql.InputObjectFieldConfig{
			Type:        filterContainer.WhereOperatorEnum,
			Description: descriptions.WhereOperatorDesc,
		},
		"path": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(graphql.String),
			Description: descriptions.WherePathDesc,
		},
		"valueInt": &graphql.InputObjectFieldConfig{
			Type:        graphql.Int,
			Description: descriptions.WhereValueIntDesc,
		},
		"valueNumber": &graphql.InputObjectFieldConfig{
			Type:        graphql.Float,
			Description: descriptions.WhereValueNumberDesc,
		},
		"valueBoolean": &graphql.InputObjectFieldConfig{
			Type:        graphql.Boolean,
			Description: descriptions.WhereValueBooleanDesc,
		},
		"valueString": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.WhereValueStringDesc,
		},
	}

	return staticFilterElements
}

func genOperatorObject() *graphql.Enum {
	enumFilterOptionsMap := graphql.EnumValueConfigMap{
		"And":              &graphql.EnumValueConfig{},
		"Or":               &graphql.EnumValueConfig{},
		"Equal":            &graphql.EnumValueConfig{},
		"Not":              &graphql.EnumValueConfig{},
		"NotEqual":         &graphql.EnumValueConfig{},
		"GreaterThan":      &graphql.EnumValueConfig{},
		"GreaterThanEqual": &graphql.EnumValueConfig{},
		"LessThan":         &graphql.EnumValueConfig{},
		"LessThanEqual":    &graphql.EnumValueConfig{},
	}

	enumFilterOptionsConf := graphql.EnumConfig{
		Name:        "WhereOperatorEnum",
		Values:      enumFilterOptionsMap,
		Description: descriptions.WhereOperatorEnumDesc,
	}

	return graphql.NewEnum(enumFilterOptionsConf)
}

// use a thunk to avoid a cyclical relationship (filters refer to filters refer to .... ad infinitum)
func genOperandsObject(filterContainer *utils.FilterContainer, staticFilterElements graphql.InputObjectConfigFieldMap) *graphql.InputObject {
	outputObject := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "WhereOperandsInpObj",
			Fields: (graphql.InputObjectConfigFieldMapThunk)(func() graphql.InputObjectConfigFieldMap {
				filterFields := genOperandsObjectFields(filterContainer, staticFilterElements)
				return filterFields
			}),
			Description: descriptions.WhereOperandsInpObjDesc,
		},
	)

	filterContainer.Operands = outputObject

	return outputObject
}

func genOperandsObjectFields(filterContainer *utils.FilterContainer, staticFilterElements graphql.InputObjectConfigFieldMap) graphql.InputObjectConfigFieldMap {
	outputFieldConfigMap := staticFilterElements

	outputFieldConfigMap["operands"] = &graphql.InputObjectFieldConfig{
		Type:        graphql.NewList(filterContainer.Operands),
		Description: descriptions.WhereOperandsInpObjDesc,
	}

	return outputFieldConfigMap
}

// generate the GroupBy filter fields
func genGroupByFilterFields() graphql.InputObjectConfigFieldMap {
	groupByFilterFields := graphql.InputObjectConfigFieldMap{
		"group": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.GroupByGroupDesc,
		},
		"count": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.GroupByCountDesc,
		},
		"sum": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.GroupBySumDesc,
		},
		"min": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.GroupByMinDesc,
		},
		"max": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.GroupByMaxDesc,
		},
		"mean": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.GroupByMeanDesc,
		},
		"median": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.GroupByMedianDesc,
		},
		"mode": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: descriptions.GroupByModeDesc,
		},
	}

	return groupByFilterFields
}
