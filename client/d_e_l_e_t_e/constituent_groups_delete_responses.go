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

// ConstituentGroupsDeleteReader is a Reader for the ConstituentGroupsDelete structure.
type ConstituentGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewConstituentGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConstituentGroupsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConstituentGroupsDeleteNoContent creates a ConstituentGroupsDeleteNoContent with default headers values
func NewConstituentGroupsDeleteNoContent() *ConstituentGroupsDeleteNoContent {
	return &ConstituentGroupsDeleteNoContent{}
}

/*
ConstituentGroupsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ConstituentGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this constituent groups delete no content response has a 2xx status code
func (o *ConstituentGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent groups delete no content response has a 3xx status code
func (o *ConstituentGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent groups delete no content response has a 4xx status code
func (o *ConstituentGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent groups delete no content response has a 5xx status code
func (o *ConstituentGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent groups delete no content response a status code equal to that given
func (o *ConstituentGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the constituent groups delete no content response
func (o *ConstituentGroupsDeleteNoContent) Code() int {
	return 204
}

func (o *ConstituentGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentGroups/{id}][%d] constituentGroupsDeleteNoContent", 204)
}

func (o *ConstituentGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentGroups/{id}][%d] constituentGroupsDeleteNoContent", 204)
}

func (o *ConstituentGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConstituentGroupsDeleteDefault creates a ConstituentGroupsDeleteDefault with default headers values
func NewConstituentGroupsDeleteDefault(code int) *ConstituentGroupsDeleteDefault {
	return &ConstituentGroupsDeleteDefault{
		_statusCode: code,
	}
}

/*
ConstituentGroupsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ConstituentGroupsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this constituent groups delete default response has a 2xx status code
func (o *ConstituentGroupsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this constituent groups delete default response has a 3xx status code
func (o *ConstituentGroupsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this constituent groups delete default response has a 4xx status code
func (o *ConstituentGroupsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this constituent groups delete default response has a 5xx status code
func (o *ConstituentGroupsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this constituent groups delete default response a status code equal to that given
func (o *ConstituentGroupsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the constituent groups delete default response
func (o *ConstituentGroupsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ConstituentGroupsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentGroups/{id}][%d] ConstituentGroups_Delete default %s", o._statusCode, payload)
}

func (o *ConstituentGroupsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentGroups/{id}][%d] ConstituentGroups_Delete default %s", o._statusCode, payload)
}

func (o *ConstituentGroupsDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ConstituentGroupsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
