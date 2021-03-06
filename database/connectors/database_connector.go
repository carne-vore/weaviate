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

package dbconnector

import (
	"context"

	"github.com/creativesoftwarefdn/weaviate/database/connectors/utils"
	"github.com/go-openapi/strfmt"
	"github.com/graphql-go/graphql"

	"github.com/creativesoftwarefdn/weaviate/database/connector_state"
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema_migrator"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi"
	"github.com/creativesoftwarefdn/weaviate/messages"
	"github.com/creativesoftwarefdn/weaviate/models"
)

// BaseConnector is the interface that all connectors should have
type BaseConnector interface {
	schema_migrator.Migrator

	Connect() error
	Init() error
	SetServerAddress(serverAddress string)
	SetSchema(s schema.Schema)
	SetMessaging(m *messages.Messaging)

	AddThing(ctx context.Context, thing *models.Thing, UUID strfmt.UUID) error
	GetThing(ctx context.Context, UUID strfmt.UUID, thingResponse *models.ThingGetResponse) error
	ListThings(ctx context.Context, first int, offset int, keyID strfmt.UUID, wheres []*connutils.WhereQuery, thingsResponse *models.ThingsListResponse) error
	UpdateThing(ctx context.Context, thing *models.Thing, UUID strfmt.UUID) error
	DeleteThing(ctx context.Context, thing *models.Thing, UUID strfmt.UUID) error
	HistoryThing(ctx context.Context, UUID strfmt.UUID, history *models.ThingHistory) error
	MoveToHistoryThing(ctx context.Context, thing *models.Thing, UUID strfmt.UUID, deleted bool) error

	AddAction(ctx context.Context, action *models.Action, UUID strfmt.UUID) error
	GetAction(ctx context.Context, UUID strfmt.UUID, actionResponse *models.ActionGetResponse) error
	ListActions(ctx context.Context, first int, offset int, keyID strfmt.UUID, wheres []*connutils.WhereQuery, actionsResponse *models.ActionsListResponse) error
	UpdateAction(ctx context.Context, action *models.Action, UUID strfmt.UUID) error
	DeleteAction(ctx context.Context, action *models.Action, UUID strfmt.UUID) error
	HistoryAction(ctx context.Context, UUID strfmt.UUID, history *models.ActionHistory) error
	MoveToHistoryAction(ctx context.Context, action *models.Action, UUID strfmt.UUID, deleted bool) error

	AddKey(ctx context.Context, key *models.Key, UUID strfmt.UUID, token string) error
	ValidateToken(ctx context.Context, UUID strfmt.UUID, key *models.KeyGetResponse) (token string, err error)
	GetKey(ctx context.Context, UUID strfmt.UUID, keyResponse *models.KeyGetResponse) error
	DeleteKey(ctx context.Context, key *models.Key, UUID strfmt.UUID) error
	GetKeyChildren(ctx context.Context, UUID strfmt.UUID, children *[]*models.KeyGetResponse) error
	UpdateKey(ctx context.Context, key *models.Key, UUID strfmt.UUID, token string) error

	GetGraph(request graphql.ResolveParams) (interface{}, error)
}

// DatabaseConnector is the interface that all DB-connectors should have
type DatabaseConnector interface {
	BaseConnector
	connector_state.Connector
	graphqlapi.Resolver

	GetThings(ctx context.Context, UUIDs []strfmt.UUID, thingResponse *models.ThingsListResponse) error
	GetActions(ctx context.Context, UUIDs []strfmt.UUID, actionResponse *models.ActionsListResponse) error
	GetKeys(ctx context.Context, UUIDs []strfmt.UUID, keyResponse *[]*models.KeyGetResponse) error
}
