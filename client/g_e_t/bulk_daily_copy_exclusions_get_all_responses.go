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

// BulkDailyCopyExclusionsGetAllReader is a Reader for the BulkDailyCopyExclusionsGetAll structure.
type BulkDailyCopyExclusionsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkDailyCopyExclusionsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkDailyCopyExclusionsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBulkDailyCopyExclusionsGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBulkDailyCopyExclusionsGetAllOK creates a BulkDailyCopyExclusionsGetAllOK with default headers values
func NewBulkDailyCopyExclusionsGetAllOK() *BulkDailyCopyExclusionsGetAllOK {
	return &BulkDailyCopyExclusionsGetAllOK{}
}

/*
BulkDailyCopyExclusionsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type BulkDailyCopyExclusionsGetAllOK struct {
	Payload []*models.BulkDailyCopyExclusion
}

// IsSuccess returns true when this bulk daily copy exclusions get all o k response has a 2xx status code
func (o *BulkDailyCopyExclusionsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bulk daily copy exclusions get all o k response has a 3xx status code
func (o *BulkDailyCopyExclusionsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk daily copy exclusions get all o k response has a 4xx status code
func (o *BulkDailyCopyExclusionsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk daily copy exclusions get all o k response has a 5xx status code
func (o *BulkDailyCopyExclusionsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk daily copy exclusions get all o k response a status code equal to that given
func (o *BulkDailyCopyExclusionsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bulk daily copy exclusions get all o k response
func (o *BulkDailyCopyExclusionsGetAllOK) Code() int {
	return 200
}

func (o *BulkDailyCopyExclusionsGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/BulkDailyCopyExclusions][%d] bulkDailyCopyExclusionsGetAllOK %s", 200, payload)
}

func (o *BulkDailyCopyExclusionsGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/BulkDailyCopyExclusions][%d] bulkDailyCopyExclusionsGetAllOK %s", 200, payload)
}

func (o *BulkDailyCopyExclusionsGetAllOK) GetPayload() []*models.BulkDailyCopyExclusion {
	return o.Payload
}

func (o *BulkDailyCopyExclusionsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBulkDailyCopyExclusionsGetAllDefault creates a BulkDailyCopyExclusionsGetAllDefault with default headers values
func NewBulkDailyCopyExclusionsGetAllDefault(code int) *BulkDailyCopyExclusionsGetAllDefault {
	return &BulkDailyCopyExclusionsGetAllDefault{
		_statusCode: code,
	}
}

/*
BulkDailyCopyExclusionsGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type BulkDailyCopyExclusionsGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this bulk daily copy exclusions get all default response has a 2xx status code
func (o *BulkDailyCopyExclusionsGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this bulk daily copy exclusions get all default response has a 3xx status code
func (o *BulkDailyCopyExclusionsGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this bulk daily copy exclusions get all default response has a 4xx status code
func (o *BulkDailyCopyExclusionsGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this bulk daily copy exclusions get all default response has a 5xx status code
func (o *BulkDailyCopyExclusionsGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this bulk daily copy exclusions get all default response a status code equal to that given
func (o *BulkDailyCopyExclusionsGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the bulk daily copy exclusions get all default response
func (o *BulkDailyCopyExclusionsGetAllDefault) Code() int {
	return o._statusCode
}

func (o *BulkDailyCopyExclusionsGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/BulkDailyCopyExclusions][%d] BulkDailyCopyExclusions_GetAll default %s", o._statusCode, payload)
}

func (o *BulkDailyCopyExclusionsGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/BulkDailyCopyExclusions][%d] BulkDailyCopyExclusions_GetAll default %s", o._statusCode, payload)
}

func (o *BulkDailyCopyExclusionsGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *BulkDailyCopyExclusionsGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
