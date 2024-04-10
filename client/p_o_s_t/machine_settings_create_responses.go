// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// MachineSettingsCreateReader is a Reader for the MachineSettingsCreate structure.
type MachineSettingsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachineSettingsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMachineSettingsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/MachineSettings] MachineSettings_Create", response, response.Code())
	}
}

// NewMachineSettingsCreateOK creates a MachineSettingsCreateOK with default headers values
func NewMachineSettingsCreateOK() *MachineSettingsCreateOK {
	return &MachineSettingsCreateOK{}
}

/*
MachineSettingsCreateOK describes a response with status code 200, with default header values.

OK
*/
type MachineSettingsCreateOK struct {
	Payload *models.MachineSetting
}

// IsSuccess returns true when this machine settings create o k response has a 2xx status code
func (o *MachineSettingsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this machine settings create o k response has a 3xx status code
func (o *MachineSettingsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this machine settings create o k response has a 4xx status code
func (o *MachineSettingsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this machine settings create o k response has a 5xx status code
func (o *MachineSettingsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this machine settings create o k response a status code equal to that given
func (o *MachineSettingsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the machine settings create o k response
func (o *MachineSettingsCreateOK) Code() int {
	return 200
}

func (o *MachineSettingsCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/MachineSettings][%d] machineSettingsCreateOK  %+v", 200, o.Payload)
}

func (o *MachineSettingsCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/MachineSettings][%d] machineSettingsCreateOK  %+v", 200, o.Payload)
}

func (o *MachineSettingsCreateOK) GetPayload() *models.MachineSetting {
	return o.Payload
}

func (o *MachineSettingsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
