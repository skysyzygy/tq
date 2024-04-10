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

// DesignsGetReader is a Reader for the DesignsGet structure.
type DesignsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesignsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDesignsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/Designs/{id}] Designs_Get", response, response.Code())
	}
}

// NewDesignsGetOK creates a DesignsGetOK with default headers values
func NewDesignsGetOK() *DesignsGetOK {
	return &DesignsGetOK{}
}

/*
DesignsGetOK describes a response with status code 200, with default header values.

OK
*/
type DesignsGetOK struct {
	Payload *models.Design
}

// IsSuccess returns true when this designs get o k response has a 2xx status code
func (o *DesignsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this designs get o k response has a 3xx status code
func (o *DesignsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this designs get o k response has a 4xx status code
func (o *DesignsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this designs get o k response has a 5xx status code
func (o *DesignsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this designs get o k response a status code equal to that given
func (o *DesignsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the designs get o k response
func (o *DesignsGetOK) Code() int {
	return 200
}

func (o *DesignsGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/Designs/{id}][%d] designsGetOK  %+v", 200, o.Payload)
}

func (o *DesignsGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/Designs/{id}][%d] designsGetOK  %+v", 200, o.Payload)
}

func (o *DesignsGetOK) GetPayload() *models.Design {
	return o.Payload
}

func (o *DesignsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Design)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
