package glpi_api_client

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_GLPI(t *testing.T) {
	glpiEndpoint, err := url.Parse("http://localhost")
	if err != nil {
		panic(err)
	}
	glpiClient := NewGLPIClient(GlpiClientConfig{
		ApiEndpoint: *glpiEndpoint,
		AppToken:    "1qNmrctAia7ZvFP20wT2GhuFx0o6NskjgjjAW863",
		AuthUser:    AuthUserClient{},
	})
	err = glpiClient.InitSession()
	if err != nil {
		panic(err)
	}
	ticketId, err := glpiClient.CreateTicket(CreateTicket{
		Name:         "Test Ticket 2",
		Content:      "Content test ticket.",
		Status:       1,
		Urgency:      1,
		DisableNotif: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Create ticket, id:", ticketId)
	readTickets, err := glpiClient.ReadAllTickets()
	if err != nil {
		panic(err)
	}
	for _, value := range readTickets {
		fmt.Println(value)
	}
}

func Test_UpdateTicket(t *testing.T) {
	glpiEndpoint, err := url.Parse("http://localhost")
	if err != nil {
		panic(err)
	}
	glpiClient := NewGLPIClient(GlpiClientConfig{
		ApiEndpoint: *glpiEndpoint,
		AppToken:    "1qNmrctAia7ZvFP20wT2GhuFx0o6NskjgjjAW863",
		AuthUser:    AuthUserClient{},
	})
	err = glpiClient.InitSession()
	if err != nil {
		panic(err)
	}
	err = glpiClient.UpdateTicket(33, CreateTicket{
		Name:         "Test Ticket 2",
		Content:      "Content test ticket.",
		Status:       1,
		Urgency:      1,
		DisableNotif: true,
	})
	if err != nil {
		panic(err)
	}
}
