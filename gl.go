package vantagepoint

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// GLSummary represents a general ledger summary record from the GLSummary
// endpoint, containing aggregated account balances for a given period.
type GLSummary struct {
	Account             string  `json:"Account,omitempty"`
	Name                string  `json:"Name,omitempty"`
	Period              int     `json:"Period,omitempty"`
	Org                 string  `json:"Org,omitempty"`
	Amount              float64 `json:"Amount,omitempty"`
	CBAmount            float64 `json:"CBAmount,omitempty"`
	CreditAmount        float64 `json:"CreditAmount,omitempty"`
	DebitAmount         float64 `json:"DebitAmount,omitempty"`
	CBCreditAmount      float64 `json:"CBCreditAmount,omitempty"`
	CBDebitAmount       float64 `json:"CBDebitAmount,omitempty"`
	TransType           string  `json:"TransType,omitempty"`
	Status              string  `json:"Status,omitempty"`
	Type                int     `json:"Type,omitempty"`
	CashBasisAccount    string  `json:"CashBasisAccount,omitempty"`
	Detail              string  `json:"Detail,omitempty"`
	FasAccount          string  `json:"FasAccount,omitempty"`
	AccountCurrencyCode string  `json:"AccountCurrencyCode,omitempty"`
	GlobalAccount       string  `json:"GlobalAccount,omitempty"`
}

