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

// BookingsAddDocumentReader is a Reader for the BookingsAddDocument structure.
type BookingsAddDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookingsAddDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookingsAddDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /EventsManagement/Bookings/{bookingId}/Documents] Bookings_AddDocument", response, response.Code())
	}
}

// NewBookingsAddDocumentOK creates a BookingsAddDocumentOK with default headers values
func NewBookingsAddDocumentOK() *BookingsAddDocumentOK {
	return &BookingsAddDocumentOK{}
}

/*
BookingsAddDocumentOK describes a response with status code 200, with default header values.

OK
*/
type BookingsAddDocumentOK struct {
	Payload *models.DocumentSummary
}

// IsSuccess returns true when this bookings add document o k response has a 2xx status code
func (o *BookingsAddDocumentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bookings add document o k response has a 3xx status code
func (o *BookingsAddDocumentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bookings add document o k response has a 4xx status code
func (o *BookingsAddDocumentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bookings add document o k response has a 5xx status code
func (o *BookingsAddDocumentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bookings add document o k response a status code equal to that given
func (o *BookingsAddDocumentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bookings add document o k response
func (o *BookingsAddDocumentOK) Code() int {
	return 200
}

func (o *BookingsAddDocumentOK) Error() string {
	return fmt.Sprintf("[POST /EventsManagement/Bookings/{bookingId}/Documents][%d] bookingsAddDocumentOK  %+v", 200, o.Payload)
}

func (o *BookingsAddDocumentOK) String() string {
	return fmt.Sprintf("[POST /EventsManagement/Bookings/{bookingId}/Documents][%d] bookingsAddDocumentOK  %+v", 200, o.Payload)
}

func (o *BookingsAddDocumentOK) GetPayload() *models.DocumentSummary {
	return o.Payload
}

func (o *BookingsAddDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
