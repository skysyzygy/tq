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

// ConstituentsSearchReader is a Reader for the ConstituentsSearch structure.
type ConstituentsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Constituents/Search] Constituents_Search", response, response.Code())
	}
}

// NewConstituentsSearchOK creates a ConstituentsSearchOK with default headers values
func NewConstituentsSearchOK() *ConstituentsSearchOK {
	return &ConstituentsSearchOK{}
}

/*
ConstituentsSearchOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentsSearchOK struct {
	Payload *models.ConstituentSearchResponse
}

// IsSuccess returns true when this constituents search o k response has a 2xx status code
func (o *ConstituentsSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituents search o k response has a 3xx status code
func (o *ConstituentsSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituents search o k response has a 4xx status code
func (o *ConstituentsSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituents search o k response has a 5xx status code
func (o *ConstituentsSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituents search o k response a status code equal to that given
func (o *ConstituentsSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituents search o k response
func (o *ConstituentsSearchOK) Code() int {
	return 200
}

func (o *ConstituentsSearchOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Constituents/Search][%d] constituentsSearchOK  %+v", 200, o.Payload)
}

func (o *ConstituentsSearchOK) String() string {
	return fmt.Sprintf("[GET /CRM/Constituents/Search][%d] constituentsSearchOK  %+v", 200, o.Payload)
}

func (o *ConstituentsSearchOK) GetPayload() *models.ConstituentSearchResponse {
	return o.Payload
}

func (o *ConstituentsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
