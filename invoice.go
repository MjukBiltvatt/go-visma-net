package vismanet

import "encoding/json"

// ResponseInvoice is an invoice as represented in a response from the Visma.net API
type ResponseInvoice struct {
	DocumentDueDate           Time                `json:"documentDueDate"`
	CashDiscountDate          Time                `json:"cashDiscountDate"`
	ExternalReference         string              `json:"externalReference"`
	IsRotRutDeductible        bool                `json:"isRotRutDeductible"`
	ExchangeRate              float64             `json:"exchangeRate"`
	DunningLetterDate         Time                `json:"dunningLetterDate"`
	DuningLetterLevel         float64             `json:"dunningLetterLevel"`
	SendToAutoInvoice         bool                `json:"sendToAutoInvoice"`
	RoundingDiff              float64             `json:"roundingDiff"`
	StartDate                 Time                `json:"startDate"`
	EndDate                   Time                `json:"endDate"`
	AccountingCostRef         string              `json:"accountingCostRef"`
	OriginatorDocRef          string              `json:"originatorDocRef"`
	ContractDocRef            string              `json:"contractDocRef"`
	ExcludeDebtCollection     bool                `json:"excludeDebtCollection"`
	Hold                      bool                `json:"hold"`
	DiscountTotal             float64             `json:"discountTotal"`
	DiscountTotalInCurrency   float64             `json:"discountTotalInCurrency"`
	DetailTotal               float64             `json:"detailTotal"`
	DetailTotalInCurrency     float64             `json:"detailTotalInCurrency"`
	VatTaxableTotal           float64             `json:"vatTaxableTotal"`
	VatTaxableTotalInCurrency float64             `json:"vatTaxableTotalInCurrency"`
	VatExemptTotal            float64             `json:"vatExemptTotal"`
	VatExemptTotalInCurrency  float64             `json:"vatExemptTotalInCurrency"`
	SalesPersonID             int                 `json:"salesPersonID"` //TODO: double check this field name is correct
	SalesPersonDescription    string              `json:"salesPersonDescr"`
	PaymentReference          string              `json:"paymentReference"`
	DontPrint                 bool                `json:"dontPrint"`
	DontEmail                 bool                `json:"dontEmail"`
	Revoked                   bool                `json:"revoked"`
	DocumentType              string              `json:"documentType"`
	ReferenceNumber           string              `json:"referenceNumber"`
	PostPeriod                string              `json:"postPeriod"`
	FinancialPeriod           string              `json:"financialPeriod"`
	ClosedFinancialPeriod     string              `json:"closedFinancialPeriod"`
	DocumentDate              Time                `json:"documentDate"`
	OrigInvoiceDate           Time                `json:"origInvoiceDate"`
	Status                    string              `json:"status"`
	CurrencyID                string              `json:"currencyId"`
	Amount                    float64             `json:"amount"`
	AmountInCurrency          float64             `json:"amountInCurrency"`
	Balance                   float64             `json:"balance"`
	BalanceInCurrency         float64             `json:"balanceInCurrency"`
	CashDiscount              float64             `json:"cashDiscount"`
	CashDiscountInCurrency    float64             `json:"cashDiscountInCurrency"`
	CustomerRefNumber         string              `json:"customerRefNumber"`
	InvoiceText               string              `json:"invoiceText"`
	LastModifiedDateTime      Time                `json:"lastModifiedDateTime"`
	CreatedDateTime           Time                `json:"createdDateTime"`
	Note                      string              `json:"note"`
	VatTotal                  float64             `json:"vatTotal"`
	VatTotalInCurrency        float64             `json:"vatTotalInCurrency"`
	CashAccount               string              `json:"cashAccount"`
	CustomerProject           string              `json:"customerProject"`
	ErrorInfo                 string              `json:"errorInfo"`
	InvoiceAddress            Address             `json:"invoiceAddress"`
	Attachments               []Attachment        `json:"attachments"`
	Contact                   IDNameEntity        `json:"contact"`
	Customer                  NumberNameEntity    `json:"customer"`
	ChildRecord               NumberNameEntity    `json:"childRecord"`
	PaymentMethod             IDDescriptionEntity `json:"paymentMethod"`
	CreditTerms               IDDescriptionEntity `json:"creditTerms"`
	CustomerVatZone           IDDescriptionEntity `json:"customerVatZone"`
	SalesPerson               IDDescriptionEntity `json:"salesPerson"`
	Project                   IDDescriptionEntity `json:"project"`
	Metadata                  Metadata            `json:"metadata"`

	DebtCollection struct {
		CaseNumber string `json:"caseNumber"`
		CaseType   string `json:"caseType"`
		CaseStatus string `json:"caseStatus"`
		CaseURL    string `json:"caseUrl"`
	} `json:"debtCollection"`

	DirectDebitMandate struct {
		ID                 string `json:"id"`
		MandateID          string `json:"mandateId"`
		MandateDescription string `json:"mandateDescription"`
	} `json:"directDebitMandate"`

	InvoiceContact struct {
		ContactID       int    `json:"contactId"`
		BusinessName    string `json:"businessName"`
		Attention       string `json:"attention"`
		Email           string `json:"email"`
		Phone1          string `json:"phone1"`
		OverrideContact bool   `json:"overrideContact"`
	} `json:"invoiceContact"`

	TaxDetails []struct {
		TaxID         string  `json:"taxId"`
		RecordID      int     `json:"recordId"`
		VatRate       float64 `json:"vatRate"`
		TaxableAmount float64 `json:"taxableAmount"`
		VatAmount     float64 `json:"vatAmount"`
		VatID         struct {
			Number      string `json:"number"`
			Description string `json:"description"`
		} `json:"vatId"`
	} `json:"taxDetails"`

	InvoiceLines []struct {
		TermStartDate            Time                `json:"termStartDate"`
		TermEndDate              Time                `json:"termEndDate"`
		IsRotRutDeductible       bool                `json:"isRotRutDeductible"`
		ItemType                 string              `json:"itemType"`
		DeductableAmount         float64             `json:"deductableAmount"`
		LineType                 string              `json:"lineType"`
		StartDate                Time                `json:"startDate"`
		EndDate                  Time                `json:"endDate"`
		SOShipmentNumber         string              `json:"soShipmentNbr"`
		SOOrderType              string              `json:"soOrderType"`
		SOOrderNumber            string              `json:"soOrderNbr"`
		LineNumber               int                 `json:"lineNumber"`
		InventoryNumber          string              `json:"inventoryNumber"`
		Description              string              `json:"description"`
		Quantity                 float64             `json:"quantity"`
		UnitPrice                float64             `json:"unitPrice"`
		UnitPriceInCurrency      float64             `json:"unitPriceInCurrency"`
		ManualAmount             float64             `json:"manualAmount"`
		ManualAmountInCurrency   float64             `json:"manualAmountInCurrency"`
		Amount                   float64             `json:"amount"`
		Cost                     float64             `json:"cost"`
		AmountInCurrency         float64             `json:"amountInCurrency"`
		UOM                      string              `json:"uom"`
		DiscountPercent          float64             `json:"discountPercent"`
		DiscountAmount           float64             `json:"discountAmount"`
		DiscountAmountInCurrency float64             `json:"discountAmountInCurrency"`
		ManualDiscount           bool                `json:"manualDiscount"`
		SalesPerson              string              `json:"salesperson"`
		DeferralSchedule         int                 `json:"deferralSchedule"`
		DeferralCode             string              `json:"deferralCode"`
		DiscountCode             string              `json:"discountCode"`
		Note                     string              `json:"note"`
		ExternalLink             string              `json:"externalLink"`
		Attachments              []Attachment        `json:"attachments"`
		ProjectTask              IDDescriptionEntity `json:"projectTask"`
		VatCode                  IDDescriptionEntity `json:"vatCode"`
		Seller                   IDDescriptionEntity `json:"seller"`
		BranchNumber             NumberNameEntity    `json:"branchNumber"`

		Subaccount struct {
			SubaccountNumber     string    `json:"subaccountNumber"`
			SubaccountID         int       `json:"subaccountId"`
			Description          string    `json:"description"`
			Active               bool      `json:"active"`
			ErrorInfo            string    `json:"errorInfo"`
			LastModifiedDateTime Time      `json:"lastModifiedDateTime"`
			Segments             []Segment `json:"segments"`
			Metadata             Metadata  `json:"metadata"`
		} `json:"subaccount"`

		Account struct {
			Type          string `json:"type"`
			ExternalCode1 string `json:"externalCode1"`
			ExternalCode2 string `json:"externalCode2"`
			Number        string `json:"number"`
			Description   string `json:"description"`
		} `json:"account"`

		TypeOfWork struct {
			RutRotType  string `json:"rutRotType"`
			Description string `json:"description"`
			XMLTag      string `json:"xmlTag"`
		} `json:"typeOfWork"`
	} `json:"invoiceLines"`

	Applications []struct {
		DocType           string  `json:"docType"`
		CustomerCD        string  `json:"customerCD"`
		RefNbr            string  `json:"refNbr"`
		AmountPaid        float64 `json:"amountPaid"`
		CashDiscountTaken float64 `json:"cashDiscountTaken"`
		Balance           float64 `json:"balance"`
		PendingPPD        bool    `json:"pendingPPD"`
		Released          bool    `json:"released"`
		Hold              bool    `json:"hold"`
		Voided            bool    `json:"voided"`
		PPDCrMemoRefNbr   string  `json:"ppdCrMemoRefNbr"`
		PaymentRef        string  `json:"paymentRef"`
		Status            string  `json:"status"`
		ApplicationDate   Time    `json:"applicationDate"`
		ApplicationPeriod string  `json:"applicationPeriod"`
		InvoiceText       string  `json:"invoiceText"`
	} `json:"applications"`

	Location struct {
		ID        string `json:"id"`
		CountryID string `json:"countryId"`
		Name      string `json:"name"`
	} `json:"location"`
}

