package vismanet

import (
	"os"
	"testing"
)

func TestPutAttachmentV1(t *testing.T) {
	id := os.Getenv("TEST_ATTACHMENT_ID")
	req := testClient.NewPutAttachmentV1Request()
	req.SetPathParams(PutAttachmentV1PathParams{id})
	req.SetBody(RequestAttachment{SendToAutoInvoice: true})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.ResourceID() == "" {
		t.Errorf("Expected non-empty resource ID, got %s", resp.ResourceID())
	}
}
