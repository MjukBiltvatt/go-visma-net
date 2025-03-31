A client library for the Visma Net Service API. For a better understanding of how it works, see the [official documentation](https://integration.visma.net/API-index/) for the API.

Inspired by [github.com/omniboost/go-visma.net](https://github.com/omniboost/go-visma.net).

# Setup
```
go get github.com/MjukBiltvatt/go-visma-net
```

# Usage

## Creating a client

```go
// Create an oath configuration, in this example we're using the client credentials flow
oauthConf := clientcredentials.Config{
	ClientID:     "isv_your_client_id",
	ClientSecret: "your_client_secret",
	TokenURL:     "https://connect.visma.com/connect/token",
	Scopes:       []string{
		"vismanet_erp_service_api:create",
		"vismanet_erp_service_api:delete",
		"vismanet_erp_service_api:read",
		"vismanet_erp_service_api:update"
	},
	EndpointParams: map[string][]string{
		"tenant_id": {"your_tenant_id"},
	},
}

// Create a http client using the oauth configuration
httpClient := oauthConf.Client(context.Background())

// Create a Visma.net client using the http client
client := vismanet.NewClient(httpClient)
```

## Attachment

### PUT `/controller/api/v1/attachment/{attachmentId}`

Update an existing attachment with the specified ID.

```go
req := testClient.NewPutAttachmentV1Request()
req.SetPathParams(PutAttachmentV1PathParams{AttachmentID: "attachment_id"})
req.SetBody(RequestAttachment{SendToAutoInvoice: true})
if err != nil {
	fmt.Println("Error updating attachment:", err)
}
fmt.Println("Attachment updated successfully:", resp.ResourceID())
```

## Customer

### GET `/controller/api/v1/customer/{customerCd}`

Get a customer with a specific customer number.

```go
req := client.NewGetCustomerV1Request()
req.SetPathParams(GetCustomerV1PathParams{
	customerCD: "customer_number",
})
resp, err := req.Do()
if err != nil {
	if resp.StatusCode() == http.StatusNotFound {
		fmt.Println("Customer not found")
	} else {
		fmt.Println("Error getting customer:", err)
	}
}
fmt.Println("Customer retrieved successfully:", resp.Customer.Number)
```

### POST `/controller/api/v1/customer`

Create a new customer.

```go
req := client.NewPostCustomerV1Request()
req.SetBody(RequestCustomer{
	Name: "Test",
	MainAddress: &RequestNestedAddress{
		AddressLine1: "Testgatan 1",
		City:         "Testdalen",
		PostalCode:   "12345",
		CountryID:    "SE",
	},
})
resp, err := req.Do()
if err != nil {
	fmt.Println("Error creating customer:", err)
}
fmt.Println("Customer created successfully:", resp.ResourceID())
```

### PUT `/controller/api/v1/customer/{customerCd}`

Update an existing customer with the specified customer number.

```go
req := testClient.NewPutCustomerV1Request()
req.SetPathParams(PutCustomerV1PathParams{
	customerCD: "customer_number",
})
req.SetBody(RequestCustomer{
	Name: "John Doe",
})
resp, err := req.Do()
if err != nil {
	fmt.Println("Error updating customer:", err)
}
fmt.Println("Customer updated successfully:", resp.ResourceID())
```

## CustomerInvoice

### GET `/controller/api/v1/customerinvoice/{invoiceNumber}`

Get a customer invoice with a specific invoice number.

```go
req := testClient.NewGetCustomerInvoiceV1Request()
req.SetPathParams(GetCustomerInvoiceV1PathParams{
	invoiceNumber: "invoice_number",
})
resp, err := req.Do()
if err != nil {
	if resp.StatusCode() == http.StatusNotFound {
		fmt.Println("Invoice not found")
	} else {
		fmt.Println("Error getting invoice:", err)
	}
}
fmt.Println("Invoice retrieved successfully:", resp.Customer.Number)
```

### DELETE `/controller/api/v1/customerinvoice/{invoiceNumber}`

Delete a customer invoice with a specific invoice number.

```go
req := testClient.NewGetCustomerInvoiceV1Request()
req.SetPathParams(GetCustomerInvoiceV1PathParams{
	invoiceNumber: "invoice_number",
})
resp, err := req.Do()
if err != nil {
	if resp.StatusCode() == http.StatusNotFound {
		fmt.Println("Invoice not found")
	} else {
		fmt.Println("Error deleting invoice:", err)
	}
}
fmt.Println("Invoice deleted successfully")
```

### POST `/controller/api/v1/customerinvoice/{invoiceNumber}/attachment`

Upload a file as an attachment to a customer invoice.

```go
req := testClient.NewPostCustomerInvoiceAttachmentV1Request()
req.SetPathParams(PostCustomerInvoiceAttachmentV1PathParams{
	InvoiceNumber: os.Getenv("TEST_CUSTOMER_INVOICE_NUMBER"),
})
req.SetBody(FileUploadBody{
	Files: []File{
		{
			Name:    "Test.txt",
			Content: []byte("Hello world"),
		},
	},
})
resp, err := req.Do()
if err != nil {
	fmt.Println("Error creating attachment:", err)
}
fmt.Println("Attachment created successfully:", resp.ResourceID())
```

## CustomerInvoiceV2

### POST `/controller/api/v2/customerinvoice`

Create a new customer invoice.

```go
req := testClient.NewPostCustomerInvoiceV2Request()
req.SetBody(RequestInvoice{
	CurrencyID:     "SEK",
	CustomerNumber: StringValue(os.Getenv("TEST_CUSTOMER_CD")),
	InvoiceLines: []RequestInvoiceLine{
		{
			Operation:           "Insert",
			ItemType:            "Service",
			BranchNumber:        "1",
			Description:         "Test",
			Quantity:            1,
			UnitPriceInCurrency: 100,
			AccountNumber:       "3015",
			Subaccount: []RequestSegment{
				{
					SegmentID:    1,
					SegmentValue: "00",
				},
				{
					SegmentID:    2,
					SegmentValue: "000",
				},
				{
					SegmentID:    3,
					SegmentValue: "00",
				},
			},
		},
	},
})
resp, err := req.Do()
if err != nil {
	fmt.Println("Error creating invoice:", err)
}
fmt.Println("Invoice created successfully:", resp.ResourceID())
```

# Testing

Set up your testing environment by copying the contents of `.env.template` to a `.env` file. After that you can use [godotenv](https://github.com/joho/godotenv) to set the environment variables required for testing.

Run all tests:
```
$ godotenv go test
```