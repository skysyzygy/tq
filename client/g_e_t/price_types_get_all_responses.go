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

// PriceTypesGetAllReader is a Reader for the PriceTypesGetAll structure.
type PriceTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPriceTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/PriceTypes] PriceTypes_GetAll", response, response.Code())
	}
}

// NewPriceTypesGetAllOK creates a PriceTypesGetAllOK with default headers values
func NewPriceTypesGetAllOK() *PriceTypesGetAllOK {
	return &PriceTypesGetAllOK{}
}

/*
PriceTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PriceTypesGetAllOK struct {
	Payload []*models.PriceType
}

// IsSuccess returns true when this price types get all o k response has a 2xx status code
func (o *PriceTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price types get all o k response has a 3xx status code
func (o *PriceTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price types get all o k response has a 4xx status code
func (o *PriceTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this price types get all o k response has a 5xx status code
func (o *PriceTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this price types get all o k response a status code equal to that given
func (o *PriceTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the price types get all o k response
func (o *PriceTypesGetAllOK) Code() int {
	return 200
}

func (o *PriceTypesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /TXN/PriceTypes][%d] priceTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *PriceTypesGetAllOK) String() string {
	return fmt.Sprintf("[GET /TXN/PriceTypes][%d] priceTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *PriceTypesGetAllOK) GetPayload() []*models.PriceType {
	return o.Payload
}

func (o *PriceTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
