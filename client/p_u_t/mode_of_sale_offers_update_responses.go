// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ModeOfSaleOffersUpdateReader is a Reader for the ModeOfSaleOffersUpdate structure.
type ModeOfSaleOffersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModeOfSaleOffersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModeOfSaleOffersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /TXN/ModeOfSaleOffers/{modeOfSaleOfferId}] ModeOfSaleOffers_Update", response, response.Code())
	}
}

// NewModeOfSaleOffersUpdateOK creates a ModeOfSaleOffersUpdateOK with default headers values
func NewModeOfSaleOffersUpdateOK() *ModeOfSaleOffersUpdateOK {
	return &ModeOfSaleOffersUpdateOK{}
}

/*
ModeOfSaleOffersUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ModeOfSaleOffersUpdateOK struct {
	Payload *models.ModeOfSaleOffer
}

// IsSuccess returns true when this mode of sale offers update o k response has a 2xx status code
func (o *ModeOfSaleOffersUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mode of sale offers update o k response has a 3xx status code
func (o *ModeOfSaleOffersUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mode of sale offers update o k response has a 4xx status code
func (o *ModeOfSaleOffersUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mode of sale offers update o k response has a 5xx status code
func (o *ModeOfSaleOffersUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mode of sale offers update o k response a status code equal to that given
func (o *ModeOfSaleOffersUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mode of sale offers update o k response
func (o *ModeOfSaleOffersUpdateOK) Code() int {
	return 200
}

func (o *ModeOfSaleOffersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /TXN/ModeOfSaleOffers/{modeOfSaleOfferId}][%d] modeOfSaleOffersUpdateOK  %+v", 200, o.Payload)
}

func (o *ModeOfSaleOffersUpdateOK) String() string {
	return fmt.Sprintf("[PUT /TXN/ModeOfSaleOffers/{modeOfSaleOfferId}][%d] modeOfSaleOffersUpdateOK  %+v", 200, o.Payload)
}

func (o *ModeOfSaleOffersUpdateOK) GetPayload() *models.ModeOfSaleOffer {
	return o.Payload
}

func (o *ModeOfSaleOffersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModeOfSaleOffer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
