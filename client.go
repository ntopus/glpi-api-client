package glpi_api_client

import (
	"bytes"
	"net/http"
	"net/url"
	"path"
)

const headerContentType = "Content-Type"
const headerAppToken = "App-Token"
const headerSessionToken = "Session-Token"
const headerAuthorization = "Authorization"

const applicationJson = "application/json"
const authHeaderValue = "Basic Z2xwaTpnbHBp"

const errorUnexpectedResponse = "unexpected response"

const initSessionUrlPath = "initSession"
const ticketUrlPath = "ticket"

type InputItem struct {
	Input interface{} `json:"input"`
}

type GLPIClient struct {
	apiEndpoint  url.URL
	appToken     string
	sessionToken string
}

func NewGLPIClient(config GlpiClientConfig) *GLPIClient {
	apiEndpoint := config.ApiEndpoint
	apiEndpoint.Path = path.Join(apiEndpoint.Path, "api")
	return &GLPIClient{
		apiEndpoint: apiEndpoint,
		appToken:    config.AppToken,
	}
}

func (glpiClient *GLPIClient) InitSession() error {
	initSessionEndpoint := glpiClient.createEndpoint(initSessionUrlPath)
	resp, err := glpiClient.makeAndDoRequest(http.MethodGet, initSessionEndpoint.String(), nil)
	if err != nil {
		return err
	}
	value, err := getSessionToken(resp)
	if err != nil {
		return err
	}
	glpiClient.sessionToken = value
	return nil
}

func (glpiClient *GLPIClient) AddTicket(ticket CreateTicket) (int, error) {
	json, err := getInputJson(ticket)
	if err != nil {
		return 0, err
	}
	createTicketEndpoint := glpiClient.createEndpoint(ticketUrlPath)
	resp, err := glpiClient.makeAndDoRequest(http.MethodPost, createTicketEndpoint.String(), json)
	if err != nil {
		return 0, err
	}
	return getCreateTicketResponseID(resp)
}

func (glpiClient *GLPIClient) GetTickets() ([]ReadTicket, error) {
	readTicketsEndpoint := glpiClient.createEndpoint(ticketUrlPath)
	resp, err := glpiClient.makeAndDoRequest(http.MethodGet, readTicketsEndpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	return getReadTickets(resp)
}

func (glpiClient *GLPIClient) makeAndDoRequest(method, endpoint string, body []byte) (*http.Response, error) {
	var req *http.Request
	var err error
	if len(body) > 0 {
		reader := bytes.NewReader(body)
		req, err = http.NewRequest(method, endpoint, reader)
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}
	if err != nil {
		return nil, err
	}
	glpiClient.addRequestHeaders(req)
	resp, err := doRequest(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (glpiClient *GLPIClient) createEndpoint(pathEndpoint string) url.URL {
	initSessionEndpoint := glpiClient.apiEndpoint
	initSessionEndpoint.Path = path.Join(glpiClient.apiEndpoint.Path, pathEndpoint)
	return initSessionEndpoint
}

func (glpiClient *GLPIClient) addRequestHeaders(req *http.Request) {
	req.Header.Add(headerContentType, applicationJson)
	req.Header.Add(headerAppToken, glpiClient.appToken)
	req.Header.Add(headerAuthorization, authHeaderValue)
	if glpiClient.sessionToken != "" {
		req.Header.Add(headerSessionToken, glpiClient.sessionToken)
	}
}
