// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// AliasTypesDeleteReader is a Reader for the AliasTypesDelete structure.
type AliasTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AliasTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAliasTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAliasTypesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAliasTypesDeleteNoContent creates a AliasTypesDeleteNoContent with default headers values
func NewAliasTypesDeleteNoContent() *AliasTypesDeleteNoContent {
	return &AliasTypesDeleteNoContent{}
}

/*
AliasTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type AliasTypesDeleteNoContent struct {
}

// IsSuccess returns true when this alias types delete no content response has a 2xx status code
func (o *AliasTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alias types delete no content response has a 3xx status code
func (o *AliasTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alias types delete no content response has a 4xx status code
func (o *AliasTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this alias types delete no content response has a 5xx status code
func (o *AliasTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this alias types delete no content response a status code equal to that given
func (o *AliasTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the alias types delete no content response
func (o *AliasTypesDeleteNoContent) Code() int {
	return 204
}

func (o *AliasTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/AliasTypes/{id}][%d] aliasTypesDeleteNoContent", 204)
}

func (o *AliasTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/AliasTypes/{id}][%d] aliasTypesDeleteNoContent", 204)
}

func (o *AliasTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAliasTypesDeleteDefault creates a AliasTypesDeleteDefault with default headers values
func NewAliasTypesDeleteDefault(code int) *AliasTypesDeleteDefault {
	return &AliasTypesDeleteDefault{
		_statusCode: code,
	}
}

/*
AliasTypesDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type AliasTypesDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this alias types delete default response has a 2xx status code
func (o *AliasTypesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this alias types delete default response has a 3xx status code
func (o *AliasTypesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this alias types delete default response has a 4xx status code
func (o *AliasTypesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this alias types delete default response has a 5xx status code
func (o *AliasTypesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this alias types delete default response a status code equal to that given
func (o *AliasTypesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the alias types delete default response
func (o *AliasTypesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *AliasTypesDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/AliasTypes/{id}][%d] AliasTypes_Delete default %s", o._statusCode, payload)
}

func (o *AliasTypesDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/AliasTypes/{id}][%d] AliasTypes_Delete default %s", o._statusCode, payload)
}

func (o *AliasTypesDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AliasTypesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
