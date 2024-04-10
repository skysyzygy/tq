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

// PackagesGetSeatsReader is a Reader for the PackagesGetSeats structure.
type PackagesGetSeatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackagesGetSeatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackagesGetSeatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/Packages/{packageId}/Seats] Packages_GetSeats", response, response.Code())
	}
}

// NewPackagesGetSeatsOK creates a PackagesGetSeatsOK with default headers values
func NewPackagesGetSeatsOK() *PackagesGetSeatsOK {
	return &PackagesGetSeatsOK{}
}

/*
PackagesGetSeatsOK describes a response with status code 200, with default header values.

OK
*/
type PackagesGetSeatsOK struct {
	Payload []*models.Seat
}

// IsSuccess returns true when this packages get seats o k response has a 2xx status code
func (o *PackagesGetSeatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this packages get seats o k response has a 3xx status code
func (o *PackagesGetSeatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this packages get seats o k response has a 4xx status code
func (o *PackagesGetSeatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this packages get seats o k response has a 5xx status code
func (o *PackagesGetSeatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this packages get seats o k response a status code equal to that given
func (o *PackagesGetSeatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the packages get seats o k response
func (o *PackagesGetSeatsOK) Code() int {
	return 200
}

func (o *PackagesGetSeatsOK) Error() string {
	return fmt.Sprintf("[GET /TXN/Packages/{packageId}/Seats][%d] packagesGetSeatsOK  %+v", 200, o.Payload)
}

func (o *PackagesGetSeatsOK) String() string {
	return fmt.Sprintf("[GET /TXN/Packages/{packageId}/Seats][%d] packagesGetSeatsOK  %+v", 200, o.Payload)
}

func (o *PackagesGetSeatsOK) GetPayload() []*models.Seat {
	return o.Payload
}

func (o *PackagesGetSeatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
