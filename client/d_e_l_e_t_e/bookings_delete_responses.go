// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// BookingsDeleteReader is a Reader for the BookingsDelete structure.
type BookingsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookingsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewBookingsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /EventsManagement/Bookings/{bookingId}] Bookings_Delete", response, response.Code())
	}
}

// NewBookingsDeleteNoContent creates a BookingsDeleteNoContent with default headers values
func NewBookingsDeleteNoContent() *BookingsDeleteNoContent {
	return &BookingsDeleteNoContent{}
}

/*
BookingsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type BookingsDeleteNoContent struct {
}

// IsSuccess returns true when this bookings delete no content response has a 2xx status code
func (o *BookingsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bookings delete no content response has a 3xx status code
func (o *BookingsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookings delete no content response has a 4xx status code
func (o *BookingsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookings delete no content response has a 5xx status code
func (o *BookingsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this bookings delete no content response a status code equal to that given
func (o *BookingsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the bookings delete no content response
func (o *BookingsDeleteNoContent) Code() int {
	return 204
}

func (o *BookingsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /EventsManagement/Bookings/{bookingId}][%d] bookingsDeleteNoContent ", 204)
}

func (o *BookingsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /EventsManagement/Bookings/{bookingId}][%d] bookingsDeleteNoContent ", 204)
}

func (o *BookingsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
