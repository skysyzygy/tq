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

// CartAddPackageItemReader is a Reader for the CartAddPackageItem structure.
type CartAddPackageItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartAddPackageItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCartAddPackageItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Web/Cart/{sessionKey}/Packages/Fixed] Cart_AddPackageItem", response, response.Code())
	}
}

// NewCartAddPackageItemOK creates a CartAddPackageItemOK with default headers values
func NewCartAddPackageItemOK() *CartAddPackageItemOK {
	return &CartAddPackageItemOK{}
}

/*
CartAddPackageItemOK describes a response with status code 200, with default header values.

OK
*/
type CartAddPackageItemOK struct {
	Payload *models.AddPackageItemResponse
}

// IsSuccess returns true when this cart add package item o k response has a 2xx status code
func (o *CartAddPackageItemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart add package item o k response has a 3xx status code
func (o *CartAddPackageItemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart add package item o k response has a 4xx status code
func (o *CartAddPackageItemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart add package item o k response has a 5xx status code
func (o *CartAddPackageItemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cart add package item o k response a status code equal to that given
func (o *CartAddPackageItemOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cart add package item o k response
func (o *CartAddPackageItemOK) Code() int {
	return 200
}

func (o *CartAddPackageItemOK) Error() string {
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Packages/Fixed][%d] cartAddPackageItemOK  %+v", 200, o.Payload)
}

func (o *CartAddPackageItemOK) String() string {
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Packages/Fixed][%d] cartAddPackageItemOK  %+v", 200, o.Payload)
}

func (o *CartAddPackageItemOK) GetPayload() *models.AddPackageItemResponse {
	return o.Payload
}

func (o *CartAddPackageItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddPackageItemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
