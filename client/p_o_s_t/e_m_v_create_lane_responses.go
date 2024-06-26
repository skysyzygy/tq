// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// EMVCreateLaneReader is a Reader for the EMVCreateLane structure.
type EMVCreateLaneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EMVCreateLaneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEMVCreateLaneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEMVCreateLaneDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEMVCreateLaneOK creates a EMVCreateLaneOK with default headers values
func NewEMVCreateLaneOK() *EMVCreateLaneOK {
	return &EMVCreateLaneOK{}
}

/*
EMVCreateLaneOK describes a response with status code 200, with default header values.

OK
*/
type EMVCreateLaneOK struct {
	Payload *models.Lane
}

// IsSuccess returns true when this e m v create lane o k response has a 2xx status code
func (o *EMVCreateLaneOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this e m v create lane o k response has a 3xx status code
func (o *EMVCreateLaneOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this e m v create lane o k response has a 4xx status code
func (o *EMVCreateLaneOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this e m v create lane o k response has a 5xx status code
func (o *EMVCreateLaneOK) IsServerError() bool {
	return false
}

// IsCode returns true when this e m v create lane o k response a status code equal to that given
func (o *EMVCreateLaneOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the e m v create lane o k response
func (o *EMVCreateLaneOK) Code() int {
	return 200
}

func (o *EMVCreateLaneOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /PaymentGateway/EMV/TriPosLanes][%d] eMVCreateLaneOK %s", 200, payload)
}

func (o *EMVCreateLaneOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /PaymentGateway/EMV/TriPosLanes][%d] eMVCreateLaneOK %s", 200, payload)
}

func (o *EMVCreateLaneOK) GetPayload() *models.Lane {
	return o.Payload
}

func (o *EMVCreateLaneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Lane)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEMVCreateLaneDefault creates a EMVCreateLaneDefault with default headers values
func NewEMVCreateLaneDefault(code int) *EMVCreateLaneDefault {
	return &EMVCreateLaneDefault{
		_statusCode: code,
	}
}

/*
EMVCreateLaneDefault describes a response with status code -1, with default header values.

Error
*/
type EMVCreateLaneDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this e m v create lane default response has a 2xx status code
func (o *EMVCreateLaneDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this e m v create lane default response has a 3xx status code
func (o *EMVCreateLaneDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this e m v create lane default response has a 4xx status code
func (o *EMVCreateLaneDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this e m v create lane default response has a 5xx status code
func (o *EMVCreateLaneDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this e m v create lane default response a status code equal to that given
func (o *EMVCreateLaneDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the e m v create lane default response
func (o *EMVCreateLaneDefault) Code() int {
	return o._statusCode
}

func (o *EMVCreateLaneDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /PaymentGateway/EMV/TriPosLanes][%d] EMV_CreateLane default %s", o._statusCode, payload)
}

func (o *EMVCreateLaneDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /PaymentGateway/EMV/TriPosLanes][%d] EMV_CreateLane default %s", o._statusCode, payload)
}

func (o *EMVCreateLaneDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *EMVCreateLaneDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
