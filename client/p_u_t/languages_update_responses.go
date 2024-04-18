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

// LanguagesUpdateReader is a Reader for the LanguagesUpdate structure.
type LanguagesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LanguagesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLanguagesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/Languages/{id}] Languages_Update", response, response.Code())
	}
}

// NewLanguagesUpdateOK creates a LanguagesUpdateOK with default headers values
func NewLanguagesUpdateOK() *LanguagesUpdateOK {
	return &LanguagesUpdateOK{}
}

/*
LanguagesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type LanguagesUpdateOK struct {
	Payload *models.Language
}

// IsSuccess returns true when this languages update o k response has a 2xx status code
func (o *LanguagesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this languages update o k response has a 3xx status code
func (o *LanguagesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this languages update o k response has a 4xx status code
func (o *LanguagesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this languages update o k response has a 5xx status code
func (o *LanguagesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this languages update o k response a status code equal to that given
func (o *LanguagesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the languages update o k response
func (o *LanguagesUpdateOK) Code() int {
	return 200
}

func (o *LanguagesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/Languages/{id}][%d] languagesUpdateOK  %+v", 200, o.Payload)
}

func (o *LanguagesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/Languages/{id}][%d] languagesUpdateOK  %+v", 200, o.Payload)
}

func (o *LanguagesUpdateOK) GetPayload() *models.Language {
	return o.Payload
}

func (o *LanguagesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Language)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}