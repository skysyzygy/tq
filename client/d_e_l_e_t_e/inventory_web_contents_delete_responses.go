// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
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
		return nil, runtime.NewAPIError("[DELETE /Txn/InventoryWebContents/{inventoryWebContentId}] InventoryWebContents_Delete", response, response.Code())
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
	return fmt.Sprintf("[DELETE /Txn/InventoryWebContents/{inventoryWebContentId}][%d] inventoryWebContentsDeleteNoContent ", 204)
}

func (o *InventoryWebContentsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /Txn/InventoryWebContents/{inventoryWebContentId}][%d] inventoryWebContentsDeleteNoContent ", 204)
}

func (o *InventoryWebContentsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}