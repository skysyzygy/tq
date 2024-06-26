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

// HoldCodeCategoriesGetReader is a Reader for the HoldCodeCategoriesGet structure.
type HoldCodeCategoriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HoldCodeCategoriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHoldCodeCategoriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHoldCodeCategoriesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHoldCodeCategoriesGetOK creates a HoldCodeCategoriesGetOK with default headers values
func NewHoldCodeCategoriesGetOK() *HoldCodeCategoriesGetOK {
	return &HoldCodeCategoriesGetOK{}
}

/*
HoldCodeCategoriesGetOK describes a response with status code 200, with default header values.

OK
*/
type HoldCodeCategoriesGetOK struct {
	Payload *models.HoldCodeCategory
}

// IsSuccess returns true when this hold code categories get o k response has a 2xx status code
func (o *HoldCodeCategoriesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hold code categories get o k response has a 3xx status code
func (o *HoldCodeCategoriesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hold code categories get o k response has a 4xx status code
func (o *HoldCodeCategoriesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hold code categories get o k response has a 5xx status code
func (o *HoldCodeCategoriesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hold code categories get o k response a status code equal to that given
func (o *HoldCodeCategoriesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the hold code categories get o k response
func (o *HoldCodeCategoriesGetOK) Code() int {
	return 200
}

func (o *HoldCodeCategoriesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/HoldCodeCategories/{id}][%d] holdCodeCategoriesGetOK %s", 200, payload)
}

func (o *HoldCodeCategoriesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/HoldCodeCategories/{id}][%d] holdCodeCategoriesGetOK %s", 200, payload)
}

func (o *HoldCodeCategoriesGetOK) GetPayload() *models.HoldCodeCategory {
	return o.Payload
}

func (o *HoldCodeCategoriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HoldCodeCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHoldCodeCategoriesGetDefault creates a HoldCodeCategoriesGetDefault with default headers values
func NewHoldCodeCategoriesGetDefault(code int) *HoldCodeCategoriesGetDefault {
	return &HoldCodeCategoriesGetDefault{
		_statusCode: code,
	}
}

/*
HoldCodeCategoriesGetDefault describes a response with status code -1, with default header values.

Error
*/
type HoldCodeCategoriesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this hold code categories get default response has a 2xx status code
func (o *HoldCodeCategoriesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hold code categories get default response has a 3xx status code
func (o *HoldCodeCategoriesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hold code categories get default response has a 4xx status code
func (o *HoldCodeCategoriesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hold code categories get default response has a 5xx status code
func (o *HoldCodeCategoriesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hold code categories get default response a status code equal to that given
func (o *HoldCodeCategoriesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the hold code categories get default response
func (o *HoldCodeCategoriesGetDefault) Code() int {
	return o._statusCode
}

func (o *HoldCodeCategoriesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/HoldCodeCategories/{id}][%d] HoldCodeCategories_Get default %s", o._statusCode, payload)
}

func (o *HoldCodeCategoriesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/HoldCodeCategories/{id}][%d] HoldCodeCategories_Get default %s", o._statusCode, payload)
}

func (o *HoldCodeCategoriesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *HoldCodeCategoriesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
