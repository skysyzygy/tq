// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// CartAddSubPackageItemReader is a Reader for the CartAddSubPackageItem structure.
type CartAddSubPackageItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartAddSubPackageItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCartAddSubPackageItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCartAddSubPackageItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCartAddSubPackageItemOK creates a CartAddSubPackageItemOK with default headers values
func NewCartAddSubPackageItemOK() *CartAddSubPackageItemOK {
	return &CartAddSubPackageItemOK{}
}

/*
CartAddSubPackageItemOK describes a response with status code 200, with default header values.

OK
*/
type CartAddSubPackageItemOK struct {
	Payload *models.AddSubPackageItemResponse
}

// IsSuccess returns true when this cart add sub package item o k response has a 2xx status code
func (o *CartAddSubPackageItemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart add sub package item o k response has a 3xx status code
func (o *CartAddSubPackageItemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart add sub package item o k response has a 4xx status code
func (o *CartAddSubPackageItemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart add sub package item o k response has a 5xx status code
func (o *CartAddSubPackageItemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cart add sub package item o k response a status code equal to that given
func (o *CartAddSubPackageItemOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cart add sub package item o k response
func (o *CartAddSubPackageItemOK) Code() int {
	return 200
}

func (o *CartAddSubPackageItemOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Packages/Super][%d] cartAddSubPackageItemOK %s", 200, payload)
}

func (o *CartAddSubPackageItemOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Packages/Super][%d] cartAddSubPackageItemOK %s", 200, payload)
}

func (o *CartAddSubPackageItemOK) GetPayload() *models.AddSubPackageItemResponse {
	return o.Payload
}

func (o *CartAddSubPackageItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddSubPackageItemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCartAddSubPackageItemDefault creates a CartAddSubPackageItemDefault with default headers values
func NewCartAddSubPackageItemDefault(code int) *CartAddSubPackageItemDefault {
	return &CartAddSubPackageItemDefault{
		_statusCode: code,
	}
}

/*
CartAddSubPackageItemDefault describes a response with status code -1, with default header values.

Error
*/
type CartAddSubPackageItemDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this cart add sub package item default response has a 2xx status code
func (o *CartAddSubPackageItemDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cart add sub package item default response has a 3xx status code
func (o *CartAddSubPackageItemDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cart add sub package item default response has a 4xx status code
func (o *CartAddSubPackageItemDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cart add sub package item default response has a 5xx status code
func (o *CartAddSubPackageItemDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cart add sub package item default response a status code equal to that given
func (o *CartAddSubPackageItemDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cart add sub package item default response
func (o *CartAddSubPackageItemDefault) Code() int {
	return o._statusCode
}

func (o *CartAddSubPackageItemDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Packages/Super][%d] Cart_AddSubPackageItem default %s", o._statusCode, payload)
}

func (o *CartAddSubPackageItemDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Packages/Super][%d] Cart_AddSubPackageItem default %s", o._statusCode, payload)
}

func (o *CartAddSubPackageItemDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CartAddSubPackageItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
