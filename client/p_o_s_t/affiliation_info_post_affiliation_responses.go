// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// AffiliationInfoPostAffiliationReader is a Reader for the AffiliationInfoPostAffiliation structure.
type AffiliationInfoPostAffiliationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AffiliationInfoPostAffiliationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAffiliationInfoPostAffiliationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAffiliationInfoPostAffiliationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAffiliationInfoPostAffiliationOK creates a AffiliationInfoPostAffiliationOK with default headers values
func NewAffiliationInfoPostAffiliationOK() *AffiliationInfoPostAffiliationOK {
	return &AffiliationInfoPostAffiliationOK{}
}

/*
AffiliationInfoPostAffiliationOK describes a response with status code 200, with default header values.

OK
*/
type AffiliationInfoPostAffiliationOK struct {
	Payload *models.AffiliationInfo
}

// IsSuccess returns true when this affiliation info post affiliation o k response has a 2xx status code
func (o *AffiliationInfoPostAffiliationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this affiliation info post affiliation o k response has a 3xx status code
func (o *AffiliationInfoPostAffiliationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this affiliation info post affiliation o k response has a 4xx status code
func (o *AffiliationInfoPostAffiliationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this affiliation info post affiliation o k response has a 5xx status code
func (o *AffiliationInfoPostAffiliationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this affiliation info post affiliation o k response a status code equal to that given
func (o *AffiliationInfoPostAffiliationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the affiliation info post affiliation o k response
func (o *AffiliationInfoPostAffiliationOK) Code() int {
	return 200
}

func (o *AffiliationInfoPostAffiliationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/AffiliationInfo][%d] affiliationInfoPostAffiliationOK %s", 200, payload)
}

func (o *AffiliationInfoPostAffiliationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/AffiliationInfo][%d] affiliationInfoPostAffiliationOK %s", 200, payload)
}

func (o *AffiliationInfoPostAffiliationOK) GetPayload() *models.AffiliationInfo {
	return o.Payload
}

func (o *AffiliationInfoPostAffiliationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AffiliationInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAffiliationInfoPostAffiliationDefault creates a AffiliationInfoPostAffiliationDefault with default headers values
func NewAffiliationInfoPostAffiliationDefault(code int) *AffiliationInfoPostAffiliationDefault {
	return &AffiliationInfoPostAffiliationDefault{
		_statusCode: code,
	}
}

/*
AffiliationInfoPostAffiliationDefault describes a response with status code -1, with default header values.

Error
*/
type AffiliationInfoPostAffiliationDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this affiliation info post affiliation default response has a 2xx status code
func (o *AffiliationInfoPostAffiliationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this affiliation info post affiliation default response has a 3xx status code
func (o *AffiliationInfoPostAffiliationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this affiliation info post affiliation default response has a 4xx status code
func (o *AffiliationInfoPostAffiliationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this affiliation info post affiliation default response has a 5xx status code
func (o *AffiliationInfoPostAffiliationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this affiliation info post affiliation default response a status code equal to that given
func (o *AffiliationInfoPostAffiliationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the affiliation info post affiliation default response
func (o *AffiliationInfoPostAffiliationDefault) Code() int {
	return o._statusCode
}

func (o *AffiliationInfoPostAffiliationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/AffiliationInfo][%d] AffiliationInfo_PostAffiliation default %s", o._statusCode, payload)
}

func (o *AffiliationInfoPostAffiliationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/AffiliationInfo][%d] AffiliationInfo_PostAffiliation default %s", o._statusCode, payload)
}

func (o *AffiliationInfoPostAffiliationDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AffiliationInfoPostAffiliationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
