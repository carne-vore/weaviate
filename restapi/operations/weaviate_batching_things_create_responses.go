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
// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateBatchingThingsCreateOKCode is the HTTP code returned for type WeaviateBatchingThingsCreateOK
const WeaviateBatchingThingsCreateOKCode int = 200

/*WeaviateBatchingThingsCreateOK Things created.

swagger:response weaviateBatchingThingsCreateOK
*/
type WeaviateBatchingThingsCreateOK struct {

	/*
	  In: Body
	*/
	Payload []*models.ThingsGetResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingThingsCreateOK creates WeaviateBatchingThingsCreateOK with default headers values
func NewWeaviateBatchingThingsCreateOK() *WeaviateBatchingThingsCreateOK {

	return &WeaviateBatchingThingsCreateOK{}
}

// WithPayload adds the payload to the weaviate batching things create o k response
func (o *WeaviateBatchingThingsCreateOK) WithPayload(payload []*models.ThingsGetResponse) *WeaviateBatchingThingsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching things create o k response
func (o *WeaviateBatchingThingsCreateOK) SetPayload(payload []*models.ThingsGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingThingsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.ThingsGetResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// WeaviateBatchingThingsCreateAcceptedCode is the HTTP code returned for type WeaviateBatchingThingsCreateAccepted
const WeaviateBatchingThingsCreateAcceptedCode int = 202

/*WeaviateBatchingThingsCreateAccepted Successfully received.

swagger:response weaviateBatchingThingsCreateAccepted
*/
type WeaviateBatchingThingsCreateAccepted struct {

	/*
	  In: Body
	*/
	Payload []*models.ThingsGetResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingThingsCreateAccepted creates WeaviateBatchingThingsCreateAccepted with default headers values
func NewWeaviateBatchingThingsCreateAccepted() *WeaviateBatchingThingsCreateAccepted {

	return &WeaviateBatchingThingsCreateAccepted{}
}

// WithPayload adds the payload to the weaviate batching things create accepted response
func (o *WeaviateBatchingThingsCreateAccepted) WithPayload(payload []*models.ThingsGetResponse) *WeaviateBatchingThingsCreateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching things create accepted response
func (o *WeaviateBatchingThingsCreateAccepted) SetPayload(payload []*models.ThingsGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingThingsCreateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.ThingsGetResponse, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// WeaviateBatchingThingsCreateUnauthorizedCode is the HTTP code returned for type WeaviateBatchingThingsCreateUnauthorized
const WeaviateBatchingThingsCreateUnauthorizedCode int = 401

/*WeaviateBatchingThingsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateBatchingThingsCreateUnauthorized
*/
type WeaviateBatchingThingsCreateUnauthorized struct {
}

// NewWeaviateBatchingThingsCreateUnauthorized creates WeaviateBatchingThingsCreateUnauthorized with default headers values
func NewWeaviateBatchingThingsCreateUnauthorized() *WeaviateBatchingThingsCreateUnauthorized {

	return &WeaviateBatchingThingsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateBatchingThingsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateBatchingThingsCreateForbiddenCode is the HTTP code returned for type WeaviateBatchingThingsCreateForbidden
const WeaviateBatchingThingsCreateForbiddenCode int = 403

/*WeaviateBatchingThingsCreateForbidden The used API-key has insufficient permissions.

swagger:response weaviateBatchingThingsCreateForbidden
*/
type WeaviateBatchingThingsCreateForbidden struct {
}

// NewWeaviateBatchingThingsCreateForbidden creates WeaviateBatchingThingsCreateForbidden with default headers values
func NewWeaviateBatchingThingsCreateForbidden() *WeaviateBatchingThingsCreateForbidden {

	return &WeaviateBatchingThingsCreateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateBatchingThingsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateBatchingThingsCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateBatchingThingsCreateUnprocessableEntity
const WeaviateBatchingThingsCreateUnprocessableEntityCode int = 422

/*WeaviateBatchingThingsCreateUnprocessableEntity Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateBatchingThingsCreateUnprocessableEntity
*/
type WeaviateBatchingThingsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateBatchingThingsCreateUnprocessableEntity creates WeaviateBatchingThingsCreateUnprocessableEntity with default headers values
func NewWeaviateBatchingThingsCreateUnprocessableEntity() *WeaviateBatchingThingsCreateUnprocessableEntity {

	return &WeaviateBatchingThingsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate batching things create unprocessable entity response
func (o *WeaviateBatchingThingsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateBatchingThingsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate batching things create unprocessable entity response
func (o *WeaviateBatchingThingsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateBatchingThingsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
