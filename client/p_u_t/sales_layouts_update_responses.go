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

// SalesLayoutsUpdateReader is a Reader for the SalesLayoutsUpdate structure.
type SalesLayoutsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesLayoutsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSalesLayoutsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSalesLayoutsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesLayoutsUpdateOK creates a SalesLayoutsUpdateOK with default headers values
func NewSalesLayoutsUpdateOK() *SalesLayoutsUpdateOK {
	return &SalesLayoutsUpdateOK{}
}

/*
SalesLayoutsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type SalesLayoutsUpdateOK struct {
	Payload *models.SalesLayout
}

// IsSuccess returns true when this sales layouts update o k response has a 2xx status code
func (o *SalesLayoutsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sales layouts update o k response has a 3xx status code
func (o *SalesLayoutsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sales layouts update o k response has a 4xx status code
func (o *SalesLayoutsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sales layouts update o k response has a 5xx status code
func (o *SalesLayoutsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sales layouts update o k response a status code equal to that given
func (o *SalesLayoutsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sales layouts update o k response
func (o *SalesLayoutsUpdateOK) Code() int {
	return 200
}

func (o *SalesLayoutsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/SalesLayouts/Setup/{salesLayoutId}][%d] salesLayoutsUpdateOK %s", 200, payload)
}

func (o *SalesLayoutsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/SalesLayouts/Setup/{salesLayoutId}][%d] salesLayoutsUpdateOK %s", 200, payload)
}

func (o *SalesLayoutsUpdateOK) GetPayload() *models.SalesLayout {
	return o.Payload
}

func (o *SalesLayoutsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesLayout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesLayoutsUpdateDefault creates a SalesLayoutsUpdateDefault with default headers values
func NewSalesLayoutsUpdateDefault(code int) *SalesLayoutsUpdateDefault {
	return &SalesLayoutsUpdateDefault{
		_statusCode: code,
	}
}

/*
SalesLayoutsUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type SalesLayoutsUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this sales layouts update default response has a 2xx status code
func (o *SalesLayoutsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sales layouts update default response has a 3xx status code
func (o *SalesLayoutsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sales layouts update default response has a 4xx status code
func (o *SalesLayoutsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sales layouts update default response has a 5xx status code
func (o *SalesLayoutsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sales layouts update default response a status code equal to that given
func (o *SalesLayoutsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sales layouts update default response
func (o *SalesLayoutsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *SalesLayoutsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/SalesLayouts/Setup/{salesLayoutId}][%d] SalesLayouts_Update default %s", o._statusCode, payload)
}

func (o *SalesLayoutsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/SalesLayouts/Setup/{salesLayoutId}][%d] SalesLayouts_Update default %s", o._statusCode, payload)
}

func (o *SalesLayoutsUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SalesLayoutsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
