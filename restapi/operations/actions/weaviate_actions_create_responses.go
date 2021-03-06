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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateActionsCreateOKCode is the HTTP code returned for type WeaviateActionsCreateOK
const WeaviateActionsCreateOKCode int = 200

/*WeaviateActionsCreateOK Action created

swagger:response weaviateActionsCreateOK
*/
type WeaviateActionsCreateOK struct {

	/*
	  In: Body
	*/
	Payload *models.ActionGetResponse `json:"body,omitempty"`
}

// NewWeaviateActionsCreateOK creates WeaviateActionsCreateOK with default headers values
func NewWeaviateActionsCreateOK() *WeaviateActionsCreateOK {

	return &WeaviateActionsCreateOK{}
}

// WithPayload adds the payload to the weaviate actions create o k response
func (o *WeaviateActionsCreateOK) WithPayload(payload *models.ActionGetResponse) *WeaviateActionsCreateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions create o k response
func (o *WeaviateActionsCreateOK) SetPayload(payload *models.ActionGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsCreateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsCreateAcceptedCode is the HTTP code returned for type WeaviateActionsCreateAccepted
const WeaviateActionsCreateAcceptedCode int = 202

/*WeaviateActionsCreateAccepted Successfully received. No guarantees are made that the Action is persisted.

swagger:response weaviateActionsCreateAccepted
*/
type WeaviateActionsCreateAccepted struct {

	/*
	  In: Body
	*/
	Payload *models.ActionGetResponse `json:"body,omitempty"`
}

// NewWeaviateActionsCreateAccepted creates WeaviateActionsCreateAccepted with default headers values
func NewWeaviateActionsCreateAccepted() *WeaviateActionsCreateAccepted {

	return &WeaviateActionsCreateAccepted{}
}

// WithPayload adds the payload to the weaviate actions create accepted response
func (o *WeaviateActionsCreateAccepted) WithPayload(payload *models.ActionGetResponse) *WeaviateActionsCreateAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions create accepted response
func (o *WeaviateActionsCreateAccepted) SetPayload(payload *models.ActionGetResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsCreateAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateActionsCreateUnauthorizedCode is the HTTP code returned for type WeaviateActionsCreateUnauthorized
const WeaviateActionsCreateUnauthorizedCode int = 401

/*WeaviateActionsCreateUnauthorized Unauthorized or invalid credentials.

swagger:response weaviateActionsCreateUnauthorized
*/
type WeaviateActionsCreateUnauthorized struct {
}

// NewWeaviateActionsCreateUnauthorized creates WeaviateActionsCreateUnauthorized with default headers values
func NewWeaviateActionsCreateUnauthorized() *WeaviateActionsCreateUnauthorized {

	return &WeaviateActionsCreateUnauthorized{}
}

// WriteResponse to the client
func (o *WeaviateActionsCreateUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// WeaviateActionsCreateForbiddenCode is the HTTP code returned for type WeaviateActionsCreateForbidden
const WeaviateActionsCreateForbiddenCode int = 403

/*WeaviateActionsCreateForbidden The used API-key has insufficient permissions.

swagger:response weaviateActionsCreateForbidden
*/
type WeaviateActionsCreateForbidden struct {
}

// NewWeaviateActionsCreateForbidden creates WeaviateActionsCreateForbidden with default headers values
func NewWeaviateActionsCreateForbidden() *WeaviateActionsCreateForbidden {

	return &WeaviateActionsCreateForbidden{}
}

// WriteResponse to the client
func (o *WeaviateActionsCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// WeaviateActionsCreateUnprocessableEntityCode is the HTTP code returned for type WeaviateActionsCreateUnprocessableEntity
const WeaviateActionsCreateUnprocessableEntityCode int = 422

/*WeaviateActionsCreateUnprocessableEntity Request body contains well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?

swagger:response weaviateActionsCreateUnprocessableEntity
*/
type WeaviateActionsCreateUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewWeaviateActionsCreateUnprocessableEntity creates WeaviateActionsCreateUnprocessableEntity with default headers values
func NewWeaviateActionsCreateUnprocessableEntity() *WeaviateActionsCreateUnprocessableEntity {

	return &WeaviateActionsCreateUnprocessableEntity{}
}

// WithPayload adds the payload to the weaviate actions create unprocessable entity response
func (o *WeaviateActionsCreateUnprocessableEntity) WithPayload(payload *models.ErrorResponse) *WeaviateActionsCreateUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate actions create unprocessable entity response
func (o *WeaviateActionsCreateUnprocessableEntity) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateActionsCreateUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
