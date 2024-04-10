// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// PackageTypesUpdateReader is a Reader for the PackageTypesUpdate structure.
type PackageTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PackageTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPackageTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/PackageTypes/{id}] PackageTypes_Update", response, response.Code())
	}
}

// NewPackageTypesUpdateOK creates a PackageTypesUpdateOK with default headers values
func NewPackageTypesUpdateOK() *PackageTypesUpdateOK {
	return &PackageTypesUpdateOK{}
}

/*
PackageTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PackageTypesUpdateOK struct {
	Payload *models.PackageType
}

// IsSuccess returns true when this package types update o k response has a 2xx status code
func (o *PackageTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this package types update o k response has a 3xx status code
func (o *PackageTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this package types update o k response has a 4xx status code
func (o *PackageTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this package types update o k response has a 5xx status code
func (o *PackageTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this package types update o k response a status code equal to that given
func (o *PackageTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the package types update o k response
func (o *PackageTypesUpdateOK) Code() int {
	return 200
}

func (o *PackageTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/PackageTypes/{id}][%d] packageTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *PackageTypesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/PackageTypes/{id}][%d] packageTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *PackageTypesUpdateOK) GetPayload() *models.PackageType {
	return o.Payload
}

func (o *PackageTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PackageType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
