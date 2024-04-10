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

// InventoryWebContentsCreateReader is a Reader for the InventoryWebContentsCreate structure.
type InventoryWebContentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryWebContentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInventoryWebContentsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Txn/InventoryWebContents] InventoryWebContents_Create", response, response.Code())
	}
}

// NewInventoryWebContentsCreateOK creates a InventoryWebContentsCreateOK with default headers values
func NewInventoryWebContentsCreateOK() *InventoryWebContentsCreateOK {
	return &InventoryWebContentsCreateOK{}
}

/*
InventoryWebContentsCreateOK describes a response with status code 200, with default header values.

OK
*/
type InventoryWebContentsCreateOK struct {
	Payload *models.InventoryWebContent
}

// IsSuccess returns true when this inventory web contents create o k response has a 2xx status code
func (o *InventoryWebContentsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory web contents create o k response has a 3xx status code
func (o *InventoryWebContentsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory web contents create o k response has a 4xx status code
func (o *InventoryWebContentsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory web contents create o k response has a 5xx status code
func (o *InventoryWebContentsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory web contents create o k response a status code equal to that given
func (o *InventoryWebContentsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the inventory web contents create o k response
func (o *InventoryWebContentsCreateOK) Code() int {
	return 200
}

func (o *InventoryWebContentsCreateOK) Error() string {
	return fmt.Sprintf("[POST /Txn/InventoryWebContents][%d] inventoryWebContentsCreateOK  %+v", 200, o.Payload)
}

func (o *InventoryWebContentsCreateOK) String() string {
	return fmt.Sprintf("[POST /Txn/InventoryWebContents][%d] inventoryWebContentsCreateOK  %+v", 200, o.Payload)
}

func (o *InventoryWebContentsCreateOK) GetPayload() *models.InventoryWebContent {
	return o.Payload
}

func (o *InventoryWebContentsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryWebContent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
