// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ResourceTypesDeleteReader is a Reader for the ResourceTypesDelete structure.
type ResourceTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewResourceTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /EventsManagement/ResourceTypes/{id}] ResourceTypes_Delete", response, response.Code())
	}
}

// NewResourceTypesDeleteNoContent creates a ResourceTypesDeleteNoContent with default headers values
func NewResourceTypesDeleteNoContent() *ResourceTypesDeleteNoContent {
	return &ResourceTypesDeleteNoContent{}
}

/*
ResourceTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ResourceTypesDeleteNoContent struct {
}

// IsSuccess returns true when this resource types delete no content response has a 2xx status code
func (o *ResourceTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource types delete no content response has a 3xx status code
func (o *ResourceTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource types delete no content response has a 4xx status code
func (o *ResourceTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource types delete no content response has a 5xx status code
func (o *ResourceTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this resource types delete no content response a status code equal to that given
func (o *ResourceTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the resource types delete no content response
func (o *ResourceTypesDeleteNoContent) Code() int {
	return 204
}

func (o *ResourceTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /EventsManagement/ResourceTypes/{id}][%d] resourceTypesDeleteNoContent ", 204)
}

func (o *ResourceTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /EventsManagement/ResourceTypes/{id}][%d] resourceTypesDeleteNoContent ", 204)
}

func (o *ResourceTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
