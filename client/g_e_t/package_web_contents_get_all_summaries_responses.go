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

// PackageWebContentsGetAllSummariesReader is a Reader for the PackageWebContentsGetAllSummaries structure.
type PackageWebContentsGetAllSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackageWebContentsGetAllSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackageWebContentsGetAllSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Txn/PackageWebContents/Summary] PackageWebContents_GetAllSummaries", response, response.Code())
	}
}

// NewPackageWebContentsGetAllSummariesOK creates a PackageWebContentsGetAllSummariesOK with default headers values
func NewPackageWebContentsGetAllSummariesOK() *PackageWebContentsGetAllSummariesOK {
	return &PackageWebContentsGetAllSummariesOK{}
}

/*
PackageWebContentsGetAllSummariesOK describes a response with status code 200, with default header values.

OK
*/
type PackageWebContentsGetAllSummariesOK struct {
	Payload []*models.PackageWebContentSummary
}

// IsSuccess returns true when this package web contents get all summaries o k response has a 2xx status code
func (o *PackageWebContentsGetAllSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this package web contents get all summaries o k response has a 3xx status code
func (o *PackageWebContentsGetAllSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this package web contents get all summaries o k response has a 4xx status code
func (o *PackageWebContentsGetAllSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this package web contents get all summaries o k response has a 5xx status code
func (o *PackageWebContentsGetAllSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this package web contents get all summaries o k response a status code equal to that given
func (o *PackageWebContentsGetAllSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the package web contents get all summaries o k response
func (o *PackageWebContentsGetAllSummariesOK) Code() int {
	return 200
}

func (o *PackageWebContentsGetAllSummariesOK) Error() string {
	return fmt.Sprintf("[GET /Txn/PackageWebContents/Summary][%d] packageWebContentsGetAllSummariesOK  %+v", 200, o.Payload)
}

func (o *PackageWebContentsGetAllSummariesOK) String() string {
	return fmt.Sprintf("[GET /Txn/PackageWebContents/Summary][%d] packageWebContentsGetAllSummariesOK  %+v", 200, o.Payload)
}

func (o *PackageWebContentsGetAllSummariesOK) GetPayload() []*models.PackageWebContentSummary {
	return o.Payload
}

func (o *PackageWebContentsGetAllSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
