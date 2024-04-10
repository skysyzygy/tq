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

// InternalGetConstituentMiniSnapshotReader is a Reader for the InternalGetConstituentMiniSnapshot structure.
type InternalGetConstituentMiniSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalGetConstituentMiniSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalGetConstituentMiniSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Internal/ConstituentMiniSnapshot/{constituentId}] Internal_GetConstituentMiniSnapshot", response, response.Code())
	}
}

// NewInternalGetConstituentMiniSnapshotOK creates a InternalGetConstituentMiniSnapshotOK with default headers values
func NewInternalGetConstituentMiniSnapshotOK() *InternalGetConstituentMiniSnapshotOK {
	return &InternalGetConstituentMiniSnapshotOK{}
}

/*
InternalGetConstituentMiniSnapshotOK describes a response with status code 200, with default header values.

OK
*/
type InternalGetConstituentMiniSnapshotOK struct {
	Payload *models.ConstituentMiniSnapshot
}

// IsSuccess returns true when this internal get constituent mini snapshot o k response has a 2xx status code
func (o *InternalGetConstituentMiniSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal get constituent mini snapshot o k response has a 3xx status code
func (o *InternalGetConstituentMiniSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal get constituent mini snapshot o k response has a 4xx status code
func (o *InternalGetConstituentMiniSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal get constituent mini snapshot o k response has a 5xx status code
func (o *InternalGetConstituentMiniSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this internal get constituent mini snapshot o k response a status code equal to that given
func (o *InternalGetConstituentMiniSnapshotOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the internal get constituent mini snapshot o k response
func (o *InternalGetConstituentMiniSnapshotOK) Code() int {
	return 200
}

func (o *InternalGetConstituentMiniSnapshotOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Internal/ConstituentMiniSnapshot/{constituentId}][%d] internalGetConstituentMiniSnapshotOK  %+v", 200, o.Payload)
}

func (o *InternalGetConstituentMiniSnapshotOK) String() string {
	return fmt.Sprintf("[GET /CRM/Internal/ConstituentMiniSnapshot/{constituentId}][%d] internalGetConstituentMiniSnapshotOK  %+v", 200, o.Payload)
}

func (o *InternalGetConstituentMiniSnapshotOK) GetPayload() *models.ConstituentMiniSnapshot {
	return o.Payload
}

func (o *InternalGetConstituentMiniSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentMiniSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
