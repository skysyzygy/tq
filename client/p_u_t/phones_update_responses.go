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

// PhonesUpdateReader is a Reader for the PhonesUpdate structure.
type PhonesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhonesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhonesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /CRM/Phones/{phoneId}] Phones_Update", response, response.Code())
	}
}

// NewPhonesUpdateOK creates a PhonesUpdateOK with default headers values
func NewPhonesUpdateOK() *PhonesUpdateOK {
	return &PhonesUpdateOK{}
}

/*
PhonesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PhonesUpdateOK struct {
	Payload *models.Phone
}

// IsSuccess returns true when this phones update o k response has a 2xx status code
func (o *PhonesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this phones update o k response has a 3xx status code
func (o *PhonesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this phones update o k response has a 4xx status code
func (o *PhonesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this phones update o k response has a 5xx status code
func (o *PhonesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this phones update o k response a status code equal to that given
func (o *PhonesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the phones update o k response
func (o *PhonesUpdateOK) Code() int {
	return 200
}

func (o *PhonesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /CRM/Phones/{phoneId}][%d] phonesUpdateOK  %+v", 200, o.Payload)
}

func (o *PhonesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /CRM/Phones/{phoneId}][%d] phonesUpdateOK  %+v", 200, o.Payload)
}

func (o *PhonesUpdateOK) GetPayload() *models.Phone {
	return o.Payload
}

func (o *PhonesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Phone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
