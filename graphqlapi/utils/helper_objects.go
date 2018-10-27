package utils

import (
	"github.com/graphql-go/graphql"
)

// GraphQLNetworkFieldContents contains all objects regarding GraphQL fields
type GraphQLNetworkFieldContents struct {
	NetworkGetObject     *graphql.Object // Object containing all fields for GraphQL Network Get schema generation
	NetworkGetMetaObject *graphql.Object // Object containing all fields for GraphQL Network GetMeta schema generation
	NetworkFetchObject   *graphql.Object // Object containing all fields for GraphQL Network Fetch schema generation
}

// FilterContainer contains all objects regarding GraphQL filters
type FilterContainer struct {
	WhereOperatorEnum    *graphql.Enum                   // Object containing all fields for the Where filter
	LocalFilterOptions   map[string]*graphql.InputObject // Object containing all fields for Local filters
	NetworkFilterOptions map[string]*graphql.InputObject // Object containing all fields for Network filters
}
