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

// BookingCategoriesCreateReader is a Reader for the BookingCategoriesCreate structure.
type BookingCategoriesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookingCategoriesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookingCategoriesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/BookingCategories] BookingCategories_Create", response, response.Code())
	}
}

// NewBookingCategoriesCreateOK creates a BookingCategoriesCreateOK with default headers values
func NewBookingCategoriesCreateOK() *BookingCategoriesCreateOK {
	return &BookingCategoriesCreateOK{}
}

/*
BookingCategoriesCreateOK describes a response with status code 200, with default header values.

OK
*/
type BookingCategoriesCreateOK struct {
	Payload *models.BookingCategory
}

// IsSuccess returns true when this booking categories create o k response has a 2xx status code
func (o *BookingCategoriesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this booking categories create o k response has a 3xx status code
func (o *BookingCategoriesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this booking categories create o k response has a 4xx status code
func (o *BookingCategoriesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this booking categories create o k response has a 5xx status code
func (o *BookingCategoriesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this booking categories create o k response a status code equal to that given
func (o *BookingCategoriesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the booking categories create o k response
func (o *BookingCategoriesCreateOK) Code() int {
	return 200
}

func (o *BookingCategoriesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/BookingCategories][%d] bookingCategoriesCreateOK  %+v", 200, o.Payload)
}

func (o *BookingCategoriesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/BookingCategories][%d] bookingCategoriesCreateOK  %+v", 200, o.Payload)
}

func (o *BookingCategoriesCreateOK) GetPayload() *models.BookingCategory {
	return o.Payload
}

func (o *BookingCategoriesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BookingCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
