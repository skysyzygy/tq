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

// ContactPointCategoriesGetReader is a Reader for the ContactPointCategoriesGet structure.
type ContactPointCategoriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPointCategoriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactPointCategoriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ContactPointCategories/{id}] ContactPointCategories_Get", response, response.Code())
	}
}

// NewContactPointCategoriesGetOK creates a ContactPointCategoriesGetOK with default headers values
func NewContactPointCategoriesGetOK() *ContactPointCategoriesGetOK {
	return &ContactPointCategoriesGetOK{}
}

/*
ContactPointCategoriesGetOK describes a response with status code 200, with default header values.

OK
*/
type ContactPointCategoriesGetOK struct {
	Payload *models.ContactPointCategory
}

// IsSuccess returns true when this contact point categories get o k response has a 2xx status code
func (o *ContactPointCategoriesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact point categories get o k response has a 3xx status code
func (o *ContactPointCategoriesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact point categories get o k response has a 4xx status code
func (o *ContactPointCategoriesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact point categories get o k response has a 5xx status code
func (o *ContactPointCategoriesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact point categories get o k response a status code equal to that given
func (o *ContactPointCategoriesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact point categories get o k response
func (o *ContactPointCategoriesGetOK) Code() int {
	return 200
}

func (o *ContactPointCategoriesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ContactPointCategories/{id}][%d] contactPointCategoriesGetOK  %+v", 200, o.Payload)
}

func (o *ContactPointCategoriesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ContactPointCategories/{id}][%d] contactPointCategoriesGetOK  %+v", 200, o.Payload)
}

func (o *ContactPointCategoriesGetOK) GetPayload() *models.ContactPointCategory {
	return o.Payload
}

func (o *ContactPointCategoriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactPointCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
