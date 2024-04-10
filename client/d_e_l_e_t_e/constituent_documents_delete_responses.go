// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ConstituentDocumentsDeleteReader is a Reader for the ConstituentDocumentsDelete structure.
type ConstituentDocumentsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentDocumentsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewConstituentDocumentsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /CRM/Documents/{documentId}] ConstituentDocuments_Delete", response, response.Code())
	}
}

// NewConstituentDocumentsDeleteNoContent creates a ConstituentDocumentsDeleteNoContent with default headers values
func NewConstituentDocumentsDeleteNoContent() *ConstituentDocumentsDeleteNoContent {
	return &ConstituentDocumentsDeleteNoContent{}
}

/*
ConstituentDocumentsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ConstituentDocumentsDeleteNoContent struct {
}

// IsSuccess returns true when this constituent documents delete no content response has a 2xx status code
func (o *ConstituentDocumentsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent documents delete no content response has a 3xx status code
func (o *ConstituentDocumentsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent documents delete no content response has a 4xx status code
func (o *ConstituentDocumentsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent documents delete no content response has a 5xx status code
func (o *ConstituentDocumentsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent documents delete no content response a status code equal to that given
func (o *ConstituentDocumentsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the constituent documents delete no content response
func (o *ConstituentDocumentsDeleteNoContent) Code() int {
	return 204
}

func (o *ConstituentDocumentsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /CRM/Documents/{documentId}][%d] constituentDocumentsDeleteNoContent ", 204)
}

func (o *ConstituentDocumentsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /CRM/Documents/{documentId}][%d] constituentDocumentsDeleteNoContent ", 204)
}

func (o *ConstituentDocumentsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
