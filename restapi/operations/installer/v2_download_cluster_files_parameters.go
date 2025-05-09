// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewV2DownloadClusterFilesParams creates a new V2DownloadClusterFilesParams object
//
// There are no default values defined in the spec.
func NewV2DownloadClusterFilesParams() V2DownloadClusterFilesParams {

	return V2DownloadClusterFilesParams{}
}

// V2DownloadClusterFilesParams contains all the bound params for the v2 download cluster files operation
// typically these are obtained from a http.Request
//
// swagger:parameters V2DownloadClusterFiles
type V2DownloadClusterFilesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The cluster that owns the file that should be downloaded.
	  Required: true
	  In: path
	*/
	ClusterID strfmt.UUID
	/*The software version of the discovery agent that is downloading the file.
	  In: header
	*/
	DiscoveryAgentVersion *string
	/*The file to be downloaded.
	  Required: true
	  In: query
	*/
	FileName string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV2DownloadClusterFilesParams() beforehand.
func (o *V2DownloadClusterFilesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rClusterID, rhkClusterID, _ := route.Params.GetOK("cluster_id")
	if err := o.bindClusterID(rClusterID, rhkClusterID, route.Formats); err != nil {
		res = append(res, err)
	}

	if err := o.bindDiscoveryAgentVersion(r.Header[http.CanonicalHeaderKey("discovery_agent_version")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qFileName, qhkFileName, _ := qs.GetOK("file_name")
	if err := o.bindFileName(qFileName, qhkFileName, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClusterID binds and validates parameter ClusterID from path.
func (o *V2DownloadClusterFilesParams) bindClusterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("cluster_id", "path", "strfmt.UUID", raw)
	}
	o.ClusterID = *(value.(*strfmt.UUID))

	if err := o.validateClusterID(formats); err != nil {
		return err
	}

	return nil
}

// validateClusterID carries on validations for parameter ClusterID
func (o *V2DownloadClusterFilesParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.FormatOf("cluster_id", "path", "uuid", o.ClusterID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindDiscoveryAgentVersion binds and validates parameter DiscoveryAgentVersion from header.
func (o *V2DownloadClusterFilesParams) bindDiscoveryAgentVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.DiscoveryAgentVersion = &raw

	return nil
}

// bindFileName binds and validates parameter FileName from query.
func (o *V2DownloadClusterFilesParams) bindFileName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("file_name", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("file_name", "query", raw); err != nil {
		return err
	}
	o.FileName = raw

	if err := o.validateFileName(formats); err != nil {
		return err
	}

	return nil
}

// validateFileName carries on validations for parameter FileName
func (o *V2DownloadClusterFilesParams) validateFileName(formats strfmt.Registry) error {

	if err := validate.EnumCase("file_name", "query", o.FileName, []interface{}{"bootstrap.ign", "master.ign", "metadata.json", "worker.ign", "install-config.yaml", "custom_manifests.json", "custom_manifests.yaml", "arbiter.ign"}, true); err != nil {
		return err
	}

	return nil
}
