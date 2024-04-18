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

// InventoryContactPermissionTypesGetAllReader is a Reader for the InventoryContactPermissionTypesGetAll structure.
type InventoryContactPermissionTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryContactPermissionTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryContactPermissionTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/InventoryContactPermissionTypes] InventoryContactPermissionTypes_GetAll", response, response.Code())
	}
}

// NewInventoryContactPermissionTypesGetAllOK creates a InventoryContactPermissionTypesGetAllOK with default headers values
func NewInventoryContactPermissionTypesGetAllOK() *InventoryContactPermissionTypesGetAllOK {
	return &InventoryContactPermissionTypesGetAllOK{}
}

/*
InventoryContactPermissionTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type InventoryContactPermissionTypesGetAllOK struct {
	Payload []*models.InventoryContactPermissionType
}

// IsSuccess returns true when this inventory contact permission types get all o k response has a 2xx status code
func (o *InventoryContactPermissionTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory contact permission types get all o k response has a 3xx status code
func (o *InventoryContactPermissionTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory contact permission types get all o k response has a 4xx status code
func (o *InventoryContactPermissionTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory contact permission types get all o k response has a 5xx status code
func (o *InventoryContactPermissionTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory contact permission types get all o k response a status code equal to that given
func (o *InventoryContactPermissionTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the inventory contact permission types get all o k response
func (o *InventoryContactPermissionTypesGetAllOK) Code() int {
	return 200
}

func (o *InventoryContactPermissionTypesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /TXN/InventoryContactPermissionTypes][%d] inventoryContactPermissionTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *InventoryContactPermissionTypesGetAllOK) String() string {
	return fmt.Sprintf("[GET /TXN/InventoryContactPermissionTypes][%d] inventoryContactPermissionTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *InventoryContactPermissionTypesGetAllOK) GetPayload() []*models.InventoryContactPermissionType {
	return o.Payload
}

func (o *InventoryContactPermissionTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}