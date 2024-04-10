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

// ConstituentsSchedulePurgeReader is a Reader for the ConstituentsSchedulePurge structure.
type ConstituentsSchedulePurgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentsSchedulePurgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentsSchedulePurgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /CRM/Constituents/{constituentId}/Purge/Schedule] Constituents_SchedulePurge", response, response.Code())
	}
}

// NewConstituentsSchedulePurgeOK creates a ConstituentsSchedulePurgeOK with default headers values
func NewConstituentsSchedulePurgeOK() *ConstituentsSchedulePurgeOK {
	return &ConstituentsSchedulePurgeOK{}
}

/*
ConstituentsSchedulePurgeOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentsSchedulePurgeOK struct {
	Payload *models.ConstituentSnapshot
}

// IsSuccess returns true when this constituents schedule purge o k response has a 2xx status code
func (o *ConstituentsSchedulePurgeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituents schedule purge o k response has a 3xx status code
func (o *ConstituentsSchedulePurgeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituents schedule purge o k response has a 4xx status code
func (o *ConstituentsSchedulePurgeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituents schedule purge o k response has a 5xx status code
func (o *ConstituentsSchedulePurgeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituents schedule purge o k response a status code equal to that given
func (o *ConstituentsSchedulePurgeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituents schedule purge o k response
func (o *ConstituentsSchedulePurgeOK) Code() int {
	return 200
}

func (o *ConstituentsSchedulePurgeOK) Error() string {
	return fmt.Sprintf("[POST /CRM/Constituents/{constituentId}/Purge/Schedule][%d] constituentsSchedulePurgeOK  %+v", 200, o.Payload)
}

func (o *ConstituentsSchedulePurgeOK) String() string {
	return fmt.Sprintf("[POST /CRM/Constituents/{constituentId}/Purge/Schedule][%d] constituentsSchedulePurgeOK  %+v", 200, o.Payload)
}

func (o *ConstituentsSchedulePurgeOK) GetPayload() *models.ConstituentSnapshot {
	return o.Payload
}

func (o *ConstituentsSchedulePurgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
