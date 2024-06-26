// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// ContactPermissionTypesDeleteReader is a Reader for the ContactPermissionTypesDelete structure.
type ContactPermissionTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPermissionTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContactPermissionTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContactPermissionTypesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContactPermissionTypesDeleteNoContent creates a ContactPermissionTypesDeleteNoContent with default headers values
func NewContactPermissionTypesDeleteNoContent() *ContactPermissionTypesDeleteNoContent {
	return &ContactPermissionTypesDeleteNoContent{}
}

/*
ContactPermissionTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ContactPermissionTypesDeleteNoContent struct {
}

// IsSuccess returns true when this contact permission types delete no content response has a 2xx status code
func (o *ContactPermissionTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact permission types delete no content response has a 3xx status code
func (o *ContactPermissionTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact permission types delete no content response has a 4xx status code
func (o *ContactPermissionTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact permission types delete no content response has a 5xx status code
func (o *ContactPermissionTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this contact permission types delete no content response a status code equal to that given
func (o *ContactPermissionTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the contact permission types delete no content response
func (o *ContactPermissionTypesDeleteNoContent) Code() int {
	return 204
}

func (o *ContactPermissionTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ContactPermissionTypes/{id}][%d] contactPermissionTypesDeleteNoContent", 204)
}

func (o *ContactPermissionTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ContactPermissionTypes/{id}][%d] contactPermissionTypesDeleteNoContent", 204)
}

func (o *ContactPermissionTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContactPermissionTypesDeleteDefault creates a ContactPermissionTypesDeleteDefault with default headers values
func NewContactPermissionTypesDeleteDefault(code int) *ContactPermissionTypesDeleteDefault {
	return &ContactPermissionTypesDeleteDefault{
		_statusCode: code,
	}
}

/*
ContactPermissionTypesDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ContactPermissionTypesDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this contact permission types delete default response has a 2xx status code
func (o *ContactPermissionTypesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contact permission types delete default response has a 3xx status code
func (o *ContactPermissionTypesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contact permission types delete default response has a 4xx status code
func (o *ContactPermissionTypesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contact permission types delete default response has a 5xx status code
func (o *ContactPermissionTypesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contact permission types delete default response a status code equal to that given
func (o *ContactPermissionTypesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contact permission types delete default response
func (o *ContactPermissionTypesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ContactPermissionTypesDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ContactPermissionTypes/{id}][%d] ContactPermissionTypes_Delete default %s", o._statusCode, payload)
}

func (o *ContactPermissionTypesDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ContactPermissionTypes/{id}][%d] ContactPermissionTypes_Delete default %s", o._statusCode, payload)
}

func (o *ContactPermissionTypesDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ContactPermissionTypesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
