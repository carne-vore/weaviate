package graphqlapi

import ()

const WeaviateObjDesc = "Location of the root query"
const WeaviateLocalDesc = "Query a local Weaviate instance"
const WeaviateNetworkDesc = "Query a Weaviate network"

// LOCAL
const LocalObjDesc = "Type of query on the local Weaviate"

// Get
const LocalGetActionsDesc = "Get Actions on the Local Weaviate"
const LocalGetThingsDesc = "Get Things on the Local Weaviate"

const LocalGetObjDesc = "Type of Get function to get Things or Actions on the Local Weaviate"
const LocalGetDesc = "Get Things or Actions on the local weaviate"

const LocalGetActionsObjDesc = "Type of Actions i.e. Actions classes to Get on the Local Weaviate"
const LocalGetThingsObjDesc = "Type of Things i.e. Things classes to Get on the Local Weaviate"

const LocalGetClassUUIDDesc = "UUID of the thing or action given by the local Weaviate instance"

// GetMeta
const LocalGetMetaActionsDesc = "Get Meta information about Actions on the Local Weaviate"
const LocalGetMetaThingsDesc = "Get Meta information about Things on the Local Weaviate"

const LocalGetMetaObjDesc = "Type of Get function to get meta information about Things or Actions on the Local Weaviate"
const LocalGetMetaDesc = "Query to Get Meta information about the data in the local Weaviate instance"

const GetMetaPropertyTypeDesc string = "Datatype of the property"
const GetMetaPropertyCountDesc string = "Total amount of found instances"
const GetMetaPropertyTopOccurrencesDesc string = "Object for the most frequent property values"
const GetMetaPropertyTopOccurrencesValueDesc string = "The most frequently occurring value of this property in the dataset"
const GetMetaPropertyTopOccurrencesOccursDesc string = "Number of occurrence of this property value"
const GetMetaPropertyLowestDesc string = "Lowest value found in the dataset for this property"
const GetMetaPropertyHighestDesc string = "Highest value found in the dataset for this property"
const GetMetaPropertyAverageDesc string = "Average value found in the dataset for this property"
const GetMetaPropertySumDesc string = "Sum of values found in the dataset for this property"
const GetMetaPropertyObjectDesc string = "Object for property meta information"

// NETWORK
const NetworkWeaviateDesc = "Object field for network Weaviate instance: "
const NetworkObjDesc = "Type of query on the Weaviate network"

// Get
const NetworkGetObjDesc = "Get Things or Actions on a Weaviate in the network"
const NetworkGetActionsDesc = "Get Actions from the network"
const NetworkGetThingsDesc = "Get Things from the network"
const NetworkGetWeaviateObjDesc = "Container for Get Things and Actions fields for network Weaviate instance: "
const NetworkGetDesc = "Get Things or Actions on the network of weaviate"

const NetworkGetWeaviateActionsObjDesc = "Type of Actions i.e. Actions classes to Get from the Network"

const NetworkGetClassUUIDDesc = "UUID of the thing or action given by the Weaviate network"

const NetworkGetWeaviateThingsObjDesc = "Type of Things i.e. Things classes to Get from the Network"

// GetMeta
const NetworkGetMetaObjDesc = "Get meta information about Things or Actions on a Weaviate in the network"
const NetworkGetMetaActionsDesc = "Get Meta information about Actions on a Weaviate in the network"
const NetworkGetMetaThingsDesc = "Get Meta information about Things on a Weaviate in a network"
const NetworkGetMetaWeaviateObjDesc = "Container for GetMeta Things and Actions fields for network Weaviate instance: "
const NetworkGetMetaDesc = "Query to Get Meta information about the data in the Network Weaviate"

const NetworkGetMetaWeaviateThingsObjDesc = "Type of Things i.e. Things classes to GetMeta information of on a Weaviate in the Network"
const NetworkGetMetaWeaviateActionsObjDesc = "Type of Actions i.e. Actions classes to GetMeta information of on a Weaviate in the Network"

const LocalGetMetaThingsObjDesc = "Type of Things i.e. Things classes to GetMeta information of on the Local Weaviate"
const LocalGetMetaActionsObjDesc = "Type of Actions i.e. Actions classes to GetMeta information of on the Local Weaviate"

const GetMetaMetaPropertyDesc = "Meta information about a class object and its (filtered) objects"
const GetMetaPropertyDesc = "Meta information about the property "

const GetMetaClassPropertyTotalTrueDesc = "The amount of times this property's value is true in the dataset"
const GetMetaClassPropertyPercentageTrueDesc = "Percentage of boolean values that is true"
const GetMetaClassPropertyPointingToDesc = "Which other classes the object property is pointing to"
const GetMetaClassMetaCountDesc = "Total amount of found instances"
const GetMetaClassMetaObjDesc = "Meta information about a class object and its (filtered) objects"

// Fetch
const NetworkFetchActionsDesc = "Fetch Actions from the network"
const NetworkFetchThingsDesc = "Fetch Things from the network"
const NetworkFetchFuzzyDesc = "Perform a Fuzzy Fetch, with only a ontology value, on the Network" // TODO - rewrite this one
const NetworkFetchObjDesc = "Type of network fetch: e.g. Things, Actions"                         // TODO - rewrite this one

const NetworkFetchActionBeaconDesc = "Beacon as an Action result found in the Network"                                         // TODO
const NetworkFetchActionCertaintyDesc = "Certainty of beacon result found in the Network has expected ontology characterisics" // TODO
const NetworkFetchActionsObjDesc = "Type of Actions i.e. classes to fetch on the network"                                      // TODO

