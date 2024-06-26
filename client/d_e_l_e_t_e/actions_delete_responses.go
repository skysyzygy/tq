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

// ActionsDeleteReader is a Reader for the ActionsDelete structure.
type ActionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewActionsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActionsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActionsDeleteNoContent creates a ActionsDeleteNoContent with default headers values
func NewActionsDeleteNoContent() *ActionsDeleteNoContent {
	return &ActionsDeleteNoContent{}
}

/*
ActionsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ActionsDeleteNoContent struct {
}

// IsSuccess returns true when this actions delete no content response has a 2xx status code
func (o *ActionsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this actions delete no content response has a 3xx status code
func (o *ActionsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this actions delete no content response has a 4xx status code
func (o *ActionsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this actions delete no content response has a 5xx status code
func (o *ActionsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this actions delete no content response a status code equal to that given
func (o *ActionsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the actions delete no content response
func (o *ActionsDeleteNoContent) Code() int {
	return 204
}

func (o *ActionsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /CRM/Actions/{actionId}][%d] actionsDeleteNoContent", 204)
}

func (o *ActionsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /CRM/Actions/{actionId}][%d] actionsDeleteNoContent", 204)
}

func (o *ActionsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionsDeleteDefault creates a ActionsDeleteDefault with default headers values
func NewActionsDeleteDefault(code int) *ActionsDeleteDefault {
	return &ActionsDeleteDefault{
		_statusCode: code,
	}
}

/*
ActionsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ActionsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this actions delete default response has a 2xx status code
func (o *ActionsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this actions delete default response has a 3xx status code
func (o *ActionsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this actions delete default response has a 4xx status code
func (o *ActionsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this actions delete default response has a 5xx status code
func (o *ActionsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this actions delete default response a status code equal to that given
func (o *ActionsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the actions delete default response
func (o *ActionsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ActionsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /CRM/Actions/{actionId}][%d] Actions_Delete default %s", o._statusCode, payload)
}

func (o *ActionsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /CRM/Actions/{actionId}][%d] Actions_Delete default %s", o._statusCode, payload)
}

func (o *ActionsDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ActionsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
