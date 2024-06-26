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

// DocumentCategoriesGetReader is a Reader for the DocumentCategoriesGet structure.
type DocumentCategoriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocumentCategoriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDocumentCategoriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDocumentCategoriesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDocumentCategoriesGetOK creates a DocumentCategoriesGetOK with default headers values
func NewDocumentCategoriesGetOK() *DocumentCategoriesGetOK {
	return &DocumentCategoriesGetOK{}
}

/*
DocumentCategoriesGetOK describes a response with status code 200, with default header values.

OK
*/
type DocumentCategoriesGetOK struct {
	Payload *models.DocumentCategory
}

// IsSuccess returns true when this document categories get o k response has a 2xx status code
func (o *DocumentCategoriesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this document categories get o k response has a 3xx status code
func (o *DocumentCategoriesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this document categories get o k response has a 4xx status code
func (o *DocumentCategoriesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this document categories get o k response has a 5xx status code
func (o *DocumentCategoriesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this document categories get o k response a status code equal to that given
func (o *DocumentCategoriesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the document categories get o k response
func (o *DocumentCategoriesGetOK) Code() int {
	return 200
}

func (o *DocumentCategoriesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/DocumentCategories/{id}][%d] documentCategoriesGetOK %s", 200, payload)
}

func (o *DocumentCategoriesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/DocumentCategories/{id}][%d] documentCategoriesGetOK %s", 200, payload)
}

func (o *DocumentCategoriesGetOK) GetPayload() *models.DocumentCategory {
	return o.Payload
}

func (o *DocumentCategoriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocumentCategoriesGetDefault creates a DocumentCategoriesGetDefault with default headers values
func NewDocumentCategoriesGetDefault(code int) *DocumentCategoriesGetDefault {
	return &DocumentCategoriesGetDefault{
		_statusCode: code,
	}
}

/*
DocumentCategoriesGetDefault describes a response with status code -1, with default header values.

Error
*/
type DocumentCategoriesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this document categories get default response has a 2xx status code
func (o *DocumentCategoriesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this document categories get default response has a 3xx status code
func (o *DocumentCategoriesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this document categories get default response has a 4xx status code
func (o *DocumentCategoriesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this document categories get default response has a 5xx status code
func (o *DocumentCategoriesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this document categories get default response a status code equal to that given
func (o *DocumentCategoriesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the document categories get default response
func (o *DocumentCategoriesGetDefault) Code() int {
	return o._statusCode
}

func (o *DocumentCategoriesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/DocumentCategories/{id}][%d] DocumentCategories_Get default %s", o._statusCode, payload)
}

func (o *DocumentCategoriesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/DocumentCategories/{id}][%d] DocumentCategories_Get default %s", o._statusCode, payload)
}

func (o *DocumentCategoriesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DocumentCategoriesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
