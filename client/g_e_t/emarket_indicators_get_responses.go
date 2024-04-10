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

// EmarketIndicatorsGetReader is a Reader for the EmarketIndicatorsGet structure.
type EmarketIndicatorsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmarketIndicatorsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmarketIndicatorsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/EmarketIndicators/{id}] EmarketIndicators_Get", response, response.Code())
	}
}

// NewEmarketIndicatorsGetOK creates a EmarketIndicatorsGetOK with default headers values
func NewEmarketIndicatorsGetOK() *EmarketIndicatorsGetOK {
	return &EmarketIndicatorsGetOK{}
}

/*
EmarketIndicatorsGetOK describes a response with status code 200, with default header values.

OK
*/
type EmarketIndicatorsGetOK struct {
	Payload *models.EmarketIndicator
}

// IsSuccess returns true when this emarket indicators get o k response has a 2xx status code
func (o *EmarketIndicatorsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this emarket indicators get o k response has a 3xx status code
func (o *EmarketIndicatorsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emarket indicators get o k response has a 4xx status code
func (o *EmarketIndicatorsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this emarket indicators get o k response has a 5xx status code
func (o *EmarketIndicatorsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this emarket indicators get o k response a status code equal to that given
func (o *EmarketIndicatorsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the emarket indicators get o k response
func (o *EmarketIndicatorsGetOK) Code() int {
	return 200
}

func (o *EmarketIndicatorsGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/EmarketIndicators/{id}][%d] emarketIndicatorsGetOK  %+v", 200, o.Payload)
}

func (o *EmarketIndicatorsGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/EmarketIndicators/{id}][%d] emarketIndicatorsGetOK  %+v", 200, o.Payload)
}

func (o *EmarketIndicatorsGetOK) GetPayload() *models.EmarketIndicator {
	return o.Payload
}

func (o *EmarketIndicatorsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmarketIndicator)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
