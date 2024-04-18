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

// ConstituentTypeAffiliatesUpdateReader is a Reader for the ConstituentTypeAffiliatesUpdate structure.
type ConstituentTypeAffiliatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentTypeAffiliatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentTypeAffiliatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/ConstituentTypeAffiliates/{id}] ConstituentTypeAffiliates_Update", response, response.Code())
	}
}

// NewConstituentTypeAffiliatesUpdateOK creates a ConstituentTypeAffiliatesUpdateOK with default headers values
func NewConstituentTypeAffiliatesUpdateOK() *ConstituentTypeAffiliatesUpdateOK {
	return &ConstituentTypeAffiliatesUpdateOK{}
}

/*
ConstituentTypeAffiliatesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentTypeAffiliatesUpdateOK struct {
	Payload *models.ConstituentTypeAffiliate
}

// IsSuccess returns true when this constituent type affiliates update o k response has a 2xx status code
func (o *ConstituentTypeAffiliatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent type affiliates update o k response has a 3xx status code
func (o *ConstituentTypeAffiliatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent type affiliates update o k response has a 4xx status code
func (o *ConstituentTypeAffiliatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent type affiliates update o k response has a 5xx status code
func (o *ConstituentTypeAffiliatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent type affiliates update o k response a status code equal to that given
func (o *ConstituentTypeAffiliatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituent type affiliates update o k response
func (o *ConstituentTypeAffiliatesUpdateOK) Code() int {
	return 200
}

func (o *ConstituentTypeAffiliatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/ConstituentTypeAffiliates/{id}][%d] constituentTypeAffiliatesUpdateOK  %+v", 200, o.Payload)
}

func (o *ConstituentTypeAffiliatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/ConstituentTypeAffiliates/{id}][%d] constituentTypeAffiliatesUpdateOK  %+v", 200, o.Payload)
}

func (o *ConstituentTypeAffiliatesUpdateOK) GetPayload() *models.ConstituentTypeAffiliate {
	return o.Payload
}

func (o *ConstituentTypeAffiliatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentTypeAffiliate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}