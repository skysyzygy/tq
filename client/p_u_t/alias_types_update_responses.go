// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// AliasTypesUpdateReader is a Reader for the AliasTypesUpdate structure.
type AliasTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AliasTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAliasTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/AliasTypes/{id}] AliasTypes_Update", response, response.Code())
	}
}

// NewAliasTypesUpdateOK creates a AliasTypesUpdateOK with default headers values
func NewAliasTypesUpdateOK() *AliasTypesUpdateOK {
	return &AliasTypesUpdateOK{}
}

/*
AliasTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type AliasTypesUpdateOK struct {
	Payload *models.AliasType
}

// IsSuccess returns true when this alias types update o k response has a 2xx status code
func (o *AliasTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this alias types update o k response has a 3xx status code
func (o *AliasTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this alias types update o k response has a 4xx status code
func (o *AliasTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this alias types update o k response has a 5xx status code
func (o *AliasTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this alias types update o k response a status code equal to that given
func (o *AliasTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the alias types update o k response
func (o *AliasTypesUpdateOK) Code() int {
	return 200
}

func (o *AliasTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/AliasTypes/{id}][%d] aliasTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *AliasTypesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/AliasTypes/{id}][%d] aliasTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *AliasTypesUpdateOK) GetPayload() *models.AliasType {
	return o.Payload
}

func (o *AliasTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AliasType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
