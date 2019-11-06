package glpi_api_client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func getSessionToken(resp *http.Response) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	responseJson := make(map[string]interface{})
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return "", err
	}
	value, ok := responseJson["session_token"].(string)
	if !ok {
		return "", errors.New(errorUnexpectedResponse)
	}
	return value, nil
}

func doRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

func getCreateTicketResponseID(resp *http.Response) (int, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var response CreateTicketResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}
	return response.Id, nil
}

func getReadTickets(resp *http.Response) ([]ReadTicket, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var readTickets []ReadTicket
	err = json.Unmarshal(body, &readTickets)
	if err != nil {
		return nil, err
	}
	return readTickets, nil
}

type CreateTicketResponse struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

func getInputJson(ticket CreateTicket) ([]byte, error) {
	inputItem := InputItem{Input: ticket}
	inputJson, err := json.Marshal(inputItem)
	if err != nil {
		return nil, err
	}
	return inputJson, nil
}
