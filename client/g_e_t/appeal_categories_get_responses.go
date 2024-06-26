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

// AppealCategoriesGetReader is a Reader for the AppealCategoriesGet structure.
type AppealCategoriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppealCategoriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppealCategoriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAppealCategoriesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppealCategoriesGetOK creates a AppealCategoriesGetOK with default headers values
func NewAppealCategoriesGetOK() *AppealCategoriesGetOK {
	return &AppealCategoriesGetOK{}
}

/*
AppealCategoriesGetOK describes a response with status code 200, with default header values.

OK
*/
type AppealCategoriesGetOK struct {
	Payload *models.AppealCategory
}

// IsSuccess returns true when this appeal categories get o k response has a 2xx status code
func (o *AppealCategoriesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this appeal categories get o k response has a 3xx status code
func (o *AppealCategoriesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this appeal categories get o k response has a 4xx status code
func (o *AppealCategoriesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this appeal categories get o k response has a 5xx status code
func (o *AppealCategoriesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this appeal categories get o k response a status code equal to that given
func (o *AppealCategoriesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the appeal categories get o k response
func (o *AppealCategoriesGetOK) Code() int {
	return 200
}

func (o *AppealCategoriesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories/{id}][%d] appealCategoriesGetOK %s", 200, payload)
}

func (o *AppealCategoriesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories/{id}][%d] appealCategoriesGetOK %s", 200, payload)
}

func (o *AppealCategoriesGetOK) GetPayload() *models.AppealCategory {
	return o.Payload
}

func (o *AppealCategoriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppealCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppealCategoriesGetDefault creates a AppealCategoriesGetDefault with default headers values
func NewAppealCategoriesGetDefault(code int) *AppealCategoriesGetDefault {
	return &AppealCategoriesGetDefault{
		_statusCode: code,
	}
}

/*
AppealCategoriesGetDefault describes a response with status code -1, with default header values.

Error
*/
type AppealCategoriesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this appeal categories get default response has a 2xx status code
func (o *AppealCategoriesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this appeal categories get default response has a 3xx status code
func (o *AppealCategoriesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this appeal categories get default response has a 4xx status code
func (o *AppealCategoriesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this appeal categories get default response has a 5xx status code
func (o *AppealCategoriesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this appeal categories get default response a status code equal to that given
func (o *AppealCategoriesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the appeal categories get default response
func (o *AppealCategoriesGetDefault) Code() int {
	return o._statusCode
}

func (o *AppealCategoriesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories/{id}][%d] AppealCategories_Get default %s", o._statusCode, payload)
}

func (o *AppealCategoriesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories/{id}][%d] AppealCategories_Get default %s", o._statusCode, payload)
}

func (o *AppealCategoriesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AppealCategoriesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
