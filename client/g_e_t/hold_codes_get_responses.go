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

// HoldCodesGetReader is a Reader for the HoldCodesGet structure.
type HoldCodesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HoldCodesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHoldCodesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHoldCodesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHoldCodesGetOK creates a HoldCodesGetOK with default headers values
func NewHoldCodesGetOK() *HoldCodesGetOK {
	return &HoldCodesGetOK{}
}

/*
HoldCodesGetOK describes a response with status code 200, with default header values.

OK
*/
type HoldCodesGetOK struct {
	Payload *models.HoldCode
}

// IsSuccess returns true when this hold codes get o k response has a 2xx status code
func (o *HoldCodesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hold codes get o k response has a 3xx status code
func (o *HoldCodesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hold codes get o k response has a 4xx status code
func (o *HoldCodesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hold codes get o k response has a 5xx status code
func (o *HoldCodesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hold codes get o k response a status code equal to that given
func (o *HoldCodesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hold codes get o k response
func (o *HoldCodesGetOK) Code() int {
	return 200
}

func (o *HoldCodesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/HoldCodes/{holdCodeId}][%d] holdCodesGetOK %s", 200, payload)
}

func (o *HoldCodesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/HoldCodes/{holdCodeId}][%d] holdCodesGetOK %s", 200, payload)
}

func (o *HoldCodesGetOK) GetPayload() *models.HoldCode {
	return o.Payload
}

func (o *HoldCodesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HoldCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHoldCodesGetDefault creates a HoldCodesGetDefault with default headers values
func NewHoldCodesGetDefault(code int) *HoldCodesGetDefault {
	return &HoldCodesGetDefault{
		_statusCode: code,
	}
}

/*
HoldCodesGetDefault describes a response with status code -1, with default header values.

Error
*/
type HoldCodesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this hold codes get default response has a 2xx status code
func (o *HoldCodesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hold codes get default response has a 3xx status code
func (o *HoldCodesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hold codes get default response has a 4xx status code
func (o *HoldCodesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hold codes get default response has a 5xx status code
func (o *HoldCodesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hold codes get default response a status code equal to that given
func (o *HoldCodesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hold codes get default response
func (o *HoldCodesGetDefault) Code() int {
	return o._statusCode
}

func (o *HoldCodesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/HoldCodes/{holdCodeId}][%d] HoldCodes_Get default %s", o._statusCode, payload)
}

func (o *HoldCodesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/HoldCodes/{holdCodeId}][%d] HoldCodes_Get default %s", o._statusCode, payload)
}

func (o *HoldCodesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *HoldCodesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
