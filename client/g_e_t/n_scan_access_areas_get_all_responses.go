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

// NScanAccessAreasGetAllReader is a Reader for the NScanAccessAreasGetAll structure.
type NScanAccessAreasGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NScanAccessAreasGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNScanAccessAreasGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNScanAccessAreasGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNScanAccessAreasGetAllOK creates a NScanAccessAreasGetAllOK with default headers values
func NewNScanAccessAreasGetAllOK() *NScanAccessAreasGetAllOK {
	return &NScanAccessAreasGetAllOK{}
}

/*
NScanAccessAreasGetAllOK describes a response with status code 200, with default header values.

OK
*/
type NScanAccessAreasGetAllOK struct {
	Payload []*models.NScanAccessArea
}

// IsSuccess returns true when this n scan access areas get all o k response has a 2xx status code
func (o *NScanAccessAreasGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this n scan access areas get all o k response has a 3xx status code
func (o *NScanAccessAreasGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this n scan access areas get all o k response has a 4xx status code
func (o *NScanAccessAreasGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this n scan access areas get all o k response has a 5xx status code
func (o *NScanAccessAreasGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this n scan access areas get all o k response a status code equal to that given
func (o *NScanAccessAreasGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the n scan access areas get all o k response
func (o *NScanAccessAreasGetAllOK) Code() int {
	return 200
}

func (o *NScanAccessAreasGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/NScanAccessAreas][%d] nScanAccessAreasGetAllOK %s", 200, payload)
}

func (o *NScanAccessAreasGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/NScanAccessAreas][%d] nScanAccessAreasGetAllOK %s", 200, payload)
}

func (o *NScanAccessAreasGetAllOK) GetPayload() []*models.NScanAccessArea {
	return o.Payload
}

func (o *NScanAccessAreasGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNScanAccessAreasGetAllDefault creates a NScanAccessAreasGetAllDefault with default headers values
func NewNScanAccessAreasGetAllDefault(code int) *NScanAccessAreasGetAllDefault {
	return &NScanAccessAreasGetAllDefault{
		_statusCode: code,
	}
}

/*
NScanAccessAreasGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type NScanAccessAreasGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this n scan access areas get all default response has a 2xx status code
func (o *NScanAccessAreasGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this n scan access areas get all default response has a 3xx status code
func (o *NScanAccessAreasGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this n scan access areas get all default response has a 4xx status code
func (o *NScanAccessAreasGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this n scan access areas get all default response has a 5xx status code
func (o *NScanAccessAreasGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this n scan access areas get all default response a status code equal to that given
func (o *NScanAccessAreasGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the n scan access areas get all default response
func (o *NScanAccessAreasGetAllDefault) Code() int {
	return o._statusCode
}

func (o *NScanAccessAreasGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/NScanAccessAreas][%d] NScanAccessAreas_GetAll default %s", o._statusCode, payload)
}

func (o *NScanAccessAreasGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/NScanAccessAreas][%d] NScanAccessAreas_GetAll default %s", o._statusCode, payload)
}

func (o *NScanAccessAreasGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *NScanAccessAreasGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
