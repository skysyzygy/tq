// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CustomExecuteLocalProcedureWithMultipleResultSetsReader is a Reader for the CustomExecuteLocalProcedureWithMultipleResultSets structure.
type CustomExecuteLocalProcedureWithMultipleResultSetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomExecuteLocalProcedureWithMultipleResultSetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Custom/Execute/MultipleResultSets] Custom_ExecuteLocalProcedureWithMultipleResultSets", response, response.Code())
	}
}

// NewCustomExecuteLocalProcedureWithMultipleResultSetsOK creates a CustomExecuteLocalProcedureWithMultipleResultSetsOK with default headers values
func NewCustomExecuteLocalProcedureWithMultipleResultSetsOK() *CustomExecuteLocalProcedureWithMultipleResultSetsOK {
	return &CustomExecuteLocalProcedureWithMultipleResultSetsOK{}
}

/*
CustomExecuteLocalProcedureWithMultipleResultSetsOK describes a response with status code 200, with default header values.

OK
*/
type CustomExecuteLocalProcedureWithMultipleResultSetsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this custom execute local procedure with multiple result sets o k response has a 2xx status code
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom execute local procedure with multiple result sets o k response has a 3xx status code
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom execute local procedure with multiple result sets o k response has a 4xx status code
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom execute local procedure with multiple result sets o k response has a 5xx status code
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom execute local procedure with multiple result sets o k response a status code equal to that given
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the custom execute local procedure with multiple result sets o k response
func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) Code() int {
	return 200
}

func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) Error() string {
	return fmt.Sprintf("[POST /Custom/Execute/MultipleResultSets][%d] customExecuteLocalProcedureWithMultipleResultSetsOK  %+v", 200, o.Payload)
}

func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) String() string {
	return fmt.Sprintf("[POST /Custom/Execute/MultipleResultSets][%d] customExecuteLocalProcedureWithMultipleResultSetsOK  %+v", 200, o.Payload)
}

func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomExecuteLocalProcedureWithMultipleResultSetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}