package automox

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

const serversURL = "/api/servers"

// ServersService is an interface for interacting with the server endpoints
// of the Automox API
type ServersService interface {
	List(context.Context) (Servers, error)
	Get(context.Context, int64) (*ServerDetails, error)
	GetPackages(context.Context, int64) (*Packages, error)
	GetCommandQueue(context.Context, int64) (*CommandQueue, error)
}

// ServersClient facilitates requests with the Automox servers
type ServersClient struct {
	client *Client
}

// List all tasks assigned to a given ticket ID
func (c *ServersClient) List(ctx context.Context) (Servers, error) {
	// TODO: Here....
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.apiURL,
		Path:   serversURL,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Servers{}
	_, err = c.client.makeRequest(req, res)
	if err != nil {
		return nil, err
	}

	return *res, nil
}

// Get a specific Server ticket by Server ID.
func (c *ServersClient) Get(ctx context.Context, id int64) (*ServerDetails, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.apiURL,
		Path:   fmt.Sprintf("%s/%d", serversURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &ServerDetails{}
	if _, err := c.client.makeRequest(req, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetPackages retrieves the list of packages installed on a server
func (c *ServersClient) GetPackages(ctx context.Context, id int64) (*Packages, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.apiURL,
		Path:   fmt.Sprintf("%s/%d/packages", serversURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &Packages{}
	if _, err := c.client.makeRequest(req, res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetCommandQueue returns the queue of upcoming commands for the specified device
func (c *ServersClient) GetCommandQueue(ctx context.Context, id int64) (*CommandQueue, error) {
	url := &url.URL{
		Scheme: "https",
		Host:   c.client.apiURL,
		Path:   fmt.Sprintf("%s/%d/queues", serversURL, id),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	res := &CommandQueue{}
	if _, err := c.client.makeRequest(req, res); err != nil {
		return nil, err
	}
	return res, nil
}
