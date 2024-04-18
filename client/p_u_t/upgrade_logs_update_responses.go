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

// UpgradeLogsUpdateReader is a Reader for the UpgradeLogsUpdate structure.
type UpgradeLogsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpgradeLogsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpgradeLogsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /Admin/UpgradeLogs/{upgradeLogId}] UpgradeLogs_Update", response, response.Code())
	}
}

// NewUpgradeLogsUpdateOK creates a UpgradeLogsUpdateOK with default headers values
func NewUpgradeLogsUpdateOK() *UpgradeLogsUpdateOK {
	return &UpgradeLogsUpdateOK{}
}

/*
UpgradeLogsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type UpgradeLogsUpdateOK struct {
	Payload *models.UpgradeLog
}

// IsSuccess returns true when this upgrade logs update o k response has a 2xx status code
func (o *UpgradeLogsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upgrade logs update o k response has a 3xx status code
func (o *UpgradeLogsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upgrade logs update o k response has a 4xx status code
func (o *UpgradeLogsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upgrade logs update o k response has a 5xx status code
func (o *UpgradeLogsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upgrade logs update o k response a status code equal to that given
func (o *UpgradeLogsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upgrade logs update o k response
func (o *UpgradeLogsUpdateOK) Code() int {
	return 200
}

func (o *UpgradeLogsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /Admin/UpgradeLogs/{upgradeLogId}][%d] upgradeLogsUpdateOK  %+v", 200, o.Payload)
}

func (o *UpgradeLogsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /Admin/UpgradeLogs/{upgradeLogId}][%d] upgradeLogsUpdateOK  %+v", 200, o.Payload)
}

func (o *UpgradeLogsUpdateOK) GetPayload() *models.UpgradeLog {
	return o.Payload
}

func (o *UpgradeLogsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpgradeLog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}