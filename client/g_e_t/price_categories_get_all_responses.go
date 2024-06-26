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

// PriceCategoriesGetAllReader is a Reader for the PriceCategoriesGetAll structure.
type PriceCategoriesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceCategoriesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPriceCategoriesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPriceCategoriesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPriceCategoriesGetAllOK creates a PriceCategoriesGetAllOK with default headers values
func NewPriceCategoriesGetAllOK() *PriceCategoriesGetAllOK {
	return &PriceCategoriesGetAllOK{}
}

/*
PriceCategoriesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PriceCategoriesGetAllOK struct {
	Payload []*models.PriceCategory
}

// IsSuccess returns true when this price categories get all o k response has a 2xx status code
func (o *PriceCategoriesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price categories get all o k response has a 3xx status code
func (o *PriceCategoriesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price categories get all o k response has a 4xx status code
func (o *PriceCategoriesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this price categories get all o k response has a 5xx status code
func (o *PriceCategoriesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this price categories get all o k response a status code equal to that given
func (o *PriceCategoriesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the price categories get all o k response
func (o *PriceCategoriesGetAllOK) Code() int {
	return 200
}

func (o *PriceCategoriesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PriceCategories][%d] priceCategoriesGetAllOK %s", 200, payload)
}

func (o *PriceCategoriesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PriceCategories][%d] priceCategoriesGetAllOK %s", 200, payload)
}

func (o *PriceCategoriesGetAllOK) GetPayload() []*models.PriceCategory {
	return o.Payload
}

func (o *PriceCategoriesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPriceCategoriesGetAllDefault creates a PriceCategoriesGetAllDefault with default headers values
func NewPriceCategoriesGetAllDefault(code int) *PriceCategoriesGetAllDefault {
	return &PriceCategoriesGetAllDefault{
		_statusCode: code,
	}
}

/*
PriceCategoriesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type PriceCategoriesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this price categories get all default response has a 2xx status code
func (o *PriceCategoriesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this price categories get all default response has a 3xx status code
func (o *PriceCategoriesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this price categories get all default response has a 4xx status code
func (o *PriceCategoriesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this price categories get all default response has a 5xx status code
func (o *PriceCategoriesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this price categories get all default response a status code equal to that given
func (o *PriceCategoriesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the price categories get all default response
func (o *PriceCategoriesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *PriceCategoriesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PriceCategories][%d] PriceCategories_GetAll default %s", o._statusCode, payload)
}

func (o *PriceCategoriesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PriceCategories][%d] PriceCategories_GetAll default %s", o._statusCode, payload)
}

func (o *PriceCategoriesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PriceCategoriesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
