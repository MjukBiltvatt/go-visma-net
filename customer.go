package vismanet

import "encoding/json"

// ResponseCustomer is a customer as represented in a response from the Visma.net API
type ResponseCustomer struct {
	InternalID                   int                  `json:"internalId"`
	Number                       string               `json:"number"`
	Name                         string               `json:"name"`
	Status                       string               `json:"status"`
	AccountReference             string               `json:"accountReference"`
	ConsolidateInvoices          bool                 `json:"consolidateInvoices"`
	CurrencyID                   string               `json:"currencyId"`
	CreditVerification           string               `json:"creditVerification"`
	CreditLimit                  float64              `json:"creditLimit"`
	CreditDaysPastDue            int                  `json:"creditDaysPastDue"`
	PrintInvoices                bool                 `json:"printInvoices"`
	AcceptAutoInvoices           bool                 `json:"acceptAutoInvoices"`
	SendInvoicesByEmail          bool                 `json:"sendInvoicesByEmail"`
	SendDunningLettersViaEmail   bool                 `json:"sendDunningLettersViaEMail"`
	PrintDunningLetters          bool                 `json:"printDunningLetters"`
	PrintStatements              bool                 `json:"printStatements"`
	SendStatementsByEmail        bool                 `json:"sendStatementsByEmail"`
	PrintMultiCurrencyStatements bool                 `json:"printMultiCurrencyStatements"`
	StatementType                string               `json:"statementType"`
	VatRegistrationID            string               `json:"vatRegistrationId"`
	CorporateID                  string               `json:"corporateId"`
	InvoiceToDefaultLocation     bool                 `json:"invoiceToDefaultLocation"`
	DefaultPaymentMethodID       string               `json:"defaultPaymentMethodId"`
	ExcludeDebtCollection        bool                 `json:"excludeDebtCollection"`
	LastModifiedDateTime         Time                 `json:"lastModifiedDateTime"`
	CreatedDateTime              Time                 `json:"createdDateTime"`
	MainAddress                  *Address             `json:"mainAddress"`
	MainContact                  *Contact             `json:"mainContact"`
	InvoiceAddress               *Address             `json:"invoiceAddress"`
	InvoiceContact               *Contact             `json:"invoiceContact"`
	DeliveryAddress              *Address             `json:"deliveryAddress"`
	DeliveryContact              *Contact             `json:"deliveryContact"`
	CustomerClass                *IDDescriptionEntity `json:"customerClass"`
	CreditTerms                  *IDDescriptionEntity `json:"creditTerms"`
	VatZone                      *IDDescriptionEntity `json:"vatZone"`

	Location struct {
		ID        string `json:"id"`
		CountryID string `json:"countryId"`
		Name      string `json:"name"`
	} `json:"location"`

	GLAccounts struct {
		CustomerLedgerAccount    Account    `json:"customerLedgerAccount"`
		CustomerLedgerSubaccount Subaccount `json:"customerLedgerSubaccount"`
		SalesAccount             Account    `json:"salesAccount"`
		SalesNonTaxableAccount   Account    `json:"salesNonTaxableAccount"`
		SalesSubaccount          Subaccount `json:"salesSubaccount"`
		CashDiscountAccount      Account    `json:"cashDiscountAccount"`
		CashDiscountSubaccount   Subaccount `json:"cashDiscountSubaccount"`
		PrepaymentAccount        Account    `json:"prepaymentAccount"`
		PrepaymentSubaccount     Subaccount `json:"prepaymentSubaccount"`
	} `json:"glAccounts"`

	Attributes []struct {
		//TODO
	} `json:"attributes"`

	PaymentMethods []struct {
		//TODO
	} `json:"paymentMethods"`
}

