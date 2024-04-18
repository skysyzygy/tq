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

// TriPOSCloudConfigurationsGetAllReader is a Reader for the TriPOSCloudConfigurationsGetAll structure.
type TriPOSCloudConfigurationsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriPOSCloudConfigurationsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTriPOSCloudConfigurationsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/TriPOSCloudConfigurations] TriPOSCloudConfigurations_GetAll", response, response.Code())
	}
}

// NewTriPOSCloudConfigurationsGetAllOK creates a TriPOSCloudConfigurationsGetAllOK with default headers values
func NewTriPOSCloudConfigurationsGetAllOK() *TriPOSCloudConfigurationsGetAllOK {
	return &TriPOSCloudConfigurationsGetAllOK{}
}

/*
TriPOSCloudConfigurationsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type TriPOSCloudConfigurationsGetAllOK struct {
	Payload []*models.TriPOSCloudConfiguration
}

// IsSuccess returns true when this tri p o s cloud configurations get all o k response has a 2xx status code
func (o *TriPOSCloudConfigurationsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tri p o s cloud configurations get all o k response has a 3xx status code
func (o *TriPOSCloudConfigurationsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tri p o s cloud configurations get all o k response has a 4xx status code
func (o *TriPOSCloudConfigurationsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tri p o s cloud configurations get all o k response has a 5xx status code
func (o *TriPOSCloudConfigurationsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tri p o s cloud configurations get all o k response a status code equal to that given
func (o *TriPOSCloudConfigurationsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the tri p o s cloud configurations get all o k response
func (o *TriPOSCloudConfigurationsGetAllOK) Code() int {
	return 200
}

func (o *TriPOSCloudConfigurationsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/TriPOSCloudConfigurations][%d] triPOSCloudConfigurationsGetAllOK  %+v", 200, o.Payload)
}

func (o *TriPOSCloudConfigurationsGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/TriPOSCloudConfigurations][%d] triPOSCloudConfigurationsGetAllOK  %+v", 200, o.Payload)
}

func (o *TriPOSCloudConfigurationsGetAllOK) GetPayload() []*models.TriPOSCloudConfiguration {
	return o.Payload
}

func (o *TriPOSCloudConfigurationsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}