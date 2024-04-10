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

// PriceTypeUserGroupsGetReader is a Reader for the PriceTypeUserGroupsGet structure.
type PriceTypeUserGroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceTypeUserGroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPriceTypeUserGroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/PriceTypeUserGroups/{priceTypeUserGroupId}] PriceTypeUserGroups_Get", response, response.Code())
	}
}

// NewPriceTypeUserGroupsGetOK creates a PriceTypeUserGroupsGetOK with default headers values
func NewPriceTypeUserGroupsGetOK() *PriceTypeUserGroupsGetOK {
	return &PriceTypeUserGroupsGetOK{}
}

/*
PriceTypeUserGroupsGetOK describes a response with status code 200, with default header values.

OK
*/
type PriceTypeUserGroupsGetOK struct {
	Payload *models.PriceTypeUserGroup
}

// IsSuccess returns true when this price type user groups get o k response has a 2xx status code
func (o *PriceTypeUserGroupsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price type user groups get o k response has a 3xx status code
func (o *PriceTypeUserGroupsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price type user groups get o k response has a 4xx status code
func (o *PriceTypeUserGroupsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this price type user groups get o k response has a 5xx status code
func (o *PriceTypeUserGroupsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this price type user groups get o k response a status code equal to that given
func (o *PriceTypeUserGroupsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the price type user groups get o k response
func (o *PriceTypeUserGroupsGetOK) Code() int {
	return 200
}

func (o *PriceTypeUserGroupsGetOK) Error() string {
	return fmt.Sprintf("[GET /TXN/PriceTypeUserGroups/{priceTypeUserGroupId}][%d] priceTypeUserGroupsGetOK  %+v", 200, o.Payload)
}

func (o *PriceTypeUserGroupsGetOK) String() string {
	return fmt.Sprintf("[GET /TXN/PriceTypeUserGroups/{priceTypeUserGroupId}][%d] priceTypeUserGroupsGetOK  %+v", 200, o.Payload)
}

func (o *PriceTypeUserGroupsGetOK) GetPayload() *models.PriceTypeUserGroup {
	return o.Payload
}

func (o *PriceTypeUserGroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PriceTypeUserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
