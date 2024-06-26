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

// AppealCategoriesGetAllReader is a Reader for the AppealCategoriesGetAll structure.
type AppealCategoriesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppealCategoriesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppealCategoriesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAppealCategoriesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAppealCategoriesGetAllOK creates a AppealCategoriesGetAllOK with default headers values
func NewAppealCategoriesGetAllOK() *AppealCategoriesGetAllOK {
	return &AppealCategoriesGetAllOK{}
}

/*
AppealCategoriesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type AppealCategoriesGetAllOK struct {
	Payload []*models.AppealCategory
}

// IsSuccess returns true when this appeal categories get all o k response has a 2xx status code
func (o *AppealCategoriesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this appeal categories get all o k response has a 3xx status code
func (o *AppealCategoriesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this appeal categories get all o k response has a 4xx status code
func (o *AppealCategoriesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this appeal categories get all o k response has a 5xx status code
func (o *AppealCategoriesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this appeal categories get all o k response a status code equal to that given
func (o *AppealCategoriesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the appeal categories get all o k response
func (o *AppealCategoriesGetAllOK) Code() int {
	return 200
}

func (o *AppealCategoriesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories][%d] appealCategoriesGetAllOK %s", 200, payload)
}

func (o *AppealCategoriesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories][%d] appealCategoriesGetAllOK %s", 200, payload)
}

func (o *AppealCategoriesGetAllOK) GetPayload() []*models.AppealCategory {
	return o.Payload
}

func (o *AppealCategoriesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAppealCategoriesGetAllDefault creates a AppealCategoriesGetAllDefault with default headers values
func NewAppealCategoriesGetAllDefault(code int) *AppealCategoriesGetAllDefault {
	return &AppealCategoriesGetAllDefault{
		_statusCode: code,
	}
}

/*
AppealCategoriesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type AppealCategoriesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this appeal categories get all default response has a 2xx status code
func (o *AppealCategoriesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this appeal categories get all default response has a 3xx status code
func (o *AppealCategoriesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this appeal categories get all default response has a 4xx status code
func (o *AppealCategoriesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this appeal categories get all default response has a 5xx status code
func (o *AppealCategoriesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this appeal categories get all default response a status code equal to that given
func (o *AppealCategoriesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the appeal categories get all default response
func (o *AppealCategoriesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *AppealCategoriesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories][%d] AppealCategories_GetAll default %s", o._statusCode, payload)
}

func (o *AppealCategoriesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AppealCategories][%d] AppealCategories_GetAll default %s", o._statusCode, payload)
}

func (o *AppealCategoriesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AppealCategoriesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
