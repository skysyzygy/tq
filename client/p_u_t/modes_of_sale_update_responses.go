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

// ModesOfSaleUpdateReader is a Reader for the ModesOfSaleUpdate structure.
type ModesOfSaleUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModesOfSaleUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModesOfSaleUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewModesOfSaleUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModesOfSaleUpdateOK creates a ModesOfSaleUpdateOK with default headers values
func NewModesOfSaleUpdateOK() *ModesOfSaleUpdateOK {
	return &ModesOfSaleUpdateOK{}
}

/*
ModesOfSaleUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ModesOfSaleUpdateOK struct {
	Payload *models.ModeOfSale
}

// IsSuccess returns true when this modes of sale update o k response has a 2xx status code
func (o *ModesOfSaleUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modes of sale update o k response has a 3xx status code
func (o *ModesOfSaleUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modes of sale update o k response has a 4xx status code
func (o *ModesOfSaleUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modes of sale update o k response has a 5xx status code
func (o *ModesOfSaleUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modes of sale update o k response a status code equal to that given
func (o *ModesOfSaleUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the modes of sale update o k response
func (o *ModesOfSaleUpdateOK) Code() int {
	return 200
}

func (o *ModesOfSaleUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/ModesOfSale/{modeOfSaleId}][%d] modesOfSaleUpdateOK %s", 200, payload)
}

func (o *ModesOfSaleUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/ModesOfSale/{modeOfSaleId}][%d] modesOfSaleUpdateOK %s", 200, payload)
}

func (o *ModesOfSaleUpdateOK) GetPayload() *models.ModeOfSale {
	return o.Payload
}

func (o *ModesOfSaleUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModeOfSale)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModesOfSaleUpdateDefault creates a ModesOfSaleUpdateDefault with default headers values
func NewModesOfSaleUpdateDefault(code int) *ModesOfSaleUpdateDefault {
	return &ModesOfSaleUpdateDefault{
		_statusCode: code,
	}
}

/*
ModesOfSaleUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type ModesOfSaleUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this modes of sale update default response has a 2xx status code
func (o *ModesOfSaleUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this modes of sale update default response has a 3xx status code
func (o *ModesOfSaleUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this modes of sale update default response has a 4xx status code
func (o *ModesOfSaleUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this modes of sale update default response has a 5xx status code
func (o *ModesOfSaleUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this modes of sale update default response a status code equal to that given
func (o *ModesOfSaleUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the modes of sale update default response
func (o *ModesOfSaleUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ModesOfSaleUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/ModesOfSale/{modeOfSaleId}][%d] ModesOfSale_Update default %s", o._statusCode, payload)
}

func (o *ModesOfSaleUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/ModesOfSale/{modeOfSaleId}][%d] ModesOfSale_Update default %s", o._statusCode, payload)
}

func (o *ModesOfSaleUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ModesOfSaleUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
