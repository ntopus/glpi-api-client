package glpi_api_client

import "net/url"

type GlpiClientConfig struct {
	ApiEndpoint url.URL
	AppToken    string
	AuthUser    AuthUserClient
}

type AuthUserClient struct {
	Username string
	Password string
}
