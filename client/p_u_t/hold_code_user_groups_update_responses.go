// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// HoldCodeUserGroupsUpdateReader is a Reader for the HoldCodeUserGroupsUpdate structure.
type HoldCodeUserGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HoldCodeUserGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHoldCodeUserGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHoldCodeUserGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHoldCodeUserGroupsUpdateOK creates a HoldCodeUserGroupsUpdateOK with default headers values
func NewHoldCodeUserGroupsUpdateOK() *HoldCodeUserGroupsUpdateOK {
	return &HoldCodeUserGroupsUpdateOK{}
}

/*
HoldCodeUserGroupsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type HoldCodeUserGroupsUpdateOK struct {
	Payload *models.HoldCodeUserGroup
}

// IsSuccess returns true when this hold code user groups update o k response has a 2xx status code
func (o *HoldCodeUserGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hold code user groups update o k response has a 3xx status code
func (o *HoldCodeUserGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hold code user groups update o k response has a 4xx status code
func (o *HoldCodeUserGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hold code user groups update o k response has a 5xx status code
func (o *HoldCodeUserGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hold code user groups update o k response a status code equal to that given
func (o *HoldCodeUserGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hold code user groups update o k response
func (o *HoldCodeUserGroupsUpdateOK) Code() int {
	return 200
}

func (o *HoldCodeUserGroupsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/HoldCodeUserGroups/{holdCodeUserGroupId}][%d] holdCodeUserGroupsUpdateOK %s", 200, payload)
}

func (o *HoldCodeUserGroupsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/HoldCodeUserGroups/{holdCodeUserGroupId}][%d] holdCodeUserGroupsUpdateOK %s", 200, payload)
}

func (o *HoldCodeUserGroupsUpdateOK) GetPayload() *models.HoldCodeUserGroup {
	return o.Payload
}

func (o *HoldCodeUserGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HoldCodeUserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHoldCodeUserGroupsUpdateDefault creates a HoldCodeUserGroupsUpdateDefault with default headers values
func NewHoldCodeUserGroupsUpdateDefault(code int) *HoldCodeUserGroupsUpdateDefault {
	return &HoldCodeUserGroupsUpdateDefault{
		_statusCode: code,
	}
}

/*
HoldCodeUserGroupsUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type HoldCodeUserGroupsUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this hold code user groups update default response has a 2xx status code
func (o *HoldCodeUserGroupsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hold code user groups update default response has a 3xx status code
func (o *HoldCodeUserGroupsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hold code user groups update default response has a 4xx status code
func (o *HoldCodeUserGroupsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hold code user groups update default response has a 5xx status code
func (o *HoldCodeUserGroupsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hold code user groups update default response a status code equal to that given
func (o *HoldCodeUserGroupsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hold code user groups update default response
func (o *HoldCodeUserGroupsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *HoldCodeUserGroupsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/HoldCodeUserGroups/{holdCodeUserGroupId}][%d] HoldCodeUserGroups_Update default %s", o._statusCode, payload)
}

func (o *HoldCodeUserGroupsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/HoldCodeUserGroups/{holdCodeUserGroupId}][%d] HoldCodeUserGroups_Update default %s", o._statusCode, payload)
}

func (o *HoldCodeUserGroupsUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *HoldCodeUserGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
