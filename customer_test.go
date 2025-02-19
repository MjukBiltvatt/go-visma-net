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
	if err != nil {
		t.Error(err)
	} else if resp.Customer.Number != customerCD {
		t.Errorf("Expected customer number %s, got %s", customerCD, resp.Customer.Number)
	}
}

func TestPostCustomerV1(t *testing.T) {
	req := testClient.NewPostCustomerV1Request()
	req.SetBody(CustomerRequestBody{Name: "Test"})
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	} else if resp.ResourceID() == "" {
		t.Errorf("Expected non-empty resource ID, got %s", resp.ResourceID())
	}
}
