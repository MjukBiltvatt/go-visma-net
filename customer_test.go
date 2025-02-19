package vismanet

import (
	"os"
	"testing"
)

func TestGetCustomerV1(t *testing.T) {
	customerCD := os.Getenv("TEST_CUSTOMER_CD")
	req := testClient.NewGetCustomerV1Request()
	req.SetPathParams(GetCustomerV1PathParams{customerCD})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.Customer.Number != customerCD {
		t.Errorf("Expected customer number %s, got %s", customerCD, resp.Customer.Number)
	}
}

func TestPostCustomerV1(t *testing.T) {
	req := testClient.NewPostCustomerV1Request()
	req.SetBody(RequestCustomer{
		Name: "Test",
		MainAddress: &RequestNestedAddress{
			AddressLine1: "123 Test St",
			City:         "Testville",
			PostalCode:   "12345",
			CountryID:    "US",
		},
	})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.ResourceID() == "" {
		t.Errorf("Expected non-empty resource ID, got %s", resp.ResourceID())
	}
}