// RequestCustomer is a customer as represented in a request to the Visma.net API
type RequestCustomer struct {
	Number                       *StringValue               `json:"number,omitempty"`
	Name                         StringValue                `json:"name"`
	Status                       *StringValue               `json:"status,omitempty"`
	AccountReference             *StringValue               `json:"accountReference,omitempty"`
	NumberOfEmployees            *IntValue                  `json:"numberOfEmployees,omitempty"`
	ParentRecordNumber           *StringValue               `json:"parentRecordNumber,omitempty"`
	CurrencyID                   *StringValue               `json:"currencyId,omitempty"`
	CreditLimit                  *IntValue                  `json:"creditLimit,omitempty"`
	CreditDaysPastDue            *IntValue                  `json:"creditDaysPastDue,omitempty"`
	CustomerClassID              *StringValue               `json:"customerClassId,omitempty"`
	CreditTermsID                *StringValue               `json:"creditTermsId,omitempty"`
	PrintInvoices                *BoolValue                 `json:"printInvoices,omitempty"`
	AcceptAutoInvoices           *BoolValue                 `json:"acceptAutoInvoices,omitempty"`
	SendInvoicesByEmail          *BoolValue                 `json:"sendInvoicesByEmail,omitempty"`
	SendDunningLettersViaEmail   *BoolValue                 `json:"sendDunningLettersViaEMail,omitempty"`
	PrintDunningLetters          *BoolValue                 `json:"printDunningLetters,omitempty"`
	PrintStatements              *BoolValue                 `json:"printStatements,omitempty"`
	SendStatementsByEmail        *BoolValue                 `json:"sendStatementsByEmail,omitempty"`
	PrintMultiCurrencyStatements *BoolValue                 `json:"printMultiCurrencyStatements,omitempty"`
	InvoiceToDefaultLocation     *BoolValue                 `json:"invoiceToDefaultLocation,omitempty"`
	VatRegistrationID            *StringValue               `json:"vatRegistrationId,omitempty"`
	CorporateID                  *StringValue               `json:"corporateId,omitempty"`
	VatZoneID                    *StringValue               `json:"vatZoneId,omitempty"`
	GLN                          *StringValue               `json:"gln,omitempty"`
	Note                         *StringValue               `json:"note,omitempty"`
	MainAddress                  *RequestNestedAddress      `json:"mainAddress,omitempty"`
	MainContact                  *RequestNestedContact      `json:"mainContact,omitempty"`
	CreditVerification           *StringValue               `json:"creditVerification,omitempty"`
	InvoiceAddress               *RequestNestedAddress      `json:"invoiceAddress,omitempty"`
	InvoiceContact               *RequestNestedContact      `json:"invoiceContact,omitempty"`
	StatementType                *StringValue               `json:"statementType,omitempty"`
	DeliveryAddress              *RequestNestedAddress      `json:"deliveryAddress,omitempty"`
	DeliveryContact              *RequestNestedContact      `json:"deliveryContact,omitempty"`
	PriceClassID                 *StringValue               `json:"priceClassId,omitempty"`
	OverrideNumberSeries         *BoolValue                 `json:"overrideNumberSeries,omitempty"`
	ExcludeDebtCollection        *BoolValue                 `json:"excludeDebtCollection,omitempty"`
	GLAccounts                   *RequestCustomerGLAccounts `json:"glAccounts,omitempty"`
	//TODO: implement overrideWithClassValues? bool with no value wrapper...
	//TODO: implement eInvoiceContract?
	//TODO: implement defaultPaymentMethod?
	//TODO: implement directDebitLines?
	//TODO: implement attributeLines?
}

// RequestCustomerGLAccounts is a customer's GL accounts as represented in a request to the Visma.net API
type RequestCustomerGLAccounts struct {
	CustomerLedgerAccount    *StringValue      `json:"customerLedgerAccount,omitempty"`
	CustomerLedgerSubaccount *[]RequestSegment `json:"customerLedgerSubaccount,omitempty"`
	SalesAccount             *StringValue      `json:"salesAccount,omitempty"`
	SalesNonTaxableAccount   *StringValue      `json:"salesNonTaxableAccount,omitempty"`
	SalesEuAccount           *StringValue      `json:"salesEuAccount,omitempty"`
	SalesExportAccount       *StringValue      `json:"salesExportAccount,omitempty"`
	SalesSubaccount          *[]RequestSegment `json:"salesSubaccount,omitempty"`
	DiscountAccount          *StringValue      `json:"discountAccount,omitempty"`
	DiscountSubaccount       *[]RequestSegment `json:"discountSubaccount,omitempty"`
	FreightAccount           *StringValue      `json:"freightAccount,omitempty"`
	FreightSubaccount        *[]RequestSegment `json:"freightSubaccount,omitempty"`
	CashDiscountAccount      *StringValue      `json:"cashDiscountAccount,omitempty"`
	CashDiscountSubaccount   *[]RequestSegment `json:"cashDiscountSubaccount,omitempty"`
	PrepaymentAccount        *StringValue      `json:"prepaymentAccount,omitempty"`
	PrepaymentSubaccount     *[]RequestSegment `json:"prepaymentSubaccount,omitempty"`
}

