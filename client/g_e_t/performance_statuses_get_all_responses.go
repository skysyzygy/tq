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

// PerformanceStatusesGetAllReader is a Reader for the PerformanceStatusesGetAll structure.
type PerformanceStatusesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceStatusesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceStatusesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/PerformanceStatuses] PerformanceStatuses_GetAll", response, response.Code())
	}
}

// NewPerformanceStatusesGetAllOK creates a PerformanceStatusesGetAllOK with default headers values
func NewPerformanceStatusesGetAllOK() *PerformanceStatusesGetAllOK {
	return &PerformanceStatusesGetAllOK{}
}

/*
PerformanceStatusesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceStatusesGetAllOK struct {
	Payload []*models.PerformanceStatus
}

// IsSuccess returns true when this performance statuses get all o k response has a 2xx status code
func (o *PerformanceStatusesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance statuses get all o k response has a 3xx status code
func (o *PerformanceStatusesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance statuses get all o k response has a 4xx status code
func (o *PerformanceStatusesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance statuses get all o k response has a 5xx status code
func (o *PerformanceStatusesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance statuses get all o k response a status code equal to that given
func (o *PerformanceStatusesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance statuses get all o k response
func (o *PerformanceStatusesGetAllOK) Code() int {
	return 200
}

func (o *PerformanceStatusesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/PerformanceStatuses][%d] performanceStatusesGetAllOK  %+v", 200, o.Payload)
}

func (o *PerformanceStatusesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/PerformanceStatuses][%d] performanceStatusesGetAllOK  %+v", 200, o.Payload)
}

func (o *PerformanceStatusesGetAllOK) GetPayload() []*models.PerformanceStatus {
	return o.Payload
}

func (o *PerformanceStatusesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}