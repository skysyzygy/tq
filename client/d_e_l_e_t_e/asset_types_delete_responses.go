// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AssetTypesDeleteReader is a Reader for the AssetTypesDelete structure.
type AssetTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssetTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAssetTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/AssetTypes/{id}] AssetTypes_Delete", response, response.Code())
	}
}

// NewAssetTypesDeleteNoContent creates a AssetTypesDeleteNoContent with default headers values
func NewAssetTypesDeleteNoContent() *AssetTypesDeleteNoContent {
	return &AssetTypesDeleteNoContent{}
}

/*
AssetTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type AssetTypesDeleteNoContent struct {
}

// IsSuccess returns true when this asset types delete no content response has a 2xx status code
func (o *AssetTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this asset types delete no content response has a 3xx status code
func (o *AssetTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this asset types delete no content response has a 4xx status code
func (o *AssetTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this asset types delete no content response has a 5xx status code
func (o *AssetTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this asset types delete no content response a status code equal to that given
func (o *AssetTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the asset types delete no content response
func (o *AssetTypesDeleteNoContent) Code() int {
	return 204
}

func (o *AssetTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/AssetTypes/{id}][%d] assetTypesDeleteNoContent ", 204)
}

func (o *AssetTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/AssetTypes/{id}][%d] assetTypesDeleteNoContent ", 204)
}

func (o *AssetTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
