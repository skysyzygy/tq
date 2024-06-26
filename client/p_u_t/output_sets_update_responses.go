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

// OutputSetsUpdateReader is a Reader for the OutputSetsUpdate structure.
type OutputSetsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OutputSetsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOutputSetsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOutputSetsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOutputSetsUpdateOK creates a OutputSetsUpdateOK with default headers values
func NewOutputSetsUpdateOK() *OutputSetsUpdateOK {
	return &OutputSetsUpdateOK{}
}

/*
OutputSetsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type OutputSetsUpdateOK struct {
	Payload *models.OutputSet
}

// IsSuccess returns true when this output sets update o k response has a 2xx status code
func (o *OutputSetsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this output sets update o k response has a 3xx status code
func (o *OutputSetsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this output sets update o k response has a 4xx status code
func (o *OutputSetsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this output sets update o k response has a 5xx status code
func (o *OutputSetsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this output sets update o k response a status code equal to that given
func (o *OutputSetsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the output sets update o k response
func (o *OutputSetsUpdateOK) Code() int {
	return 200
}

func (o *OutputSetsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Reporting/OutputSets/{outputSetId}][%d] outputSetsUpdateOK %s", 200, payload)
}

func (o *OutputSetsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Reporting/OutputSets/{outputSetId}][%d] outputSetsUpdateOK %s", 200, payload)
}

func (o *OutputSetsUpdateOK) GetPayload() *models.OutputSet {
	return o.Payload
}

func (o *OutputSetsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutputSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutputSetsUpdateDefault creates a OutputSetsUpdateDefault with default headers values
func NewOutputSetsUpdateDefault(code int) *OutputSetsUpdateDefault {
	return &OutputSetsUpdateDefault{
		_statusCode: code,
	}
}

/*
OutputSetsUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type OutputSetsUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this output sets update default response has a 2xx status code
func (o *OutputSetsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this output sets update default response has a 3xx status code
func (o *OutputSetsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this output sets update default response has a 4xx status code
func (o *OutputSetsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this output sets update default response has a 5xx status code
func (o *OutputSetsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this output sets update default response a status code equal to that given
func (o *OutputSetsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the output sets update default response
func (o *OutputSetsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *OutputSetsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Reporting/OutputSets/{outputSetId}][%d] OutputSets_Update default %s", o._statusCode, payload)
}

func (o *OutputSetsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Reporting/OutputSets/{outputSetId}][%d] OutputSets_Update default %s", o._statusCode, payload)
}

func (o *OutputSetsUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *OutputSetsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
