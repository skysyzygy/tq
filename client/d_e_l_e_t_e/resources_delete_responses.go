// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ResourcesDeleteReader is a Reader for the ResourcesDelete structure.
type ResourcesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourcesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewResourcesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /EventsManagement/Resources/{id}] Resources_Delete", response, response.Code())
	}
}

// NewResourcesDeleteNoContent creates a ResourcesDeleteNoContent with default headers values
func NewResourcesDeleteNoContent() *ResourcesDeleteNoContent {
	return &ResourcesDeleteNoContent{}
}

/*
ResourcesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ResourcesDeleteNoContent struct {
}

// IsSuccess returns true when this resources delete no content response has a 2xx status code
func (o *ResourcesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resources delete no content response has a 3xx status code
func (o *ResourcesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resources delete no content response has a 4xx status code
func (o *ResourcesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this resources delete no content response has a 5xx status code
func (o *ResourcesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this resources delete no content response a status code equal to that given
func (o *ResourcesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the resources delete no content response
func (o *ResourcesDeleteNoContent) Code() int {
	return 204
}

func (o *ResourcesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /EventsManagement/Resources/{id}][%d] resourcesDeleteNoContent ", 204)
}

func (o *ResourcesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /EventsManagement/Resources/{id}][%d] resourcesDeleteNoContent ", 204)
}

func (o *ResourcesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
