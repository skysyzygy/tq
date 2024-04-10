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

// OriginsUpdateReader is a Reader for the OriginsUpdate structure.
type OriginsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OriginsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOriginsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/Origins/{id}] Origins_Update", response, response.Code())
	}
}

// NewOriginsUpdateOK creates a OriginsUpdateOK with default headers values
func NewOriginsUpdateOK() *OriginsUpdateOK {
	return &OriginsUpdateOK{}
}

/*
OriginsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type OriginsUpdateOK struct {
	Payload *models.Origin
}

// IsSuccess returns true when this origins update o k response has a 2xx status code
func (o *OriginsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this origins update o k response has a 3xx status code
func (o *OriginsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this origins update o k response has a 4xx status code
func (o *OriginsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this origins update o k response has a 5xx status code
func (o *OriginsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this origins update o k response a status code equal to that given
func (o *OriginsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the origins update o k response
func (o *OriginsUpdateOK) Code() int {
	return 200
}

func (o *OriginsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/Origins/{id}][%d] originsUpdateOK  %+v", 200, o.Payload)
}

func (o *OriginsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/Origins/{id}][%d] originsUpdateOK  %+v", 200, o.Payload)
}

func (o *OriginsUpdateOK) GetPayload() *models.Origin {
	return o.Payload
}

func (o *OriginsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Origin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
