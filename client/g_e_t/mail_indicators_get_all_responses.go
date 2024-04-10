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

// MailIndicatorsGetAllReader is a Reader for the MailIndicatorsGetAll structure.
type MailIndicatorsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MailIndicatorsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMailIndicatorsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/MailIndicators] MailIndicators_GetAll", response, response.Code())
	}
}

// NewMailIndicatorsGetAllOK creates a MailIndicatorsGetAllOK with default headers values
func NewMailIndicatorsGetAllOK() *MailIndicatorsGetAllOK {
	return &MailIndicatorsGetAllOK{}
}

/*
MailIndicatorsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type MailIndicatorsGetAllOK struct {
	Payload []*models.MailIndicator
}

// IsSuccess returns true when this mail indicators get all o k response has a 2xx status code
func (o *MailIndicatorsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mail indicators get all o k response has a 3xx status code
func (o *MailIndicatorsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mail indicators get all o k response has a 4xx status code
func (o *MailIndicatorsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mail indicators get all o k response has a 5xx status code
func (o *MailIndicatorsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mail indicators get all o k response a status code equal to that given
func (o *MailIndicatorsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mail indicators get all o k response
func (o *MailIndicatorsGetAllOK) Code() int {
	return 200
}

func (o *MailIndicatorsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/MailIndicators][%d] mailIndicatorsGetAllOK  %+v", 200, o.Payload)
}

func (o *MailIndicatorsGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/MailIndicators][%d] mailIndicatorsGetAllOK  %+v", 200, o.Payload)
}

func (o *MailIndicatorsGetAllOK) GetPayload() []*models.MailIndicator {
	return o.Payload
}

func (o *MailIndicatorsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
