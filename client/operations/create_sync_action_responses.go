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

// CreateSyncActionReader is a Reader for the CreateSyncAction structure.
type CreateSyncActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSyncActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCreateSyncActionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateSyncActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewCreateSyncActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateSyncActionNoContent creates a CreateSyncActionNoContent with default headers values
func NewCreateSyncActionNoContent() *CreateSyncActionNoContent {
	return &CreateSyncActionNoContent{}
}

/*CreateSyncActionNoContent handles this case with default header values.

The update was successful
*/
type CreateSyncActionNoContent struct {
}

func (o *CreateSyncActionNoContent) Error() string {
	return fmt.Sprintf("[PUT /actions][%d] createSyncActionNoContent ", 204)
}

func (o *CreateSyncActionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSyncActionBadRequest creates a CreateSyncActionBadRequest with default headers values
func NewCreateSyncActionBadRequest() *CreateSyncActionBadRequest {
	return &CreateSyncActionBadRequest{}
}

/*CreateSyncActionBadRequest handles this case with default header values.

The action cannot be executed due to bad input
*/
type CreateSyncActionBadRequest struct {
	Payload *models.Error
}

func (o *CreateSyncActionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /actions][%d] createSyncActionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateSyncActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateSyncActionDefault creates a CreateSyncActionDefault with default headers values
func NewCreateSyncActionDefault(code int) *CreateSyncActionDefault {
	return &CreateSyncActionDefault{
		_statusCode: code,
	}
}

/*CreateSyncActionDefault handles this case with default header values.

Internal Server Error
*/
type CreateSyncActionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the create sync action default response
func (o *CreateSyncActionDefault) Code() int {
	return o._statusCode
}

func (o *CreateSyncActionDefault) Error() string {
	return fmt.Sprintf("[PUT /actions][%d] createSyncAction default  %+v", o._statusCode, o.Payload)
}

func (o *CreateSyncActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
