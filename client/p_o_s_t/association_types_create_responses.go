// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// AssociationTypesCreateReader is a Reader for the AssociationTypesCreate structure.
type AssociationTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssociationTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssociationTypesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/AssociationTypes] AssociationTypes_Create", response, response.Code())
	}
}

// NewAssociationTypesCreateOK creates a AssociationTypesCreateOK with default headers values
func NewAssociationTypesCreateOK() *AssociationTypesCreateOK {
	return &AssociationTypesCreateOK{}
}

/*
AssociationTypesCreateOK describes a response with status code 200, with default header values.

OK
*/
type AssociationTypesCreateOK struct {
	Payload *models.AssociationType
}

// IsSuccess returns true when this association types create o k response has a 2xx status code
func (o *AssociationTypesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this association types create o k response has a 3xx status code
func (o *AssociationTypesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this association types create o k response has a 4xx status code
func (o *AssociationTypesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this association types create o k response has a 5xx status code
func (o *AssociationTypesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this association types create o k response a status code equal to that given
func (o *AssociationTypesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the association types create o k response
func (o *AssociationTypesCreateOK) Code() int {
	return 200
}

func (o *AssociationTypesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/AssociationTypes][%d] associationTypesCreateOK  %+v", 200, o.Payload)
}

func (o *AssociationTypesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/AssociationTypes][%d] associationTypesCreateOK  %+v", 200, o.Payload)
}

func (o *AssociationTypesCreateOK) GetPayload() *models.AssociationType {
	return o.Payload
}

func (o *AssociationTypesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssociationType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}