// LaborDetail represents a posted labor transaction from the LaborDetail
// endpoint, containing hours, amounts, and billing information for a single
// labor entry.
type LaborDetail struct {
	Period                            int     `json:"Period,omitempty"`
	PostSeq                           int     `json:"PostSeq,omitempty"`
	PKey                              string  `json:"PKey,omitempty"`
	WBS1                              string  `json:"WBS1,omitempty"`
	WBS2                              string  `json:"WBS2,omitempty"`
	WBS3                              string  `json:"WBS3,omitempty"`
	LaborCode                         string  `json:"LaborCode,omitempty"`
	Employee                          string  `json:"Employee,omitempty"`
	TransType                         string  `json:"TransType,omitempty"`
	TransDate                         string  `json:"TransDate,omitempty"`
	Name                              string  `json:"Name,omitempty"`
	RegHrs                            float64 `json:"RegHrs,omitempty"`
	OvtHrs                            float64 `json:"OvtHrs,omitempty"`
	SpecialOvtHrs                     float64 `json:"SpecialOvtHrs,omitempty"`
	RegAmt                            float64 `json:"RegAmt,omitempty"`
	OvtAmt                            float64 `json:"OvtAmt,omitempty"`
	SpecialOvtAmt                     float64 `json:"SpecialOvtAmt,omitempty"`
	BillExt                           float64 `json:"BillExt,omitempty"`
	Rate                              float64 `json:"Rate,omitempty"`
	OvtPct                            float64 `json:"OvtPct,omitempty"`
	OvtRate                           float64 `json:"OvtRate,omitempty"`
	SpecialOvtPct                     float64 `json:"SpecialOvtPct,omitempty"`
	SpecialOvtRate                    float64 `json:"SpecialOvtRate,omitempty"`
	EmType                            string  `json:"EmType,omitempty"`
	Pool                              int     `json:"Pool,omitempty"`
	Category                          int     `json:"Category,omitempty"`
	EmOrg                             string  `json:"EmOrg,omitempty"`
	PrOrg                             string  `json:"PrOrg,omitempty"`
	ChargeType                        string  `json:"ChargeType,omitempty"`
	RateType                          string  `json:"RateType,omitempty"`
	DebitLedgerMiscPKey               string  `json:"DebitLedgerMiscPKey,omitempty"`
	CreditLedgerMiscPKey              string  `json:"CreditLedgerMiscPKey,omitempty"`
	SuppressBill                      string  `json:"SuppressBill,omitempty"`
	BillStatus                        string  `json:"BillStatus,omitempty"`
	Comment                           string  `json:"Comment,omitempty"`
	BilledWBS1                        string  `json:"BilledWBS1,omitempty"`
	BilledWBS2                        string  `json:"BilledWBS2,omitempty"`
	BilledWBS3                        string  `json:"BilledWBS3,omitempty"`
	BilledInvoice                     string  `json:"BilledInvoice,omitempty"`
	BilledPeriod                      int     `json:"BilledPeriod,omitempty"`
	XferWBS1                          string  `json:"XferWBS1,omitempty"`
	XferWBS2                          string  `json:"XferWBS2,omitempty"`
	XferWBS3                          string  `json:"XferWBS3,omitempty"`
	XferLaborCode                     string  `json:"XferLaborCode,omitempty"`
	ProjectCost                       string  `json:"ProjectCost,omitempty"`
	Locale                            string  `json:"Locale,omitempty"`
	RegAmtProjectCurrency             float64 `json:"RegAmtProjectCurrency,omitempty"`
	OvtAmtProjectCurrency             float64 `json:"OvtAmtProjectCurrency,omitempty"`
	SpecialOvtAmtProjectCurrency      float64 `json:"SpecialOvtAmtProjectCurrency,omitempty"`
	RateProjectCurrency               float64 `json:"RateProjectCurrency,omitempty"`
	OvtRateProjectCurrency            float64 `json:"OvtRateProjectCurrency,omitempty"`
	SpecialOvtRateProjectCurrency     float64 `json:"SpecialOvtRateProjectCurrency,omitempty"`
	ProjectExchangeInfo               string  `json:"ProjectExchangeInfo,omitempty"`
	RegAmtBillingCurrency             float64 `json:"RegAmtBillingCurrency,omitempty"`
	OvtAmtBillingCurrency             float64 `json:"OvtAmtBillingCurrency,omitempty"`
	SpecialOvtAmtBillingCurrency      float64 `json:"SpecialOvtAmtBillingCurrency,omitempty"`
	RateBillingCurrency               float64 `json:"RateBillingCurrency,omitempty"`
	OvtRateBillingCurrency            float64 `json:"OvtRateBillingCurrency,omitempty"`
	SpecialOvtRateBillingCurrency     float64 `json:"SpecialOvtRateBillingCurrency,omitempty"`
	BillingExchangeInfo               string  `json:"BillingExchangeInfo,omitempty"`
	RegAmtEmployeeCurrency            float64 `json:"RegAmtEmployeeCurrency,omitempty"`
	OvtAmtEmployeeCurrency            float64 `json:"OvtAmtEmployeeCurrency,omitempty"`
	SpecialOvtAmtEmployeeCurrency     float64 `json:"SpecialOvtAmtEmployeeCurrency,omitempty"`
	RateEmployeeCurrency              float64 `json:"RateEmployeeCurrency,omitempty"`
	OvtRateEmployeeCurrency           float64 `json:"OvtRateEmployeeCurrency,omitempty"`
	SpecialOvtRateEmployeeCurrency    float64 `json:"SpecialOvtRateEmployeeCurrency,omitempty"`
	ExchangeInfo                      string  `json:"ExchangeInfo,omitempty"`
	EMCurrencyCode                    string  `json:"EMCurrencyCode,omitempty"`
	Payrate                           float64 `json:"Payrate,omitempty"`
	PayOvtPct                         float64 `json:"PayOvtPct,omitempty"`
	PaySpecialOvtPct                  float64 `json:"PaySpecialOvtPct,omitempty"`
	NonBill                           string  `json:"NonBill,omitempty"`
	InvoiceStatus                     string  `json:"InvoiceStatus,omitempty"`
	Rowversion                        string  `json:"Rowversion,omitempty"`
	RealizationAmountEmployeeCurrency float64 `json:"RealizationAmountEmployeeCurrency,omitempty"`
	RealizationAmountProjectCurrency  float64 `json:"RealizationAmountProjectCurrency,omitempty"`
	RealizationAmountBillingCurrency  float64 `json:"RealizationAmountBillingCurrency,omitempty"`
	ModUser                           string  `json:"ModUser,omitempty"`
	AuthorizedBy                      string  `json:"AuthorizedBy,omitempty"`
	BillTaxCodeOverride               string  `json:"BillTaxCodeOverride,omitempty"`
	BillTax2CodeOverride              string  `json:"BillTax2CodeOverride,omitempty"`
	SelPeriod                         int     `json:"SelPeriod,omitempty"`
	SelPostSeq                        int     `json:"SelPostSeq,omitempty"`
	SelOvtPeriod                      int     `json:"SelOvtPeriod,omitempty"`
	SelOvtPostSeq                     int     `json:"SelOvtPostSeq,omitempty"`
	WrittenOffPeriod                  int     `json:"WrittenOffPeriod,omitempty"`
	CostRateTableUsed                 string  `json:"CostRateTableUsed,omitempty"`
	TimekeeperEndDate                 string  `json:"TimekeeperEndDate,omitempty"`
	TransferredPeriod                 int     `json:"TransferredPeriod,omitempty"`
	TransferredBillStatus             string  `json:"TransferredBillStatus,omitempty"`
	XferCategory                      int     `json:"XferCategory,omitempty"`
	TLInternalKey                     string  `json:"TLInternalKey,omitempty"`
	TLProcessed                       string  `json:"TLProcessed,omitempty"`
	LastPSAExportDate                 string  `json:"LastPSAExportDate,omitempty"`
}

