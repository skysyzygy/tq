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

// BulkDailyCopyExclusionsGetReader is a Reader for the BulkDailyCopyExclusionsGet structure.
type BulkDailyCopyExclusionsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkDailyCopyExclusionsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBulkDailyCopyExclusionsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/BulkDailyCopyExclusions/{bulkDailyCopyExclusionId}] BulkDailyCopyExclusions_Get", response, response.Code())
	}
}

// NewBulkDailyCopyExclusionsGetOK creates a BulkDailyCopyExclusionsGetOK with default headers values
func NewBulkDailyCopyExclusionsGetOK() *BulkDailyCopyExclusionsGetOK {
	return &BulkDailyCopyExclusionsGetOK{}
}

/*
BulkDailyCopyExclusionsGetOK describes a response with status code 200, with default header values.

OK
*/
type BulkDailyCopyExclusionsGetOK struct {
	Payload *models.BulkDailyCopyExclusion
}

// IsSuccess returns true when this bulk daily copy exclusions get o k response has a 2xx status code
func (o *BulkDailyCopyExclusionsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this bulk daily copy exclusions get o k response has a 3xx status code
func (o *BulkDailyCopyExclusionsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this bulk daily copy exclusions get o k response has a 4xx status code
func (o *BulkDailyCopyExclusionsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this bulk daily copy exclusions get o k response has a 5xx status code
func (o *BulkDailyCopyExclusionsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this bulk daily copy exclusions get o k response a status code equal to that given
func (o *BulkDailyCopyExclusionsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the bulk daily copy exclusions get o k response
func (o *BulkDailyCopyExclusionsGetOK) Code() int {
	return 200
}

func (o *BulkDailyCopyExclusionsGetOK) Error() string {
	return fmt.Sprintf("[GET /TXN/BulkDailyCopyExclusions/{bulkDailyCopyExclusionId}][%d] bulkDailyCopyExclusionsGetOK  %+v", 200, o.Payload)
}

func (o *BulkDailyCopyExclusionsGetOK) String() string {
	return fmt.Sprintf("[GET /TXN/BulkDailyCopyExclusions/{bulkDailyCopyExclusionId}][%d] bulkDailyCopyExclusionsGetOK  %+v", 200, o.Payload)
}

func (o *BulkDailyCopyExclusionsGetOK) GetPayload() *models.BulkDailyCopyExclusion {
	return o.Payload
}

func (o *BulkDailyCopyExclusionsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkDailyCopyExclusion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
