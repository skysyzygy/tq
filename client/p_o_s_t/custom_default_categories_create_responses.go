// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// CustomDefaultCategoriesCreateReader is a Reader for the CustomDefaultCategoriesCreate structure.
type CustomDefaultCategoriesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomDefaultCategoriesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomDefaultCategoriesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/CustomDefaultCategories] CustomDefaultCategories_Create", response, response.Code())
	}
}

// NewCustomDefaultCategoriesCreateOK creates a CustomDefaultCategoriesCreateOK with default headers values
func NewCustomDefaultCategoriesCreateOK() *CustomDefaultCategoriesCreateOK {
	return &CustomDefaultCategoriesCreateOK{}
}

/*
CustomDefaultCategoriesCreateOK describes a response with status code 200, with default header values.

OK
*/
type CustomDefaultCategoriesCreateOK struct {
	Payload *models.CustomDefaultCategory
}

// IsSuccess returns true when this custom default categories create o k response has a 2xx status code
func (o *CustomDefaultCategoriesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom default categories create o k response has a 3xx status code
func (o *CustomDefaultCategoriesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom default categories create o k response has a 4xx status code
func (o *CustomDefaultCategoriesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom default categories create o k response has a 5xx status code
func (o *CustomDefaultCategoriesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom default categories create o k response a status code equal to that given
func (o *CustomDefaultCategoriesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the custom default categories create o k response
func (o *CustomDefaultCategoriesCreateOK) Code() int {
	return 200
}

func (o *CustomDefaultCategoriesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/CustomDefaultCategories][%d] customDefaultCategoriesCreateOK  %+v", 200, o.Payload)
}

func (o *CustomDefaultCategoriesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/CustomDefaultCategories][%d] customDefaultCategoriesCreateOK  %+v", 200, o.Payload)
}

func (o *CustomDefaultCategoriesCreateOK) GetPayload() *models.CustomDefaultCategory {
	return o.Payload
}

func (o *CustomDefaultCategoriesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomDefaultCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
