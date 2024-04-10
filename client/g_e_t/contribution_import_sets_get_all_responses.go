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

// ContributionImportSetsGetAllReader is a Reader for the ContributionImportSetsGetAll structure.
type ContributionImportSetsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContributionImportSetsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContributionImportSetsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ContributionImportSets] ContributionImportSets_GetAll", response, response.Code())
	}
}

// NewContributionImportSetsGetAllOK creates a ContributionImportSetsGetAllOK with default headers values
func NewContributionImportSetsGetAllOK() *ContributionImportSetsGetAllOK {
	return &ContributionImportSetsGetAllOK{}
}

/*
ContributionImportSetsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ContributionImportSetsGetAllOK struct {
	Payload []*models.ContributionImportSet
}

// IsSuccess returns true when this contribution import sets get all o k response has a 2xx status code
func (o *ContributionImportSetsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contribution import sets get all o k response has a 3xx status code
func (o *ContributionImportSetsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contribution import sets get all o k response has a 4xx status code
func (o *ContributionImportSetsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contribution import sets get all o k response has a 5xx status code
func (o *ContributionImportSetsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contribution import sets get all o k response a status code equal to that given
func (o *ContributionImportSetsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contribution import sets get all o k response
func (o *ContributionImportSetsGetAllOK) Code() int {
	return 200
}

func (o *ContributionImportSetsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ContributionImportSets][%d] contributionImportSetsGetAllOK  %+v", 200, o.Payload)
}

func (o *ContributionImportSetsGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ContributionImportSets][%d] contributionImportSetsGetAllOK  %+v", 200, o.Payload)
}

func (o *ContributionImportSetsGetAllOK) GetPayload() []*models.ContributionImportSet {
	return o.Payload
}

func (o *ContributionImportSetsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
