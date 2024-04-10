// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CustomDefaultsDeleteReader is a Reader for the CustomDefaultsDelete structure.
type CustomDefaultsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomDefaultsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCustomDefaultsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/CustomDefaults/{id}] CustomDefaults_Delete", response, response.Code())
	}
}

// NewCustomDefaultsDeleteNoContent creates a CustomDefaultsDeleteNoContent with default headers values
func NewCustomDefaultsDeleteNoContent() *CustomDefaultsDeleteNoContent {
	return &CustomDefaultsDeleteNoContent{}
}

/*
CustomDefaultsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type CustomDefaultsDeleteNoContent struct {
}

// IsSuccess returns true when this custom defaults delete no content response has a 2xx status code
func (o *CustomDefaultsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom defaults delete no content response has a 3xx status code
func (o *CustomDefaultsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom defaults delete no content response has a 4xx status code
func (o *CustomDefaultsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom defaults delete no content response has a 5xx status code
func (o *CustomDefaultsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this custom defaults delete no content response a status code equal to that given
func (o *CustomDefaultsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the custom defaults delete no content response
func (o *CustomDefaultsDeleteNoContent) Code() int {
	return 204
}

func (o *CustomDefaultsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/CustomDefaults/{id}][%d] customDefaultsDeleteNoContent ", 204)
}

func (o *CustomDefaultsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/CustomDefaults/{id}][%d] customDefaultsDeleteNoContent ", 204)
}

func (o *CustomDefaultsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
