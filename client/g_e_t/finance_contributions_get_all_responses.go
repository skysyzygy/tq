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

// FinanceContributionsGetAllReader is a Reader for the FinanceContributionsGetAll structure.
type FinanceContributionsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FinanceContributionsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFinanceContributionsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Finance/Contributions] FinanceContributions_GetAll", response, response.Code())
	}
}

// NewFinanceContributionsGetAllOK creates a FinanceContributionsGetAllOK with default headers values
func NewFinanceContributionsGetAllOK() *FinanceContributionsGetAllOK {
	return &FinanceContributionsGetAllOK{}
}

/*
FinanceContributionsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type FinanceContributionsGetAllOK struct {
	Payload []*models.Contribution
}

// IsSuccess returns true when this finance contributions get all o k response has a 2xx status code
func (o *FinanceContributionsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this finance contributions get all o k response has a 3xx status code
func (o *FinanceContributionsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this finance contributions get all o k response has a 4xx status code
func (o *FinanceContributionsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this finance contributions get all o k response has a 5xx status code
func (o *FinanceContributionsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this finance contributions get all o k response a status code equal to that given
func (o *FinanceContributionsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the finance contributions get all o k response
func (o *FinanceContributionsGetAllOK) Code() int {
	return 200
}

func (o *FinanceContributionsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /Finance/Contributions][%d] financeContributionsGetAllOK  %+v", 200, o.Payload)
}

func (o *FinanceContributionsGetAllOK) String() string {
	return fmt.Sprintf("[GET /Finance/Contributions][%d] financeContributionsGetAllOK  %+v", 200, o.Payload)
}

func (o *FinanceContributionsGetAllOK) GetPayload() []*models.Contribution {
	return o.Payload
}

func (o *FinanceContributionsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
