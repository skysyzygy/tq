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

// TitlesGetAllReader is a Reader for the TitlesGetAll structure.
type TitlesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TitlesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTitlesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/Titles] Titles_GetAll", response, response.Code())
	}
}

// NewTitlesGetAllOK creates a TitlesGetAllOK with default headers values
func NewTitlesGetAllOK() *TitlesGetAllOK {
	return &TitlesGetAllOK{}
}

/*
TitlesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type TitlesGetAllOK struct {
	Payload []*models.Title
}

// IsSuccess returns true when this titles get all o k response has a 2xx status code
func (o *TitlesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this titles get all o k response has a 3xx status code
func (o *TitlesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this titles get all o k response has a 4xx status code
func (o *TitlesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this titles get all o k response has a 5xx status code
func (o *TitlesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this titles get all o k response a status code equal to that given
func (o *TitlesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the titles get all o k response
func (o *TitlesGetAllOK) Code() int {
	return 200
}

func (o *TitlesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /TXN/Titles][%d] titlesGetAllOK  %+v", 200, o.Payload)
}

func (o *TitlesGetAllOK) String() string {
	return fmt.Sprintf("[GET /TXN/Titles][%d] titlesGetAllOK  %+v", 200, o.Payload)
}

func (o *TitlesGetAllOK) GetPayload() []*models.Title {
	return o.Payload
}

func (o *TitlesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}