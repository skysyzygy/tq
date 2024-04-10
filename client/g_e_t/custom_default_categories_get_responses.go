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

// CustomDefaultCategoriesGetReader is a Reader for the CustomDefaultCategoriesGet structure.
type CustomDefaultCategoriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomDefaultCategoriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomDefaultCategoriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/CustomDefaultCategories/{id}] CustomDefaultCategories_Get", response, response.Code())
	}
}

// NewCustomDefaultCategoriesGetOK creates a CustomDefaultCategoriesGetOK with default headers values
func NewCustomDefaultCategoriesGetOK() *CustomDefaultCategoriesGetOK {
	return &CustomDefaultCategoriesGetOK{}
}

/*
CustomDefaultCategoriesGetOK describes a response with status code 200, with default header values.

OK
*/
type CustomDefaultCategoriesGetOK struct {
	Payload *models.CustomDefaultCategory
}

// IsSuccess returns true when this custom default categories get o k response has a 2xx status code
func (o *CustomDefaultCategoriesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom default categories get o k response has a 3xx status code
func (o *CustomDefaultCategoriesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom default categories get o k response has a 4xx status code
func (o *CustomDefaultCategoriesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom default categories get o k response has a 5xx status code
func (o *CustomDefaultCategoriesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom default categories get o k response a status code equal to that given
func (o *CustomDefaultCategoriesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the custom default categories get o k response
func (o *CustomDefaultCategoriesGetOK) Code() int {
	return 200
}

func (o *CustomDefaultCategoriesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/CustomDefaultCategories/{id}][%d] customDefaultCategoriesGetOK  %+v", 200, o.Payload)
}

func (o *CustomDefaultCategoriesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/CustomDefaultCategories/{id}][%d] customDefaultCategoriesGetOK  %+v", 200, o.Payload)
}

func (o *CustomDefaultCategoriesGetOK) GetPayload() *models.CustomDefaultCategory {
	return o.Payload
}

func (o *CustomDefaultCategoriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomDefaultCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
