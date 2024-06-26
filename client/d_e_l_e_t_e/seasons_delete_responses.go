// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// SeasonsDeleteReader is a Reader for the SeasonsDelete structure.
type SeasonsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SeasonsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSeasonsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSeasonsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSeasonsDeleteNoContent creates a SeasonsDeleteNoContent with default headers values
func NewSeasonsDeleteNoContent() *SeasonsDeleteNoContent {
	return &SeasonsDeleteNoContent{}
}

/*
SeasonsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type SeasonsDeleteNoContent struct {
}

// IsSuccess returns true when this seasons delete no content response has a 2xx status code
func (o *SeasonsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this seasons delete no content response has a 3xx status code
func (o *SeasonsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this seasons delete no content response has a 4xx status code
func (o *SeasonsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this seasons delete no content response has a 5xx status code
func (o *SeasonsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this seasons delete no content response a status code equal to that given
func (o *SeasonsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the seasons delete no content response
func (o *SeasonsDeleteNoContent) Code() int {
	return 204
}

func (o *SeasonsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/Seasons/{id}][%d] seasonsDeleteNoContent", 204)
}

func (o *SeasonsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/Seasons/{id}][%d] seasonsDeleteNoContent", 204)
}

func (o *SeasonsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSeasonsDeleteDefault creates a SeasonsDeleteDefault with default headers values
func NewSeasonsDeleteDefault(code int) *SeasonsDeleteDefault {
	return &SeasonsDeleteDefault{
		_statusCode: code,
	}
}

/*
SeasonsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type SeasonsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this seasons delete default response has a 2xx status code
func (o *SeasonsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this seasons delete default response has a 3xx status code
func (o *SeasonsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this seasons delete default response has a 4xx status code
func (o *SeasonsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this seasons delete default response has a 5xx status code
func (o *SeasonsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this seasons delete default response a status code equal to that given
func (o *SeasonsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the seasons delete default response
func (o *SeasonsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SeasonsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/Seasons/{id}][%d] Seasons_Delete default %s", o._statusCode, payload)
}

func (o *SeasonsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /ReferenceData/Seasons/{id}][%d] Seasons_Delete default %s", o._statusCode, payload)
}

func (o *SeasonsDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SeasonsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
