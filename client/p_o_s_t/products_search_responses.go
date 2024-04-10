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

// ProductsSearchReader is a Reader for the ProductsSearch structure.
type ProductsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /TXN/Products/Search] Products_Search", response, response.Code())
	}
}

// NewProductsSearchOK creates a ProductsSearchOK with default headers values
func NewProductsSearchOK() *ProductsSearchOK {
	return &ProductsSearchOK{}
}

/*
ProductsSearchOK describes a response with status code 200, with default header values.

OK
*/
type ProductsSearchOK struct {
	Payload []*models.ProductSearchResponse
}

// IsSuccess returns true when this products search o k response has a 2xx status code
func (o *ProductsSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this products search o k response has a 3xx status code
func (o *ProductsSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this products search o k response has a 4xx status code
func (o *ProductsSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this products search o k response has a 5xx status code
func (o *ProductsSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this products search o k response a status code equal to that given
func (o *ProductsSearchOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the products search o k response
func (o *ProductsSearchOK) Code() int {
	return 200
}

func (o *ProductsSearchOK) Error() string {
	return fmt.Sprintf("[POST /TXN/Products/Search][%d] productsSearchOK  %+v", 200, o.Payload)
}

func (o *ProductsSearchOK) String() string {
	return fmt.Sprintf("[POST /TXN/Products/Search][%d] productsSearchOK  %+v", 200, o.Payload)
}

func (o *ProductsSearchOK) GetPayload() []*models.ProductSearchResponse {
	return o.Payload
}

func (o *ProductsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
