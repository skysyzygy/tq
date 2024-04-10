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

// CountriesCreateReader is a Reader for the CountriesCreate structure.
type CountriesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountriesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountriesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/Countries] Countries_Create", response, response.Code())
	}
}

// NewCountriesCreateOK creates a CountriesCreateOK with default headers values
func NewCountriesCreateOK() *CountriesCreateOK {
	return &CountriesCreateOK{}
}

/*
CountriesCreateOK describes a response with status code 200, with default header values.

OK
*/
type CountriesCreateOK struct {
	Payload *models.Country
}

// IsSuccess returns true when this countries create o k response has a 2xx status code
func (o *CountriesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this countries create o k response has a 3xx status code
func (o *CountriesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this countries create o k response has a 4xx status code
func (o *CountriesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this countries create o k response has a 5xx status code
func (o *CountriesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this countries create o k response a status code equal to that given
func (o *CountriesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the countries create o k response
func (o *CountriesCreateOK) Code() int {
	return 200
}

func (o *CountriesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/Countries][%d] countriesCreateOK  %+v", 200, o.Payload)
}

func (o *CountriesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/Countries][%d] countriesCreateOK  %+v", 200, o.Payload)
}

func (o *CountriesCreateOK) GetPayload() *models.Country {
	return o.Payload
}

func (o *CountriesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Country)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
