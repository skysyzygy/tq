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

// ContributionImportSetsCreateReader is a Reader for the ContributionImportSetsCreate structure.
type ContributionImportSetsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContributionImportSetsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContributionImportSetsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContributionImportSetsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContributionImportSetsCreateOK creates a ContributionImportSetsCreateOK with default headers values
func NewContributionImportSetsCreateOK() *ContributionImportSetsCreateOK {
	return &ContributionImportSetsCreateOK{}
}

/*
ContributionImportSetsCreateOK describes a response with status code 200, with default header values.

OK
*/
type ContributionImportSetsCreateOK struct {
	Payload *models.ContributionImportSet
}

// IsSuccess returns true when this contribution import sets create o k response has a 2xx status code
func (o *ContributionImportSetsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contribution import sets create o k response has a 3xx status code
func (o *ContributionImportSetsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contribution import sets create o k response has a 4xx status code
func (o *ContributionImportSetsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contribution import sets create o k response has a 5xx status code
func (o *ContributionImportSetsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contribution import sets create o k response a status code equal to that given
func (o *ContributionImportSetsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contribution import sets create o k response
func (o *ContributionImportSetsCreateOK) Code() int {
	return 200
}

func (o *ContributionImportSetsCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ContributionImportSets][%d] contributionImportSetsCreateOK %s", 200, payload)
}

func (o *ContributionImportSetsCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ContributionImportSets][%d] contributionImportSetsCreateOK %s", 200, payload)
}

func (o *ContributionImportSetsCreateOK) GetPayload() *models.ContributionImportSet {
	return o.Payload
}

func (o *ContributionImportSetsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContributionImportSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContributionImportSetsCreateDefault creates a ContributionImportSetsCreateDefault with default headers values
func NewContributionImportSetsCreateDefault(code int) *ContributionImportSetsCreateDefault {
	return &ContributionImportSetsCreateDefault{
		_statusCode: code,
	}
}

/*
ContributionImportSetsCreateDefault describes a response with status code -1, with default header values.

Error
*/
type ContributionImportSetsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this contribution import sets create default response has a 2xx status code
func (o *ContributionImportSetsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contribution import sets create default response has a 3xx status code
func (o *ContributionImportSetsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contribution import sets create default response has a 4xx status code
func (o *ContributionImportSetsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contribution import sets create default response has a 5xx status code
func (o *ContributionImportSetsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contribution import sets create default response a status code equal to that given
func (o *ContributionImportSetsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contribution import sets create default response
func (o *ContributionImportSetsCreateDefault) Code() int {
	return o._statusCode
}

func (o *ContributionImportSetsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ContributionImportSets][%d] ContributionImportSets_Create default %s", o._statusCode, payload)
}

func (o *ContributionImportSetsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ContributionImportSets][%d] ContributionImportSets_Create default %s", o._statusCode, payload)
}

func (o *ContributionImportSetsCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ContributionImportSetsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
