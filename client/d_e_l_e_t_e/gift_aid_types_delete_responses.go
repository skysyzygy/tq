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

// GiftAidTypesDeleteReader is a Reader for the GiftAidTypesDelete structure.
type GiftAidTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GiftAidTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGiftAidTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGiftAidTypesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGiftAidTypesDeleteNoContent creates a GiftAidTypesDeleteNoContent with default headers values
func NewGiftAidTypesDeleteNoContent() *GiftAidTypesDeleteNoContent {
	return &GiftAidTypesDeleteNoContent{}
}

/*
GiftAidTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type GiftAidTypesDeleteNoContent struct {
}

// IsSuccess returns true when this gift aid types delete no content response has a 2xx status code
func (o *GiftAidTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this gift aid types delete no content response has a 3xx status code
func (o *GiftAidTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this gift aid types delete no content response has a 4xx status code
func (o *GiftAidTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this gift aid types delete no content response has a 5xx status code
func (o *GiftAidTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this gift aid types delete no content response a status code equal to that given
func (o *GiftAidTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the gift aid types delete no content response
func (o *GiftAidTypesDeleteNoContent) Code() int {
	return 204
}

func (o *GiftAidTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/GiftAidTypes/{id}][%d] giftAidTypesDeleteNoContent", 204)
}

func (o *GiftAidTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/GiftAidTypes/{id}][%d] giftAidTypesDeleteNoContent", 204)
}

func (o *GiftAidTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGiftAidTypesDeleteDefault creates a GiftAidTypesDeleteDefault with default headers values
func NewGiftAidTypesDeleteDefault(code int) *GiftAidTypesDeleteDefault {
	return &GiftAidTypesDeleteDefault{
		_statusCode: code,
	}
}

/*
GiftAidTypesDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type GiftAidTypesDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this gift aid types delete default response has a 2xx status code
func (o *GiftAidTypesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this gift aid types delete default response has a 3xx status code
func (o *GiftAidTypesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this gift aid types delete default response has a 4xx status code
func (o *GiftAidTypesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this gift aid types delete default response has a 5xx status code
func (o *GiftAidTypesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this gift aid types delete default response a status code equal to that given
func (o *GiftAidTypesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the gift aid types delete default response
func (o *GiftAidTypesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *GiftAidTypesDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/GiftAidTypes/{id}][%d] GiftAidTypes_Delete default %s", o._statusCode, payload)
}

func (o *GiftAidTypesDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/GiftAidTypes/{id}][%d] GiftAidTypes_Delete default %s", o._statusCode, payload)
}

func (o *GiftAidTypesDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *GiftAidTypesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
