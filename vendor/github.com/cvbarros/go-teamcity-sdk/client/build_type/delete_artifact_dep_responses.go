// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteArtifactDepReader is a Reader for the DeleteArtifactDep structure.
type DeleteArtifactDepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteArtifactDepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteArtifactDepDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteArtifactDepDefault creates a DeleteArtifactDepDefault with default headers values
func NewDeleteArtifactDepDefault(code int) *DeleteArtifactDepDefault {
	return &DeleteArtifactDepDefault{
		_statusCode: code,
	}
}

/*DeleteArtifactDepDefault handles this case with default header values.

successful operation
*/
type DeleteArtifactDepDefault struct {
	_statusCode int
}

// Code gets the status code for the delete artifact dep default response
func (o *DeleteArtifactDepDefault) Code() int {
	return o._statusCode
}

func (o *DeleteArtifactDepDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator}][%d] deleteArtifactDep default ", o._statusCode)
}

func (o *DeleteArtifactDepDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}