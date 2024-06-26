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

// TemplateCategoriesUpdateReader is a Reader for the TemplateCategoriesUpdate structure.
type TemplateCategoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplateCategoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplateCategoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTemplateCategoriesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTemplateCategoriesUpdateOK creates a TemplateCategoriesUpdateOK with default headers values
func NewTemplateCategoriesUpdateOK() *TemplateCategoriesUpdateOK {
	return &TemplateCategoriesUpdateOK{}
}

/*
TemplateCategoriesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type TemplateCategoriesUpdateOK struct {
	Payload *models.TemplateCategory
}

// IsSuccess returns true when this template categories update o k response has a 2xx status code
func (o *TemplateCategoriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this template categories update o k response has a 3xx status code
func (o *TemplateCategoriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this template categories update o k response has a 4xx status code
func (o *TemplateCategoriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this template categories update o k response has a 5xx status code
func (o *TemplateCategoriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this template categories update o k response a status code equal to that given
func (o *TemplateCategoriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the template categories update o k response
func (o *TemplateCategoriesUpdateOK) Code() int {
	return 200
}

func (o *TemplateCategoriesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/TemplateCategories/{id}][%d] templateCategoriesUpdateOK %s", 200, payload)
}

func (o *TemplateCategoriesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/TemplateCategories/{id}][%d] templateCategoriesUpdateOK %s", 200, payload)
}

func (o *TemplateCategoriesUpdateOK) GetPayload() *models.TemplateCategory {
	return o.Payload
}

func (o *TemplateCategoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplateCategoriesUpdateDefault creates a TemplateCategoriesUpdateDefault with default headers values
func NewTemplateCategoriesUpdateDefault(code int) *TemplateCategoriesUpdateDefault {
	return &TemplateCategoriesUpdateDefault{
		_statusCode: code,
	}
}

/*
TemplateCategoriesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type TemplateCategoriesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this template categories update default response has a 2xx status code
func (o *TemplateCategoriesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this template categories update default response has a 3xx status code
func (o *TemplateCategoriesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this template categories update default response has a 4xx status code
func (o *TemplateCategoriesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this template categories update default response has a 5xx status code
func (o *TemplateCategoriesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this template categories update default response a status code equal to that given
func (o *TemplateCategoriesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the template categories update default response
func (o *TemplateCategoriesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TemplateCategoriesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/TemplateCategories/{id}][%d] TemplateCategories_Update default %s", o._statusCode, payload)
}

func (o *TemplateCategoriesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/TemplateCategories/{id}][%d] TemplateCategories_Update default %s", o._statusCode, payload)
}

func (o *TemplateCategoriesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *TemplateCategoriesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
