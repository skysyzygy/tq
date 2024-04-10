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

// AffiliationsGetAllReader is a Reader for the AffiliationsGetAll structure.
type AffiliationsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AffiliationsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAffiliationsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Affiliations] Affiliations_GetAll", response, response.Code())
	}
}

// NewAffiliationsGetAllOK creates a AffiliationsGetAllOK with default headers values
func NewAffiliationsGetAllOK() *AffiliationsGetAllOK {
	return &AffiliationsGetAllOK{}
}

/*
AffiliationsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type AffiliationsGetAllOK struct {
	Payload []*models.Affiliation
}

// IsSuccess returns true when this affiliations get all o k response has a 2xx status code
func (o *AffiliationsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this affiliations get all o k response has a 3xx status code
func (o *AffiliationsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this affiliations get all o k response has a 4xx status code
func (o *AffiliationsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this affiliations get all o k response has a 5xx status code
func (o *AffiliationsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this affiliations get all o k response a status code equal to that given
func (o *AffiliationsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the affiliations get all o k response
func (o *AffiliationsGetAllOK) Code() int {
	return 200
}

func (o *AffiliationsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Affiliations][%d] affiliationsGetAllOK  %+v", 200, o.Payload)
}

func (o *AffiliationsGetAllOK) String() string {
	return fmt.Sprintf("[GET /CRM/Affiliations][%d] affiliationsGetAllOK  %+v", 200, o.Payload)
}

func (o *AffiliationsGetAllOK) GetPayload() []*models.Affiliation {
	return o.Payload
}

func (o *AffiliationsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
