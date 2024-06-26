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

// BulkCopySetsGetAllReader is a Reader for the BulkCopySetsGetAll structure.
type BulkCopySetsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkCopySetsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkCopySetsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBulkCopySetsGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkCopySetsGetAllOK creates a BulkCopySetsGetAllOK with default headers values
func NewBulkCopySetsGetAllOK() *BulkCopySetsGetAllOK {
	return &BulkCopySetsGetAllOK{}
}

/*
BulkCopySetsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type BulkCopySetsGetAllOK struct {
	Payload []*models.BulkCopySet
}

// IsSuccess returns true when this bulk copy sets get all o k response has a 2xx status code
func (o *BulkCopySetsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bulk copy sets get all o k response has a 3xx status code
func (o *BulkCopySetsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk copy sets get all o k response has a 4xx status code
func (o *BulkCopySetsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk copy sets get all o k response has a 5xx status code
func (o *BulkCopySetsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk copy sets get all o k response a status code equal to that given
func (o *BulkCopySetsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bulk copy sets get all o k response
func (o *BulkCopySetsGetAllOK) Code() int {
	return 200
}

func (o *BulkCopySetsGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/BulkCopySets][%d] bulkCopySetsGetAllOK %s", 200, payload)
}

func (o *BulkCopySetsGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/BulkCopySets][%d] bulkCopySetsGetAllOK %s", 200, payload)
}

func (o *BulkCopySetsGetAllOK) GetPayload() []*models.BulkCopySet {
	return o.Payload
}

func (o *BulkCopySetsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkCopySetsGetAllDefault creates a BulkCopySetsGetAllDefault with default headers values
func NewBulkCopySetsGetAllDefault(code int) *BulkCopySetsGetAllDefault {
	return &BulkCopySetsGetAllDefault{
		_statusCode: code,
	}
}

/*
BulkCopySetsGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type BulkCopySetsGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this bulk copy sets get all default response has a 2xx status code
func (o *BulkCopySetsGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bulk copy sets get all default response has a 3xx status code
func (o *BulkCopySetsGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bulk copy sets get all default response has a 4xx status code
func (o *BulkCopySetsGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bulk copy sets get all default response has a 5xx status code
func (o *BulkCopySetsGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bulk copy sets get all default response a status code equal to that given
func (o *BulkCopySetsGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the bulk copy sets get all default response
func (o *BulkCopySetsGetAllDefault) Code() int {
	return o._statusCode
}

func (o *BulkCopySetsGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/BulkCopySets][%d] BulkCopySets_GetAll default %s", o._statusCode, payload)
}

func (o *BulkCopySetsGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/BulkCopySets][%d] BulkCopySets_GetAll default %s", o._statusCode, payload)
}

func (o *BulkCopySetsGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *BulkCopySetsGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