// PSALedgerEntry represents a single posted transaction from the PSALedger
// endpoint, covering labor, expense, AP, and miscellaneous transaction types.
type PSALedgerEntry struct {
	Batch                             string  `json:"Batch,omitempty"`
	Period                            int     `json:"Period,omitempty"`
	PostSeq                           int     `json:"PostSeq,omitempty"`
	PKey                              string  `json:"PKey,omitempty"`
	WBS1                              string  `json:"WBS1,omitempty"`
	WBS2                              string  `json:"WBS2,omitempty"`
	WBS3                              string  `json:"WBS3,omitempty"`
	Account                           string  `json:"Account,omitempty"`
	Org                               string  `json:"Org,omitempty"`
	TransType                         string  `json:"TransType,omitempty"`
	SubType                           string  `json:"SubType,omitempty"`
	RefNo                             string  `json:"RefNo,omitempty"`
	TransDate                         string  `json:"TransDate,omitempty"`
	Desc1                             string  `json:"Desc1,omitempty"`
	Desc2                             string  `json:"Desc2,omitempty"`
	Amount                            float64 `json:"Amount,omitempty"`
	CBAmount                          float64 `json:"CBAmount,omitempty"`
	BillExt                           float64 `json:"BillExt,omitempty"`
	ProjectCost                       string  `json:"ProjectCost,omitempty"`
	AutoEntry                         string  `json:"AutoEntry,omitempty"`
	SuppressBill                      string  `json:"SuppressBill,omitempty"`
	BillStatus                        string  `json:"BillStatus,omitempty"`
	SkipGL                            string  `json:"SkipGL,omitempty"`
	BankCode                          string  `json:"BankCode,omitempty"`
	Invoice                           string  `json:"Invoice,omitempty"`
	InvoiceSection                    string  `json:"InvoiceSection,omitempty"`
	Employee                          string  `json:"Employee,omitempty"`
	Vendor                            string  `json:"Vendor,omitempty"`
	Line                              int     `json:"Line,omitempty"`
	PartialPayment                    float64 `json:"PartialPayment,omitempty"`
	Discount                          float64 `json:"Discount,omitempty"`
	Voucher                           string  `json:"Voucher,omitempty"`
	BilledWBS1                        string  `json:"BilledWBS1,omitempty"`
	BilledWBS2                        string  `json:"BilledWBS2,omitempty"`
	BilledWBS3                        string  `json:"BilledWBS3,omitempty"`
	BilledInvoice                     string  `json:"BilledInvoice,omitempty"`
	BilledPeriod                      int     `json:"BilledPeriod,omitempty"`
	Unit                              string  `json:"Unit,omitempty"`
	UnitTable                         string  `json:"UnitTable,omitempty"`
	UnitQuantity                      float64 `json:"UnitQuantity,omitempty"`
	UnitCostRate                      float64 `json:"UnitCostRate,omitempty"`
	UnitBillingRate                   float64 `json:"UnitBillingRate,omitempty"`
	UnitBillExt                       float64 `json:"UnitBillExt,omitempty"`
	XferWBS1                          string  `json:"XferWBS1,omitempty"`
	XferWBS2                          string  `json:"XferWBS2,omitempty"`
	XferWBS3                          string  `json:"XferWBS3,omitempty"`
	XferAccount                       string  `json:"XferAccount,omitempty"`
	TaxCode                           string  `json:"TaxCode,omitempty"`
	TaxBasis                          float64 `json:"TaxBasis,omitempty"`
	TaxCBBasis                        float64 `json:"TaxCBBasis,omitempty"`
	WrittenOffPeriod                  int     `json:"WrittenOffPeriod,omitempty"`
	TransactionAmount                 float64 `json:"TransactionAmount,omitempty"`
	TransactionCurrencyCode           string  `json:"TransactionCurrencyCode,omitempty"`
	ExchangeInfo                      string  `json:"ExchangeInfo,omitempty"`
	AmountProjectCurrency             float64 `json:"AmountProjectCurrency,omitempty"`
	ProjectExchangeInfo               string  `json:"ProjectExchangeInfo,omitempty"`
	AmountBillingCurrency             float64 `json:"AmountBillingCurrency,omitempty"`
	BillingExchangeInfo               string  `json:"BillingExchangeInfo,omitempty"`
	AutoEntryAmount                   float64 `json:"AutoEntryAmount,omitempty"`
	AutoEntryExchangeInfo             string  `json:"AutoEntryExchangeInfo,omitempty"`
	AutoEntryOrg                      string  `json:"AutoEntryOrg,omitempty"`
	AutoEntryAccount                  string  `json:"AutoEntryAccount,omitempty"`
	AmountSourceCurrency              float64 `json:"AmountSourceCurrency,omitempty"`
	SourceExchangeInfo                string  `json:"SourceExchangeInfo,omitempty"`
	PONumber                          string  `json:"PONumber,omitempty"`
	UnitCostRateBillingCurrency       float64 `json:"UnitCostRateBillingCurrency,omitempty"`
	LinkCompany                       string  `json:"LinkCompany,omitempty"`
	LinkWBS1                          string  `json:"LinkWBS1,omitempty"`
	LinkWBS2                          string  `json:"LinkWBS2,omitempty"`
	LinkWBS3                          string  `json:"LinkWBS3,omitempty"`
	Diary                             string  `json:"Diary,omitempty"`
	DiaryNo                           int     `json:"DiaryNo,omitempty"`
	ModUser                           string  `json:"ModUser,omitempty"`
	AuthorizedBy                      string  `json:"AuthorizedBy,omitempty"`
	NonBill                           string  `json:"NonBill,omitempty"`
	CreditMemoRefNo                   string  `json:"CreditMemoRefNo,omitempty"`
	InvoiceStatus                     string  `json:"InvoiceStatus,omitempty"`
	OriginatingVendor                 string  `json:"OriginatingVendor,omitempty"`
	BillTaxCodeOverride               string  `json:"BillTaxCodeOverride,omitempty"`
	BillTax2CodeOverride              string  `json:"BillTax2CodeOverride,omitempty"`
	GainsAndLossesType                string  `json:"GainsAndLossesType,omitempty"`
	TransferredPeriod                 int     `json:"TransferredPeriod,omitempty"`
	TransferredBillStatus             string  `json:"TransferredBillStatus,omitempty"`
	RealizationAmountEmployeeCurrency float64 `json:"RealizationAmountEmployeeCurrency,omitempty"`
	RealizationAmountProjectCurrency  float64 `json:"RealizationAmountProjectCurrency,omitempty"`
	RealizationAmountBillingCurrency  float64 `json:"RealizationAmountBillingCurrency,omitempty"`
	DiscountFunctionalCurrency        float64 `json:"DiscountFunctionalCurrency,omitempty"`
	OriginalAmountSourceCurrency      float64 `json:"OriginalAmountSourceCurrency,omitempty"`
	OriginalPaymentCurrencyCode       string  `json:"OriginalPaymentCurrencyCode,omitempty"`
	PreInvoice                        string  `json:"PreInvoice,omitempty"`
	EmOrg                             string  `json:"EmOrg,omitempty"`
	LastPSAExportDate                 string  `json:"LastPSAExportDate,omitempty"`
}

