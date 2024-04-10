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

// CumulativeGivingReceiptsGetAllReader is a Reader for the CumulativeGivingReceiptsGetAll structure.
type CumulativeGivingReceiptsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CumulativeGivingReceiptsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCumulativeGivingReceiptsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/CumulativeGivingReceipts] CumulativeGivingReceipts_GetAll", response, response.Code())
	}
}

// NewCumulativeGivingReceiptsGetAllOK creates a CumulativeGivingReceiptsGetAllOK with default headers values
func NewCumulativeGivingReceiptsGetAllOK() *CumulativeGivingReceiptsGetAllOK {
	return &CumulativeGivingReceiptsGetAllOK{}
}

/*
CumulativeGivingReceiptsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type CumulativeGivingReceiptsGetAllOK struct {
	Payload []*models.CumulativeGivingReceipt
}

// IsSuccess returns true when this cumulative giving receipts get all o k response has a 2xx status code
func (o *CumulativeGivingReceiptsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cumulative giving receipts get all o k response has a 3xx status code
func (o *CumulativeGivingReceiptsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cumulative giving receipts get all o k response has a 4xx status code
func (o *CumulativeGivingReceiptsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cumulative giving receipts get all o k response has a 5xx status code
func (o *CumulativeGivingReceiptsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cumulative giving receipts get all o k response a status code equal to that given
func (o *CumulativeGivingReceiptsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cumulative giving receipts get all o k response
func (o *CumulativeGivingReceiptsGetAllOK) Code() int {
	return 200
}

func (o *CumulativeGivingReceiptsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /CRM/CumulativeGivingReceipts][%d] cumulativeGivingReceiptsGetAllOK  %+v", 200, o.Payload)
}

func (o *CumulativeGivingReceiptsGetAllOK) String() string {
	return fmt.Sprintf("[GET /CRM/CumulativeGivingReceipts][%d] cumulativeGivingReceiptsGetAllOK  %+v", 200, o.Payload)
}

func (o *CumulativeGivingReceiptsGetAllOK) GetPayload() []*models.CumulativeGivingReceipt {
	return o.Payload
}

func (o *CumulativeGivingReceiptsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