// RequestInvoice is an invoice as represented in requests to the Visma.net API
type RequestInvoice struct {
	PaymentMethodID                   StringValue                   `json:"paymentMethodId,omitempty"`
	CreditTermsID                     StringValue                   `json:"creditTermsId,omitempty"`
	CurrencyID                        StringValue                   `json:"currencyId,omitempty"`
	CustomerRefNumber                 StringValue                   `json:"customerRefNumber,omitempty"`
	CashDiscountDate                  TimeValue                     `json:"cashDiscountDate,omitempty"`
	DocumentDueDate                   TimeValue                     `json:"documentDueDate,omitempty"`
	ExternalReference                 StringValue                   `json:"externalReference,omitempty"`
	CustomerProject                   StringValue                   `json:"customerProject,omitempty"`
	ExchangeRate                      FloatValue                    `json:"exchangeRate,omitempty"`
	DomesticServiceDeductibleDocument BoolValue                     `json:"domesticServiceDeductibleDocument,omitempty"`
	RotRutDetails                     *RequestInvoiceRotRutDetails  `json:"rotRutDetails,omitempty"`
	PaymentReference                  StringValue                   `json:"paymentReference,omitempty"`
	Contact                           IntValue                      `json:"contact,omitempty"` //What even is this?
	Project                           StringValue                   `json:"project,omitempty"`
	TaxDetailLines                    []RequestInvoiceTaxDetailLine `json:"taxDetailLines,omitempty"`
	InvoiceLines                      []RequestInvoiceLine          `json:"invoiceLines,omitempty"`
	SendToAutoInvoice                 BoolValue                     `json:"sendToAutoInvoice,omitempty"`
	CustomerVatZoneID                 StringValue                   `json:"customerVatZoneId,omitempty"`
	BillingAddress                    *RequestNestedAddress         `json:"billingAddress,omitempty"`
	InvoiceContact                    *RequestInvoiceContact        `json:"invoiceContact,omitempty"`
	StartDate                         TimeValue                     `json:"startDate,omitempty"`
	EndDate                           TimeValue                     `json:"endDate,omitempty"`
	AccountingCostRef                 StringValue                   `json:"accountingCostRef,omitempty"`
	OriginatorDocRef                  StringValue                   `json:"originatorDocRef,omitempty"`
	ContractDocRef                    StringValue                   `json:"contractDocRef,omitempty"`
	Revoked                           BoolValue                     `json:"revoked,omitempty"`
	OverrideNumberSeries              BoolValue                     `json:"overrideNumberSeries,omitempty"`
	DirectDebitMandateID              StringValue                   `json:"directDebitMandateId,omitempty"`
	ExcludeDebtCollection             BoolValue                     `json:"excludeDebtCollection,omitempty"`
	DebtCollectionCaseNumber          StringValue                   `json:"debtCollectionCaseNbr,omitempty"`
	DebtCollectionCaseStatus          StringValue                   `json:"debtCollectionCaseStatus,omitempty"`
	DebtCollectionCaseType            StringValue                   `json:"debtCollectionCaseType,omitempty"`
	DebtCollectionCaseURL             StringValue                   `json:"debtCollectionCaseUrl,omitempty"`
	ReferenceNumber                   StringValue                   `json:"referenceNumber,omitempty"`
	CustomerNumber                    StringValue                   `json:"customerNumber"`
	ChildCustomerNumber               StringValue                   `json:"childCustomerNumber,omitempty"`
	ConsolidateInvoices               BoolValue                     `json:"consolidateInvoices,omitempty"`
	DocumentDate                      TimeValue                     `json:"documentDate,omitempty"`
	OriginalDocumentDate              TimeValue                     `json:"origDocumentDate,omitempty"`
	Hold                              BoolValue                     `json:"hold,omitempty"`
	PostPeriod                        StringValue                   `json:"postPeriod,omitempty"`
	FinancialPeriod                   StringValue                   `json:"financialPeriod,omitempty"`
	InvoiceText                       StringValue                   `json:"invoiceText,omitempty"`
	LocationID                        StringValue                   `json:"locationId,omitempty"`
	SalesPersonID                     IntValue                      `json:"salesPersonId,omitempty"`
	Note                              StringValue                   `json:"note,omitempty"`
	BranchNumber                      StringValue                   `json:"branchNumber,omitempty"`
	CashAccount                       StringValue                   `json:"cashAccount,omitempty"`
	DontPrint                         BoolValue                     `json:"dontPrint,omitempty"`
	DontEmail                         BoolValue                     `json:"dontEmail,omitempty"`
}

