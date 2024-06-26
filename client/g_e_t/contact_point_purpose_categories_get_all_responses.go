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

// ContactPointPurposeCategoriesGetAllReader is a Reader for the ContactPointPurposeCategoriesGetAll structure.
type ContactPointPurposeCategoriesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPointPurposeCategoriesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactPointPurposeCategoriesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContactPointPurposeCategoriesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContactPointPurposeCategoriesGetAllOK creates a ContactPointPurposeCategoriesGetAllOK with default headers values
func NewContactPointPurposeCategoriesGetAllOK() *ContactPointPurposeCategoriesGetAllOK {
	return &ContactPointPurposeCategoriesGetAllOK{}
}

/*
ContactPointPurposeCategoriesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ContactPointPurposeCategoriesGetAllOK struct {
	Payload []*models.PurposeCategory
}

// IsSuccess returns true when this contact point purpose categories get all o k response has a 2xx status code
func (o *ContactPointPurposeCategoriesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact point purpose categories get all o k response has a 3xx status code
func (o *ContactPointPurposeCategoriesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact point purpose categories get all o k response has a 4xx status code
func (o *ContactPointPurposeCategoriesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact point purpose categories get all o k response has a 5xx status code
func (o *ContactPointPurposeCategoriesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact point purpose categories get all o k response a status code equal to that given
func (o *ContactPointPurposeCategoriesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact point purpose categories get all o k response
func (o *ContactPointPurposeCategoriesGetAllOK) Code() int {
	return 200
}

func (o *ContactPointPurposeCategoriesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ContactPointPurposeCategories][%d] contactPointPurposeCategoriesGetAllOK %s", 200, payload)
}

func (o *ContactPointPurposeCategoriesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ContactPointPurposeCategories][%d] contactPointPurposeCategoriesGetAllOK %s", 200, payload)
}

func (o *ContactPointPurposeCategoriesGetAllOK) GetPayload() []*models.PurposeCategory {
	return o.Payload
}

func (o *ContactPointPurposeCategoriesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContactPointPurposeCategoriesGetAllDefault creates a ContactPointPurposeCategoriesGetAllDefault with default headers values
func NewContactPointPurposeCategoriesGetAllDefault(code int) *ContactPointPurposeCategoriesGetAllDefault {
	return &ContactPointPurposeCategoriesGetAllDefault{
		_statusCode: code,
	}
}

/*
ContactPointPurposeCategoriesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type ContactPointPurposeCategoriesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this contact point purpose categories get all default response has a 2xx status code
func (o *ContactPointPurposeCategoriesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contact point purpose categories get all default response has a 3xx status code
func (o *ContactPointPurposeCategoriesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contact point purpose categories get all default response has a 4xx status code
func (o *ContactPointPurposeCategoriesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contact point purpose categories get all default response has a 5xx status code
func (o *ContactPointPurposeCategoriesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contact point purpose categories get all default response a status code equal to that given
func (o *ContactPointPurposeCategoriesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contact point purpose categories get all default response
func (o *ContactPointPurposeCategoriesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *ContactPointPurposeCategoriesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ContactPointPurposeCategories][%d] ContactPointPurposeCategories_GetAll default %s", o._statusCode, payload)
}

func (o *ContactPointPurposeCategoriesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/ContactPointPurposeCategories][%d] ContactPointPurposeCategories_GetAll default %s", o._statusCode, payload)
}

func (o *ContactPointPurposeCategoriesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ContactPointPurposeCategoriesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
