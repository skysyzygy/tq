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

// PackagesGetSeatFeesReader is a Reader for the PackagesGetSeatFees structure.
type PackagesGetSeatFeesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackagesGetSeatFeesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackagesGetSeatFeesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/Packages/{packageId}/SeatFees] Packages_GetSeatFees", response, response.Code())
	}
}

// NewPackagesGetSeatFeesOK creates a PackagesGetSeatFeesOK with default headers values
func NewPackagesGetSeatFeesOK() *PackagesGetSeatFeesOK {
	return &PackagesGetSeatFeesOK{}
}

/*
PackagesGetSeatFeesOK describes a response with status code 200, with default header values.

OK
*/
type PackagesGetSeatFeesOK struct {
	Payload []*models.ProductSeatFee
}

// IsSuccess returns true when this packages get seat fees o k response has a 2xx status code
func (o *PackagesGetSeatFeesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packages get seat fees o k response has a 3xx status code
func (o *PackagesGetSeatFeesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packages get seat fees o k response has a 4xx status code
func (o *PackagesGetSeatFeesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packages get seat fees o k response has a 5xx status code
func (o *PackagesGetSeatFeesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packages get seat fees o k response a status code equal to that given
func (o *PackagesGetSeatFeesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packages get seat fees o k response
func (o *PackagesGetSeatFeesOK) Code() int {
	return 200
}

func (o *PackagesGetSeatFeesOK) Error() string {
	return fmt.Sprintf("[GET /TXN/Packages/{packageId}/SeatFees][%d] packagesGetSeatFeesOK  %+v", 200, o.Payload)
}

func (o *PackagesGetSeatFeesOK) String() string {
	return fmt.Sprintf("[GET /TXN/Packages/{packageId}/SeatFees][%d] packagesGetSeatFeesOK  %+v", 200, o.Payload)
}

func (o *PackagesGetSeatFeesOK) GetPayload() []*models.ProductSeatFee {
	return o.Payload
}

func (o *PackagesGetSeatFeesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}