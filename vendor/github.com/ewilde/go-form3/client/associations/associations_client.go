// Code generated by go-swagger; DO NOT EDIT.

package associations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new associations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for associations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeletePayportID deletes service association
*/
func (a *Client) DeletePayportID(params *DeletePayportIDParams) (*DeletePayportIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePayportIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePayportID",
		Method:             "DELETE",
		PathPattern:        "/payport/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePayportIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePayportIDNoContent), nil

}

/*
DeleteStarlingID deletes organisation association
*/
func (a *Client) DeleteStarlingID(params *DeleteStarlingIDParams) (*DeleteStarlingIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteStarlingIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteStarlingID",
		Method:             "DELETE",
		PathPattern:        "/starling/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteStarlingIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteStarlingIDNoContent), nil

}

/*
GetPayport lists all organisation associations
*/
func (a *Client) GetPayport(params *GetPayportParams) (*GetPayportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPayportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPayport",
		Method:             "GET",
		PathPattern:        "/payport",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPayportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPayportOK), nil

}

/*
GetPayportID fetches service association
*/
func (a *Client) GetPayportID(params *GetPayportIDParams) (*GetPayportIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPayportIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPayportID",
		Method:             "GET",
		PathPattern:        "/payport/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPayportIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPayportIDOK), nil

}

/*
GetStarling lists all organisation associations
*/
func (a *Client) GetStarling(params *GetStarlingParams) (*GetStarlingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStarlingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStarling",
		Method:             "GET",
		PathPattern:        "/starling",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStarlingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStarlingOK), nil

}

/*
GetStarlingID fetches organisation association
*/
func (a *Client) GetStarlingID(params *GetStarlingIDParams) (*GetStarlingIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStarlingIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStarlingID",
		Method:             "GET",
		PathPattern:        "/starling/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStarlingIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStarlingIDOK), nil

}

/*
PostPayport creates payport service association
*/
func (a *Client) PostPayport(params *PostPayportParams) (*PostPayportCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPayportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPayport",
		Method:             "POST",
		PathPattern:        "/payport",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPayportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPayportCreated), nil

}

/*
PostStarling creates organisation association
*/
func (a *Client) PostStarling(params *PostStarlingParams) (*PostStarlingCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStarlingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStarling",
		Method:             "POST",
		PathPattern:        "/starling",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStarlingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStarlingCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
