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

// SeasonTypesCreateReader is a Reader for the SeasonTypesCreate structure.
type SeasonTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SeasonTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSeasonTypesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/SeasonTypes] SeasonTypes_Create", response, response.Code())
	}
}

// NewSeasonTypesCreateOK creates a SeasonTypesCreateOK with default headers values
func NewSeasonTypesCreateOK() *SeasonTypesCreateOK {
	return &SeasonTypesCreateOK{}
}

/*
SeasonTypesCreateOK describes a response with status code 200, with default header values.

OK
*/
type SeasonTypesCreateOK struct {
	Payload *models.SeasonType
}

// IsSuccess returns true when this season types create o k response has a 2xx status code
func (o *SeasonTypesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this season types create o k response has a 3xx status code
func (o *SeasonTypesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this season types create o k response has a 4xx status code
func (o *SeasonTypesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this season types create o k response has a 5xx status code
func (o *SeasonTypesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this season types create o k response a status code equal to that given
func (o *SeasonTypesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the season types create o k response
func (o *SeasonTypesCreateOK) Code() int {
	return 200
}

func (o *SeasonTypesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/SeasonTypes][%d] seasonTypesCreateOK  %+v", 200, o.Payload)
}

func (o *SeasonTypesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/SeasonTypes][%d] seasonTypesCreateOK  %+v", 200, o.Payload)
}

func (o *SeasonTypesCreateOK) GetPayload() *models.SeasonType {
	return o.Payload
}

func (o *SeasonTypesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SeasonType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}