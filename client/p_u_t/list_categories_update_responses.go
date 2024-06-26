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

// ListCategoriesUpdateReader is a Reader for the ListCategoriesUpdate structure.
type ListCategoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCategoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCategoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListCategoriesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListCategoriesUpdateOK creates a ListCategoriesUpdateOK with default headers values
func NewListCategoriesUpdateOK() *ListCategoriesUpdateOK {
	return &ListCategoriesUpdateOK{}
}

/*
ListCategoriesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ListCategoriesUpdateOK struct {
	Payload *models.ListCategory
}

// IsSuccess returns true when this list categories update o k response has a 2xx status code
func (o *ListCategoriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list categories update o k response has a 3xx status code
func (o *ListCategoriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list categories update o k response has a 4xx status code
func (o *ListCategoriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list categories update o k response has a 5xx status code
func (o *ListCategoriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list categories update o k response a status code equal to that given
func (o *ListCategoriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list categories update o k response
func (o *ListCategoriesUpdateOK) Code() int {
	return 200
}

func (o *ListCategoriesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ListCategories/{id}][%d] listCategoriesUpdateOK %s", 200, payload)
}

func (o *ListCategoriesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ListCategories/{id}][%d] listCategoriesUpdateOK %s", 200, payload)
}

func (o *ListCategoriesUpdateOK) GetPayload() *models.ListCategory {
	return o.Payload
}

func (o *ListCategoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCategoriesUpdateDefault creates a ListCategoriesUpdateDefault with default headers values
func NewListCategoriesUpdateDefault(code int) *ListCategoriesUpdateDefault {
	return &ListCategoriesUpdateDefault{
		_statusCode: code,
	}
}

/*
ListCategoriesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type ListCategoriesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this list categories update default response has a 2xx status code
func (o *ListCategoriesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list categories update default response has a 3xx status code
func (o *ListCategoriesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list categories update default response has a 4xx status code
func (o *ListCategoriesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list categories update default response has a 5xx status code
func (o *ListCategoriesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list categories update default response a status code equal to that given
func (o *ListCategoriesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list categories update default response
func (o *ListCategoriesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ListCategoriesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ListCategories/{id}][%d] ListCategories_Update default %s", o._statusCode, payload)
}

func (o *ListCategoriesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ListCategories/{id}][%d] ListCategories_Update default %s", o._statusCode, payload)
}

func (o *ListCategoriesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ListCategoriesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
