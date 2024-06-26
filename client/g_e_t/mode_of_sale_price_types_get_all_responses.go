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

// ModeOfSalePriceTypesGetAllReader is a Reader for the ModeOfSalePriceTypesGetAll structure.
type ModeOfSalePriceTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModeOfSalePriceTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModeOfSalePriceTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewModeOfSalePriceTypesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModeOfSalePriceTypesGetAllOK creates a ModeOfSalePriceTypesGetAllOK with default headers values
func NewModeOfSalePriceTypesGetAllOK() *ModeOfSalePriceTypesGetAllOK {
	return &ModeOfSalePriceTypesGetAllOK{}
}

/*
ModeOfSalePriceTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ModeOfSalePriceTypesGetAllOK struct {
	Payload []*models.ModeOfSalePriceType
}

// IsSuccess returns true when this mode of sale price types get all o k response has a 2xx status code
func (o *ModeOfSalePriceTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mode of sale price types get all o k response has a 3xx status code
func (o *ModeOfSalePriceTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mode of sale price types get all o k response has a 4xx status code
func (o *ModeOfSalePriceTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mode of sale price types get all o k response has a 5xx status code
func (o *ModeOfSalePriceTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mode of sale price types get all o k response a status code equal to that given
func (o *ModeOfSalePriceTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mode of sale price types get all o k response
func (o *ModeOfSalePriceTypesGetAllOK) Code() int {
	return 200
}

func (o *ModeOfSalePriceTypesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/ModeOfSalePriceTypes][%d] modeOfSalePriceTypesGetAllOK %s", 200, payload)
}

func (o *ModeOfSalePriceTypesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/ModeOfSalePriceTypes][%d] modeOfSalePriceTypesGetAllOK %s", 200, payload)
}

func (o *ModeOfSalePriceTypesGetAllOK) GetPayload() []*models.ModeOfSalePriceType {
	return o.Payload
}

func (o *ModeOfSalePriceTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModeOfSalePriceTypesGetAllDefault creates a ModeOfSalePriceTypesGetAllDefault with default headers values
func NewModeOfSalePriceTypesGetAllDefault(code int) *ModeOfSalePriceTypesGetAllDefault {
	return &ModeOfSalePriceTypesGetAllDefault{
		_statusCode: code,
	}
}

/*
ModeOfSalePriceTypesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type ModeOfSalePriceTypesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this mode of sale price types get all default response has a 2xx status code
func (o *ModeOfSalePriceTypesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this mode of sale price types get all default response has a 3xx status code
func (o *ModeOfSalePriceTypesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this mode of sale price types get all default response has a 4xx status code
func (o *ModeOfSalePriceTypesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this mode of sale price types get all default response has a 5xx status code
func (o *ModeOfSalePriceTypesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this mode of sale price types get all default response a status code equal to that given
func (o *ModeOfSalePriceTypesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the mode of sale price types get all default response
func (o *ModeOfSalePriceTypesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *ModeOfSalePriceTypesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/ModeOfSalePriceTypes][%d] ModeOfSalePriceTypes_GetAll default %s", o._statusCode, payload)
}

func (o *ModeOfSalePriceTypesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/ModeOfSalePriceTypes][%d] ModeOfSalePriceTypes_GetAll default %s", o._statusCode, payload)
}

func (o *ModeOfSalePriceTypesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ModeOfSalePriceTypesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
