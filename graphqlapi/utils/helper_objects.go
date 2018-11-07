package utils

import (
	"github.com/graphql-go/graphql"
)

// GraphQLNetworkFieldContents contains all objects regarding GraphQL fields
type GraphQLNetworkFieldContents struct {
	NetworkGetObject        *graphql.Object // Object containing all fields for GraphQL Network Get schema generation
	NetworkGetMetaObject    *graphql.Object // Object containing all fields for GraphQL Network GetMeta schema generation
	NetworkFetchObject      *graphql.Object // Object containing all fields for GraphQL Network Fetch schema generation
	NetworkIntrospectObject *graphql.Object // Object containing all fields for GraphQL Network Introspect schema generation
}

// FilterContainer contains all objects regarding GraphQL filters. Some filter elements are declared as global variables in the prototype, this struct achieves the same goal.
type FilterContainer struct {
	WhereOperatorEnum                           *graphql.Enum // Object containing all fields for the Where filter
	Operands                                    *graphql.InputObject
	LocalFilterOptions                          map[string]*graphql.InputObject // Object containing all fields for Local filters
	NetworkFilterOptions                        map[string]*graphql.InputObject // Object containing all fields for Network filters
	FetchThingsActionsWhereFilterArgConf        *graphql.ArgumentConfig         // Object containing the Where filter fields for Fetch Things and Actions
	IntrospectThingsActionsWhereFilterArgConf   *graphql.ArgumentConfig         // Object containing the Where filter fields for Introspect Things and Actions
	WeaviateNetworkWhereNameKeywordsInpObj      *graphql.InputObject            // Object containing a global filter element
	WeaviateNetworkIntrospectPropertiesObjField *graphql.Field                  // Object containing a global filter element
} // TODO: rename localfilter to localgetandgetmetafilt... and networkfilteroptions to networkgetandgetmetafilteroptions
