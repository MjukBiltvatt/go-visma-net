package vismanet

import (
	"context"
	"os"
	"strings"
	"testing"

	"golang.org/x/oauth2/clientcredentials"
)

var testClient *Client

func TestMain(m *testing.M) {
	oauthConf := clientcredentials.Config{
		ClientID:     os.Getenv("TEST_CLIENT_ID"),
		ClientSecret: os.Getenv("TEST_CLIENT_SECRET"),
		TokenURL:     os.Getenv("TEST_TOKEN_URL"),
		Scopes:       strings.Split(os.Getenv("TEST_SCOPES"), " "),
		EndpointParams: map[string][]string{
			"tenant_id": {os.Getenv("TEST_TENANT_ID")},
		},
	}
	httpClient := oauthConf.Client(context.Background())
	testClient = NewClient(httpClient)
	testClient.Debug = os.Getenv("DEBUG") == "true"
	m.Run()
}
