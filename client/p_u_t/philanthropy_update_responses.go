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

// PhilanthropyUpdateReader is a Reader for the PhilanthropyUpdate structure.
type PhilanthropyUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhilanthropyUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhilanthropyUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /CRM/Philanthropy/{philanthropyEntryId}] Philanthropy_Update", response, response.Code())
	}
}

// NewPhilanthropyUpdateOK creates a PhilanthropyUpdateOK with default headers values
func NewPhilanthropyUpdateOK() *PhilanthropyUpdateOK {
	return &PhilanthropyUpdateOK{}
}

/*
PhilanthropyUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PhilanthropyUpdateOK struct {
	Payload *models.PhilanthropyEntry
}

// IsSuccess returns true when this philanthropy update o k response has a 2xx status code
func (o *PhilanthropyUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this philanthropy update o k response has a 3xx status code
func (o *PhilanthropyUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this philanthropy update o k response has a 4xx status code
func (o *PhilanthropyUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this philanthropy update o k response has a 5xx status code
func (o *PhilanthropyUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this philanthropy update o k response a status code equal to that given
func (o *PhilanthropyUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the philanthropy update o k response
func (o *PhilanthropyUpdateOK) Code() int {
	return 200
}

func (o *PhilanthropyUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /CRM/Philanthropy/{philanthropyEntryId}][%d] philanthropyUpdateOK  %+v", 200, o.Payload)
}

func (o *PhilanthropyUpdateOK) String() string {
	return fmt.Sprintf("[PUT /CRM/Philanthropy/{philanthropyEntryId}][%d] philanthropyUpdateOK  %+v", 200, o.Payload)
}

func (o *PhilanthropyUpdateOK) GetPayload() *models.PhilanthropyEntry {
	return o.Payload
}

func (o *PhilanthropyUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhilanthropyEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
