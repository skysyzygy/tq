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

// ModesOfSaleGetSummariesReader is a Reader for the ModesOfSaleGetSummaries structure.
type ModesOfSaleGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModesOfSaleGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModesOfSaleGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/ModesOfSale/Summary] ModesOfSale_GetSummaries", response, response.Code())
	}
}

// NewModesOfSaleGetSummariesOK creates a ModesOfSaleGetSummariesOK with default headers values
func NewModesOfSaleGetSummariesOK() *ModesOfSaleGetSummariesOK {
	return &ModesOfSaleGetSummariesOK{}
}

/*
ModesOfSaleGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type ModesOfSaleGetSummariesOK struct {
	Payload []*models.ModeOfSaleSummary
}

// IsSuccess returns true when this modes of sale get summaries o k response has a 2xx status code
func (o *ModesOfSaleGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modes of sale get summaries o k response has a 3xx status code
func (o *ModesOfSaleGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modes of sale get summaries o k response has a 4xx status code
func (o *ModesOfSaleGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modes of sale get summaries o k response has a 5xx status code
func (o *ModesOfSaleGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modes of sale get summaries o k response a status code equal to that given
func (o *ModesOfSaleGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the modes of sale get summaries o k response
func (o *ModesOfSaleGetSummariesOK) Code() int {
	return 200
}

func (o *ModesOfSaleGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /TXN/ModesOfSale/Summary][%d] modesOfSaleGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ModesOfSaleGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /TXN/ModesOfSale/Summary][%d] modesOfSaleGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ModesOfSaleGetSummariesOK) GetPayload() []*models.ModeOfSaleSummary {
	return o.Payload
}

func (o *ModesOfSaleGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
