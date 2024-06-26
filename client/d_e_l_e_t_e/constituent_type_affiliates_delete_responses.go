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

// ConstituentTypeAffiliatesDeleteReader is a Reader for the ConstituentTypeAffiliatesDelete structure.
type ConstituentTypeAffiliatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentTypeAffiliatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewConstituentTypeAffiliatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConstituentTypeAffiliatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConstituentTypeAffiliatesDeleteNoContent creates a ConstituentTypeAffiliatesDeleteNoContent with default headers values
func NewConstituentTypeAffiliatesDeleteNoContent() *ConstituentTypeAffiliatesDeleteNoContent {
	return &ConstituentTypeAffiliatesDeleteNoContent{}
}

/*
ConstituentTypeAffiliatesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ConstituentTypeAffiliatesDeleteNoContent struct {
}

// IsSuccess returns true when this constituent type affiliates delete no content response has a 2xx status code
func (o *ConstituentTypeAffiliatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent type affiliates delete no content response has a 3xx status code
func (o *ConstituentTypeAffiliatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent type affiliates delete no content response has a 4xx status code
func (o *ConstituentTypeAffiliatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent type affiliates delete no content response has a 5xx status code
func (o *ConstituentTypeAffiliatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent type affiliates delete no content response a status code equal to that given
func (o *ConstituentTypeAffiliatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the constituent type affiliates delete no content response
func (o *ConstituentTypeAffiliatesDeleteNoContent) Code() int {
	return 204
}

func (o *ConstituentTypeAffiliatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentTypeAffiliates/{id}][%d] constituentTypeAffiliatesDeleteNoContent", 204)
}

func (o *ConstituentTypeAffiliatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentTypeAffiliates/{id}][%d] constituentTypeAffiliatesDeleteNoContent", 204)
}

func (o *ConstituentTypeAffiliatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConstituentTypeAffiliatesDeleteDefault creates a ConstituentTypeAffiliatesDeleteDefault with default headers values
func NewConstituentTypeAffiliatesDeleteDefault(code int) *ConstituentTypeAffiliatesDeleteDefault {
	return &ConstituentTypeAffiliatesDeleteDefault{
		_statusCode: code,
	}
}

/*
ConstituentTypeAffiliatesDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ConstituentTypeAffiliatesDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this constituent type affiliates delete default response has a 2xx status code
func (o *ConstituentTypeAffiliatesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this constituent type affiliates delete default response has a 3xx status code
func (o *ConstituentTypeAffiliatesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this constituent type affiliates delete default response has a 4xx status code
func (o *ConstituentTypeAffiliatesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this constituent type affiliates delete default response has a 5xx status code
func (o *ConstituentTypeAffiliatesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this constituent type affiliates delete default response a status code equal to that given
func (o *ConstituentTypeAffiliatesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the constituent type affiliates delete default response
func (o *ConstituentTypeAffiliatesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ConstituentTypeAffiliatesDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentTypeAffiliates/{id}][%d] ConstituentTypeAffiliates_Delete default %s", o._statusCode, payload)
}

func (o *ConstituentTypeAffiliatesDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentTypeAffiliates/{id}][%d] ConstituentTypeAffiliates_Delete default %s", o._statusCode, payload)
}

func (o *ConstituentTypeAffiliatesDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ConstituentTypeAffiliatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
