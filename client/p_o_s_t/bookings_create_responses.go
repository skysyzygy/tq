// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// BookingsCreateReader is a Reader for the BookingsCreate structure.
type BookingsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookingsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookingsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBookingsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBookingsCreateOK creates a BookingsCreateOK with default headers values
func NewBookingsCreateOK() *BookingsCreateOK {
	return &BookingsCreateOK{}
}

/*
BookingsCreateOK describes a response with status code 200, with default header values.

OK
*/
type BookingsCreateOK struct {
	Payload *models.Booking
}

// IsSuccess returns true when this bookings create o k response has a 2xx status code
func (o *BookingsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bookings create o k response has a 3xx status code
func (o *BookingsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookings create o k response has a 4xx status code
func (o *BookingsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookings create o k response has a 5xx status code
func (o *BookingsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bookings create o k response a status code equal to that given
func (o *BookingsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bookings create o k response
func (o *BookingsCreateOK) Code() int {
	return 200
}

func (o *BookingsCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /EventsManagement/Bookings][%d] bookingsCreateOK %s", 200, payload)
}

func (o *BookingsCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /EventsManagement/Bookings][%d] bookingsCreateOK %s", 200, payload)
}

func (o *BookingsCreateOK) GetPayload() *models.Booking {
	return o.Payload
}

func (o *BookingsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Booking)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookingsCreateDefault creates a BookingsCreateDefault with default headers values
func NewBookingsCreateDefault(code int) *BookingsCreateDefault {
	return &BookingsCreateDefault{
		_statusCode: code,
	}
}

/*
BookingsCreateDefault describes a response with status code -1, with default header values.

Error
*/
type BookingsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this bookings create default response has a 2xx status code
func (o *BookingsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bookings create default response has a 3xx status code
func (o *BookingsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bookings create default response has a 4xx status code
func (o *BookingsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bookings create default response has a 5xx status code
func (o *BookingsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bookings create default response a status code equal to that given
func (o *BookingsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the bookings create default response
func (o *BookingsCreateDefault) Code() int {
	return o._statusCode
}

func (o *BookingsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /EventsManagement/Bookings][%d] Bookings_Create default %s", o._statusCode, payload)
}

func (o *BookingsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /EventsManagement/Bookings][%d] Bookings_Create default %s", o._statusCode, payload)
}

func (o *BookingsCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *BookingsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