// GLSummaryQuery extends Query with GL-specific parameters for column
// summation and grouping. Use NewGLSummaryQuery to create one.
type GLSummaryQuery struct {
	*Query
	sumColumns []string
	groupBy    []string
}

// NewGLSummaryQuery creates a new GLSummaryQuery with an embedded Query
// builder. All standard Query methods (Limit, Filter, Fields, etc.) are
// available via embedding.
func NewGLSummaryQuery() *GLSummaryQuery {
	return &GLSummaryQuery{
		Query: NewQuery(),
	}
}

// SumColumns specifies which numeric columns to aggregate in the GL summary
// response. These are sent as the sumColumns query parameter.
func (q *GLSummaryQuery) SumColumns(cols ...string) *GLSummaryQuery {
	q.sumColumns = append(q.sumColumns, cols...)
	return q
}

// GroupBy specifies which columns to group the GL summary results by. These
// are sent as the groupBy query parameter.
func (q *GLSummaryQuery) GroupBy(cols ...string) *GLSummaryQuery {
	q.groupBy = append(q.groupBy, cols...)
	return q
}

// Values encodes the GLSummaryQuery into url.Values, including the base Query
// parameters plus the GL-specific sumColumns and groupBy parameters.
func (q *GLSummaryQuery) Values() url.Values {
	var v url.Values
	if q.Query != nil {
		v = q.Query.Values()
	} else {
		v = make(url.Values)
	}
	if len(q.sumColumns) > 0 {
		v.Set("sumColumns", strings.Join(q.sumColumns, ","))
	}
	if len(q.groupBy) > 0 {
		v.Set("groupBy", strings.Join(q.groupBy, ","))
	}
	return v
}

