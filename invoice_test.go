package vismanet

import (
	"os"
	"testing"
)

func TestGetInvoiceV1(t *testing.T) {
	invoiceNumber := os.Getenv("TEST_INVOICE_NUMBER")
	req := testClient.NewGetInvoiceV1Request()
	req.SetPathParams(GetInvoiceV1PathParams{invoiceNumber})
	resp, err := req.Do()
	debugDumpResponse(testClient, resp)
	if err != nil {
		t.Error(err)
	} else if resp.Invoices[0].ReferenceNumber != invoiceNumber {
		t.Errorf("Expected invoice number %s, got %s", invoiceNumber, resp.Invoices[0].ReferenceNumber)
	}
}
