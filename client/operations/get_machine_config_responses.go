// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mcastelino/fireclient/client/models"
)

// GetMachineConfigReader is a Reader for the GetMachineConfig structure.
type GetMachineConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMachineConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMachineConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMachineConfigOK creates a GetMachineConfigOK with default headers values
func NewGetMachineConfigOK() *GetMachineConfigOK {
	return &GetMachineConfigOK{}
}

/*GetMachineConfigOK handles this case with default header values.

OK
*/
type GetMachineConfigOK struct {
	Payload *models.MachineConfiguration
}

func (o *GetMachineConfigOK) Error() string {
	return fmt.Sprintf("[GET /machine-config][%d] getMachineConfigOK  %+v", 200, o.Payload)
}

func (o *GetMachineConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineConfigDefault creates a GetMachineConfigDefault with default headers values
func NewGetMachineConfigDefault(code int) *GetMachineConfigDefault {
	return &GetMachineConfigDefault{
		_statusCode: code,
	}
}

/*GetMachineConfigDefault handles this case with default header values.

Internal server error
*/
type GetMachineConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get machine config default response
func (o *GetMachineConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetMachineConfigDefault) Error() string {
	return fmt.Sprintf("[GET /machine-config][%d] GetMachineConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetMachineConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}