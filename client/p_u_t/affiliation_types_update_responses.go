// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// AffiliationTypesUpdateReader is a Reader for the AffiliationTypesUpdate structure.
type AffiliationTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AffiliationTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAffiliationTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAffiliationTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAffiliationTypesUpdateOK creates a AffiliationTypesUpdateOK with default headers values
func NewAffiliationTypesUpdateOK() *AffiliationTypesUpdateOK {
	return &AffiliationTypesUpdateOK{}
}

/*
AffiliationTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type AffiliationTypesUpdateOK struct {
	Payload *models.AffiliationType
}

// IsSuccess returns true when this affiliation types update o k response has a 2xx status code
func (o *AffiliationTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this affiliation types update o k response has a 3xx status code
func (o *AffiliationTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this affiliation types update o k response has a 4xx status code
func (o *AffiliationTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this affiliation types update o k response has a 5xx status code
func (o *AffiliationTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this affiliation types update o k response a status code equal to that given
func (o *AffiliationTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the affiliation types update o k response
func (o *AffiliationTypesUpdateOK) Code() int {
	return 200
}

func (o *AffiliationTypesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/AffiliationTypes/{id}][%d] affiliationTypesUpdateOK %s", 200, payload)
}

func (o *AffiliationTypesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/AffiliationTypes/{id}][%d] affiliationTypesUpdateOK %s", 200, payload)
}

func (o *AffiliationTypesUpdateOK) GetPayload() *models.AffiliationType {
	return o.Payload
}

func (o *AffiliationTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AffiliationType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAffiliationTypesUpdateDefault creates a AffiliationTypesUpdateDefault with default headers values
func NewAffiliationTypesUpdateDefault(code int) *AffiliationTypesUpdateDefault {
	return &AffiliationTypesUpdateDefault{
		_statusCode: code,
	}
}

/*
AffiliationTypesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type AffiliationTypesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this affiliation types update default response has a 2xx status code
func (o *AffiliationTypesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this affiliation types update default response has a 3xx status code
func (o *AffiliationTypesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this affiliation types update default response has a 4xx status code
func (o *AffiliationTypesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this affiliation types update default response has a 5xx status code
func (o *AffiliationTypesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this affiliation types update default response a status code equal to that given
func (o *AffiliationTypesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the affiliation types update default response
func (o *AffiliationTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *AffiliationTypesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/AffiliationTypes/{id}][%d] AffiliationTypes_Update default %s", o._statusCode, payload)
}

func (o *AffiliationTypesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/AffiliationTypes/{id}][%d] AffiliationTypes_Update default %s", o._statusCode, payload)
}

func (o *AffiliationTypesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AffiliationTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
