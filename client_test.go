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
		AppToken:    "CPyc2169FpYjZXZNKjDg4woylV8hONVyNYGnP5B3",
		AuthUser: AuthUserClient{
			Username: "glpi",
			Password: "glpi",
		},
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
		Impact:       1,
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
	ticket, err := glpiClient.ReadTicket(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(*ticket)
	err = glpiClient.AddFollowupTicket(1, AddFollowupTicket{
		IsPrivate:      "1",
		RequestTypesId: "6",
		Content:        "Followup Added with Golang!!!",
	})
	if err != nil {
		panic(err)
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
