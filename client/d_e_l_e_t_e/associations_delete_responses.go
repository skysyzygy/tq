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

// AssociationsDeleteReader is a Reader for the AssociationsDelete structure.
type AssociationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssociationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAssociationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAssociationsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAssociationsDeleteNoContent creates a AssociationsDeleteNoContent with default headers values
func NewAssociationsDeleteNoContent() *AssociationsDeleteNoContent {
	return &AssociationsDeleteNoContent{}
}

/*
AssociationsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type AssociationsDeleteNoContent struct {
}

// IsSuccess returns true when this associations delete no content response has a 2xx status code
func (o *AssociationsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this associations delete no content response has a 3xx status code
func (o *AssociationsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this associations delete no content response has a 4xx status code
func (o *AssociationsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this associations delete no content response has a 5xx status code
func (o *AssociationsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this associations delete no content response a status code equal to that given
func (o *AssociationsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the associations delete no content response
func (o *AssociationsDeleteNoContent) Code() int {
	return 204
}

func (o *AssociationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /CRM/Associations/{associationId}][%d] associationsDeleteNoContent", 204)
}

func (o *AssociationsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /CRM/Associations/{associationId}][%d] associationsDeleteNoContent", 204)
}

func (o *AssociationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAssociationsDeleteDefault creates a AssociationsDeleteDefault with default headers values
func NewAssociationsDeleteDefault(code int) *AssociationsDeleteDefault {
	return &AssociationsDeleteDefault{
		_statusCode: code,
	}
}

/*
AssociationsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type AssociationsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this associations delete default response has a 2xx status code
func (o *AssociationsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this associations delete default response has a 3xx status code
func (o *AssociationsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this associations delete default response has a 4xx status code
func (o *AssociationsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this associations delete default response has a 5xx status code
func (o *AssociationsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this associations delete default response a status code equal to that given
func (o *AssociationsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the associations delete default response
func (o *AssociationsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *AssociationsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /CRM/Associations/{associationId}][%d] Associations_Delete default %s", o._statusCode, payload)
}

func (o *AssociationsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /CRM/Associations/{associationId}][%d] Associations_Delete default %s", o._statusCode, payload)
}

func (o *AssociationsDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AssociationsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
