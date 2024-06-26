// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// ProductKeywordsGetKeywordsReader is a Reader for the ProductKeywordsGetKeywords structure.
type ProductKeywordsGetKeywordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductKeywordsGetKeywordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductKeywordsGetKeywordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProductKeywordsGetKeywordsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProductKeywordsGetKeywordsOK creates a ProductKeywordsGetKeywordsOK with default headers values
func NewProductKeywordsGetKeywordsOK() *ProductKeywordsGetKeywordsOK {
	return &ProductKeywordsGetKeywordsOK{}
}

/*
ProductKeywordsGetKeywordsOK describes a response with status code 200, with default header values.

OK
*/
type ProductKeywordsGetKeywordsOK struct {
	Payload []*models.KeywordResult
}

// IsSuccess returns true when this product keywords get keywords o k response has a 2xx status code
func (o *ProductKeywordsGetKeywordsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this product keywords get keywords o k response has a 3xx status code
func (o *ProductKeywordsGetKeywordsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this product keywords get keywords o k response has a 4xx status code
func (o *ProductKeywordsGetKeywordsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this product keywords get keywords o k response has a 5xx status code
func (o *ProductKeywordsGetKeywordsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this product keywords get keywords o k response a status code equal to that given
func (o *ProductKeywordsGetKeywordsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the product keywords get keywords o k response
func (o *ProductKeywordsGetKeywordsOK) Code() int {
	return 200
}

func (o *ProductKeywordsGetKeywordsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/ProductKeywords][%d] productKeywordsGetKeywordsOK %s", 200, payload)
}

func (o *ProductKeywordsGetKeywordsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/ProductKeywords][%d] productKeywordsGetKeywordsOK %s", 200, payload)
}

func (o *ProductKeywordsGetKeywordsOK) GetPayload() []*models.KeywordResult {
	return o.Payload
}

func (o *ProductKeywordsGetKeywordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductKeywordsGetKeywordsDefault creates a ProductKeywordsGetKeywordsDefault with default headers values
func NewProductKeywordsGetKeywordsDefault(code int) *ProductKeywordsGetKeywordsDefault {
	return &ProductKeywordsGetKeywordsDefault{
		_statusCode: code,
	}
}

/*
ProductKeywordsGetKeywordsDefault describes a response with status code -1, with default header values.

Error
*/
type ProductKeywordsGetKeywordsDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this product keywords get keywords default response has a 2xx status code
func (o *ProductKeywordsGetKeywordsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this product keywords get keywords default response has a 3xx status code
func (o *ProductKeywordsGetKeywordsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this product keywords get keywords default response has a 4xx status code
func (o *ProductKeywordsGetKeywordsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this product keywords get keywords default response has a 5xx status code
func (o *ProductKeywordsGetKeywordsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this product keywords get keywords default response a status code equal to that given
func (o *ProductKeywordsGetKeywordsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the product keywords get keywords default response
func (o *ProductKeywordsGetKeywordsDefault) Code() int {
	return o._statusCode
}

func (o *ProductKeywordsGetKeywordsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/ProductKeywords][%d] ProductKeywords_GetKeywords default %s", o._statusCode, payload)
}

func (o *ProductKeywordsGetKeywordsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/ProductKeywords][%d] ProductKeywords_GetKeywords default %s", o._statusCode, payload)
}

func (o *ProductKeywordsGetKeywordsDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ProductKeywordsGetKeywordsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
