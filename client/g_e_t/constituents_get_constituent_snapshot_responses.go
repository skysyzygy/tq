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

// ConstituentsGetConstituentSnapshotReader is a Reader for the ConstituentsGetConstituentSnapshot structure.
type ConstituentsGetConstituentSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentsGetConstituentSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentsGetConstituentSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConstituentsGetConstituentSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConstituentsGetConstituentSnapshotOK creates a ConstituentsGetConstituentSnapshotOK with default headers values
func NewConstituentsGetConstituentSnapshotOK() *ConstituentsGetConstituentSnapshotOK {
	return &ConstituentsGetConstituentSnapshotOK{}
}

/*
ConstituentsGetConstituentSnapshotOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentsGetConstituentSnapshotOK struct {
	Payload *models.ConstituentSnapshot
}

// IsSuccess returns true when this constituents get constituent snapshot o k response has a 2xx status code
func (o *ConstituentsGetConstituentSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituents get constituent snapshot o k response has a 3xx status code
func (o *ConstituentsGetConstituentSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituents get constituent snapshot o k response has a 4xx status code
func (o *ConstituentsGetConstituentSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituents get constituent snapshot o k response has a 5xx status code
func (o *ConstituentsGetConstituentSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituents get constituent snapshot o k response a status code equal to that given
func (o *ConstituentsGetConstituentSnapshotOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituents get constituent snapshot o k response
func (o *ConstituentsGetConstituentSnapshotOK) Code() int {
	return 200
}

func (o *ConstituentsGetConstituentSnapshotOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/{constituentId}/Snapshot][%d] constituentsGetConstituentSnapshotOK %s", 200, payload)
}

func (o *ConstituentsGetConstituentSnapshotOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/{constituentId}/Snapshot][%d] constituentsGetConstituentSnapshotOK %s", 200, payload)
}

func (o *ConstituentsGetConstituentSnapshotOK) GetPayload() *models.ConstituentSnapshot {
	return o.Payload
}

func (o *ConstituentsGetConstituentSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentSnapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConstituentsGetConstituentSnapshotDefault creates a ConstituentsGetConstituentSnapshotDefault with default headers values
func NewConstituentsGetConstituentSnapshotDefault(code int) *ConstituentsGetConstituentSnapshotDefault {
	return &ConstituentsGetConstituentSnapshotDefault{
		_statusCode: code,
	}
}

/*
ConstituentsGetConstituentSnapshotDefault describes a response with status code -1, with default header values.

Error
*/
type ConstituentsGetConstituentSnapshotDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this constituents get constituent snapshot default response has a 2xx status code
func (o *ConstituentsGetConstituentSnapshotDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this constituents get constituent snapshot default response has a 3xx status code
func (o *ConstituentsGetConstituentSnapshotDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this constituents get constituent snapshot default response has a 4xx status code
func (o *ConstituentsGetConstituentSnapshotDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this constituents get constituent snapshot default response has a 5xx status code
func (o *ConstituentsGetConstituentSnapshotDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this constituents get constituent snapshot default response a status code equal to that given
func (o *ConstituentsGetConstituentSnapshotDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the constituents get constituent snapshot default response
func (o *ConstituentsGetConstituentSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *ConstituentsGetConstituentSnapshotDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/{constituentId}/Snapshot][%d] Constituents_GetConstituentSnapshot default %s", o._statusCode, payload)
}

func (o *ConstituentsGetConstituentSnapshotDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/Constituents/{constituentId}/Snapshot][%d] Constituents_GetConstituentSnapshot default %s", o._statusCode, payload)
}

func (o *ConstituentsGetConstituentSnapshotDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ConstituentsGetConstituentSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
