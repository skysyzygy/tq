// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// BatchTypesDeleteReader is a Reader for the BatchTypesDelete structure.
type BatchTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BatchTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewBatchTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/BatchTypes/{id}] BatchTypes_Delete", response, response.Code())
	}
}

// NewBatchTypesDeleteNoContent creates a BatchTypesDeleteNoContent with default headers values
func NewBatchTypesDeleteNoContent() *BatchTypesDeleteNoContent {
	return &BatchTypesDeleteNoContent{}
}

/*
BatchTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type BatchTypesDeleteNoContent struct {
}

// IsSuccess returns true when this batch types delete no content response has a 2xx status code
func (o *BatchTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this batch types delete no content response has a 3xx status code
func (o *BatchTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this batch types delete no content response has a 4xx status code
func (o *BatchTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this batch types delete no content response has a 5xx status code
func (o *BatchTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this batch types delete no content response a status code equal to that given
func (o *BatchTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the batch types delete no content response
func (o *BatchTypesDeleteNoContent) Code() int {
	return 204
}

func (o *BatchTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/BatchTypes/{id}][%d] batchTypesDeleteNoContent ", 204)
}

func (o *BatchTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/BatchTypes/{id}][%d] batchTypesDeleteNoContent ", 204)
}

func (o *BatchTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
