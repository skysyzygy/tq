// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// SubLineItemStatusesDeleteReader is a Reader for the SubLineItemStatusesDelete structure.
type SubLineItemStatusesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubLineItemStatusesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSubLineItemStatusesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSubLineItemStatusesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSubLineItemStatusesDeleteNoContent creates a SubLineItemStatusesDeleteNoContent with default headers values
func NewSubLineItemStatusesDeleteNoContent() *SubLineItemStatusesDeleteNoContent {
	return &SubLineItemStatusesDeleteNoContent{}
}

/*
SubLineItemStatusesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type SubLineItemStatusesDeleteNoContent struct {
}

// IsSuccess returns true when this sub line item statuses delete no content response has a 2xx status code
func (o *SubLineItemStatusesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sub line item statuses delete no content response has a 3xx status code
func (o *SubLineItemStatusesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sub line item statuses delete no content response has a 4xx status code
func (o *SubLineItemStatusesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this sub line item statuses delete no content response has a 5xx status code
func (o *SubLineItemStatusesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this sub line item statuses delete no content response a status code equal to that given
func (o *SubLineItemStatusesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the sub line item statuses delete no content response
func (o *SubLineItemStatusesDeleteNoContent) Code() int {
	return 204
}

func (o *SubLineItemStatusesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/SubLineItemStatuses/{id}][%d] subLineItemStatusesDeleteNoContent", 204)
}

func (o *SubLineItemStatusesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/SubLineItemStatuses/{id}][%d] subLineItemStatusesDeleteNoContent", 204)
}

func (o *SubLineItemStatusesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubLineItemStatusesDeleteDefault creates a SubLineItemStatusesDeleteDefault with default headers values
func NewSubLineItemStatusesDeleteDefault(code int) *SubLineItemStatusesDeleteDefault {
	return &SubLineItemStatusesDeleteDefault{
		_statusCode: code,
	}
}

/*
SubLineItemStatusesDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type SubLineItemStatusesDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this sub line item statuses delete default response has a 2xx status code
func (o *SubLineItemStatusesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sub line item statuses delete default response has a 3xx status code
func (o *SubLineItemStatusesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sub line item statuses delete default response has a 4xx status code
func (o *SubLineItemStatusesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sub line item statuses delete default response has a 5xx status code
func (o *SubLineItemStatusesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sub line item statuses delete default response a status code equal to that given
func (o *SubLineItemStatusesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sub line item statuses delete default response
func (o *SubLineItemStatusesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SubLineItemStatusesDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/SubLineItemStatuses/{id}][%d] SubLineItemStatuses_Delete default %s", o._statusCode, payload)
}

func (o *SubLineItemStatusesDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/SubLineItemStatuses/{id}][%d] SubLineItemStatuses_Delete default %s", o._statusCode, payload)
}

func (o *SubLineItemStatusesDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SubLineItemStatusesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
