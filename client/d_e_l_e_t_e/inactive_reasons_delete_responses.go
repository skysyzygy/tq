// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// InactiveReasonsDeleteReader is a Reader for the InactiveReasonsDelete structure.
type InactiveReasonsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InactiveReasonsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInactiveReasonsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/InactiveReasons/{id}] InactiveReasons_Delete", response, response.Code())
	}
}

// NewInactiveReasonsDeleteNoContent creates a InactiveReasonsDeleteNoContent with default headers values
func NewInactiveReasonsDeleteNoContent() *InactiveReasonsDeleteNoContent {
	return &InactiveReasonsDeleteNoContent{}
}

/*
InactiveReasonsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type InactiveReasonsDeleteNoContent struct {
}

// IsSuccess returns true when this inactive reasons delete no content response has a 2xx status code
func (o *InactiveReasonsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inactive reasons delete no content response has a 3xx status code
func (o *InactiveReasonsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inactive reasons delete no content response has a 4xx status code
func (o *InactiveReasonsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this inactive reasons delete no content response has a 5xx status code
func (o *InactiveReasonsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this inactive reasons delete no content response a status code equal to that given
func (o *InactiveReasonsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the inactive reasons delete no content response
func (o *InactiveReasonsDeleteNoContent) Code() int {
	return 204
}

func (o *InactiveReasonsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/InactiveReasons/{id}][%d] inactiveReasonsDeleteNoContent ", 204)
}

func (o *InactiveReasonsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/InactiveReasons/{id}][%d] inactiveReasonsDeleteNoContent ", 204)
}

func (o *InactiveReasonsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
