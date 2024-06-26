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

// PremieresUpdateReader is a Reader for the PremieresUpdate structure.
type PremieresUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PremieresUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPremieresUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPremieresUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPremieresUpdateOK creates a PremieresUpdateOK with default headers values
func NewPremieresUpdateOK() *PremieresUpdateOK {
	return &PremieresUpdateOK{}
}

/*
PremieresUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PremieresUpdateOK struct {
	Payload *models.Premiere
}

// IsSuccess returns true when this premieres update o k response has a 2xx status code
func (o *PremieresUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this premieres update o k response has a 3xx status code
func (o *PremieresUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this premieres update o k response has a 4xx status code
func (o *PremieresUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this premieres update o k response has a 5xx status code
func (o *PremieresUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this premieres update o k response a status code equal to that given
func (o *PremieresUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the premieres update o k response
func (o *PremieresUpdateOK) Code() int {
	return 200
}

func (o *PremieresUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Premieres/{id}][%d] premieresUpdateOK %s", 200, payload)
}

func (o *PremieresUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Premieres/{id}][%d] premieresUpdateOK %s", 200, payload)
}

func (o *PremieresUpdateOK) GetPayload() *models.Premiere {
	return o.Payload
}

func (o *PremieresUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Premiere)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPremieresUpdateDefault creates a PremieresUpdateDefault with default headers values
func NewPremieresUpdateDefault(code int) *PremieresUpdateDefault {
	return &PremieresUpdateDefault{
		_statusCode: code,
	}
}

/*
PremieresUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type PremieresUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this premieres update default response has a 2xx status code
func (o *PremieresUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this premieres update default response has a 3xx status code
func (o *PremieresUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this premieres update default response has a 4xx status code
func (o *PremieresUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this premieres update default response has a 5xx status code
func (o *PremieresUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this premieres update default response a status code equal to that given
func (o *PremieresUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the premieres update default response
func (o *PremieresUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PremieresUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Premieres/{id}][%d] Premieres_Update default %s", o._statusCode, payload)
}

func (o *PremieresUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Premieres/{id}][%d] Premieres_Update default %s", o._statusCode, payload)
}

func (o *PremieresUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PremieresUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
