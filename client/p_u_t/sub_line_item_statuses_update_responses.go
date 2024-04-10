// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// SubLineItemStatusesUpdateReader is a Reader for the SubLineItemStatusesUpdate structure.
type SubLineItemStatusesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubLineItemStatusesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubLineItemStatusesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/SubLineItemStatuses/{id}] SubLineItemStatuses_Update", response, response.Code())
	}
}

// NewSubLineItemStatusesUpdateOK creates a SubLineItemStatusesUpdateOK with default headers values
func NewSubLineItemStatusesUpdateOK() *SubLineItemStatusesUpdateOK {
	return &SubLineItemStatusesUpdateOK{}
}

/*
SubLineItemStatusesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type SubLineItemStatusesUpdateOK struct {
	Payload *models.SubLineItemStatus
}

// IsSuccess returns true when this sub line item statuses update o k response has a 2xx status code
func (o *SubLineItemStatusesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sub line item statuses update o k response has a 3xx status code
func (o *SubLineItemStatusesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sub line item statuses update o k response has a 4xx status code
func (o *SubLineItemStatusesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sub line item statuses update o k response has a 5xx status code
func (o *SubLineItemStatusesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sub line item statuses update o k response a status code equal to that given
func (o *SubLineItemStatusesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sub line item statuses update o k response
func (o *SubLineItemStatusesUpdateOK) Code() int {
	return 200
}

func (o *SubLineItemStatusesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/SubLineItemStatuses/{id}][%d] subLineItemStatusesUpdateOK  %+v", 200, o.Payload)
}

func (o *SubLineItemStatusesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/SubLineItemStatuses/{id}][%d] subLineItemStatusesUpdateOK  %+v", 200, o.Payload)
}

func (o *SubLineItemStatusesUpdateOK) GetPayload() *models.SubLineItemStatus {
	return o.Payload
}

func (o *SubLineItemStatusesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubLineItemStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
