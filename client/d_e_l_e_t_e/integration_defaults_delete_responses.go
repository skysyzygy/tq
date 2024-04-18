// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IntegrationDefaultsDeleteReader is a Reader for the IntegrationDefaultsDelete structure.
type IntegrationDefaultsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationDefaultsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIntegrationDefaultsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/IntegrationDefaults/{id}] IntegrationDefaults_Delete", response, response.Code())
	}
}

// NewIntegrationDefaultsDeleteNoContent creates a IntegrationDefaultsDeleteNoContent with default headers values
func NewIntegrationDefaultsDeleteNoContent() *IntegrationDefaultsDeleteNoContent {
	return &IntegrationDefaultsDeleteNoContent{}
}

/*
IntegrationDefaultsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type IntegrationDefaultsDeleteNoContent struct {
}

// IsSuccess returns true when this integration defaults delete no content response has a 2xx status code
func (o *IntegrationDefaultsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this integration defaults delete no content response has a 3xx status code
func (o *IntegrationDefaultsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this integration defaults delete no content response has a 4xx status code
func (o *IntegrationDefaultsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this integration defaults delete no content response has a 5xx status code
func (o *IntegrationDefaultsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this integration defaults delete no content response a status code equal to that given
func (o *IntegrationDefaultsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the integration defaults delete no content response
func (o *IntegrationDefaultsDeleteNoContent) Code() int {
	return 204
}

func (o *IntegrationDefaultsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/IntegrationDefaults/{id}][%d] integrationDefaultsDeleteNoContent ", 204)
}

func (o *IntegrationDefaultsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/IntegrationDefaults/{id}][%d] integrationDefaultsDeleteNoContent ", 204)
}

func (o *IntegrationDefaultsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}