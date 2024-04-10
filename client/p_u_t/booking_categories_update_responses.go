// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// BookingCategoriesUpdateReader is a Reader for the BookingCategoriesUpdate structure.
type BookingCategoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookingCategoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookingCategoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/BookingCategories/{id}] BookingCategories_Update", response, response.Code())
	}
}

// NewBookingCategoriesUpdateOK creates a BookingCategoriesUpdateOK with default headers values
func NewBookingCategoriesUpdateOK() *BookingCategoriesUpdateOK {
	return &BookingCategoriesUpdateOK{}
}

/*
BookingCategoriesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type BookingCategoriesUpdateOK struct {
	Payload *models.BookingCategory
}

// IsSuccess returns true when this booking categories update o k response has a 2xx status code
func (o *BookingCategoriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this booking categories update o k response has a 3xx status code
func (o *BookingCategoriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this booking categories update o k response has a 4xx status code
func (o *BookingCategoriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this booking categories update o k response has a 5xx status code
func (o *BookingCategoriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this booking categories update o k response a status code equal to that given
func (o *BookingCategoriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the booking categories update o k response
func (o *BookingCategoriesUpdateOK) Code() int {
	return 200
}

func (o *BookingCategoriesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/BookingCategories/{id}][%d] bookingCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *BookingCategoriesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/BookingCategories/{id}][%d] bookingCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *BookingCategoriesUpdateOK) GetPayload() *models.BookingCategory {
	return o.Payload
}

func (o *BookingCategoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BookingCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
