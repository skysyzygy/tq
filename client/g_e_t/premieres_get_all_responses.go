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

// PremieresGetAllReader is a Reader for the PremieresGetAll structure.
type PremieresGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PremieresGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPremieresGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPremieresGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPremieresGetAllOK creates a PremieresGetAllOK with default headers values
func NewPremieresGetAllOK() *PremieresGetAllOK {
	return &PremieresGetAllOK{}
}

/*
PremieresGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PremieresGetAllOK struct {
	Payload []*models.Premiere
}

// IsSuccess returns true when this premieres get all o k response has a 2xx status code
func (o *PremieresGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this premieres get all o k response has a 3xx status code
func (o *PremieresGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this premieres get all o k response has a 4xx status code
func (o *PremieresGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this premieres get all o k response has a 5xx status code
func (o *PremieresGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this premieres get all o k response a status code equal to that given
func (o *PremieresGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the premieres get all o k response
func (o *PremieresGetAllOK) Code() int {
	return 200
}

func (o *PremieresGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Premieres][%d] premieresGetAllOK %s", 200, payload)
}

func (o *PremieresGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Premieres][%d] premieresGetAllOK %s", 200, payload)
}

func (o *PremieresGetAllOK) GetPayload() []*models.Premiere {
	return o.Payload
}

func (o *PremieresGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPremieresGetAllDefault creates a PremieresGetAllDefault with default headers values
func NewPremieresGetAllDefault(code int) *PremieresGetAllDefault {
	return &PremieresGetAllDefault{
		_statusCode: code,
	}
}

/*
PremieresGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type PremieresGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this premieres get all default response has a 2xx status code
func (o *PremieresGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this premieres get all default response has a 3xx status code
func (o *PremieresGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this premieres get all default response has a 4xx status code
func (o *PremieresGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this premieres get all default response has a 5xx status code
func (o *PremieresGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this premieres get all default response a status code equal to that given
func (o *PremieresGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the premieres get all default response
func (o *PremieresGetAllDefault) Code() int {
	return o._statusCode
}

func (o *PremieresGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Premieres][%d] Premieres_GetAll default %s", o._statusCode, payload)
}

func (o *PremieresGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Premieres][%d] Premieres_GetAll default %s", o._statusCode, payload)
}

func (o *PremieresGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PremieresGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
