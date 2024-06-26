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

// TemplateCategoriesGetSummariesReader is a Reader for the TemplateCategoriesGetSummaries structure.
type TemplateCategoriesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplateCategoriesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplateCategoriesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTemplateCategoriesGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTemplateCategoriesGetSummariesOK creates a TemplateCategoriesGetSummariesOK with default headers values
func NewTemplateCategoriesGetSummariesOK() *TemplateCategoriesGetSummariesOK {
	return &TemplateCategoriesGetSummariesOK{}
}

/*
TemplateCategoriesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type TemplateCategoriesGetSummariesOK struct {
	Payload []*models.TemplateCategorySummary
}

// IsSuccess returns true when this template categories get summaries o k response has a 2xx status code
func (o *TemplateCategoriesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this template categories get summaries o k response has a 3xx status code
func (o *TemplateCategoriesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this template categories get summaries o k response has a 4xx status code
func (o *TemplateCategoriesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this template categories get summaries o k response has a 5xx status code
func (o *TemplateCategoriesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this template categories get summaries o k response a status code equal to that given
func (o *TemplateCategoriesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the template categories get summaries o k response
func (o *TemplateCategoriesGetSummariesOK) Code() int {
	return 200
}

func (o *TemplateCategoriesGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/TemplateCategories/Summary][%d] templateCategoriesGetSummariesOK %s", 200, payload)
}

func (o *TemplateCategoriesGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/TemplateCategories/Summary][%d] templateCategoriesGetSummariesOK %s", 200, payload)
}

func (o *TemplateCategoriesGetSummariesOK) GetPayload() []*models.TemplateCategorySummary {
	return o.Payload
}

func (o *TemplateCategoriesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplateCategoriesGetSummariesDefault creates a TemplateCategoriesGetSummariesDefault with default headers values
func NewTemplateCategoriesGetSummariesDefault(code int) *TemplateCategoriesGetSummariesDefault {
	return &TemplateCategoriesGetSummariesDefault{
		_statusCode: code,
	}
}

/*
TemplateCategoriesGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type TemplateCategoriesGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this template categories get summaries default response has a 2xx status code
func (o *TemplateCategoriesGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this template categories get summaries default response has a 3xx status code
func (o *TemplateCategoriesGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this template categories get summaries default response has a 4xx status code
func (o *TemplateCategoriesGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this template categories get summaries default response has a 5xx status code
func (o *TemplateCategoriesGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this template categories get summaries default response a status code equal to that given
func (o *TemplateCategoriesGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the template categories get summaries default response
func (o *TemplateCategoriesGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *TemplateCategoriesGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/TemplateCategories/Summary][%d] TemplateCategories_GetSummaries default %s", o._statusCode, payload)
}

func (o *TemplateCategoriesGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/TemplateCategories/Summary][%d] TemplateCategories_GetSummaries default %s", o._statusCode, payload)
}

func (o *TemplateCategoriesGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *TemplateCategoriesGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
