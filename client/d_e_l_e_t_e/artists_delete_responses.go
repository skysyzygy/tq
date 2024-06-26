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

// ArtistsDeleteReader is a Reader for the ArtistsDelete structure.
type ArtistsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArtistsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewArtistsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewArtistsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArtistsDeleteNoContent creates a ArtistsDeleteNoContent with default headers values
func NewArtistsDeleteNoContent() *ArtistsDeleteNoContent {
	return &ArtistsDeleteNoContent{}
}

/*
ArtistsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ArtistsDeleteNoContent struct {
}

// IsSuccess returns true when this artists delete no content response has a 2xx status code
func (o *ArtistsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this artists delete no content response has a 3xx status code
func (o *ArtistsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this artists delete no content response has a 4xx status code
func (o *ArtistsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this artists delete no content response has a 5xx status code
func (o *ArtistsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this artists delete no content response a status code equal to that given
func (o *ArtistsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the artists delete no content response
func (o *ArtistsDeleteNoContent) Code() int {
	return 204
}

func (o *ArtistsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /TXN/Artists/{artistId}][%d] artistsDeleteNoContent", 204)
}

func (o *ArtistsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /TXN/Artists/{artistId}][%d] artistsDeleteNoContent", 204)
}

func (o *ArtistsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewArtistsDeleteDefault creates a ArtistsDeleteDefault with default headers values
func NewArtistsDeleteDefault(code int) *ArtistsDeleteDefault {
	return &ArtistsDeleteDefault{
		_statusCode: code,
	}
}

/*
ArtistsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type ArtistsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this artists delete default response has a 2xx status code
func (o *ArtistsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this artists delete default response has a 3xx status code
func (o *ArtistsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this artists delete default response has a 4xx status code
func (o *ArtistsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this artists delete default response has a 5xx status code
func (o *ArtistsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this artists delete default response a status code equal to that given
func (o *ArtistsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the artists delete default response
func (o *ArtistsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ArtistsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /TXN/Artists/{artistId}][%d] Artists_Delete default %s", o._statusCode, payload)
}

func (o *ArtistsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /TXN/Artists/{artistId}][%d] Artists_Delete default %s", o._statusCode, payload)
}

func (o *ArtistsDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ArtistsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
