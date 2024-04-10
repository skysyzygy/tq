// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LoginTypesDeleteReader is a Reader for the LoginTypesDelete structure.
type LoginTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewLoginTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/LoginTypes/{id}] LoginTypes_Delete", response, response.Code())
	}
}

// NewLoginTypesDeleteNoContent creates a LoginTypesDeleteNoContent with default headers values
func NewLoginTypesDeleteNoContent() *LoginTypesDeleteNoContent {
	return &LoginTypesDeleteNoContent{}
}

/*
LoginTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type LoginTypesDeleteNoContent struct {
}

// IsSuccess returns true when this login types delete no content response has a 2xx status code
func (o *LoginTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this login types delete no content response has a 3xx status code
func (o *LoginTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this login types delete no content response has a 4xx status code
func (o *LoginTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this login types delete no content response has a 5xx status code
func (o *LoginTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this login types delete no content response a status code equal to that given
func (o *LoginTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the login types delete no content response
func (o *LoginTypesDeleteNoContent) Code() int {
	return 204
}

func (o *LoginTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/LoginTypes/{id}][%d] loginTypesDeleteNoContent ", 204)
}

func (o *LoginTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/LoginTypes/{id}][%d] loginTypesDeleteNoContent ", 204)
}

func (o *LoginTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
