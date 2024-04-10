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

// ZoneGroupsGetReader is a Reader for the ZoneGroupsGet structure.
type ZoneGroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ZoneGroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewZoneGroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ZoneGroups/{id}] ZoneGroups_Get", response, response.Code())
	}
}

// NewZoneGroupsGetOK creates a ZoneGroupsGetOK with default headers values
func NewZoneGroupsGetOK() *ZoneGroupsGetOK {
	return &ZoneGroupsGetOK{}
}

/*
ZoneGroupsGetOK describes a response with status code 200, with default header values.

OK
*/
type ZoneGroupsGetOK struct {
	Payload *models.ZoneGroup
}

// IsSuccess returns true when this zone groups get o k response has a 2xx status code
func (o *ZoneGroupsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this zone groups get o k response has a 3xx status code
func (o *ZoneGroupsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this zone groups get o k response has a 4xx status code
func (o *ZoneGroupsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this zone groups get o k response has a 5xx status code
func (o *ZoneGroupsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this zone groups get o k response a status code equal to that given
func (o *ZoneGroupsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the zone groups get o k response
func (o *ZoneGroupsGetOK) Code() int {
	return 200
}

func (o *ZoneGroupsGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ZoneGroups/{id}][%d] zoneGroupsGetOK  %+v", 200, o.Payload)
}

func (o *ZoneGroupsGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ZoneGroups/{id}][%d] zoneGroupsGetOK  %+v", 200, o.Payload)
}

func (o *ZoneGroupsGetOK) GetPayload() *models.ZoneGroup {
	return o.Payload
}

func (o *ZoneGroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZoneGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
