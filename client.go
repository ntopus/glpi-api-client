package glpi_api_client

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

const headerContentType = "Content-Type"
const headerAppToken = "App-Token"
const headerSessionToken = "Session-Token"
const headerAuthorization = "Authorization"

const valueApplicationJson = "application/json"
const valuePrefixAuthHeader = "Basic "

const errorUnexpectedResponse = "unexpected response"

const urlInitSessionPath = "initSession"
const urlTicketPath = "ticket"
const urlTicketFollowup = "ticketFollowup"

type InputItem struct {
	Input interface{} `json:"input"`
}

type GLPIClient struct {
	apiEndpoint     url.URL
	appToken        string
	sessionToken    string
	authHeaderValue string
}

func NewGLPIClient(config GlpiClientConfig) *GLPIClient {
	apiEndpoint := config.ApiEndpoint
	apiEndpoint.Path = path.Join(apiEndpoint.Path, "api")
	authHeaderValue := base64.StdEncoding.EncodeToString([]byte(config.AuthUser.Username + ":" + config.AuthUser.Password))
	return &GLPIClient{
		apiEndpoint:     apiEndpoint,
		appToken:        config.AppToken,
		authHeaderValue: authHeaderValue,
	}
}

func (glpiClient *GLPIClient) InitSession() error {
	initSessionEndpoint := glpiClient.createEndpoint(urlInitSessionPath)
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

func (glpiClient *GLPIClient) CreateTicket(ticket CreateTicket) (int, error) {
	json, err := getInputJson(ticket)
	if err != nil {
		return 0, err
	}
	createTicketEndpoint := glpiClient.createEndpoint(urlTicketPath)
	resp, err := glpiClient.makeAndDoRequest(http.MethodPost, createTicketEndpoint.String(), json)
	if err != nil {
		return 0, err
	}
	return getCreateTicketResponseID(resp)
}

func (glpiClient *GLPIClient) UpdateTicket(id int, ticket CreateTicket) error {
	json, err := getInputJson(ticket)
	if err != nil {
		return err
	}
	idString := strconv.Itoa(id)
	urlPath := path.Join(urlTicketPath, idString)
	createTicketEndpoint := glpiClient.createEndpoint(urlPath)
	resp, err := glpiClient.makeAndDoRequest(http.MethodPut, createTicketEndpoint.String(), json)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func (glpiClient *GLPIClient) AddFollowupTicket(id int, ticket AddFollowupTicket) error {
	json, err := getInputJson(ticket)
	if err != nil {
		return err
	}
	idString := strconv.Itoa(id)
	urlPath := path.Join(urlTicketPath, idString, urlTicketFollowup)
	createTicketEndpoint := glpiClient.createEndpoint(urlPath)
	resp, err := glpiClient.makeAndDoRequest(http.MethodPost, createTicketEndpoint.String(), json)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

//TODO: Implementar paginação.
func (glpiClient *GLPIClient) ReadAllTickets() ([]ReadTicket, error) {
	readTicketsEndpoint := glpiClient.createEndpoint(urlTicketPath)
	resp, err := glpiClient.makeAndDoRequest(http.MethodGet, readTicketsEndpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	return getReadAllTickets(resp)
}

func (glpiClient *GLPIClient) ReadTicket(id int) (*ReadTicket, error) {
	idString := strconv.Itoa(id)
	urlPath := path.Join(urlTicketPath, idString)
	readTicketsEndpoint := glpiClient.createEndpoint(urlPath)
	resp, err := glpiClient.makeAndDoRequest(http.MethodGet, readTicketsEndpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	return getReadTicket(resp)
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
	resp, err := (&http.Client{}).Do(req)
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
	req.Header.Add(headerContentType, valueApplicationJson)
	req.Header.Add(headerAppToken, glpiClient.appToken)
	req.Header.Add(headerAuthorization, valuePrefixAuthHeader+glpiClient.authHeaderValue)
	if glpiClient.sessionToken != "" {
		req.Header.Add(headerSessionToken, glpiClient.sessionToken)
	}
}
