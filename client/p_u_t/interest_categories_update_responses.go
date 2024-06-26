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

// InterestCategoriesUpdateReader is a Reader for the InterestCategoriesUpdate structure.
type InterestCategoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterestCategoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterestCategoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInterestCategoriesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInterestCategoriesUpdateOK creates a InterestCategoriesUpdateOK with default headers values
func NewInterestCategoriesUpdateOK() *InterestCategoriesUpdateOK {
	return &InterestCategoriesUpdateOK{}
}

/*
InterestCategoriesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type InterestCategoriesUpdateOK struct {
	Payload *models.InterestCategory
}

// IsSuccess returns true when this interest categories update o k response has a 2xx status code
func (o *InterestCategoriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this interest categories update o k response has a 3xx status code
func (o *InterestCategoriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this interest categories update o k response has a 4xx status code
func (o *InterestCategoriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this interest categories update o k response has a 5xx status code
func (o *InterestCategoriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this interest categories update o k response a status code equal to that given
func (o *InterestCategoriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the interest categories update o k response
func (o *InterestCategoriesUpdateOK) Code() int {
	return 200
}

func (o *InterestCategoriesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/InterestCategories/{id}][%d] interestCategoriesUpdateOK %s", 200, payload)
}

func (o *InterestCategoriesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/InterestCategories/{id}][%d] interestCategoriesUpdateOK %s", 200, payload)
}

func (o *InterestCategoriesUpdateOK) GetPayload() *models.InterestCategory {
	return o.Payload
}

func (o *InterestCategoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterestCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterestCategoriesUpdateDefault creates a InterestCategoriesUpdateDefault with default headers values
func NewInterestCategoriesUpdateDefault(code int) *InterestCategoriesUpdateDefault {
	return &InterestCategoriesUpdateDefault{
		_statusCode: code,
	}
}

/*
InterestCategoriesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type InterestCategoriesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this interest categories update default response has a 2xx status code
func (o *InterestCategoriesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this interest categories update default response has a 3xx status code
func (o *InterestCategoriesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this interest categories update default response has a 4xx status code
func (o *InterestCategoriesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this interest categories update default response has a 5xx status code
func (o *InterestCategoriesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this interest categories update default response a status code equal to that given
func (o *InterestCategoriesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the interest categories update default response
func (o *InterestCategoriesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *InterestCategoriesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/InterestCategories/{id}][%d] InterestCategories_Update default %s", o._statusCode, payload)
}

func (o *InterestCategoriesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/InterestCategories/{id}][%d] InterestCategories_Update default %s", o._statusCode, payload)
}

func (o *InterestCategoriesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *InterestCategoriesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
