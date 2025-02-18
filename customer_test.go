package vismanet

import (
	"encoding/json"
	"fmt"
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
	b, err := json.Marshal(resp.Customer)
	fmt.Println(string(b))
}
