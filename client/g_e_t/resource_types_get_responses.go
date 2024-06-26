// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ResourceTypesGetReader is a Reader for the ResourceTypesGet structure.
type ResourceTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResourceTypesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResourceTypesGetOK creates a ResourceTypesGetOK with default headers values
func NewResourceTypesGetOK() *ResourceTypesGetOK {
	return &ResourceTypesGetOK{}
}

/*
ResourceTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type ResourceTypesGetOK struct {
	Payload *models.ResourceType
}

// IsSuccess returns true when this resource types get o k response has a 2xx status code
func (o *ResourceTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource types get o k response has a 3xx status code
func (o *ResourceTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource types get o k response has a 4xx status code
func (o *ResourceTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource types get o k response has a 5xx status code
func (o *ResourceTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource types get o k response a status code equal to that given
func (o *ResourceTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource types get o k response
func (o *ResourceTypesGetOK) Code() int {
	return 200
}

func (o *ResourceTypesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /EventsManagement/ResourceTypes/{id}][%d] resourceTypesGetOK %s", 200, payload)
}

func (o *ResourceTypesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /EventsManagement/ResourceTypes/{id}][%d] resourceTypesGetOK %s", 200, payload)
}

func (o *ResourceTypesGetOK) GetPayload() *models.ResourceType {
	return o.Payload
}

func (o *ResourceTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResourceTypesGetDefault creates a ResourceTypesGetDefault with default headers values
func NewResourceTypesGetDefault(code int) *ResourceTypesGetDefault {
	return &ResourceTypesGetDefault{
		_statusCode: code,
	}
}

/*
ResourceTypesGetDefault describes a response with status code -1, with default header values.

Error
*/
type ResourceTypesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this resource types get default response has a 2xx status code
func (o *ResourceTypesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this resource types get default response has a 3xx status code
func (o *ResourceTypesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this resource types get default response has a 4xx status code
func (o *ResourceTypesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this resource types get default response has a 5xx status code
func (o *ResourceTypesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this resource types get default response a status code equal to that given
func (o *ResourceTypesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the resource types get default response
func (o *ResourceTypesGetDefault) Code() int {
	return o._statusCode
}

func (o *ResourceTypesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /EventsManagement/ResourceTypes/{id}][%d] ResourceTypes_Get default %s", o._statusCode, payload)
}

func (o *ResourceTypesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /EventsManagement/ResourceTypes/{id}][%d] ResourceTypes_Get default %s", o._statusCode, payload)
}

func (o *ResourceTypesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ResourceTypesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
