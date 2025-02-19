package vismanet

// Customer represents a customer as returned in a GET request
type Customer struct {
	InternalID                   int                `json:"internalId"`
	Number                       string             `json:"number"`
	Name                         string             `json:"name"`
	Status                       string             `json:"status"`
	MainAddress                  Address            `json:"mainAddress"`
	MainContact                  Contact            `json:"mainContact"`
	AccountReference             string             `json:"accountReference"`
	ConsolidateInvoices          bool               `json:"consolidateInvoices"`
	CustomerClass                Description        `json:"customerClass"`
	CreditTerms                  Description        `json:"creditTerms"`
	CurrencyID                   string             `json:"currencyId"`
	CreditVerification           string             `json:"creditVerification"`
	CreditLimit                  float64            `json:"creditLimit"`
	CreditDaysPastDue            int                `json:"creditDaysPastDue"`
	InvoiceAddress               Address            `json:"invoiceAddress"`
	InvoiceContact               Contact            `json:"invoiceContact"`
	PrintInvoices                bool               `json:"printInvoices"`
	AcceptAutoInvoices           bool               `json:"acceptAutoInvoices"`
	SendInvoicesByEmail          bool               `json:"sendInvoicesByEmail"`
	SendDunningLettersViaEmail   bool               `json:"sendDunningLettersViaEMail"`
	PrintDunningLetters          bool               `json:"printDunningLetters"`
	PrintStatements              bool               `json:"printStatements"`
	SendStatementsByEmail        bool               `json:"sendStatementsByEmail"`
	PrintMultiCurrencyStatements bool               `json:"printMultiCurrencyStatements"`
	StatementType                string             `json:"statementType"`
	DeliveryAddress              Address            `json:"deliveryAddress"`
	DeliveryContact              Contact            `json:"deliveryContact"`
	VatRegistrationID            string             `json:"vatRegistrationId"`
	CorporateID                  string             `json:"corporateId"`
	VatZone                      Description        `json:"vatZone"`
	Location                     Location           `json:"location"`
	Attributes                   []Attribute        `json:"attributes"`
	LastModifiedDateTime         Time               `json:"lastModifiedDateTime"`
	CreatedDateTime              Time               `json:"createdDateTime"`
	GLAccounts                   CustomerGLAccounts `json:"glAccounts"`
	InvoiceToDefaultLocation     bool               `json:"invoiceToDefaultLocation"`
	PaymentMethods               []PaymentMethod    `json:"paymentMethods"`
	DefaultPaymentMethodID       string             `json:"defaultPaymentMethodId"`
	ExcludeDebtCollection        bool               `json:"excludeDebtCollection"`
}

type CustomerGLAccounts struct {
	CustomerLedgerAccount    Account    `json:"customerLedgerAccount"`
	CustomerLedgerSubaccount Subaccount `json:"customerLedgerSubaccount"`
	SalesAccount             Account    `json:"salesAccount"`
	SalesNonTaxableAccount   Account    `json:"salesNonTaxableAccount"`
	SalesSubaccount          Subaccount `json:"salesSubaccount"`
	CashDiscountAccount      Account    `json:"cashDiscountAccount"`
	CashDiscountSubaccount   Subaccount `json:"cashDiscountSubaccount"`
	PrepaymentAccount        Account    `json:"prepaymentAccount"`
	PrepaymentSubaccount     Subaccount `json:"prepaymentSubaccount"`
}

type CustomerRequestBody struct {
	Number                       StringValue `json:"number,omitempty"`
	Name                         StringValue `json:"name"`
	Status                       StringValue `json:"status,omitempty"`
	AccountReference             StringValue `json:"accountReference,omitempty"`
	NumberOfEmployees            IntValue    `json:"numberOfEmployees,omitempty"`
	ParentRecordNumber           StringValue `json:"parentRecordNumber,omitempty"`
	CurrencyID                   StringValue `json:"currencyId,omitempty"`
	CreditLimit                  IntValue    `json:"creditLimit,omitempty"`
	CreditDaysPastDue            IntValue    `json:"creditDaysPastDue,omitempty"`
	CustomerClassID              StringValue `json:"customerClassId,omitempty"`
	CreditTermsID                StringValue `json:"creditTermsId,omitempty"`
	PrintInvoices                BoolValue   `json:"printInvoices,omitempty"`
	AcceptAutoInvoices           BoolValue   `json:"acceptAutoInvoices,omitempty"`
	SendInvoicesByEmail          BoolValue   `json:"sendInvoicesByEmail,omitempty"`
	SendDunningLettersViaEmail   BoolValue   `json:"sendDunningLettersViaEMail,omitempty"`
	PrintDunningLetters          BoolValue   `json:"printDunningLetters,omitempty"`
	PrintStatements              BoolValue   `json:"printStatements,omitempty"`
	SendStatementsByEmail        BoolValue   `json:"sendStatementsByEmail,omitempty"`
	PrintMultiCurrencyStatements BoolValue   `json:"printMultiCurrencyStatements,omitempty"`
	InvoiceToDefaultLocation     BoolValue   `json:"invoiceToDefaultLocation,omitempty"`
	VatRegistrationID            StringValue `json:"vatRegistrationId,omitempty"`
	CorporateID                  StringValue `json:"corporateId,omitempty"`
	VatZoneID                    StringValue `json:"vatZoneId,omitempty"`
	GLN                          StringValue `json:"gln,omitempty"`
	Note                         StringValue `json:"note,omitempty"`
	MainAddress                  *Address    `json:"mainAddress,omitempty"`
	MainContact                  *Contact    `json:"mainContact,omitempty"`
	CreditVerification           StringValue `json:"creditVerification,omitempty"`
	InvoiceAddress               *Address    `json:"invoiceAddress,omitempty"`
	InvoiceContact               *Contact    `json:"invoiceContact,omitempty"`
	StatementType                StringValue `json:"statementType,omitempty"`
	DeliveryAddress              *Address    `json:"deliveryAddress,omitempty"`
	DeliveryContact              *Contact    `json:"deliveryContact,omitempty"`
	PriceClassID                 StringValue `json:"priceClassId,omitempty"`
	OverrideNumberSeries         BoolValue   `json:"overrideNumberSeries,omitempty"`
	ExcludeDebtCollection        BoolValue   `json:"excludeDebtCollection,omitempty"`
	//TODO: implement overrideWithClassValues? bool with no value wrapper...
	//TODO: implement eInvoiceContract?
	//TODO: implement defaultPaymentMethod?
	//TODO: implement glAccounts?
	//TODO: implement directDebitLines?
	//TODO: implement attributeLines?
}

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
	var customer Customer
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
	Customer Customer
}

// newPostCustomerV1Request creates a new PostCustomerV1Request
func newPostCustomerV1Request(c *Client) PostCustomerV1Request {
	return PostCustomerV1Request{
		Client: c,
		Method: "POST",
		Path:   "controller/api/v1/customer",
	}
}

// PostCustomerV1Request represents a request to get a customer
type PostCustomerV1Request Request

// SetBody sets the body of the request
func (r *PostCustomerV1Request) SetBody(body CustomerRequestBody) {
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
