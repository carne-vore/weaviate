package descriptions

import ()

// NETWORK
const NetworkIntrospectDesc string = "Get Introspection information about Things, Actions and/or Beacons in a Weaviate network"
const NetworkIntrospectObjDesc string = "An object used to perform an Introspection query on a Weaviate network"

const NetworkIntrospectActionsDesc string = "Introspect Actions in a Weaviate network"
const NetworkIntrospectThingsDesc string = "Introspect Things in a Weaviate network"
const NetworkIntrospectBeaconDesc string = "Introspect Beacons in a Weaviate network"

const NetworkIntrospectWeaviateDesc string = "The Weaviate instance that the current Thing, Action or Beacon belongs to"
const NetworkIntrospectClassNameDesc string = "The name of the current Thing, Action or Beacon's class" // TODO check with @lauraham
const NetworkIntrospectCertaintyDesc string = "To filter with which certainty from 0-1"                 // what does this do?

const NetworkIntrospectActionsObjDesc string = "An object used to Introspect Actions on a Weaviate network"
const NetworkIntrospectThingsObjDesc string = "An object used to Introspect Things on a Weaviate network"
const NetworkIntrospectBeaconObjDesc string = "An object used to Introspect Beacons on a Weaviate network"

const NetworkIntrospectBeaconPropertiesDesc string = "Which properties to filter on"                           // what does this do?
const NetworkIntrospectBeaconPropertiesPropertyNameDesc string = "Which property name to filter properties on" // what does this do?

// TODO introspect beacon properties has 'Which properties to filter on' as tooltip, check this
