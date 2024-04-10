// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ListCategoriesGetReader is a Reader for the ListCategoriesGet structure.
type ListCategoriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCategoriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCategoriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ListCategories/{id}] ListCategories_Get", response, response.Code())
	}
}

// NewListCategoriesGetOK creates a ListCategoriesGetOK with default headers values
func NewListCategoriesGetOK() *ListCategoriesGetOK {
	return &ListCategoriesGetOK{}
}

/*
ListCategoriesGetOK describes a response with status code 200, with default header values.

OK
*/
type ListCategoriesGetOK struct {
	Payload *models.ListCategory
}

// IsSuccess returns true when this list categories get o k response has a 2xx status code
func (o *ListCategoriesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list categories get o k response has a 3xx status code
func (o *ListCategoriesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list categories get o k response has a 4xx status code
func (o *ListCategoriesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list categories get o k response has a 5xx status code
func (o *ListCategoriesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list categories get o k response a status code equal to that given
func (o *ListCategoriesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list categories get o k response
func (o *ListCategoriesGetOK) Code() int {
	return 200
}

func (o *ListCategoriesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ListCategories/{id}][%d] listCategoriesGetOK  %+v", 200, o.Payload)
}

func (o *ListCategoriesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ListCategories/{id}][%d] listCategoriesGetOK  %+v", 200, o.Payload)
}

func (o *ListCategoriesGetOK) GetPayload() *models.ListCategory {
	return o.Payload
}

func (o *ListCategoriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
