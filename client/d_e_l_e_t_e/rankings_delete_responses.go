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

// RankingsDeleteReader is a Reader for the RankingsDelete structure.
type RankingsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RankingsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRankingsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRankingsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRankingsDeleteNoContent creates a RankingsDeleteNoContent with default headers values
func NewRankingsDeleteNoContent() *RankingsDeleteNoContent {
	return &RankingsDeleteNoContent{}
}

/*
RankingsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type RankingsDeleteNoContent struct {
}

// IsSuccess returns true when this rankings delete no content response has a 2xx status code
func (o *RankingsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rankings delete no content response has a 3xx status code
func (o *RankingsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rankings delete no content response has a 4xx status code
func (o *RankingsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this rankings delete no content response has a 5xx status code
func (o *RankingsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this rankings delete no content response a status code equal to that given
func (o *RankingsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the rankings delete no content response
func (o *RankingsDeleteNoContent) Code() int {
	return 204
}

func (o *RankingsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /CRM/Rankings/{rankingId}][%d] rankingsDeleteNoContent", 204)
}

func (o *RankingsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /CRM/Rankings/{rankingId}][%d] rankingsDeleteNoContent", 204)
}

func (o *RankingsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRankingsDeleteDefault creates a RankingsDeleteDefault with default headers values
func NewRankingsDeleteDefault(code int) *RankingsDeleteDefault {
	return &RankingsDeleteDefault{
		_statusCode: code,
	}
}

/*
RankingsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type RankingsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this rankings delete default response has a 2xx status code
func (o *RankingsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rankings delete default response has a 3xx status code
func (o *RankingsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rankings delete default response has a 4xx status code
func (o *RankingsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rankings delete default response has a 5xx status code
func (o *RankingsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rankings delete default response a status code equal to that given
func (o *RankingsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rankings delete default response
func (o *RankingsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *RankingsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /CRM/Rankings/{rankingId}][%d] Rankings_Delete default %s", o._statusCode, payload)
}

func (o *RankingsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /CRM/Rankings/{rankingId}][%d] Rankings_Delete default %s", o._statusCode, payload)
}

func (o *RankingsDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *RankingsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