// RequestInvoiceLine is an invoice line as represented in requests to the Visma.net API
type RequestInvoiceLine struct {
	Operation                 string           `json:"operation,omitempty"`
	DiscountCode              StringValue      `json:"discountCode,omitempty"`
	DomesticServiceDeductible BoolValue        `json:"domesticServiceDeductible,omitempty"`
	ItemType                  StringValue      `json:"itemType,omitempty"`
	TypeOfWork                StringValue      `json:"typeOfWork,omitempty"`
	TaskID                    StringValue      `json:"taskId,omitempty"`
	StartDate                 TimeValue        `json:"startDate,omitempty"`
	EndDate                   TimeValue        `json:"endDate,omitempty"`
	InventoryNumber           StringValue      `json:"inventoryNumber,omitempty"`
	LineNumber                IntValue         `json:"lineNumber,omitempty"`
	Description               StringValue      `json:"description,omitempty"`
	Quantity                  FloatValue       `json:"quantity,omitempty"`
	UnitPriceInCurrency       FloatValue       `json:"unitPriceInCurrency,omitempty"`
	ManualAmountInCurrency    FloatValue       `json:"manualAmountInCurrency,omitempty"`
	AccountNumber             StringValue      `json:"accountNumber,omitempty"`
	VatCodeID                 StringValue      `json:"vatCodeId,omitempty"`
	UOM                       StringValue      `json:"uom,omitempty"`
	DiscountPercent           FloatValue       `json:"discountPercent,omitempty"`
	DiscountAmountInCurrency  FloatValue       `json:"discountAmountInCurrency,omitempty"`
	ManualDiscount            BoolValue        `json:"manualDiscount,omitempty"`
	Subaccount                []RequestSegment `json:"subaccount,omitempty"`
	SalesPerson               StringValue      `json:"salesPerson,omitempty"`
	DeferralSchedule          IntValue         `json:"deferralSchedule,omitempty"`
	DeferralCode              StringValue      `json:"deferralCode,omitempty"`
	TermStartDate             TimeValue        `json:"termStartDate,omitempty"`
	TermEndDate               TimeValue        `json:"termEndDate,omitempty"`
	Note                      StringValue      `json:"note,omitempty"`
	BranchNumber              StringValue      `json:"branchNumber,omitempty"`
}

