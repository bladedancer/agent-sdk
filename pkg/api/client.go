package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"git.ecd.axway.int/apigov/apic_agents_sdk/pkg/config"
	log "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/util/log"
)

// HTTP const definitions
const (
	GET    string = http.MethodGet
	POST   string = http.MethodPost
	PUT    string = http.MethodPut
	DELETE string = http.MethodDelete
)

// Request - the request object used when communicating to an API
type Request struct {
	Method      string
	URL         string
	QueryParams map[string]string
	Headers     map[string]string
	Body        []byte
}

// Response - the response object given back when communicating to an API
type Response struct {
	Code    int
	Body    []byte
	Headers map[string][]string
}

// Client -
type Client interface {
	Send(request Request) (*Response, error)
}

type httpClient struct {
	Client
	httpClient *http.Client
}

// NewClient - creates a new API client using the http client sent in
func NewClient(cfg config.TLSConfig, proxyURL string) Client {
	httpCli := http.DefaultClient
	if cfg != nil {
		url, err := url.Parse(proxyURL)
		if err != nil {
			log.Errorf("Error parsing proxyURL from config; creating a non-proxy client: %s", err.Error())
		}
		httpCli = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: cfg.BuildTLSConfig(),
				Proxy:           getProxyURL(url),
			},
		}
	}

	httpCli.Timeout = time.Second * 10
	return &httpClient{
		httpClient: httpCli,
	}
}

// need to provide my own function (instead of http.ProxyURL()) to handle empty url. Returning nil
// means "no proxy"
func getProxyURL(fixedURL *url.URL) func(*http.Request) (*url.URL, error) {
	return func(*http.Request) (*url.URL, error) {
		if fixedURL == nil || fixedURL.Host == "" {
			return nil, nil
		}
		return fixedURL, nil
	}
}

func (c *httpClient) getURLEncodedQueryParams(queryParams map[string]string) string {
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return params.Encode()
}

func (c *httpClient) prepareAPIRequest(request Request) (*http.Request, error) {
	requestURL := request.URL
	if len(request.QueryParams) != 0 {
		requestURL += "?" + c.getURLEncodedQueryParams(request.QueryParams)
	}
	req, err := http.NewRequest(request.Method, requestURL, bytes.NewBuffer(request.Body))
	if err != nil {
		return req, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	return req, err
}

func (c *httpClient) prepareAPIResponse(res *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(res.Body)
	response := Response{
		Code:    res.StatusCode,
		Body:    body,
		Headers: res.Header,
	}
	return &response, err
}

// Send - send the http request and returns the API Response
func (c *httpClient) Send(request Request) (*Response, error) {
	req, err := c.prepareAPIRequest(request)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return c.prepareAPIResponse(res)
}
