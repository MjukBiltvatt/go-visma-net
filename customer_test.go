package vismanet

import (
	"fmt"
	"math/rand/v2"
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
	} else if resp.IPPRequestIDHeader() == "" {
		t.Error("Expected non-empty Ipp-Request-Id header")
	} else if resp.RequestContextHeader() == "" {
		t.Error("Expected non-empty Request-Context header")
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

func TestPutCustomerV1(t *testing.T) {
	customerCD := os.Getenv("TEST_CUSTOMER_CD")
	req := testClient.NewPutCustomerV1Request()
	req.SetPathParams(PutCustomerV1PathParams{customerCD})
	req.SetBody(RequestCustomer{
		Name: StringValue(fmt.Sprintf("Testbolaget %d AB", rand.IntN(99999999))),
	})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.ResourceID() == "" {
		t.Errorf("Expected non-empty resource ID, got %s", resp.ResourceID())
	}
}
