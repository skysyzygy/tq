// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// PremieresGetReader is a Reader for the PremieresGet structure.
type PremieresGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PremieresGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPremieresGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/Premieres/{id}] Premieres_Get", response, response.Code())
	}
}

// NewPremieresGetOK creates a PremieresGetOK with default headers values
func NewPremieresGetOK() *PremieresGetOK {
	return &PremieresGetOK{}
}

/*
PremieresGetOK describes a response with status code 200, with default header values.

OK
*/
type PremieresGetOK struct {
	Payload *models.Premiere
}

// IsSuccess returns true when this premieres get o k response has a 2xx status code
func (o *PremieresGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this premieres get o k response has a 3xx status code
func (o *PremieresGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this premieres get o k response has a 4xx status code
func (o *PremieresGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this premieres get o k response has a 5xx status code
func (o *PremieresGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this premieres get o k response a status code equal to that given
func (o *PremieresGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the premieres get o k response
func (o *PremieresGetOK) Code() int {
	return 200
}

func (o *PremieresGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/Premieres/{id}][%d] premieresGetOK  %+v", 200, o.Payload)
}

func (o *PremieresGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/Premieres/{id}][%d] premieresGetOK  %+v", 200, o.Payload)
}

func (o *PremieresGetOK) GetPayload() *models.Premiere {
	return o.Payload
}

func (o *PremieresGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Premiere)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}