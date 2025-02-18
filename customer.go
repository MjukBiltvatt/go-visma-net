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
