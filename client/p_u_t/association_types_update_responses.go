// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// AssociationTypesUpdateReader is a Reader for the AssociationTypesUpdate structure.
type AssociationTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssociationTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssociationTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/AssociationTypes/{id}] AssociationTypes_Update", response, response.Code())
	}
}

// NewAssociationTypesUpdateOK creates a AssociationTypesUpdateOK with default headers values
func NewAssociationTypesUpdateOK() *AssociationTypesUpdateOK {
	return &AssociationTypesUpdateOK{}
}

/*
AssociationTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type AssociationTypesUpdateOK struct {
	Payload *models.AssociationType
}

// IsSuccess returns true when this association types update o k response has a 2xx status code
func (o *AssociationTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this association types update o k response has a 3xx status code
func (o *AssociationTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this association types update o k response has a 4xx status code
func (o *AssociationTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this association types update o k response has a 5xx status code
func (o *AssociationTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this association types update o k response a status code equal to that given
func (o *AssociationTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the association types update o k response
func (o *AssociationTypesUpdateOK) Code() int {
	return 200
}

func (o *AssociationTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/AssociationTypes/{id}][%d] associationTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *AssociationTypesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/AssociationTypes/{id}][%d] associationTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *AssociationTypesUpdateOK) GetPayload() *models.AssociationType {
	return o.Payload
}

func (o *AssociationTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssociationType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