type RequestInvoiceRotRutDetails struct {
	DistributedAutomatically BoolValue                                `json:"distributedAutomaticaly,omitempty"` //TODO: double check this field name is correct, it is spelled wrong in the API docs
	Type                     StringValue                              `json:"type,omitempty"`
	Apartment                StringValue                              `json:"apartment,omitempty"` //TODO: double check this field name is correct, it is spelled wrong in the API docs
	Estate                   StringValue                              `json:"estate,omitempty"`
	OrganizationNumber       StringValue                              `json:"organizationNbr,omitempty"`
	Distribution             *RequestInvoiceRotRutDetailsDistribution `json:"distribution,omitempty"`
}

type RequestInvoiceRotRutDetailsDistribution struct {
	Operation  string      `json:"operation,omitempty"`
	LineNumber IntValue    `json:"lineNbr,omitempty"`
	PersonalID StringValue `json:"personalId,omitempty"`
	Amount     FloatValue  `json:"amount,omitempty"`
	Extra      BoolValue   `json:"extra,omitempty"`
}

// RequestInvoiceTaxDetailLine is a tax detail line as represented in requests to the Visma.net API
type RequestInvoiceTaxDetailLine struct {
	TaxID         StringValue `json:"taxId,omitempty"`
	TaxableAmount FloatValue  `json:"taxableAmount,omitempty"`
	VatAmount     FloatValue  `json:"vatAmount,omitempty"`
}