const NetworkFetchThingBeaconDesc = "Beacon as an Thing result found in the Network"                                          // TODO
const NetworkFetchThingCertaintyDesc = "Certainty of beacon result found in the Network has expected ontology characterisics" // TODO
const NetworkFetchThingsObjDesc = "Type of Things i.e. classes to fetch on the network"                                       // TODO - refactor this + the 5 above to become 3

const NetworkFetchFuzzyBeaconDesc = "The beacon of the node that is a result from the fuzzy network fetch"                            // TODO - rewrite
const NetworkFetchFuzzyCertaintyDesc = "The certainty the node has a value matching the value searched on in the fuzzy network fetch" // TODO - rewrite
const NetworkFetchFuzzyObjDesc = "The objects what to request from this network fuzzy fetch query"                                    // TODO - rewrite

const NetworkFetchDesc = "Do a fuzzy search fetch to search Things or Actions on the network weaviate"

// Introspect
const NetworkIntrospectActionsDesc = "Actions to introspect in the network"
const NetworkIntrospectThingsDesc = "Things to introspect in the network"
const NetworkIntrospectBeaconDesc = "Beacon to introspect in the network"
const NetworkIntrospectObjDesc = "Type of object to introspect in the network" // TODO - reformulate

const NetworkIntrospectWeaviateDesc = "Weaviate node the found class in the Network Introspection search is in."
const NetworkIntrospectClassNameDesc = "To filter on which class name"
const NetworkIntrospectCertaintyDesc = "To filter with which certainty from 0-1"
const NetworkIntrospectActionsObjDesc = "Object for which Actions to introspect in the network"
const NetworkIntrospectThingsObjDesc = "Object for which Things to introspect in the network"

const NetworkIntrospectBeaconPropertiesDesc = "Which properties to filter on"
const NetworkIntrospectBeaconObjDesc = "Object for which beacon to introspect in the network"

const NetworkIntrospectBeaconPropertiesPropertyNameDesc = "Which property name to filter properties on"

const NetworkIntrospectDesc = "To fetch meta information about the ontology of Things or Actions on the network weaviate"

// FILTERS

// Where filter elements
const LocalGetWhereDesc = "Filter options for the Get search, to convert the data to the filter input"
const LocalGetMetaWhereDesc = "Filter options for the GetMeta search, to convert the data to the filter input"

const NetworkGetWhereDesc = "Filter options for the Network Get search, to convert the data to the filter input"
const NetworkGetWhereInpObjDesc = "Filter options for the Network Get search, to convert the data to the filter input"

const NetworkGetMetaWhereDesc = "Filter options for the Network GetMeta search, to convert the data to the filter input"
const NetworkGetMetaWhereInpObjDesc = "Filter options for the GetMeta search, to convert the data to the filter input"

const WhereOperandsDesc = "Operands in the 'where' filter field, is a list of objects"
const WhereOperatorDesc = "Operator in the 'where' filter field, value is one of the 'WhereOperatorEnum' object"
const WherePathDesc = "Path of from 'Things' or 'Actions' to the property name through the classes"
const WhereValueIntDesc = "Integer value that the property at the provided path will be compared to by an operator"
const WhereValueNumberDesc = "Number value that the property at the provided path will be compared to by an operator"
const WhereValueBooleanDesc = "Boolean value that the property at the provided path will be compared to by an operator"
const WhereValueStringDesc = "String value that the property at the provided path will be compared to by an operator"
const WhereOperatorEnumDesc = "Enumeration object for the 'where' filter"
const WhereOperandsInpObjDesc = "Operands in the 'where' filter field, is a list of objects"
const WhereKeywordsInpObjDesc = "Specify the value and the weight of a keyword" // remove name from const name
const WhereKeywordsValueDesc = "The value of the keyword"
const WhereKeywordsWeightDesc = "The weight of the keyword"

// Network filter elements
const NetworkTimeoutDesc = "The time in seconds after which an unresolved request automatically fails"

// Pagination filter elements
const FirstDesc = "Pagination option, show the first x results"
const AfterDesc = "Pagination option, show the results after the first x results"

// Properties and Classes filter elements (used by Fetch and Introspect Where filters)
const WherePropertiesObjDesc = "Specify which properties to filter on"
const WherePropertiesPropertyNameDesc = "Specify which property name to filter properties on"
const WherePropertiesDesc = "Specify which properties to filter on"
const WhereFirstDesc = "zeven 7" // what does this do?
const WhereCertaintyDesc = "With which certainty 0-1 to filter properties on"
const WhereNameDesc = "Specify the name of the property to filter on"
const WhereKeywordsDesc = "Specify which keywords to filter on"
const WhereObjDesc = "elf 11"
const WhereClassDesc = "Specify which classes to filter on"

// Unique Fetch filter elements
const FetchWhereInpObjDesc = "achttien 18" // might be the same as IntrospectWhereInpObjDesc
const FetchWhereFilterFieldsDesc = "Filter options for the Network GetMeta search, to convert the data to the filter input"
const FetchFuzzyValueDesc = "The ontology value to fetch Things or Actions on the Network on"
const FetchFuzzyCertaintyDesc = "The minimum certainty the nodes should match the given ontology value"

// Unique Introspect filter elements
const IntrospectWhereInpObjDesc = "Specify which classes and properties to filter on"

// GroupBy filter elements
const GroupByGroupDesc = "The property of the class to group by."
const GroupByCountDesc = "The number of instances of a property in a group."
const GroupBySumDesc = "The sum of the values of a property in a group."
const GroupByMinDesc = "The lowest occuring value of a property in a group."
const GroupByMaxDesc = "The highest occuring value of a property in a group."
const GroupByMeanDesc = "The average value of a property in a group."
const GroupByMedianDesc = "The median of a property in a group."
const GroupByModeDesc = "The mode of a property in a group."
