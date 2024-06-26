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

// ZoneGroupsGetAllReader is a Reader for the ZoneGroupsGetAll structure.
type ZoneGroupsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ZoneGroupsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewZoneGroupsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewZoneGroupsGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewZoneGroupsGetAllOK creates a ZoneGroupsGetAllOK with default headers values
func NewZoneGroupsGetAllOK() *ZoneGroupsGetAllOK {
	return &ZoneGroupsGetAllOK{}
}

/*
ZoneGroupsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ZoneGroupsGetAllOK struct {
	Payload []*models.ZoneGroup
}

// IsSuccess returns true when this zone groups get all o k response has a 2xx status code
func (o *ZoneGroupsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this zone groups get all o k response has a 3xx status code
func (o *ZoneGroupsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this zone groups get all o k response has a 4xx status code
func (o *ZoneGroupsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this zone groups get all o k response has a 5xx status code
func (o *ZoneGroupsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this zone groups get all o k response a status code equal to that given
func (o *ZoneGroupsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the zone groups get all o k response
func (o *ZoneGroupsGetAllOK) Code() int {
	return 200
}

func (o *ZoneGroupsGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ZoneGroups][%d] zoneGroupsGetAllOK %s", 200, payload)
}

func (o *ZoneGroupsGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ZoneGroups][%d] zoneGroupsGetAllOK %s", 200, payload)
}

func (o *ZoneGroupsGetAllOK) GetPayload() []*models.ZoneGroup {
	return o.Payload
}

func (o *ZoneGroupsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewZoneGroupsGetAllDefault creates a ZoneGroupsGetAllDefault with default headers values
func NewZoneGroupsGetAllDefault(code int) *ZoneGroupsGetAllDefault {
	return &ZoneGroupsGetAllDefault{
		_statusCode: code,
	}
}

/*
ZoneGroupsGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type ZoneGroupsGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this zone groups get all default response has a 2xx status code
func (o *ZoneGroupsGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this zone groups get all default response has a 3xx status code
func (o *ZoneGroupsGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this zone groups get all default response has a 4xx status code
func (o *ZoneGroupsGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this zone groups get all default response has a 5xx status code
func (o *ZoneGroupsGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this zone groups get all default response a status code equal to that given
func (o *ZoneGroupsGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the zone groups get all default response
func (o *ZoneGroupsGetAllDefault) Code() int {
	return o._statusCode
}

func (o *ZoneGroupsGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ZoneGroups][%d] ZoneGroups_GetAll default %s", o._statusCode, payload)
}

func (o *ZoneGroupsGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ZoneGroups][%d] ZoneGroups_GetAll default %s", o._statusCode, payload)
}

func (o *ZoneGroupsGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ZoneGroupsGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
