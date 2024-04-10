// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// PerformanceStatusesCreateReader is a Reader for the PerformanceStatusesCreate structure.
type PerformanceStatusesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceStatusesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceStatusesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/PerformanceStatuses] PerformanceStatuses_Create", response, response.Code())
	}
}

// NewPerformanceStatusesCreateOK creates a PerformanceStatusesCreateOK with default headers values
func NewPerformanceStatusesCreateOK() *PerformanceStatusesCreateOK {
	return &PerformanceStatusesCreateOK{}
}

/*
PerformanceStatusesCreateOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceStatusesCreateOK struct {
	Payload *models.PerformanceStatus
}

// IsSuccess returns true when this performance statuses create o k response has a 2xx status code
func (o *PerformanceStatusesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance statuses create o k response has a 3xx status code
func (o *PerformanceStatusesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance statuses create o k response has a 4xx status code
func (o *PerformanceStatusesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance statuses create o k response has a 5xx status code
func (o *PerformanceStatusesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance statuses create o k response a status code equal to that given
func (o *PerformanceStatusesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance statuses create o k response
func (o *PerformanceStatusesCreateOK) Code() int {
	return 200
}

func (o *PerformanceStatusesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/PerformanceStatuses][%d] performanceStatusesCreateOK  %+v", 200, o.Payload)
}

func (o *PerformanceStatusesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/PerformanceStatuses][%d] performanceStatusesCreateOK  %+v", 200, o.Payload)
}

func (o *PerformanceStatusesCreateOK) GetPayload() *models.PerformanceStatus {
	return o.Payload
}

func (o *PerformanceStatusesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
