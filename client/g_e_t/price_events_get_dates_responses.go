// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PriceEventsGetDatesReader is a Reader for the PriceEventsGetDates structure.
type PriceEventsGetDatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceEventsGetDatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPriceEventsGetDatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/PriceEvents/Dates] PriceEvents_GetDates", response, response.Code())
	}
}

// NewPriceEventsGetDatesOK creates a PriceEventsGetDatesOK with default headers values
func NewPriceEventsGetDatesOK() *PriceEventsGetDatesOK {
	return &PriceEventsGetDatesOK{}
}

/*
PriceEventsGetDatesOK describes a response with status code 200, with default header values.

OK
*/
type PriceEventsGetDatesOK struct {
	Payload []strfmt.DateTime
}

// IsSuccess returns true when this price events get dates o k response has a 2xx status code
func (o *PriceEventsGetDatesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price events get dates o k response has a 3xx status code
func (o *PriceEventsGetDatesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price events get dates o k response has a 4xx status code
func (o *PriceEventsGetDatesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this price events get dates o k response has a 5xx status code
func (o *PriceEventsGetDatesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this price events get dates o k response a status code equal to that given
func (o *PriceEventsGetDatesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the price events get dates o k response
func (o *PriceEventsGetDatesOK) Code() int {
	return 200
}

func (o *PriceEventsGetDatesOK) Error() string {
	return fmt.Sprintf("[GET /TXN/PriceEvents/Dates][%d] priceEventsGetDatesOK  %+v", 200, o.Payload)
}

func (o *PriceEventsGetDatesOK) String() string {
	return fmt.Sprintf("[GET /TXN/PriceEvents/Dates][%d] priceEventsGetDatesOK  %+v", 200, o.Payload)
}

func (o *PriceEventsGetDatesOK) GetPayload() []strfmt.DateTime {
	return o.Payload
}

func (o *PriceEventsGetDatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
