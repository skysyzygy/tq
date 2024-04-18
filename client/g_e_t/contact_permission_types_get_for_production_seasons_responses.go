// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ContactPermissionTypesGetForProductionSeasonsReader is a Reader for the ContactPermissionTypesGetForProductionSeasons structure.
type ContactPermissionTypesGetForProductionSeasonsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPermissionTypesGetForProductionSeasonsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactPermissionTypesGetForProductionSeasonsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ContactPermissionTypes/ByProductionSeason] ContactPermissionTypes_GetForProductionSeasons", response, response.Code())
	}
}

// NewContactPermissionTypesGetForProductionSeasonsOK creates a ContactPermissionTypesGetForProductionSeasonsOK with default headers values
func NewContactPermissionTypesGetForProductionSeasonsOK() *ContactPermissionTypesGetForProductionSeasonsOK {
	return &ContactPermissionTypesGetForProductionSeasonsOK{}
}

/*
ContactPermissionTypesGetForProductionSeasonsOK describes a response with status code 200, with default header values.

OK
*/
type ContactPermissionTypesGetForProductionSeasonsOK struct {
	Payload []*models.ContactPermissionType
}

// IsSuccess returns true when this contact permission types get for production seasons o k response has a 2xx status code
func (o *ContactPermissionTypesGetForProductionSeasonsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact permission types get for production seasons o k response has a 3xx status code
func (o *ContactPermissionTypesGetForProductionSeasonsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact permission types get for production seasons o k response has a 4xx status code
func (o *ContactPermissionTypesGetForProductionSeasonsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact permission types get for production seasons o k response has a 5xx status code
func (o *ContactPermissionTypesGetForProductionSeasonsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact permission types get for production seasons o k response a status code equal to that given
func (o *ContactPermissionTypesGetForProductionSeasonsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact permission types get for production seasons o k response
func (o *ContactPermissionTypesGetForProductionSeasonsOK) Code() int {
	return 200
}

func (o *ContactPermissionTypesGetForProductionSeasonsOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ContactPermissionTypes/ByProductionSeason][%d] contactPermissionTypesGetForProductionSeasonsOK  %+v", 200, o.Payload)
}

func (o *ContactPermissionTypesGetForProductionSeasonsOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ContactPermissionTypes/ByProductionSeason][%d] contactPermissionTypesGetForProductionSeasonsOK  %+v", 200, o.Payload)
}

func (o *ContactPermissionTypesGetForProductionSeasonsOK) GetPayload() []*models.ContactPermissionType {
	return o.Payload
}

func (o *ContactPermissionTypesGetForProductionSeasonsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}