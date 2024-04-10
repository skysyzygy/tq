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

// AffiliationsGetReader is a Reader for the AffiliationsGet structure.
type AffiliationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AffiliationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAffiliationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Affiliations/{affiliationId}] Affiliations_Get", response, response.Code())
	}
}

// NewAffiliationsGetOK creates a AffiliationsGetOK with default headers values
func NewAffiliationsGetOK() *AffiliationsGetOK {
	return &AffiliationsGetOK{}
}

/*
AffiliationsGetOK describes a response with status code 200, with default header values.

OK
*/
type AffiliationsGetOK struct {
	Payload *models.Affiliation
}

// IsSuccess returns true when this affiliations get o k response has a 2xx status code
func (o *AffiliationsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this affiliations get o k response has a 3xx status code
func (o *AffiliationsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this affiliations get o k response has a 4xx status code
func (o *AffiliationsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this affiliations get o k response has a 5xx status code
func (o *AffiliationsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this affiliations get o k response a status code equal to that given
func (o *AffiliationsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the affiliations get o k response
func (o *AffiliationsGetOK) Code() int {
	return 200
}

func (o *AffiliationsGetOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Affiliations/{affiliationId}][%d] affiliationsGetOK  %+v", 200, o.Payload)
}

func (o *AffiliationsGetOK) String() string {
	return fmt.Sprintf("[GET /CRM/Affiliations/{affiliationId}][%d] affiliationsGetOK  %+v", 200, o.Payload)
}

func (o *AffiliationsGetOK) GetPayload() *models.Affiliation {
	return o.Payload
}

func (o *AffiliationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Affiliation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
