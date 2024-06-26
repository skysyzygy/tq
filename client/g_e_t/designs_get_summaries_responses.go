// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// DesignsGetSummariesReader is a Reader for the DesignsGetSummaries structure.
type DesignsGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DesignsGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDesignsGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDesignsGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDesignsGetSummariesOK creates a DesignsGetSummariesOK with default headers values
func NewDesignsGetSummariesOK() *DesignsGetSummariesOK {
	return &DesignsGetSummariesOK{}
}

/*
DesignsGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type DesignsGetSummariesOK struct {
	Payload []*models.DesignSummary
}

// IsSuccess returns true when this designs get summaries o k response has a 2xx status code
func (o *DesignsGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this designs get summaries o k response has a 3xx status code
func (o *DesignsGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this designs get summaries o k response has a 4xx status code
func (o *DesignsGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this designs get summaries o k response has a 5xx status code
func (o *DesignsGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this designs get summaries o k response a status code equal to that given
func (o *DesignsGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the designs get summaries o k response
func (o *DesignsGetSummariesOK) Code() int {
	return 200
}

func (o *DesignsGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Designs/Summary][%d] designsGetSummariesOK %s", 200, payload)
}

func (o *DesignsGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Designs/Summary][%d] designsGetSummariesOK %s", 200, payload)
}

func (o *DesignsGetSummariesOK) GetPayload() []*models.DesignSummary {
	return o.Payload
}

func (o *DesignsGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDesignsGetSummariesDefault creates a DesignsGetSummariesDefault with default headers values
func NewDesignsGetSummariesDefault(code int) *DesignsGetSummariesDefault {
	return &DesignsGetSummariesDefault{
		_statusCode: code,
	}
}

/*
DesignsGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type DesignsGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this designs get summaries default response has a 2xx status code
func (o *DesignsGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this designs get summaries default response has a 3xx status code
func (o *DesignsGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this designs get summaries default response has a 4xx status code
func (o *DesignsGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this designs get summaries default response has a 5xx status code
func (o *DesignsGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this designs get summaries default response a status code equal to that given
func (o *DesignsGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the designs get summaries default response
func (o *DesignsGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *DesignsGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Designs/Summary][%d] Designs_GetSummaries default %s", o._statusCode, payload)
}

func (o *DesignsGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Designs/Summary][%d] Designs_GetSummaries default %s", o._statusCode, payload)
}

func (o *DesignsGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *DesignsGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
