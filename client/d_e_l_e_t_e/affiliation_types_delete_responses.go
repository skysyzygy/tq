// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AffiliationTypesDeleteReader is a Reader for the AffiliationTypesDelete structure.
type AffiliationTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AffiliationTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAffiliationTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/AffiliationTypes/{id}] AffiliationTypes_Delete", response, response.Code())
	}
}

// NewAffiliationTypesDeleteNoContent creates a AffiliationTypesDeleteNoContent with default headers values
func NewAffiliationTypesDeleteNoContent() *AffiliationTypesDeleteNoContent {
	return &AffiliationTypesDeleteNoContent{}
}

/*
AffiliationTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type AffiliationTypesDeleteNoContent struct {
}

// IsSuccess returns true when this affiliation types delete no content response has a 2xx status code
func (o *AffiliationTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this affiliation types delete no content response has a 3xx status code
func (o *AffiliationTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this affiliation types delete no content response has a 4xx status code
func (o *AffiliationTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this affiliation types delete no content response has a 5xx status code
func (o *AffiliationTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this affiliation types delete no content response a status code equal to that given
func (o *AffiliationTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the affiliation types delete no content response
func (o *AffiliationTypesDeleteNoContent) Code() int {
	return 204
}

func (o *AffiliationTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/AffiliationTypes/{id}][%d] affiliationTypesDeleteNoContent ", 204)
}

func (o *AffiliationTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/AffiliationTypes/{id}][%d] affiliationTypesDeleteNoContent ", 204)
}

func (o *AffiliationTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