func (v *RequestCustomerGLAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*v})
}

// =========================================================
// ========================== GET ==========================
// =========================================================

// newGetCustomerV1Request creates a new GetCustomerV1Request
func newGetCustomerV1Request(c *Client) GetCustomerV1Request {
	return GetCustomerV1Request{
		Client: c,
		Method: "GET",
		Path:   "controller/api/v1/customer/{{.customer_cd}}",
	}
}

// GetCustomerV1Request represents a request to get a customer
type GetCustomerV1Request Request

// SetPathParams sets the path parameters of the request
func (r *GetCustomerV1Request) SetPathParams(params GetCustomerV1PathParams) {
	r.pathParams = params
}

// Do performs the request and returns the response
func (r *GetCustomerV1Request) Do() (GetCustomerV1Response, error) {
	var customer ResponseCustomer
	resp, err := r.Client.Do((*Request)(r), &customer)
	return GetCustomerV1Response{Response{resp}, customer}, err
}

// GetCustomerV1PathParams represents the path parameters of the GetCustomerV1Request
type GetCustomerV1PathParams struct {
	CustomerCD string `schema:"customer_cd"`
}

// GetCustomerV1Response represents the response of the GetCustomerV1Request and contains the resulting customer
type GetCustomerV1Response struct {
	Response
	Customer ResponseCustomer
}

// =========================================================
// ========================== POST =========================
// =========================================================

// newPostCustomerV1Request creates a new PostCustomerV1Request
func newPostCustomerV1Request(c *Client) PostCustomerV1Request {
	return PostCustomerV1Request{
		Client: c,
		Method: "POST",
		Path:   "controller/api/v1/customer",
	}
}

// PostCustomerV1Request represents a request to create a new customer
type PostCustomerV1Request Request

// SetBody sets the body of the request
func (r *PostCustomerV1Request) SetBody(body RequestCustomer) {
	r.Body = JSONRequestBody{body}
}

// Do performs the request and returns the response
func (r *PostCustomerV1Request) Do() (PostCustomerV1Response, error) {
	resp, err := r.Client.Do((*Request)(r), nil)
	return PostCustomerV1Response{Response{resp}}, err
}

// PostCustomerV1Response represents the response of the PostCustomerV1Request
type PostCustomerV1Response struct {
	Response
}

// =========================================================
// ========================== PUT ==========================
// =========================================================

// newPutCustomerV1Request creates a new PutCustomerV1Request
func newPutCustomerV1Request(c *Client) PutCustomerV1Request {
	return PutCustomerV1Request{
		Client: c,
		Method: "PUT",
		Path:   "controller/api/v1/customer/{{.customer_cd}}",
	}
}

// PutCustomerV1Request represents a request to create a new customer
type PutCustomerV1Request Request

// SetPathParams sets the path parameters of the request
func (r *PutCustomerV1Request) SetPathParams(params PutCustomerV1PathParams) {
	r.pathParams = params
}

// SetBody sets the body of the request
func (r *PutCustomerV1Request) SetBody(body RequestCustomer) {
	r.Body = JSONRequestBody{body}
}

// Do performs the request and returns the response
func (r *PutCustomerV1Request) Do() (PutCustomerV1Response, error) {
	resp, err := r.Client.Do((*Request)(r), nil)
	return PutCustomerV1Response{Response{resp}}, err
}

// PutCustomerV1PathParams represents the path parameters of the PutCustomerV1Request
type PutCustomerV1PathParams struct {
	CustomerCD string `schema:"customer_cd"`
}

// PutCustomerV1Response represents the response of the PutCustomerV1Request
type PutCustomerV1Response struct {
	Response
}
