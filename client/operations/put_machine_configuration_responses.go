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

// PutMachineConfigurationReader is a Reader for the PutMachineConfiguration structure.
type PutMachineConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMachineConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutMachineConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutMachineConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutMachineConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutMachineConfigurationNoContent creates a PutMachineConfigurationNoContent with default headers values
func NewPutMachineConfigurationNoContent() *PutMachineConfigurationNoContent {
	return &PutMachineConfigurationNoContent{}
}

/*PutMachineConfigurationNoContent handles this case with default header values.

Machine Configuration created/updated
*/
type PutMachineConfigurationNoContent struct {
}

func (o *PutMachineConfigurationNoContent) Error() string {
	return fmt.Sprintf("[PUT /machine-config][%d] putMachineConfigurationNoContent ", 204)
}

func (o *PutMachineConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutMachineConfigurationBadRequest creates a PutMachineConfigurationBadRequest with default headers values
func NewPutMachineConfigurationBadRequest() *PutMachineConfigurationBadRequest {
	return &PutMachineConfigurationBadRequest{}
}

/*PutMachineConfigurationBadRequest handles this case with default header values.

Machine Configuration cannot be updated due to bad input
*/
type PutMachineConfigurationBadRequest struct {
	Payload *models.Error
}

func (o *PutMachineConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /machine-config][%d] putMachineConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *PutMachineConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutMachineConfigurationDefault creates a PutMachineConfigurationDefault with default headers values
func NewPutMachineConfigurationDefault(code int) *PutMachineConfigurationDefault {
	return &PutMachineConfigurationDefault{
		_statusCode: code,
	}
}

/*PutMachineConfigurationDefault handles this case with default header values.

Internal server error
*/
type PutMachineConfigurationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put machine configuration default response
func (o *PutMachineConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *PutMachineConfigurationDefault) Error() string {
	return fmt.Sprintf("[PUT /machine-config][%d] putMachineConfiguration default  %+v", o._statusCode, o.Payload)
}

func (o *PutMachineConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
