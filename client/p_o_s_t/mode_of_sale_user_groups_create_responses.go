// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ModeOfSaleUserGroupsCreateReader is a Reader for the ModeOfSaleUserGroupsCreate structure.
type ModeOfSaleUserGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModeOfSaleUserGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModeOfSaleUserGroupsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /TXN/ModeOfSaleUserGroups] ModeOfSaleUserGroups_Create", response, response.Code())
	}
}

// NewModeOfSaleUserGroupsCreateOK creates a ModeOfSaleUserGroupsCreateOK with default headers values
func NewModeOfSaleUserGroupsCreateOK() *ModeOfSaleUserGroupsCreateOK {
	return &ModeOfSaleUserGroupsCreateOK{}
}

/*
ModeOfSaleUserGroupsCreateOK describes a response with status code 200, with default header values.

OK
*/
type ModeOfSaleUserGroupsCreateOK struct {
	Payload *models.ModeOfSaleUserGroup
}

// IsSuccess returns true when this mode of sale user groups create o k response has a 2xx status code
func (o *ModeOfSaleUserGroupsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mode of sale user groups create o k response has a 3xx status code
func (o *ModeOfSaleUserGroupsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mode of sale user groups create o k response has a 4xx status code
func (o *ModeOfSaleUserGroupsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mode of sale user groups create o k response has a 5xx status code
func (o *ModeOfSaleUserGroupsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mode of sale user groups create o k response a status code equal to that given
func (o *ModeOfSaleUserGroupsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mode of sale user groups create o k response
func (o *ModeOfSaleUserGroupsCreateOK) Code() int {
	return 200
}

func (o *ModeOfSaleUserGroupsCreateOK) Error() string {
	return fmt.Sprintf("[POST /TXN/ModeOfSaleUserGroups][%d] modeOfSaleUserGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *ModeOfSaleUserGroupsCreateOK) String() string {
	return fmt.Sprintf("[POST /TXN/ModeOfSaleUserGroups][%d] modeOfSaleUserGroupsCreateOK  %+v", 200, o.Payload)
}

func (o *ModeOfSaleUserGroupsCreateOK) GetPayload() *models.ModeOfSaleUserGroup {
	return o.Payload
}

func (o *ModeOfSaleUserGroupsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModeOfSaleUserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
