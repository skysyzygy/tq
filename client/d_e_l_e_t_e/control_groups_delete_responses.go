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

// ControlGroupsDeleteReader is a Reader for the ControlGroupsDelete structure.
type ControlGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControlGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewControlGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewControlGroupsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewControlGroupsDeleteNoContent creates a ControlGroupsDeleteNoContent with default headers values
func NewControlGroupsDeleteNoContent() *ControlGroupsDeleteNoContent {
	return &ControlGroupsDeleteNoContent{}
}

/*
ControlGroupsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ControlGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this control groups delete no content response has a 2xx status code
func (o *ControlGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this control groups delete no content response has a 3xx status code
func (o *ControlGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this control groups delete no content response has a 4xx status code
func (o *ControlGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this control groups delete no content response has a 5xx status code
func (o *ControlGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this control groups delete no content response a status code equal to that given
func (o *ControlGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the control groups delete no content response
func (o *ControlGroupsDeleteNoContent) Code() int {
	return 204
}

func (o *ControlGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ControlGroups/{id}][%d] controlGroupsDeleteNoContent", 204)
}

func (o *ControlGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ControlGroups/{id}][%d] controlGroupsDeleteNoContent", 204)
}

func (o *ControlGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewControlGroupsDeleteDefault creates a ControlGroupsDeleteDefault with default headers values
func NewControlGroupsDeleteDefault(code int) *ControlGroupsDeleteDefault {
	return &ControlGroupsDeleteDefault{
		_statusCode: code,
	}
}

/*
ControlGroupsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ControlGroupsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this control groups delete default response has a 2xx status code
func (o *ControlGroupsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this control groups delete default response has a 3xx status code
func (o *ControlGroupsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this control groups delete default response has a 4xx status code
func (o *ControlGroupsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this control groups delete default response has a 5xx status code
func (o *ControlGroupsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this control groups delete default response a status code equal to that given
func (o *ControlGroupsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the control groups delete default response
func (o *ControlGroupsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ControlGroupsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ControlGroups/{id}][%d] ControlGroups_Delete default %s", o._statusCode, payload)
}

func (o *ControlGroupsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ControlGroups/{id}][%d] ControlGroups_Delete default %s", o._statusCode, payload)
}

func (o *ControlGroupsDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ControlGroupsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
