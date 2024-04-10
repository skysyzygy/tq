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

// AliasTypesCreateReader is a Reader for the AliasTypesCreate structure.
type AliasTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AliasTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAliasTypesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/AliasTypes] AliasTypes_Create", response, response.Code())
	}
}

// NewAliasTypesCreateOK creates a AliasTypesCreateOK with default headers values
func NewAliasTypesCreateOK() *AliasTypesCreateOK {
	return &AliasTypesCreateOK{}
}

/*
AliasTypesCreateOK describes a response with status code 200, with default header values.

OK
*/
type AliasTypesCreateOK struct {
	Payload *models.AliasType
}

// IsSuccess returns true when this alias types create o k response has a 2xx status code
func (o *AliasTypesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alias types create o k response has a 3xx status code
func (o *AliasTypesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alias types create o k response has a 4xx status code
func (o *AliasTypesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alias types create o k response has a 5xx status code
func (o *AliasTypesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alias types create o k response a status code equal to that given
func (o *AliasTypesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the alias types create o k response
func (o *AliasTypesCreateOK) Code() int {
	return 200
}

func (o *AliasTypesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/AliasTypes][%d] aliasTypesCreateOK  %+v", 200, o.Payload)
}

func (o *AliasTypesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/AliasTypes][%d] aliasTypesCreateOK  %+v", 200, o.Payload)
}

func (o *AliasTypesCreateOK) GetPayload() *models.AliasType {
	return o.Payload
}

func (o *AliasTypesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AliasType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
