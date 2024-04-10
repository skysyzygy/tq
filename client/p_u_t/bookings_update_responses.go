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

// BookingsUpdateReader is a Reader for the BookingsUpdate structure.
type BookingsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookingsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookingsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /EventsManagement/Bookings/{bookingId}] Bookings_Update", response, response.Code())
	}
}

// NewBookingsUpdateOK creates a BookingsUpdateOK with default headers values
func NewBookingsUpdateOK() *BookingsUpdateOK {
	return &BookingsUpdateOK{}
}

/*
BookingsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type BookingsUpdateOK struct {
	Payload *models.Booking
}

// IsSuccess returns true when this bookings update o k response has a 2xx status code
func (o *BookingsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bookings update o k response has a 3xx status code
func (o *BookingsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookings update o k response has a 4xx status code
func (o *BookingsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookings update o k response has a 5xx status code
func (o *BookingsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bookings update o k response a status code equal to that given
func (o *BookingsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bookings update o k response
func (o *BookingsUpdateOK) Code() int {
	return 200
}

func (o *BookingsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /EventsManagement/Bookings/{bookingId}][%d] bookingsUpdateOK  %+v", 200, o.Payload)
}

func (o *BookingsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /EventsManagement/Bookings/{bookingId}][%d] bookingsUpdateOK  %+v", 200, o.Payload)
}

func (o *BookingsUpdateOK) GetPayload() *models.Booking {
	return o.Payload
}

func (o *BookingsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Booking)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
