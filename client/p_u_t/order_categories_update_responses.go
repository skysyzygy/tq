// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// OrderCategoriesUpdateReader is a Reader for the OrderCategoriesUpdate structure.
type OrderCategoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderCategoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrderCategoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/OrderCategories/{id}] OrderCategories_Update", response, response.Code())
	}
}

// NewOrderCategoriesUpdateOK creates a OrderCategoriesUpdateOK with default headers values
func NewOrderCategoriesUpdateOK() *OrderCategoriesUpdateOK {
	return &OrderCategoriesUpdateOK{}
}

/*
OrderCategoriesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type OrderCategoriesUpdateOK struct {
	Payload *models.OrderCategory
}

// IsSuccess returns true when this order categories update o k response has a 2xx status code
func (o *OrderCategoriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this order categories update o k response has a 3xx status code
func (o *OrderCategoriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this order categories update o k response has a 4xx status code
func (o *OrderCategoriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this order categories update o k response has a 5xx status code
func (o *OrderCategoriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this order categories update o k response a status code equal to that given
func (o *OrderCategoriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the order categories update o k response
func (o *OrderCategoriesUpdateOK) Code() int {
	return 200
}

func (o *OrderCategoriesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/OrderCategories/{id}][%d] orderCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *OrderCategoriesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/OrderCategories/{id}][%d] orderCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *OrderCategoriesUpdateOK) GetPayload() *models.OrderCategory {
	return o.Payload
}

func (o *OrderCategoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OrderCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}