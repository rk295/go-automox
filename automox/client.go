package automox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const apiUrl = "console.automox.com"

// Client represents a new Automox API client to
// be utilized for API requests
type Client struct {
	// Context to leverage during the lifetime of the client
	Context context.Context
	// Token to use for authentication
	Token string
	// API client to utilize for making HTTP requests
	client *http.Client
	// apiURL is the base URL for the Automox API
	apiURL string
}

// Used if custom client not passed in when NewClient instantiated
func defaultHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Minute,
	}
}

// New returns a new Automox API client
func New(ctx context.Context, token string, client *http.Client) (*Client, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	if token == "" {
		return nil, missingClientConfigErr("Token")
	}

	// default to HTTP client if one is not provided
	if client == nil {
		client = defaultHTTPClient()
		client.Timeout = time.Minute * 5
	}

	return &Client{
		apiURL:  apiUrl,
		client:  client,
		Context: ctx,
		Token:   token,
	}, nil
}

// makeRequest is used internally by the Automox API client to
// make an API request and unmarshal into the response interface passed in
func (am *Client) makeRequest(r *http.Request, v interface{}) (*http.Response, error) {
	var bearer = "Bearer " + am.Token

	r.Header.Set("Accept", "application/json")
	r.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
	r.Header.Set("Strict-Transport-Security", "max-age=31536000 ; includeSubDomains")
	r.Header.Set("Authorization", bearer)

	// Replace scheme for unit tests that are using a mock server
	if os.Getenv("GO_TEST") == "1" {
		r.URL.Scheme = "http"
	}

	r.Close = true

	res, err := am.client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("error making %s request to %s", r.Method, r.URL)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		v = ErrorResponse{}
	}

	if v == nil {
		return res, nil
	}

	if os.Getenv("GO_DEBUG") == "1" {
		for k, v := range res.Header {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
	return res, json.NewDecoder(res.Body).Decode(&v)
}

// Servers is the interface between the HTTP client and the Automox servers related endpoints
func (am *Client) Servers() ServersService {
	return &ServersClient{client: am}
}
