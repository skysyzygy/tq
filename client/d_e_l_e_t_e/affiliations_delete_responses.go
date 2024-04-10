// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AffiliationsDeleteReader is a Reader for the AffiliationsDelete structure.
type AffiliationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AffiliationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAffiliationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /CRM/Affiliations/{affiliationId}] Affiliations_Delete", response, response.Code())
	}
}

// NewAffiliationsDeleteNoContent creates a AffiliationsDeleteNoContent with default headers values
func NewAffiliationsDeleteNoContent() *AffiliationsDeleteNoContent {
	return &AffiliationsDeleteNoContent{}
}

/*
AffiliationsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type AffiliationsDeleteNoContent struct {
}

// IsSuccess returns true when this affiliations delete no content response has a 2xx status code
func (o *AffiliationsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this affiliations delete no content response has a 3xx status code
func (o *AffiliationsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this affiliations delete no content response has a 4xx status code
func (o *AffiliationsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this affiliations delete no content response has a 5xx status code
func (o *AffiliationsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this affiliations delete no content response a status code equal to that given
func (o *AffiliationsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the affiliations delete no content response
func (o *AffiliationsDeleteNoContent) Code() int {
	return 204
}

func (o *AffiliationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /CRM/Affiliations/{affiliationId}][%d] affiliationsDeleteNoContent ", 204)
}

func (o *AffiliationsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /CRM/Affiliations/{affiliationId}][%d] affiliationsDeleteNoContent ", 204)
}

func (o *AffiliationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
