// Code generated by go-swagger; DO NOT EDIT.

package things

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateThingsPropertiesDeleteReader is a Reader for the WeaviateThingsPropertiesDelete structure.
type WeaviateThingsPropertiesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateThingsPropertiesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewWeaviateThingsPropertiesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateThingsPropertiesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateThingsPropertiesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewWeaviateThingsPropertiesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateThingsPropertiesDeleteNoContent creates a WeaviateThingsPropertiesDeleteNoContent with default headers values
func NewWeaviateThingsPropertiesDeleteNoContent() *WeaviateThingsPropertiesDeleteNoContent {
	return &WeaviateThingsPropertiesDeleteNoContent{}
}

/*WeaviateThingsPropertiesDeleteNoContent handles this case with default header values.

Successfully deleted.
*/
type WeaviateThingsPropertiesDeleteNoContent struct {
}

func (o *WeaviateThingsPropertiesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingId}/properties/{propertyName}][%d] weaviateThingsPropertiesDeleteNoContent ", 204)
}

func (o *WeaviateThingsPropertiesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPropertiesDeleteUnauthorized creates a WeaviateThingsPropertiesDeleteUnauthorized with default headers values
func NewWeaviateThingsPropertiesDeleteUnauthorized() *WeaviateThingsPropertiesDeleteUnauthorized {
	return &WeaviateThingsPropertiesDeleteUnauthorized{}
}

/*WeaviateThingsPropertiesDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateThingsPropertiesDeleteUnauthorized struct {
}

func (o *WeaviateThingsPropertiesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingId}/properties/{propertyName}][%d] weaviateThingsPropertiesDeleteUnauthorized ", 401)
}

func (o *WeaviateThingsPropertiesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPropertiesDeleteForbidden creates a WeaviateThingsPropertiesDeleteForbidden with default headers values
func NewWeaviateThingsPropertiesDeleteForbidden() *WeaviateThingsPropertiesDeleteForbidden {
	return &WeaviateThingsPropertiesDeleteForbidden{}
}

/*WeaviateThingsPropertiesDeleteForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateThingsPropertiesDeleteForbidden struct {
}

func (o *WeaviateThingsPropertiesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingId}/properties/{propertyName}][%d] weaviateThingsPropertiesDeleteForbidden ", 403)
}

func (o *WeaviateThingsPropertiesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateThingsPropertiesDeleteNotFound creates a WeaviateThingsPropertiesDeleteNotFound with default headers values
func NewWeaviateThingsPropertiesDeleteNotFound() *WeaviateThingsPropertiesDeleteNotFound {
	return &WeaviateThingsPropertiesDeleteNotFound{}
}

/*WeaviateThingsPropertiesDeleteNotFound handles this case with default header values.

Successful query result but no resource was found.
*/
type WeaviateThingsPropertiesDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateThingsPropertiesDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /things/{thingId}/properties/{propertyName}][%d] weaviateThingsPropertiesDeleteNotFound  %+v", 404, o.Payload)
}

func (o *WeaviateThingsPropertiesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