// ListGLSummary retrieves general ledger summary records from the GLSummary
// endpoint. The GLSummaryQuery allows specifying columns to sum and group by
// in addition to standard query parameters.
func (c *Client) ListGLSummary(ctx context.Context, q *GLSummaryQuery) ([]GLSummary, error) {
	var vals url.Values
	if q != nil {
		vals = q.Values()
	}
	var results []GLSummary
	if err := c.do(ctx, "GET", "GLSummary", vals, nil, &results); err != nil {
		return nil, fmt.Errorf("listing GL summary: %w", err)
	}
	return results, nil
}

// ListLaborDetail retrieves posted labor detail records from the LaborDetail
// endpoint. Each record represents a single posted labor transaction with
// hours, amounts, and billing information.
func (c *Client) ListLaborDetail(ctx context.Context, q *Query) ([]LaborDetail, error) {
	var results []LaborDetail
	if err := c.get(ctx, "LaborDetail", q, &results); err != nil {
		return nil, fmt.Errorf("listing labor detail: %w", err)
	}
	return results, nil
}

// ListPSALedger retrieves posted PSA ledger entries for the given transaction
// type from the PSALedger/{transactionType} endpoint. Transaction types use
// two-letter prefixes: "AP" (AP vouchers), "EX" (employee expenses), "TS"
// (timesheets), "MI" (miscellaneous expenses), "EP" (employee payments),
// "JE" (journal entries), etc.
func (c *Client) ListPSALedger(ctx context.Context, transactionType string, q *Query) ([]PSALedgerEntry, error) {
	var results []PSALedgerEntry
	if err := c.get(ctx, "PSALedger/"+transactionType, q, &results); err != nil {
		return nil, fmt.Errorf("listing PSA ledger for type %s: %w", transactionType, err)
	}
	return results, nil
}