// RequestInvoiceContact is an invoice contact nested as represented in requests to the Visma.net API
type RequestInvoiceContact struct {
	OverrideContact BoolValue   `json:"overrideContact,omitempty"`
	Name            StringValue `json:"name,omitempty"`
	Attention       StringValue `json:"attention,omitempty"`
	Email           StringValue `json:"email,omitempty"`
	Phone1          StringValue `json:"phone1,omitempty"`
}

func (c *RequestInvoiceContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(Value{*c})
}

// =========================================================
// ========================== GET ==========================
// =========================================================

// newGetInvoiceV1Request creates a new GetInvoiceV1Request
func newGetInvoiceV1Request(c *Client) GetInvoiceV1Request {
	return GetInvoiceV1Request{
		Client: c,
		Method: "GET",
		Path:   "controller/api/v1/customerinvoice/{{.invoice_number}}",
	}
}

// GetInvoiceV1Request represents a request to get a customer
type GetInvoiceV1Request Request

// SetPathParams sets the path parameters of the request
func (r *GetInvoiceV1Request) SetPathParams(params GetInvoiceV1PathParams) {
	r.pathParams = params
}

// Do performs the request and returns the response
func (r *GetInvoiceV1Request) Do() (GetInvoiceV1Response, error) {
	var invoices []ResponseInvoice
	var invoice ResponseInvoice
	resp, err := r.Client.Do((*Request)(r), &invoices, &invoice)
	if len(invoices) == 0 && invoice.ReferenceNumber != "" {
		invoices = append(invoices, invoice)
	}
	return GetInvoiceV1Response{Response{resp}, invoices}, err
}

// GetInvoiceV1PathParams represents the path parameters of the GetInvoiceV1Request
type GetInvoiceV1PathParams struct {
	InvoiceNumber string `schema:"invoice_number"`
}

// GetInvoiceV1Response represents the response of the GetInvoiceV1Request and contains the resulting customer
type GetInvoiceV1Response struct {
	Response
	Invoices []ResponseInvoice
}
