// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// ModeOfSaleSurveyQuestionsUpdateReader is a Reader for the ModeOfSaleSurveyQuestionsUpdate structure.
type ModeOfSaleSurveyQuestionsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModeOfSaleSurveyQuestionsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModeOfSaleSurveyQuestionsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewModeOfSaleSurveyQuestionsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModeOfSaleSurveyQuestionsUpdateOK creates a ModeOfSaleSurveyQuestionsUpdateOK with default headers values
func NewModeOfSaleSurveyQuestionsUpdateOK() *ModeOfSaleSurveyQuestionsUpdateOK {
	return &ModeOfSaleSurveyQuestionsUpdateOK{}
}

/*
ModeOfSaleSurveyQuestionsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ModeOfSaleSurveyQuestionsUpdateOK struct {
	Payload *models.ModeOfSaleSurveyQuestion
}

// IsSuccess returns true when this mode of sale survey questions update o k response has a 2xx status code
func (o *ModeOfSaleSurveyQuestionsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mode of sale survey questions update o k response has a 3xx status code
func (o *ModeOfSaleSurveyQuestionsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mode of sale survey questions update o k response has a 4xx status code
func (o *ModeOfSaleSurveyQuestionsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this mode of sale survey questions update o k response has a 5xx status code
func (o *ModeOfSaleSurveyQuestionsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this mode of sale survey questions update o k response a status code equal to that given
func (o *ModeOfSaleSurveyQuestionsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the mode of sale survey questions update o k response
func (o *ModeOfSaleSurveyQuestionsUpdateOK) Code() int {
	return 200
}

func (o *ModeOfSaleSurveyQuestionsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/ModeOfSaleSurveyQuestions/{modeOfSaleSurveyQuestionId}][%d] modeOfSaleSurveyQuestionsUpdateOK %s", 200, payload)
}

func (o *ModeOfSaleSurveyQuestionsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/ModeOfSaleSurveyQuestions/{modeOfSaleSurveyQuestionId}][%d] modeOfSaleSurveyQuestionsUpdateOK %s", 200, payload)
}

func (o *ModeOfSaleSurveyQuestionsUpdateOK) GetPayload() *models.ModeOfSaleSurveyQuestion {
	return o.Payload
}

func (o *ModeOfSaleSurveyQuestionsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModeOfSaleSurveyQuestion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModeOfSaleSurveyQuestionsUpdateDefault creates a ModeOfSaleSurveyQuestionsUpdateDefault with default headers values
func NewModeOfSaleSurveyQuestionsUpdateDefault(code int) *ModeOfSaleSurveyQuestionsUpdateDefault {
	return &ModeOfSaleSurveyQuestionsUpdateDefault{
		_statusCode: code,
	}
}

/*
ModeOfSaleSurveyQuestionsUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type ModeOfSaleSurveyQuestionsUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this mode of sale survey questions update default response has a 2xx status code
func (o *ModeOfSaleSurveyQuestionsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this mode of sale survey questions update default response has a 3xx status code
func (o *ModeOfSaleSurveyQuestionsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this mode of sale survey questions update default response has a 4xx status code
func (o *ModeOfSaleSurveyQuestionsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this mode of sale survey questions update default response has a 5xx status code
func (o *ModeOfSaleSurveyQuestionsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this mode of sale survey questions update default response a status code equal to that given
func (o *ModeOfSaleSurveyQuestionsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the mode of sale survey questions update default response
func (o *ModeOfSaleSurveyQuestionsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ModeOfSaleSurveyQuestionsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/ModeOfSaleSurveyQuestions/{modeOfSaleSurveyQuestionId}][%d] ModeOfSaleSurveyQuestions_Update default %s", o._statusCode, payload)
}

func (o *ModeOfSaleSurveyQuestionsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /TXN/ModeOfSaleSurveyQuestions/{modeOfSaleSurveyQuestionId}][%d] ModeOfSaleSurveyQuestions_Update default %s", o._statusCode, payload)
}

func (o *ModeOfSaleSurveyQuestionsUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ModeOfSaleSurveyQuestionsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
