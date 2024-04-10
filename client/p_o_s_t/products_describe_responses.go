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

// ProductsDescribeReader is a Reader for the ProductsDescribe structure.
type ProductsDescribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductsDescribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductsDescribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /TXN/Products/Describe] Products_Describe", response, response.Code())
	}
}

// NewProductsDescribeOK creates a ProductsDescribeOK with default headers values
func NewProductsDescribeOK() *ProductsDescribeOK {
	return &ProductsDescribeOK{}
}

/*
ProductsDescribeOK describes a response with status code 200, with default header values.

OK
*/
type ProductsDescribeOK struct {
	Payload []*models.ProductDescribeResponse
}

// IsSuccess returns true when this products describe o k response has a 2xx status code
func (o *ProductsDescribeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this products describe o k response has a 3xx status code
func (o *ProductsDescribeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this products describe o k response has a 4xx status code
func (o *ProductsDescribeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this products describe o k response has a 5xx status code
func (o *ProductsDescribeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this products describe o k response a status code equal to that given
func (o *ProductsDescribeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the products describe o k response
func (o *ProductsDescribeOK) Code() int {
	return 200
}

func (o *ProductsDescribeOK) Error() string {
	return fmt.Sprintf("[POST /TXN/Products/Describe][%d] productsDescribeOK  %+v", 200, o.Payload)
}

func (o *ProductsDescribeOK) String() string {
	return fmt.Sprintf("[POST /TXN/Products/Describe][%d] productsDescribeOK  %+v", 200, o.Payload)
}

func (o *ProductsDescribeOK) GetPayload() []*models.ProductDescribeResponse {
	return o.Payload
}

func (o *ProductsDescribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
