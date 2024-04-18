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

// ResourcesGetAllReader is a Reader for the ResourcesGetAll structure.
type ResourcesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourcesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourcesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /EventsManagement/Resources] Resources_GetAll", response, response.Code())
	}
}

// NewResourcesGetAllOK creates a ResourcesGetAllOK with default headers values
func NewResourcesGetAllOK() *ResourcesGetAllOK {
	return &ResourcesGetAllOK{}
}

/*
ResourcesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ResourcesGetAllOK struct {
	Payload []*models.Resource
}

// IsSuccess returns true when this resources get all o k response has a 2xx status code
func (o *ResourcesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resources get all o k response has a 3xx status code
func (o *ResourcesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resources get all o k response has a 4xx status code
func (o *ResourcesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resources get all o k response has a 5xx status code
func (o *ResourcesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resources get all o k response a status code equal to that given
func (o *ResourcesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resources get all o k response
func (o *ResourcesGetAllOK) Code() int {
	return 200
}

func (o *ResourcesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /EventsManagement/Resources][%d] resourcesGetAllOK  %+v", 200, o.Payload)
}

func (o *ResourcesGetAllOK) String() string {
	return fmt.Sprintf("[GET /EventsManagement/Resources][%d] resourcesGetAllOK  %+v", 200, o.Payload)
}

func (o *ResourcesGetAllOK) GetPayload() []*models.Resource {
	return o.Payload
}

func (o *ResourcesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}