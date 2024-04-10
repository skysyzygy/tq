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

// ServiceResourcesGetSummariesReader is a Reader for the ServiceResourcesGetSummaries structure.
type ServiceResourcesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceResourcesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceResourcesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ServiceResources/Summary] ServiceResources_GetSummaries", response, response.Code())
	}
}

// NewServiceResourcesGetSummariesOK creates a ServiceResourcesGetSummariesOK with default headers values
func NewServiceResourcesGetSummariesOK() *ServiceResourcesGetSummariesOK {
	return &ServiceResourcesGetSummariesOK{}
}

/*
ServiceResourcesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type ServiceResourcesGetSummariesOK struct {
	Payload []*models.ServiceResourceSummary
}

// IsSuccess returns true when this service resources get summaries o k response has a 2xx status code
func (o *ServiceResourcesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service resources get summaries o k response has a 3xx status code
func (o *ServiceResourcesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service resources get summaries o k response has a 4xx status code
func (o *ServiceResourcesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service resources get summaries o k response has a 5xx status code
func (o *ServiceResourcesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service resources get summaries o k response a status code equal to that given
func (o *ServiceResourcesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service resources get summaries o k response
func (o *ServiceResourcesGetSummariesOK) Code() int {
	return 200
}

func (o *ServiceResourcesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ServiceResources/Summary][%d] serviceResourcesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ServiceResourcesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ServiceResources/Summary][%d] serviceResourcesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ServiceResourcesGetSummariesOK) GetPayload() []*models.ServiceResourceSummary {
	return o.Payload
}

func (o *ServiceResourcesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
