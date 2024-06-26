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

// AliasTypesUpdateReader is a Reader for the AliasTypesUpdate structure.
type AliasTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AliasTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAliasTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAliasTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAliasTypesUpdateOK creates a AliasTypesUpdateOK with default headers values
func NewAliasTypesUpdateOK() *AliasTypesUpdateOK {
	return &AliasTypesUpdateOK{}
}

/*
AliasTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type AliasTypesUpdateOK struct {
	Payload *models.AliasType
}

// IsSuccess returns true when this alias types update o k response has a 2xx status code
func (o *AliasTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alias types update o k response has a 3xx status code
func (o *AliasTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alias types update o k response has a 4xx status code
func (o *AliasTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alias types update o k response has a 5xx status code
func (o *AliasTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alias types update o k response a status code equal to that given
func (o *AliasTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the alias types update o k response
func (o *AliasTypesUpdateOK) Code() int {
	return 200
}

func (o *AliasTypesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/AliasTypes/{id}][%d] aliasTypesUpdateOK %s", 200, payload)
}

func (o *AliasTypesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/AliasTypes/{id}][%d] aliasTypesUpdateOK %s", 200, payload)
}

func (o *AliasTypesUpdateOK) GetPayload() *models.AliasType {
	return o.Payload
}

func (o *AliasTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AliasType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAliasTypesUpdateDefault creates a AliasTypesUpdateDefault with default headers values
func NewAliasTypesUpdateDefault(code int) *AliasTypesUpdateDefault {
	return &AliasTypesUpdateDefault{
		_statusCode: code,
	}
}

/*
AliasTypesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type AliasTypesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this alias types update default response has a 2xx status code
func (o *AliasTypesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this alias types update default response has a 3xx status code
func (o *AliasTypesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this alias types update default response has a 4xx status code
func (o *AliasTypesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this alias types update default response has a 5xx status code
func (o *AliasTypesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this alias types update default response a status code equal to that given
func (o *AliasTypesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the alias types update default response
func (o *AliasTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *AliasTypesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/AliasTypes/{id}][%d] AliasTypes_Update default %s", o._statusCode, payload)
}

func (o *AliasTypesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/AliasTypes/{id}][%d] AliasTypes_Update default %s", o._statusCode, payload)
}

func (o *AliasTypesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AliasTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
