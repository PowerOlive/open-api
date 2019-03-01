// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netlify/open-api/go/models"
)

// NewCreateAccountParams creates a new CreateAccountParams object
// with the default values initialized.
func NewCreateAccountParams() *CreateAccountParams {
	var ()
	return &CreateAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccountParamsWithTimeout creates a new CreateAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAccountParamsWithTimeout(timeout time.Duration) *CreateAccountParams {
	var ()
	return &CreateAccountParams{

		timeout: timeout,
	}
}

// NewCreateAccountParamsWithContext creates a new CreateAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAccountParamsWithContext(ctx context.Context) *CreateAccountParams {
	var ()
	return &CreateAccountParams{

		Context: ctx,
	}
}

// NewCreateAccountParamsWithHTTPClient creates a new CreateAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAccountParamsWithHTTPClient(client *http.Client) *CreateAccountParams {
	var ()
	return &CreateAccountParams{
		HTTPClient: client,
	}
}

/*CreateAccountParams contains all the parameters to send to the API endpoint
for the create account operation typically these are written to a http.Request
*/
type CreateAccountParams struct {

	/*AccountSetup*/
	AccountSetup *models.AccountSetup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create account params
func (o *CreateAccountParams) WithTimeout(timeout time.Duration) *CreateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create account params
func (o *CreateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create account params
func (o *CreateAccountParams) WithContext(ctx context.Context) *CreateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create account params
func (o *CreateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create account params
func (o *CreateAccountParams) WithHTTPClient(client *http.Client) *CreateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create account params
func (o *CreateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountSetup adds the accountSetup to the create account params
func (o *CreateAccountParams) WithAccountSetup(accountSetup *models.AccountSetup) *CreateAccountParams {
	o.SetAccountSetup(accountSetup)
	return o
}

// SetAccountSetup adds the accountSetup to the create account params
func (o *CreateAccountParams) SetAccountSetup(accountSetup *models.AccountSetup) {
	o.AccountSetup = accountSetup
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountSetup != nil {
		if err := r.SetBodyParam(o.AccountSetup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}