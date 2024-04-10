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

// CartPrintPrintStringsReader is a Reader for the CartPrintPrintStrings structure.
type CartPrintPrintStringsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartPrintPrintStringsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCartPrintPrintStringsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Web/Cart/{sessionKey}/Print/PrintStrings] Cart_PrintPrintStrings", response, response.Code())
	}
}

// NewCartPrintPrintStringsOK creates a CartPrintPrintStringsOK with default headers values
func NewCartPrintPrintStringsOK() *CartPrintPrintStringsOK {
	return &CartPrintPrintStringsOK{}
}

/*
CartPrintPrintStringsOK describes a response with status code 200, with default header values.

OK
*/
type CartPrintPrintStringsOK struct {
	Payload *models.PrintOrderResponse
}

// IsSuccess returns true when this cart print print strings o k response has a 2xx status code
func (o *CartPrintPrintStringsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart print print strings o k response has a 3xx status code
func (o *CartPrintPrintStringsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart print print strings o k response has a 4xx status code
func (o *CartPrintPrintStringsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart print print strings o k response has a 5xx status code
func (o *CartPrintPrintStringsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cart print print strings o k response a status code equal to that given
func (o *CartPrintPrintStringsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cart print print strings o k response
func (o *CartPrintPrintStringsOK) Code() int {
	return 200
}

func (o *CartPrintPrintStringsOK) Error() string {
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Print/PrintStrings][%d] cartPrintPrintStringsOK  %+v", 200, o.Payload)
}

func (o *CartPrintPrintStringsOK) String() string {
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Print/PrintStrings][%d] cartPrintPrintStringsOK  %+v", 200, o.Payload)
}

func (o *CartPrintPrintStringsOK) GetPayload() *models.PrintOrderResponse {
	return o.Payload
}

func (o *CartPrintPrintStringsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrintOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
