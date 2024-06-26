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

// MembershipOrganizationsGetAllReader is a Reader for the MembershipOrganizationsGetAll structure.
type MembershipOrganizationsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MembershipOrganizationsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMembershipOrganizationsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMembershipOrganizationsGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMembershipOrganizationsGetAllOK creates a MembershipOrganizationsGetAllOK with default headers values
func NewMembershipOrganizationsGetAllOK() *MembershipOrganizationsGetAllOK {
	return &MembershipOrganizationsGetAllOK{}
}

/*
MembershipOrganizationsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type MembershipOrganizationsGetAllOK struct {
	Payload []*models.MembershipOrganization
}

// IsSuccess returns true when this membership organizations get all o k response has a 2xx status code
func (o *MembershipOrganizationsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this membership organizations get all o k response has a 3xx status code
func (o *MembershipOrganizationsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this membership organizations get all o k response has a 4xx status code
func (o *MembershipOrganizationsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this membership organizations get all o k response has a 5xx status code
func (o *MembershipOrganizationsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this membership organizations get all o k response a status code equal to that given
func (o *MembershipOrganizationsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the membership organizations get all o k response
func (o *MembershipOrganizationsGetAllOK) Code() int {
	return 200
}

func (o *MembershipOrganizationsGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/MembershipOrganizations][%d] membershipOrganizationsGetAllOK %s", 200, payload)
}

func (o *MembershipOrganizationsGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/MembershipOrganizations][%d] membershipOrganizationsGetAllOK %s", 200, payload)
}

func (o *MembershipOrganizationsGetAllOK) GetPayload() []*models.MembershipOrganization {
	return o.Payload
}

func (o *MembershipOrganizationsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMembershipOrganizationsGetAllDefault creates a MembershipOrganizationsGetAllDefault with default headers values
func NewMembershipOrganizationsGetAllDefault(code int) *MembershipOrganizationsGetAllDefault {
	return &MembershipOrganizationsGetAllDefault{
		_statusCode: code,
	}
}

/*
MembershipOrganizationsGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type MembershipOrganizationsGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this membership organizations get all default response has a 2xx status code
func (o *MembershipOrganizationsGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this membership organizations get all default response has a 3xx status code
func (o *MembershipOrganizationsGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this membership organizations get all default response has a 4xx status code
func (o *MembershipOrganizationsGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this membership organizations get all default response has a 5xx status code
func (o *MembershipOrganizationsGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this membership organizations get all default response a status code equal to that given
func (o *MembershipOrganizationsGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the membership organizations get all default response
func (o *MembershipOrganizationsGetAllDefault) Code() int {
	return o._statusCode
}

func (o *MembershipOrganizationsGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/MembershipOrganizations][%d] MembershipOrganizations_GetAll default %s", o._statusCode, payload)
}

func (o *MembershipOrganizationsGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/MembershipOrganizations][%d] MembershipOrganizations_GetAll default %s", o._statusCode, payload)
}

func (o *MembershipOrganizationsGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *MembershipOrganizationsGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
