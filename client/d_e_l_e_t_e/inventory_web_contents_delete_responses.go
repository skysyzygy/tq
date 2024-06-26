// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// InventoryWebContentsDeleteReader is a Reader for the InventoryWebContentsDelete structure.
type InventoryWebContentsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InventoryWebContentsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInventoryWebContentsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInventoryWebContentsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInventoryWebContentsDeleteNoContent creates a InventoryWebContentsDeleteNoContent with default headers values
func NewInventoryWebContentsDeleteNoContent() *InventoryWebContentsDeleteNoContent {
	return &InventoryWebContentsDeleteNoContent{}
}

/*
InventoryWebContentsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type InventoryWebContentsDeleteNoContent struct {
}

// IsSuccess returns true when this inventory web contents delete no content response has a 2xx status code
func (o *InventoryWebContentsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inventory web contents delete no content response has a 3xx status code
func (o *InventoryWebContentsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inventory web contents delete no content response has a 4xx status code
func (o *InventoryWebContentsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this inventory web contents delete no content response has a 5xx status code
func (o *InventoryWebContentsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this inventory web contents delete no content response a status code equal to that given
func (o *InventoryWebContentsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the inventory web contents delete no content response
func (o *InventoryWebContentsDeleteNoContent) Code() int {
	return 204
}

func (o *InventoryWebContentsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Txn/InventoryWebContents/{inventoryWebContentId}][%d] inventoryWebContentsDeleteNoContent", 204)
}

func (o *InventoryWebContentsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /Txn/InventoryWebContents/{inventoryWebContentId}][%d] inventoryWebContentsDeleteNoContent", 204)
}

func (o *InventoryWebContentsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInventoryWebContentsDeleteDefault creates a InventoryWebContentsDeleteDefault with default headers values
func NewInventoryWebContentsDeleteDefault(code int) *InventoryWebContentsDeleteDefault {
	return &InventoryWebContentsDeleteDefault{
		_statusCode: code,
	}
}

/*
InventoryWebContentsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type InventoryWebContentsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this inventory web contents delete default response has a 2xx status code
func (o *InventoryWebContentsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this inventory web contents delete default response has a 3xx status code
func (o *InventoryWebContentsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this inventory web contents delete default response has a 4xx status code
func (o *InventoryWebContentsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this inventory web contents delete default response has a 5xx status code
func (o *InventoryWebContentsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this inventory web contents delete default response a status code equal to that given
func (o *InventoryWebContentsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the inventory web contents delete default response
func (o *InventoryWebContentsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *InventoryWebContentsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /Txn/InventoryWebContents/{inventoryWebContentId}][%d] InventoryWebContents_Delete default %s", o._statusCode, payload)
}

func (o *InventoryWebContentsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /Txn/InventoryWebContents/{inventoryWebContentId}][%d] InventoryWebContents_Delete default %s", o._statusCode, payload)
}

func (o *InventoryWebContentsDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *InventoryWebContentsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
