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

// CartPrintTicketElementsReader is a Reader for the CartPrintTicketElements structure.
type CartPrintTicketElementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartPrintTicketElementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCartPrintTicketElementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Web/Cart/{sessionKey}/Print/TicketElements] Cart_PrintTicketElements", response, response.Code())
	}
}

// NewCartPrintTicketElementsOK creates a CartPrintTicketElementsOK with default headers values
func NewCartPrintTicketElementsOK() *CartPrintTicketElementsOK {
	return &CartPrintTicketElementsOK{}
}

/*
CartPrintTicketElementsOK describes a response with status code 200, with default header values.

OK
*/
type CartPrintTicketElementsOK struct {
	Payload []*models.TicketElement
}

// IsSuccess returns true when this cart print ticket elements o k response has a 2xx status code
func (o *CartPrintTicketElementsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart print ticket elements o k response has a 3xx status code
func (o *CartPrintTicketElementsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart print ticket elements o k response has a 4xx status code
func (o *CartPrintTicketElementsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart print ticket elements o k response has a 5xx status code
func (o *CartPrintTicketElementsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cart print ticket elements o k response a status code equal to that given
func (o *CartPrintTicketElementsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cart print ticket elements o k response
func (o *CartPrintTicketElementsOK) Code() int {
	return 200
}

func (o *CartPrintTicketElementsOK) Error() string {
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Print/TicketElements][%d] cartPrintTicketElementsOK  %+v", 200, o.Payload)
}

func (o *CartPrintTicketElementsOK) String() string {
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Print/TicketElements][%d] cartPrintTicketElementsOK  %+v", 200, o.Payload)
}

func (o *CartPrintTicketElementsOK) GetPayload() []*models.TicketElement {
	return o.Payload
}

func (o *CartPrintTicketElementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
