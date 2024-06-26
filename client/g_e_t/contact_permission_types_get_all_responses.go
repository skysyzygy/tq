// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ContactPermissionTypesGetAllReader is a Reader for the ContactPermissionTypesGetAll structure.
type ContactPermissionTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPermissionTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactPermissionTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContactPermissionTypesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContactPermissionTypesGetAllOK creates a ContactPermissionTypesGetAllOK with default headers values
func NewContactPermissionTypesGetAllOK() *ContactPermissionTypesGetAllOK {
	return &ContactPermissionTypesGetAllOK{}
}

/*
ContactPermissionTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ContactPermissionTypesGetAllOK struct {
	Payload []*models.ContactPermissionType
}

// IsSuccess returns true when this contact permission types get all o k response has a 2xx status code
func (o *ContactPermissionTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact permission types get all o k response has a 3xx status code
func (o *ContactPermissionTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact permission types get all o k response has a 4xx status code
func (o *ContactPermissionTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact permission types get all o k response has a 5xx status code
func (o *ContactPermissionTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact permission types get all o k response a status code equal to that given
func (o *ContactPermissionTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact permission types get all o k response
func (o *ContactPermissionTypesGetAllOK) Code() int {
	return 200
}

func (o *ContactPermissionTypesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ContactPermissionTypes][%d] contactPermissionTypesGetAllOK %s", 200, payload)
}

func (o *ContactPermissionTypesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ContactPermissionTypes][%d] contactPermissionTypesGetAllOK %s", 200, payload)
}

func (o *ContactPermissionTypesGetAllOK) GetPayload() []*models.ContactPermissionType {
	return o.Payload
}

func (o *ContactPermissionTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContactPermissionTypesGetAllDefault creates a ContactPermissionTypesGetAllDefault with default headers values
func NewContactPermissionTypesGetAllDefault(code int) *ContactPermissionTypesGetAllDefault {
	return &ContactPermissionTypesGetAllDefault{
		_statusCode: code,
	}
}

/*
ContactPermissionTypesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type ContactPermissionTypesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this contact permission types get all default response has a 2xx status code
func (o *ContactPermissionTypesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contact permission types get all default response has a 3xx status code
func (o *ContactPermissionTypesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contact permission types get all default response has a 4xx status code
func (o *ContactPermissionTypesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contact permission types get all default response has a 5xx status code
func (o *ContactPermissionTypesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contact permission types get all default response a status code equal to that given
func (o *ContactPermissionTypesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contact permission types get all default response
func (o *ContactPermissionTypesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *ContactPermissionTypesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ContactPermissionTypes][%d] ContactPermissionTypes_GetAll default %s", o._statusCode, payload)
}

func (o *ContactPermissionTypesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ContactPermissionTypes][%d] ContactPermissionTypes_GetAll default %s", o._statusCode, payload)
}

func (o *ContactPermissionTypesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ContactPermissionTypesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
