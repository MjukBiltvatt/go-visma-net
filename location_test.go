package vismanet

import (
	"fmt"
	"math/rand/v2"
	"os"
	"testing"
)

func TestPutLocationV1(t *testing.T) {
	customerCD := os.Getenv("TEST_CUSTOMER_CD")
	req := testClient.NewPutLocationV1Request()
	req.SetPathParams(PutLocationV1PathParams{customerCD, "Main"})
	req.SetBody(RequestLocation{
		LocationName: NewStringValue(fmt.Sprintf("Location %d", rand.IntN(99999999))),
	})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.ResourceID() == "" {
		t.Errorf("Expected non-empty resource ID, got %s", resp.ResourceID())
	}
}
