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

// ConstituentInactivesDeleteReader is a Reader for the ConstituentInactivesDelete structure.
type ConstituentInactivesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentInactivesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewConstituentInactivesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConstituentInactivesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConstituentInactivesDeleteNoContent creates a ConstituentInactivesDeleteNoContent with default headers values
func NewConstituentInactivesDeleteNoContent() *ConstituentInactivesDeleteNoContent {
	return &ConstituentInactivesDeleteNoContent{}
}

/*
ConstituentInactivesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ConstituentInactivesDeleteNoContent struct {
}

// IsSuccess returns true when this constituent inactives delete no content response has a 2xx status code
func (o *ConstituentInactivesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent inactives delete no content response has a 3xx status code
func (o *ConstituentInactivesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent inactives delete no content response has a 4xx status code
func (o *ConstituentInactivesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent inactives delete no content response has a 5xx status code
func (o *ConstituentInactivesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent inactives delete no content response a status code equal to that given
func (o *ConstituentInactivesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the constituent inactives delete no content response
func (o *ConstituentInactivesDeleteNoContent) Code() int {
	return 204
}

func (o *ConstituentInactivesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentInactives/{id}][%d] constituentInactivesDeleteNoContent", 204)
}

func (o *ConstituentInactivesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentInactives/{id}][%d] constituentInactivesDeleteNoContent", 204)
}

func (o *ConstituentInactivesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConstituentInactivesDeleteDefault creates a ConstituentInactivesDeleteDefault with default headers values
func NewConstituentInactivesDeleteDefault(code int) *ConstituentInactivesDeleteDefault {
	return &ConstituentInactivesDeleteDefault{
		_statusCode: code,
	}
}

/*
ConstituentInactivesDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ConstituentInactivesDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this constituent inactives delete default response has a 2xx status code
func (o *ConstituentInactivesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this constituent inactives delete default response has a 3xx status code
func (o *ConstituentInactivesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this constituent inactives delete default response has a 4xx status code
func (o *ConstituentInactivesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this constituent inactives delete default response has a 5xx status code
func (o *ConstituentInactivesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this constituent inactives delete default response a status code equal to that given
func (o *ConstituentInactivesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the constituent inactives delete default response
func (o *ConstituentInactivesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ConstituentInactivesDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentInactives/{id}][%d] ConstituentInactives_Delete default %s", o._statusCode, payload)
}

func (o *ConstituentInactivesDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentInactives/{id}][%d] ConstituentInactives_Delete default %s", o._statusCode, payload)
}

func (o *ConstituentInactivesDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ConstituentInactivesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
