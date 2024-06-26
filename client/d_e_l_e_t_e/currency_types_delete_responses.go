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

// CurrencyTypesDeleteReader is a Reader for the CurrencyTypesDelete structure.
type CurrencyTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CurrencyTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCurrencyTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCurrencyTypesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCurrencyTypesDeleteNoContent creates a CurrencyTypesDeleteNoContent with default headers values
func NewCurrencyTypesDeleteNoContent() *CurrencyTypesDeleteNoContent {
	return &CurrencyTypesDeleteNoContent{}
}

/*
CurrencyTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type CurrencyTypesDeleteNoContent struct {
}

// IsSuccess returns true when this currency types delete no content response has a 2xx status code
func (o *CurrencyTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this currency types delete no content response has a 3xx status code
func (o *CurrencyTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this currency types delete no content response has a 4xx status code
func (o *CurrencyTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this currency types delete no content response has a 5xx status code
func (o *CurrencyTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this currency types delete no content response a status code equal to that given
func (o *CurrencyTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the currency types delete no content response
func (o *CurrencyTypesDeleteNoContent) Code() int {
	return 204
}

func (o *CurrencyTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/CurrencyTypes/{id}][%d] currencyTypesDeleteNoContent", 204)
}

func (o *CurrencyTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/CurrencyTypes/{id}][%d] currencyTypesDeleteNoContent", 204)
}

func (o *CurrencyTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCurrencyTypesDeleteDefault creates a CurrencyTypesDeleteDefault with default headers values
func NewCurrencyTypesDeleteDefault(code int) *CurrencyTypesDeleteDefault {
	return &CurrencyTypesDeleteDefault{
		_statusCode: code,
	}
}

/*
CurrencyTypesDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type CurrencyTypesDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this currency types delete default response has a 2xx status code
func (o *CurrencyTypesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this currency types delete default response has a 3xx status code
func (o *CurrencyTypesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this currency types delete default response has a 4xx status code
func (o *CurrencyTypesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this currency types delete default response has a 5xx status code
func (o *CurrencyTypesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this currency types delete default response a status code equal to that given
func (o *CurrencyTypesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the currency types delete default response
func (o *CurrencyTypesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *CurrencyTypesDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/CurrencyTypes/{id}][%d] CurrencyTypes_Delete default %s", o._statusCode, payload)
}

func (o *CurrencyTypesDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/CurrencyTypes/{id}][%d] CurrencyTypes_Delete default %s", o._statusCode, payload)
}

func (o *CurrencyTypesDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CurrencyTypesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
